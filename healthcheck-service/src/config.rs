use std::env;

pub struct Config {
    pub gateway_url: String,
    pub rabbitmq_url: String,
    pub minio_url: String,
    pub minio_access_key: String,
    pub minio_secret_key: String,
    pub auth_service_host: String,
    pub auth_service_port: String,
    pub cv_service_host: String,
    pub cv_service_port: String,
    pub information_service_host: String,
    pub information_service_port: String,
    pub generator_service_host: String,
    pub generator_service_port: String,
    pub templates_service_host: String,
    pub templates_service_port: String
}

impl Config {
    pub fn from_env() -> Self {
        let gateway_host = env::var("GATEWAY_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let gateway_port = env::var("GATEWAY_SERVICE_PORT").unwrap_or_else(|_| "8000".to_string());
        let gateway_url = format!("http://{}:{}", gateway_host, gateway_port);

        let rabbitmq_url = env::var("RABBITMQ_URL").unwrap_or_else(|_| "amqp://localhost:5672/".to_string());
        let minio_url = env::var("MINIO_URL").unwrap_or_else(|_| "localhost:9000".to_string());
        let minio_access_key = env::var("MINIO_ACCESS_KEY").unwrap_or_else(|_| "minioadmin".to_string());
        let minio_secret_key = env::var("MINIO_SECRET_KEY").unwrap_or_else(|_| "minioadmin".to_string());

        let auth_service_host = env::var("AUTH_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let auth_service_port = env::var("AUTH_SERVICE_PORT").unwrap_or_else(|_| "50051".to_string());

        let cv_service_host = env::var("CV_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let cv_service_port = env::var("CV_SERVICE_PORT").unwrap_or_else(|_| "50051".to_string());

        let info_service_host = env::var("INFORMATION_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let info_service_port = env::var("INFORMATION_SERVICE_PORT").unwrap_or_else(|_| "50051".to_string());

        let generator_service_host = env::var("GENERATOR_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let generator_service_port = env::var("GENERATOR_SERVICE_PORT").unwrap_or_else(|_| "50051".to_string());

        let templates_service_host = env::var("TEMPLATES_SERVICE_HOST").unwrap_or_else(|_| "localhost".to_string());
        let templates_service_port = env::var("TEMPLATES_SERVICE_PORT").unwrap_or_else(|_| "50051".to_string());

        Config {
            gateway_url,
            rabbitmq_url,
            minio_url,
            minio_access_key,
            minio_secret_key,
            auth_service_host,
            auth_service_port,
            cv_service_host,
            cv_service_port,
            information_service_host: info_service_host,
            information_service_port: info_service_port,
            generator_service_host,
            generator_service_port,
            templates_service_host,
            templates_service_port
        }
    }
}