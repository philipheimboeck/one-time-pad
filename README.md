# one-time-pad

Store your keys encrypted for exactly one read

## Setup

### Install Go

- on [MacOSX](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos)

### Start Redis & App

```bash
docker-compose up --build
```

## Start App

```bash
go run backend/src/main.go
```

## Endpoints

* `GET /api/{key}`

  Returns the value by the key. Additionally the `X-Secret` header might be used to decrypt the value.

* `POST /api`

  Stores a new value and returns the generated key. Additionally the `X-Secret` header might be used to encrypt the value.
  
* `DELETE /api/{key}`

  Deletes a key

## Links

- https://bitbucket.org/rwirdemann/rest-apis-go
- https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
- https://security.stackexchange.com/questions/38828/how-can-i-securely-convert-a-string-password-to-a-key-used-in-aes
