# HTTPounder

Plays a list of HTTP requests against a server.
Waits for each request to respond before going to the next.

# Why not use the ___ benchmarking tool ?

This is not meant to compete with [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) or [autocannon](https://github.com/mcollina/autocannon).  The main reason I made it was to replay a list of requests, including malformed requests where the path was not a path but a URL.  Specifically automated hacking bots preforming malformed requests like
```
GET http://192.168.0.1:80/administrator/web/ HTTP/1.1
```
(note `http://192.168.0.1:80` should not be part of the first line of the HTTP header)

# Quick Start

```
go get github.com/robert-wallis/httpounder
```

## `request-list.txt`
I took a list of logs from our site that lead to a problem.  And ran a series of regular-expressions on them until I got a file that looks like this:
```
GET /
HEAD /
GET /someurl/
GET /another/url
GET http://localhost:3000/malformed-url
```

## Command Line
Now I have a server running on port 3000 and I will use `httpounder` to just replay the list.
```
httpounder -host localhost:3000 request-list.txt
```
