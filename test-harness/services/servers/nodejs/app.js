const http = require('http');

// https://nodejs.org/en/docs/guides/anatomy-of-an-http-transaction/

http.createServer((request, response) => {
  let body = [];
  request.on('error', (err) => {
    response.end("error while reading body: " + err)
}).on('data', (chunk) => {
    body.push(chunk);
}).on('end', () => {
    body = Buffer.concat(body).toString();
    
    response.on('error', (err) => {
        response.end("error while sending response: " + err)
    });

    response.end("Body length: " + body.length.toString() + " Body: " + body);
  });
}).listen(80);
