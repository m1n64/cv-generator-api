use std::fmt;
use warp::reject::Reject;

#[derive(Debug)]
pub struct GrpcError {
    message: String,
}

impl GrpcError {
    pub fn new(message: String) -> Self {
        GrpcError { message }
    }
}

impl fmt::Display for GrpcError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.message)
    }
}

impl Reject for GrpcError {}