use reqwest::Response;
use serde::Deserialize;

#[derive(Deserialize)]
struct GatewayResponse {
    message: String,
    time: String,
}


pub async fn check_gateway(gateway_url: &str) -> (String, bool, Option<String>) {
    match reqwest::get(format!("{}/ping", gateway_url)).await {
        Ok(response) if response.status().is_success() => {
            parse_gateway_response(response).await
        }
        Ok(_) => {
            ("Gateway Service".to_string(), false, None)
        }
        Err(_) => {
            ("Gateway Service".to_string(), false, None)
        }
    }
}

async fn parse_gateway_response(response: Response) -> (String, bool, Option<String>) {
    match response.json::<GatewayResponse>().await {
        Ok(data) => ("Gateway Service".to_string(), true, Some(data.time)),
        Err(_) => {
            ("Gateway Service".to_string(), false, None)
        }
    }
}