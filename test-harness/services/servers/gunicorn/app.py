def app(environ, start_response):
    body = environ["wsgi.input"].read()
    data = b"Body length: " + str(len(body)).encode() + b" Body: " + repr(body).encode()
    start_response("200 OK", [
        ("Content-Type", "text/plain"),
        ("Content-Length", str(len(data)))
    ])
    return iter([data])
