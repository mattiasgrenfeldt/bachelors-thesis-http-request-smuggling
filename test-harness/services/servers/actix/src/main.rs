use actix_web::{get, post, App, HttpResponse, HttpServer, Responder};

#[post("/")]
async fn echo1(req_body: String) -> impl Responder {
    let mut answer: String = "Body length: ".to_owned();
    answer.push_str(&req_body.len().to_string());
    answer.push_str(" Body: ");
    answer.push_str(&req_body);

    HttpResponse::Ok().body(answer)
}

#[get("/")]
async fn echo2(req_body: String) -> impl Responder {
    let mut answer: String = "Body length: ".to_owned();
    answer.push_str(&req_body.len().to_string());
    answer.push_str(" Body: ");
    answer.push_str(&req_body);

    HttpResponse::Ok().body(answer)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(echo1).service(echo2))
        .bind("0.0.0.0:80")?
        .run()
        .await
}
