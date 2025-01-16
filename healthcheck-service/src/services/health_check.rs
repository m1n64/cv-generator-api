
pub async fn check_health() -> (String, bool, Option<String>) {
    let timestamp = chrono::Utc::now().to_rfc3339();

    ("Health Check Service".to_string(), true, Some(timestamp))
}