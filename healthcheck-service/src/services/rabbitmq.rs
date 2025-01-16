use chrono::Utc;
use lapin::{Connection, ConnectionProperties};

pub async fn check_rabbitmq(rabbitmq_url: &str) -> (String, bool, Option<String>) {
    match Connection::connect(rabbitmq_url, ConnectionProperties::default()).await {
        Ok(connection) => {
            let timestamp = Utc::now().to_rfc3339();

            if let Err(err) = connection.close(0, "Normal shutdown").await {
                eprintln!("Error while closing RabbitMQ connection: {:?}", err);
            }

            ("RabbitMQ Service".to_string(), true, Some(timestamp))
        }
        Err(_) => ("RabbitMQ Service".to_string(), false, None),
    }
}