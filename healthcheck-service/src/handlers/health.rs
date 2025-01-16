use warp::{Rejection, Reply};
use serde::Serialize;
use tonic::transport::Channel;
use crate::health::health_service_client::HealthServiceClient;
use crate::services::{gateway, rabbitmq, minio, health_check};
use crate::services::grpc_service::check_service;

#[derive(Serialize)]
struct ContainerLayer {
    layer: String,
    status: bool,
}

#[derive(Serialize)]
struct ServiceStatus {
    name: String,
    status: bool,
    timestamp: Option<String>,
    containers: Option<Vec<ContainerLayer>>,
}

pub async fn check_health(
    gateway_url: String,
    rabbitmq_url: String,
    minio_config: (String, String, String),
    user_client: HealthServiceClient<Channel>,
    cv_client: HealthServiceClient<Channel>,
    info_client: HealthServiceClient<Channel>,
    generator_client: HealthServiceClient<Channel>,
    templates_client: HealthServiceClient<Channel>,
) -> Result<impl Reply, Rejection> {
    let mut results = Vec::new();

    // Current service (Healthcheck)
    let health_status = health_check::check_health().await;
    results.push(ServiceStatus {
        name: health_status.0,
        status: health_status.1,
        timestamp: health_status.2,
        containers: None,
    });

    // Gateway
    let gateway_status = gateway::check_gateway(&gateway_url).await;

    results.push(ServiceStatus {
        name: gateway_status.0,
        status: gateway_status.1,
        timestamp: gateway_status.2,
        containers: None,
    });

    // RabbitMQ
    let rabbitmq_status = rabbitmq::check_rabbitmq(&rabbitmq_url).await;
    results.push(ServiceStatus {
        name: rabbitmq_status.0,
        status: rabbitmq_status.1,
        timestamp: rabbitmq_status.2,
        containers: None,
    });

    // MinIO
    let minio_status = minio::check_minio(minio_config).await;
    results.push(ServiceStatus {
        name: minio_status.0,
        status: minio_status.1,
        timestamp: minio_status.2,
        containers: None,
    });

    // Services
    let services = vec![
        ("User Service", user_client),
        ("CV Service", cv_client),
        ("CV Information Service", info_client),
        ("CV Generator Service", generator_client),
        ("Templates Service", templates_client),
    ];

    for (name, client) in services {
        let service_status = process_service_check(name, client).await;
        results.push(service_status);
    }

    Ok(warp::reply::json(&results))
}

async fn process_service_check(
    service_name: &str,
    mut client: HealthServiceClient<Channel>,
) -> ServiceStatus {
    match check_service(service_name, client).await {
        Ok(response) => ServiceStatus {
            name: response.service_name,
            status: response.status,
            timestamp: Some(response.timestamp),
            containers: Some(vec![
                ContainerLayer {
                    layer: "database".to_string(),
                    status: response.status_db,
                },
                ContainerLayer {
                    layer: "redis".to_string(),
                    status: response.status_redis,
                },
            ]),
        },
        Err(err) => {
            eprintln!("Error calling gRPC Service {}: {:?}", service_name, err);
            ServiceStatus {
                name: service_name.to_string(),
                status: false,
                timestamp: None,
                containers: None,
            }
        }
    }
}