# aws-local-material ☁️☁️☁️

<img src="https://img.shields.io/badge/Go-v1.17-blue">

## local-s3

### Usage

```terminal
cd local-s3
docker compose up -d
```

The minio console appears when you access `http://localhost:9001`

```text
Username: user
Password: password
```

Get a list of objects under dir1/ directory.

```console
curl -XGET http://localhost:5050/api/v1/cats
```
