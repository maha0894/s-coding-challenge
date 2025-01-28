# s-coding-challenge

## Clone

```
$ git clone https://github.com/maha0894/s-coding-challenge
$ cd s-coding-challenge
```

## Run

```
$ cd s-coding-challenge
$ go mod tidy
$ go build
$ ./s-coding-challenge http
```

if the server has started succsefully you will see following in the terminal:

```
2025/01/28 23:25:28 Starting http server...
```
and following endpoints can be accessed:
```
http://localhost:8000/users/{id:[0-9]+}
http://localhost:8000/users/{id:[0-9]+}/actions
http://localhost:8000/users/referral-index
http://localhost:8000/actions/{action}/next
```