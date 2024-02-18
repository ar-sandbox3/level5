# level5

This repository exhibits the basic concept of middleware, in the context of Go-based HTTP Server.

## Utilities

This also holds utilities on dealing with JWT.

1. [`rsagen`](./cmd/rsagen)
1. [`jwtgen`](./cmd/jwtgen)
1. [`jwtval`](./cmd/jwtval)

### `rsagen`

It generates RSA keypair (private.pem and public.pem).

```console
go run cmd/rsagen/main.go
```

### `jwtgen`

It generates a JWT (token.jwt) file, that can be generated using symmetric (HS256) or assymmetric (RS256) key.

See: https://auth0.com/blog/rs256-vs-hs256-whats-the-difference/.

To run with symmetric key, you need to specify the key.

```console
go run cmd/jwtgen/main.go --symmetric --key supersecret
```

However, this program is by default generating RS256 token.

> [!NOTE]
> Please make sure to generate RSA key pair first.


```console
go run cmd/jwtgen/main.go
```

### `jwtval`

It validates the generated token.jwt using the defined key strategy (symmetric/asymmetric).

```console
go run cmd/jwtgen/main.go --symmetric --key supersecret
```

And if the token.jwt was generated using RSHA256.

```console
go run cmd/jwtgen/main.go
```
