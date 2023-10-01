# Word of Wisdom

## Task

Design and implement “Word of Wisdom” tcp server.  
* TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.  
* The choice of the POW algorithm should be explained.  
* After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.  
* Docker file should be provided both for the server and for the client that solves the POW challenge

## TODO

* Add unit and integration tests
* Add Configuration for server and client
* Add description to README

## Run

### Main flow

```sh
make built-dev
make run
```

### Unit tests

```sh
make test
```

### Linters

```sh
make lint
```

## Protocol

### Emoji protocol

#### 1. Flow

1. The Client sends a `hello` message: `👋`
2. The Server receives the `hello` message and sends a `challenge`: `🧮 ver1:7:1696187660:MTIxNjM=:todo:MQ==`
3. The Client solves the `challenge` and sends a `solution` message to the server: `👌 Mjk1ODUyMzg=`
4. The Server checks the `solution` and send a quote: `📖 Believe you can and you're halfway there`
5. Client receives the quote.
6. Done.

#### 2. Messages

| Message        | Sender | Description                                   |
|----------------|--------|-----------------------------------------------|
 | `👋`           | client | An initial message from the client            | 
 | `🧮 challenge` | server | Challenge message for the client              | 
 | `👌 solution`  | client | A message with a solution                     | 
 | `📖 quote`     | server | A message with a quote                        | 
 | `🙅`           | server | The resulting solution is incorrect           | 
| `🤦`           | server | If the solution is received before 👋 message | 
| `🤷`           | server | Unknown message                               |
