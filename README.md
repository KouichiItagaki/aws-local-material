# aws-local-material ☁️☁️☁️

<img src="https://img.shields.io/badge/Go-v1.17-blue">

## local-s3

### Usage

```
$ cd local-s3
```

```
$ docker compose up -d
```

The minio console appears when you access `http://localhost:9001`

```text
Username: user
Password: password
```

https://user-images.githubusercontent.com/58158037/160244658-19a3049b-c69f-4ce0-9bed-aa17a4b9b4be.mov

Get a list of objects under dir1/ directory.

```
$ curl -XGET http://localhost:5050/api/v1/cats
["dir1/cat.jpg"]
```

#### AWS CLI

```
$ aws configure --profile minio
AWS Access Key ID [None]: user
AWS Secret Access Key [None]: password
Default region name [None]: ap-northeast-1
Default output format [None]: json
```

```
$ aws --profile minio --endpoint-url http://localhost:9000/ s3 cp cat2.jpg s3://local-bucket/dir1/
```

## local-sqs

### Usage

```
$ cd local-sqs
```

```
$ docker compose up -d
```

#### AWS CLI

List queues

```
$ aws sqs list-queues --endpoint-url 'http://localhost:9324'
{
    "QueueUrls": [
        "http://localhost:9324/000000000000/queue1-dead-letters",
        "http://localhost:9324/000000000000/queue1"
    ]
}
```

Send message

```
$ aws sqs send-message \
    --queue-url "http://localhost:9324/000000000000/queue1" \
    --endpoint-url "http://localhost:9324" \
    --message-body "Hello ElasticMQ!"
```

Receive message

```
aws sqs receive-message \
    --queue-url "http://localhost:9324/000000000000/queue1" \
    --endpoint-url "http://localhost:9324"
```