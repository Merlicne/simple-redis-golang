# SimpleRedis

SimpleRedis is a easy-to-use Redis client written in Go. It provides a simple interface to interact with a Redis server, allowing you to perform common operations such as setting, getting, and deleting keys.

## Features

- Connect to a Redis server
- Set and get key-value pairs
- Delete keys
- Support for basic Redis commands

## Environment
- Go 1.22
- Redis 7.4.1 (helm chart redis-20.3.0)
- Kubernetes 1.30.5 (Redis Cluster used in the example)

## Configuration
```yaml
redis:
  host: redishost
  port: 6379
  password: redispassword
  db: 0
```

