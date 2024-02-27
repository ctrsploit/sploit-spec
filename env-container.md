# container envs

| Basic | Description |
| --- | --- |
| type | runtime type, NotInContainer/K8s/Containerd/Docker/... |
| kernel-version | 

| Linux Security Feature | Feature | Description |
| --- | --- | --- |
| credential | credential | uid, gid, ... |
| capability | capability | capability of pid 1 and current process |
| LSM | apparmor | enabled or not |
| | selinux | enabled or not |
| seccomp | seccomp | enabled or not |
| namespace | namespace |namespace level |
| cgroups | cgroups | cgroups type and top level subsystem |
| fs | graphdriver | graphdriver type |

| Advanced | Description |
| --- | --- |
| runtime-version | runtime version range |
| ctr-cnt | host container count |
