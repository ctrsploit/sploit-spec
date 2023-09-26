---

version: v0.3.0

---

# Sploit Specification

## Top level commands

| Commands     | Alias | Description                                              |
|--------------|-------|----------------------------------------------------------|
| env          | e     | information collection                                   |
| exploit      | x     | run a exploit                                            |
| checksec     | c     | detect vulnerabilities                                   |
| auto         | a     | auto gathering information, and detect vuls, and exploit |
| version      | -     | show sploit tool's version                               |
| spec-version | -     | show which spec does the sploit tool use                 |

## version command

execute `xsploit version`, output:

`xsploit <license: [public|pro|...]> version v0.0.1[+dev], build <gitcommit> at <datetime>`

* `license` can be public, pro, or any other value you want
* `version` follows [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html)
* `gitcommit` is the first 7 characters of the latest commit's hash
* `datetime` follows the RFC3339 format

e.g.
```
$ ctrsploit version
ctrsploit public version v0.5.0+dev, build 6165b8e at 2023-05-31T07:36:02Z
```

## suggested directory structure

According to https://github.com/golang-standards/project-layout:

* /bin: contains the built binary
    * /release
* /cmd
    * /xsploit: the cli directory
        * /env: top level command env and it's subcommands' cli.Command definition
        * /exploit: top level command exploit and it's subcommands' cli.Command definition
        * /checksec: top level command checksec and it's subcommands' cli.Command definition
        * /auto: top level command auto and it's subcommands' cli.Command definition
        * /version: top level command version and it's subcommands' cli.Command definition
* /env: env implementations
* /exploit: exploit implementations
* /checksec: checksec implementations
* /version: version implementations
* /test: Additional external test apps and test data.
* /pkg: Library code that's ok to use by external applications.

## json/text/colorful output mode

x-sploit provide 3 output options:

```
GLOBAL OPTIONS:
   --colorful  output colorfully (default: false)
   --json      output colorfully (default: false)
```

Usage

```
x-sploit --colorful subcommands
```

E.g.

![](./images/colorful.png)

Coding see [github.com/ctrsploit/sploit-spec/pkg/printer/example_test.go](https://github.com/ctrsploit/sploit-spec/blob/main/pkg/printer/example_test.go)
