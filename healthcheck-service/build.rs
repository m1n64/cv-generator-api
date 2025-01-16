fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("Starting proto compilation...");
    tonic_build::configure()
        .out_dir("src/grpc")
        .compile(&["proto/health.proto"], &["proto"])
        .expect("Failed to compile protos");
    println!("cargo:rerun-if-changed=proto/health.proto");
    Ok(())
}