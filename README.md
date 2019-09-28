# one-time-pad

Store your keys encrypted for exactly one read

## Setup

### Install Go

- on [macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos)

### Install Go Packages

```bash
go get -d -v ./...
```

### Start Redis & App

```bash
docker-compose up --build
```

When you want to start only redis, execute:

```bash
docker-compose up redis
```

## How to run the app and develop

```bash
go run backend/src/main.go
```

Debug the app with VS Code by executing the `Launch` debug script. Simply run
the `Debug: Select and Start Debugging` command and select `Launch`. Note: You
might need to install xcode command line tools (`xcode-select --install`) to
be able to run the debugger on macOS.

## Endpoints

* `GET /api/{key}`

  Returns the value by the key. Additionally the `X-Secret` header might be used to decrypt the value.

* `POST /api`

  Stores a new value and returns the generated key. Additionally the `X-Secret` header might be used to encrypt the value.
  
* `DELETE /api/{key}`

  Deletes a key

## Links

- https://bitbucket.org/rwirdemann/rest-apis-go
- https://jaxenter.de/restful-rest-api-go-golang-68845
- https://github.com/go-redis/redis
- https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
- https://security.stackexchange.com/questions/38828/how-can-i-securely-convert-a-string-password-to-a-key-used-in-aes
- https://redis.io/topics/rediscli
