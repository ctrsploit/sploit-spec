---

version: v0.2.0

---

# Sploit Specification

## Top level commands

| Commands | Alias | Description |
| --- | --- | --- |
| env | e | information collection |
| exploit | x | run a exploit |
| checksec | c | detect vulnerabilities |
| auto | a | auto gathering information, and detect vuls, and exploit |
| version | - | show sploit tool's version |
| spec-version | - | show which spec does the sploit tool use |

## version command

execute `xsploit version`, output:

`xsploit <license: [public|pro|...]> version v0.0.1[+dev], build <GitCommit> at <DateTime>`

e.g.
```
$ ctrsploit version
ctrsploit public version v0.5.0+dev, build 6165b8e at 2023-05-31T07:36:02Z
```