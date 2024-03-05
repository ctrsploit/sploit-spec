# Linux envs



| Basic          | Description                              |
| -------------- | ---------------------------------------- |
| kernel-version | 4.18.0-147.5.1.6.h841.eulerosv2r9.x86_64 |
| Compile-time   | 2024-02-05 11:39:20                      |

| Linux Security Feature                     | Feature    | Description                             |
| ------------------------------------------ | ---------- | --------------------------------------- |
| unshare                                    | unshare    | support or not                          |
| credential                                 | credential | uid, gid, ...                           |
| capability                                 | capability | capability of pid 1 and current process |
| seccomp                                    | seccomp    | enabled or not                          |
| namespace                                  | namespace  | namespace level                         |
| /proc/sys/kernel/unprivileged_bpf_disabled | Bpf        | enabled or not                          |
| xen addr leak                              | Addr       | Kernel text base addr                   |

| Advanced            | Description       |
| ------------------- | ----------------- |
| Environ             | Environ variables |
| suid list           | sudi binary list  |
| Sudi Binary Version | sudo\dpkg\pkexec  |
| Crontab list        | Crontab -l        |