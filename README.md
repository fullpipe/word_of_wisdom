# word_of_wisdom

Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge

## How to run?

Run server with one client

```bash
docker-compose up --build
```

Run client separately

```bash
docker run --rm --network="host" $(docker build -q .) client -host localhost:8080
```

## Proof Of Work

sha256 is used as the hash function. This algorithm is easy to implement and configure.

But I think for production it is better to use Script, since it is more memory-intensive. I started refactoring, but it requires quite subtle tuning of parameters depending on the given complexity https://github.com/fullpipe/word_of_wisdom/pull/1
