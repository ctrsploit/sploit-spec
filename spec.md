---

version: v0.2.1

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