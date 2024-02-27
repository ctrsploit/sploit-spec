# container envs

| Basic | Description |
| --- | --- |
| type | runtime type, NotInContainer/K8s/Containerd/Docker/... |
| graphdriver | graphdriver type |

| Linux Security Features | Description |
| --- | --- |
| capability | capability of pid 1 and current process |
| apparmor | enabled or not |
| selinux | enabled or not |
| seccomp | enabled or not |
| namespace | namespace level |
| cgroups | cgroups type and top level subsystem |

| Advanced | Description |
| --- | --- |
| version | runtime version range |
| ctr-cnt | host container count |
