# HOW TO CONTRIBUTE

## 1. Git
### 1.1 fork a repository

Fork this repo

### 1.2 clone the repository you forked
eg.
```
git clone git@github.com:ssst0n3/xsploit.git
```

### 1.3 commit and push your code
```
echo "dry run by ssst0n3" >> push.txt
git add .
git commit -m "ssst0n3's dryrun"
git push
```

### 1.4 pull request
Click the button 'Compare & pull request'

## 2. Build

### 2.1 dev env

```
make shell
```

### 2.2 Build in Container

```bash
make binary && ls -lah bin/release
```

### 2.3 Build in Local

```bash
make build
```

### 2.4 Mirror

```bash
make shell CN=1
```

or

```bash
make binary CN=1
```

### 2.5 troubleshooting

`docker: 'buildx' is not a docker command.` when execute make binary

```
apt install docker-buildx-plugin
```

If it still doesn't work, try:
1. Reinstall Docker by following the [official docker documentation](https://docs.docker.com/engine/install/)
2. Check if there is a file at `~/.docker/cli-plugins/docker-buildx`, (if there is, remove it)