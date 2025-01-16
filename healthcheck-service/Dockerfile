FROM rust:1.84 AS builder

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y protobuf-compiler

RUN cargo install cargo-watch

COPY Cargo.toml Cargo.lock ./
RUN cargo fetch
COPY . .

RUN cargo build --release

FROM debian:buster-slim AS production
COPY --from=builder /usr/src/app/target/release/health-check /usr/local/bin/health-check
EXPOSE 3030
CMD ["health-check"]

FROM rust:1.84 AS development

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y protobuf-compiler

COPY . .

RUN cargo install cargo-watch

EXPOSE 3030

CMD ["cargo", "watch", "-x", "run"]