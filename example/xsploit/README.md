---

spec-version: v0.4.0

---

# x-sploit: An example app follows sploit-spec

Just an example, do not import code under example folder  

## Usage

```
❯ ./xsploit_linux_amd64       
NAME:
   xsploit - An example sploit tool follows sploit-spec

USAGE:
   xsploit [global options] command [command options] [arguments...]

COMMANDS:
   auto         auto
   env, e       Collect information
   checksec, c  check security inside a container
   exploit, x   run a exploit
   vul, v       list vulnerabilities
   version      Show the sploit version information
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug         Output information for helping debugging sploit (default: false)
   --experimental  enable experimental feature (default: false)
   --colorful      output colorfully (default: false)
   --json          output in json format (default: false)
   --help, -h      show help
```

### env: collect env information

```
❯ ./bin/release/xsploit_linux_amd64 env     
NAME:
   xsploit env - Collect information

USAGE:
   xsploit env command [command options] [arguments...]

COMMANDS:
   auto       auto
   second, s  show the second info
   minute, m  show the minute info
   help, h    Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
   
❯ ./bin/release/xsploit_linux_amd64 env auto
second:                 32      # second of current time
minute:                 53      # second of current minute
```

### checksec: check vulnerability exists

CVE-2099-9999 exists when 2 | second . 

```
❯ ./bin/release/xsploit_linux_amd64 checksec       
NAME:
   xsploit checksec - check security inside a container

USAGE:
   xsploit checksec command [command options] [arguments...]

COMMANDS:
   auto                 auto
   CVE-2099-9999, 2099  Description of CVE-2099-9999
   help, h              Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help

❯ ./bin/release/xsploit_linux_amd64 checksec 2099
[N]  CVE-2099-9999      # Description of CVE-2099-9999
```

### exploit: run exploit

CVE-2099-9999 is a vulnerability only can be exploited by root.

```
❯ ./bin/release/xsploit_linux_amd64 exploit      
NAME:
   xsploit exploit - run a exploit

USAGE:
   xsploit exploit command [command options] [arguments...]

COMMANDS:
   auto                 auto
   CVE-2099-9999, 2099  Description of CVE-2099-9999
   help, h              Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help

❯ ./bin/release/xsploit_linux_amd64 exploit 2099
ERRO[0000] CVE-2099-9999 is not exploitable             

❯ sudo ./bin/release/xsploit_linux_amd64 exploit 2099
CVE-2099-9999 has exploited
```

### vul: list vulnerabilities supported by xsploit

```
❯ ./bin/release/xsploit_linux_amd64 vul     
NAME:
   xsploit vul - list vulnerabilities

USAGE:
   xsploit vul command [command options] [arguments...]

COMMANDS:
   CVE-2099-9999, 2099  Description of CVE-2099-9999
   help, h              Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help

❯ ./bin/release/xsploit_linux_amd64 vul 2099 
NAME:
   xsploit vul CVE-2099-9999 - Description of CVE-2099-9999

USAGE:
   xsploit vul CVE-2099-9999 command [command options] [arguments...]

COMMANDS:
   checksec, c  check vulnerability exists
   exploit, x   run exploit
   help, h      Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

### machine friendly output

```
❯ ./bin/release/xsploit_linux_amd64 --json env auto
{"minute":{"name":"minute","description":"second of current minute","result":"55"},"second":{"name":"second","description":"second of current time","result":"19"}}
```