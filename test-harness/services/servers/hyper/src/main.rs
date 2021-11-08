#![deny(warnings)]

use std::str;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server};

async fn echo(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    let whole_body = hyper::body::to_bytes(req.into_body()).await?;
    let mut answer: String = "Body length: ".to_owned();
    answer.push_str(&whole_body.len().to_string());
    answer.push_str(" Body: ");
    
    let s = match str::from_utf8(&whole_body) {
        Ok(v) => v,
        Err(_) => "error while converting request body to utf8",
    };

    answer.push_str(s);

    Ok(Response::new(Body::from(answer)))
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let addr = ([0, 0, 0, 0], 80).into();

    let service = make_service_fn(|_| async { Ok::<_, hyper::Error>(service_fn(echo)) });

    let server = Server::bind(&addr).serve(service);

    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}