package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

// SubcommandInfo holds the extracted details for a single subcommand.
type SubcommandInfo struct {
	Name    string // The package alias used in the source (e.g., cve_2019_5736)
	PkgPath string // The full import path (e.g., github.com/ctrsploit/ctrsploit/cmd/ctrsploit/vul/cve-2019-5736)
	VarName string // The variable name of the command (usually "Command")
}

// TemplateData is the data structure passed to the appTemplate.
type TemplateData struct {
	PkgPath string
	VarName string
}

// getPackageFilePaths uses `go list` to find the absolute paths of all .go files
// in a given package.
func getPackageFilePaths(pkgPath string) ([]string, error) {
	runGoList := func(format string) (string, error) {
		cmd := exec.Command("go", "list", "-f", format, pkgPath)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return "", fmt.Errorf("`go list -f \"%s\"` failed: %w\n%s", format, err, stderr.String())
		}
		return strings.TrimSpace(stdout.String()), nil
	}

	dir, err := runGoList("{{.Dir}}")
	if err != nil {
		return nil, err
	}
	if dir == "" {
		return nil, fmt.Errorf("could not find directory for package %s", pkgPath)
	}

	filesStr, err := runGoList("{{.GoFiles}}")
	if err != nil {
		return nil, err
	}

	filesStr = strings.Trim(filesStr, "[]\n\t ")
	if filesStr == "" {
		return []string{}, nil // Package contains no Go files
	}

	fileNames := strings.Split(filesStr, " ")
	paths := make([]string, len(fileNames))
	for i, name := range fileNames {
		paths[i] = filepath.Join(dir, name)
	}
	return paths, nil
}

// [CRITICAL FIX] Rewritten to match ctrsploit's actual code structure (pkg.Var, not &pkg.Var).
func extractSubcommandInfo(expr ast.Expr, importMap map[string]string) []SubcommandInfo {
	unary, ok := expr.(*ast.UnaryExpr)
	if !ok {
		return nil
	} // The top-level `Command` is `&cli.Command{...}`
	compLit, ok := unary.X.(*ast.CompositeLit)
	if !ok {
		return nil
	}

	// Find the "Subcommands" field within the CompositeLit
	for _, elt := range compLit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		if key, ok := kv.Key.(*ast.Ident); !ok || key.Name != "Subcommands" {
			continue
		}

		// Found the "Subcommands" field. From now on, we must return a non-nil slice to indicate success.
		subcommands := make([]SubcommandInfo, 0)
		slice, ok := kv.Value.(*ast.CompositeLit)
		if !ok {
			return subcommands
		} // Field exists but is not a slice literal, return empty.

		// Iterate over the elements of the Subcommands slice
		for _, subCmdExpr := range slice.Elts {
			// [CORE CHANGE] Directly handle SelectorExpr (e.g., `auto.Command`), no longer looking for UnaryExpr (`&auto.Command`).
			selector, ok := subCmdExpr.(*ast.SelectorExpr)
			if !ok {
				fmt.Printf("\t\t\t- ⚠️ Warning: Item in Subcommands list is not the expected 'pkg.Var' format. Skipping.\n")
				continue
			}

			pkgIdent, ok := selector.X.(*ast.Ident)
			if !ok {
				continue
			}

			pkgAlias := pkgIdent.Name
			varName := selector.Sel.Name
			pkgPath, found := importMap[pkgAlias]
			if !found {
				fmt.Printf("\t\t\t- ⚠️ Warning: Package alias '%s' not found in imports. Skipping.\n", pkgAlias)
				continue
			}
			subcommands = append(subcommands, SubcommandInfo{
				Name:    pkgAlias,
				PkgPath: pkgPath,
				VarName: varName,
			})
		}

		fmt.Printf("\t\t\t- ✅ Successfully extracted %d subcommands.\n", len(subcommands))
		return subcommands
	}

	// Traversed all fields but did not find "Subcommands". This is a structural mismatch.
	return nil
}

// parseSubcommandsFromSource is the main analysis function. It uses go/parser
// to build an AST and extract subcommand details without compiling any code.
func parseSubcommandsFromSource(pkgPath, varName string) ([]SubcommandInfo, error) {
	filePaths, err := getPackageFilePaths(pkgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to locate source files for package: %w", err)
	}
	if len(filePaths) == 0 {
		return nil, fmt.Errorf("no Go source files found in package '%s'", pkgPath)
	}

	fset := token.NewFileSet()
	for _, path := range filePaths {
		fmt.Printf("\t- Analyzing file: %s\n", filepath.Base(path))
		fileNode, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			log.Printf("Warning: failed to parse file %s: %v. Skipping.", path, err)
			continue
		}

		importMap := make(map[string]string)
		for _, imp := range fileNode.Imports {
			cleanPath := strings.Trim(imp.Path.Value, `"`)
			if imp.Name != nil {
				// Case: import alias "path/to/pkg"
				importMap[imp.Name.Name] = cleanPath
			} else {
				// Case: import "path/to/pkg-with-hyphens"
				parts := strings.Split(cleanPath, "/")
				baseName := parts[len(parts)-1]
				// [CRITICAL FIX] Mimic Go's behavior of converting directory names to package identifiers.
				// Replace all hyphens '-' with underscores '_' to create the correct lookup key.
				identifier := strings.ReplaceAll(baseName, "-", "_")
				importMap[identifier] = cleanPath
			}
		}

		var details []SubcommandInfo
		var found bool
		ast.Inspect(fileNode, func(n ast.Node) bool {
			decl, ok := n.(*ast.GenDecl)
			if !ok || decl.Tok != token.VAR {
				return true
			} // Continue traversal
			for _, spec := range decl.Specs {
				valSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for i, nameIdent := range valSpec.Names {
					if nameIdent.Name == varName {
						fmt.Printf("\t\t- Found variable: '%s'\n", nameIdent.Name)
						if len(valSpec.Values) > i {
							details = extractSubcommandInfo(valSpec.Values[i], importMap)
							// If extractSubcommandInfo returns non-nil, it means we've successfully processed the variable.
							if details != nil {
								found = true
								return false // Stop traversal
							}
						}
					}
				}
			}
			return true
		})

		if found {
			return details, nil
		}
	}

	return nil, fmt.Errorf("failed to find definition for variable '%s' in any file of package '%s'", pkgPath, varName)
}

// buildSubcommandApp creates a temporary Go project, generates a main.go,
// and builds the final executable for a single subcommand.
func buildSubcommandApp(info SubcommandInfo, outputDir string) (builtPath string, err error) {
	tempDir, err := os.MkdirTemp("", "cmd2app-builder-*")
	if err != nil {
		err = fmt.Errorf("failed to create build directory: %w", err)
		return
	}
	//defer os.RemoveAll(tempDir)
	fmt.Printf("\t- Temp directory: %s\n", tempDir)

	mainGoPath := filepath.Join(tempDir, "main.go")
	mainGoFile, err := os.Create(mainGoPath)
	if err != nil {
		err = fmt.Errorf("failed to create main.go: %w", err)
		return
	}

	tmpl, err := template.New("app").Parse(appTemplate)
	if err != nil {
		mainGoFile.Close()
		err = fmt.Errorf("failed to parse app template: %w", err)
		return
	}

	// Use the package alias as the default executable file name.
	outputName := info.Name
	data := TemplateData{PkgPath: info.PkgPath, VarName: info.VarName}
	if e := tmpl.Execute(mainGoFile, data); e != nil {
		mainGoFile.Close()
		err = fmt.Errorf("failed to execute app template: %w", e)
		return
	}
	mainGoFile.Close()
	fmt.Println("\t- Generated optimized main.go")

	runCmdInDir := func(dir, name string, args ...string) error {
		cmd := exec.Command(name, args...)
		cmd.Dir = dir
		// Set build environment, as the target code is platform-specific.
		cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0")
		fmt.Printf("\t- Running '%s'\n", cmd)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to run '%s %s': %w\n%s", name, strings.Join(args, " "), err, string(output))
		}
		return nil
	}

	if e := runCmdInDir(tempDir, "go", "mod", "init", "tempapp"); e != nil {
		err = fmt.Errorf("failed to initialize Go module: %w", e)
		return
	}
	if e := runCmdInDir(tempDir, "go", "mod", "tidy"); e != nil {
		err = fmt.Errorf("failed to tidy Go module: %w", e)
		return
	}

	builtPath, err = filepath.Abs(filepath.Join(outputDir, outputName))
	if err != nil {
		err = fmt.Errorf("failed to resolve absolute output path: %w", err)
		return
	}
	if e := runCmdInDir(tempDir, "go", "build", "-ldflags=-s -w", "-o", builtPath, "."); e != nil {
		err = fmt.Errorf("failed to build subcommand app: %w", e)
		return
	}
	return
}

func main() {
	app := &cli.App{
		Name:  "cmd2app",
		Usage: "Generates separate, size-optimized executables for each subcommand of a given urfave/cli.Command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output-dir",
				Aliases: []string{"d"},
				Value:   "bin/latest/",
				Usage:   "Specify the output directory for the binaries.",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() != 1 {
				return cli.Exit("Error: requires one argument.\nUsage: cmd2app <package.VarName>", 1)
			}
			fullCmdPath := c.Args().First()
			lastDotIndex := strings.LastIndex(fullCmdPath, ".")
			if lastDotIndex == -1 {
				return cli.Exit(fmt.Sprintf("Error: invalid command path '%s'. Format should be <package.VarName>", fullCmdPath), 1)
			}
			pkgPath := fullCmdPath[:lastDotIndex]
			varName := fullCmdPath[lastDotIndex+1:]

			outputDir := c.String("output-dir")

			fmt.Printf("▶️  Target Package: %s\n", pkgPath)
			fmt.Printf("▶️  Target Variable: %s\n", varName)
			fmt.Printf("▶️  Output Directory: %s\n", outputDir)

			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return cli.Exit(fmt.Sprintf("Error: failed to create output directory '%s': %v", outputDir, err), 1)
			}

			fmt.Println("🔎 Analyzing source code via AST to get subcommand details...")
			subcommands, err := parseSubcommandsFromSource(pkgPath, varName)
			if err != nil {
				return cli.Exit(fmt.Sprintf("Analysis failed: %v", err), 1)
			}

			fmt.Printf("✅ Successfully found %d subcommands.\n", len(subcommands))
			if len(subcommands) == 0 {
				fmt.Println("🟡 No subcommands found, task finished.")
				return nil
			}

			for _, subCmd := range subcommands {
				fmt.Printf("\n🔨 Building subcommand: %s (from %s)\n", subCmd.Name, subCmd.PkgPath)
				builtPath, err := buildSubcommandApp(subCmd, outputDir)
				if err != nil {
					log.Printf("❌ Build failed for '%s': %v", subCmd.Name, err)
					continue
				}
				fmt.Printf("🎉 Successfully built executable: %s\n", builtPath)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
