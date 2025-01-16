use chrono::Utc;
use minio_rsc::client::{BucketArgs};
use minio_rsc::provider::StaticProvider;
use minio_rsc::Minio;

#[derive(Debug)]
pub struct MinioError(String);

pub async fn check_minio(minio_config: (String, String, String)) -> (String, bool, Option<String>) {
    let (minio_url, access_key, secret_key) = minio_config;
    let provider = StaticProvider::new(&access_key, &secret_key, None);
    let minio = Minio::builder()
        .endpoint(&minio_url)
        .provider(provider)
        .secure(false)
        .build();

    if let Ok(client) = minio {
        if client.bucket_exists(BucketArgs::new("healthcheck")).await.is_ok() {
            let timestamp = Utc::now().to_rfc3339();

            return ("MinIO Service".to_string(), true, Some(timestamp));
        }
    }

    ("MinIO Service".to_string(), false, None)
}
