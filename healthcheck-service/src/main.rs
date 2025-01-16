mod services;
mod handlers;
mod config;
mod grpc;

use warp::Filter;
use dotenv::dotenv;
use std::env;
use tokio::join;
use tonic::transport::{Channel, Error};
use health::health_service_client::HealthServiceClient;

pub mod health {
    include!("grpc/health.rs");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("Health check service starting...");

    dotenv().ok();

    let config = config::Config::from_env();
    let port = env::var("SERVICE_PORT").unwrap_or_else(|_| "3030".to_string());

    let (user_client, cv_client, info_client, generator_client, templates_client) = join!(
        create_grpc_client(&config.auth_service_host, &config.auth_service_port),
        create_grpc_client(&config.cv_service_host, &config.cv_service_port),
        create_grpc_client(&config.information_service_host, &config.information_service_port),
        create_grpc_client(&config.generator_service_host, &config.generator_service_port),
        create_grpc_client(&config.templates_service_host, &config.templates_service_port),
    );

    let user_client = user_client?;
    let cv_client = cv_client?;
    let info_client = info_client?;
    let generator_client = generator_client?;
    let templates_client = templates_client?;

    let index_route = warp::path::end()
        .and(warp::get())
        .and_then(handlers::index::index);

    let check_route = warp::path!("check")
        .and(warp::get())
        .and(warp::any().map(move || config.gateway_url.clone()))
        .and(warp::any().map(move || config.rabbitmq_url.clone()))
        .and(warp::any().map(move || (config.minio_url.clone(), config.minio_access_key.clone(), config.minio_secret_key.clone())))
        .and(warp::any().map(move || user_client.clone()))
        .and(warp::any().map(move || cv_client.clone()))
        .and(warp::any().map(move || info_client.clone()))
        .and(warp::any().map(move || generator_client.clone()))
        .and(warp::any().map(move || templates_client.clone()))
        .and_then(handlers::health::check_health);

    let routes = index_route.or(check_route);

    println!("Health check service running on port {}", port);
    warp::serve(routes).run(([0, 0, 0, 0], port.parse::<u16>().unwrap())).await;

    Ok(())
}

pub async fn create_grpc_client(
    host: &str,
    port: &str,
) -> Result<HealthServiceClient<Channel>, Error> {
    let endpoint = format!("http://{}:{}", host, port);
    HealthServiceClient::connect(endpoint).await
}