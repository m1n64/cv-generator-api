use warp::{Rejection, Reply};
use tokio::fs;

pub async fn index() -> Result<impl Reply, Rejection> {
    let html_file_path = "templates/index.html";

    match fs::read_to_string(html_file_path).await {
        Ok(contents) => Ok(warp::reply::html(contents)),
        Err(err) => {
            eprintln!("Error reading HTML file: {:?}", err);
            Err(warp::reject::not_found())
        }
    }
}