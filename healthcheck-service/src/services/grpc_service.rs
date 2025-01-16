use tonic::Request;
use tonic::transport::Channel;
use warp::{Rejection};
use crate::health::{CheckRequest, CheckResponse};
use crate::health::health_service_client::HealthServiceClient;
use crate::services::errors::GrpcError;

pub async fn check_service(
    service_name: &str,
    mut client: HealthServiceClient<Channel>,
) -> Result<CheckResponse, Rejection> {
    let request = Request::new(CheckRequest {
        service: service_name.to_string(),
    });

    match client.check(request).await {
        Ok(response) => Ok(response.into_inner()),
        Err(err) => {
            eprintln!("Error calling gRPC service: {:?}", err);
            Err(warp::reject::custom(GrpcError::new(err.to_string())))
        }
    }
}