### video-thumbnail [![Build Status](https://github.com/better-than-yours/video-thumbnail/workflows/backend/badge.svg)](https://github.com/better-than-yours/video-thumbnail/backend) [![Go Report Card](https://goreportcard.com/badge/github.com/better-than-yours/video-thumbnail)](https://goreportcard.com/report/github.com/better-than-yours/video-thumbnail)

### deps
```sh 
$ go mod tidy && go get -u
```

### dev deps
```sh
$ apt-get update && apt-get -y install --no-install-recommends ffmpeg libavcodec-dev libavutil-dev libavformat-dev libswscale-dev
```

### secrets

```sh
VAULT_ROLE_ID=
VAULT_SECRET_ID=
```

### vault

```sh
$ vault auth enable approle
$ vault write auth/approle/role/dratini-role secret_id_ttl=0 token_policies=common-policy,dratini-policy
$ vault read auth/approle/role/dratini-role/role-id
$ vault write -f auth/approle/role/dratini-role/secret-id
```