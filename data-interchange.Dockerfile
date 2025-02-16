FROM rust:alpine

WORKDIR /data-interchange

COPY ./ ./

RUN cargo check \  
cargo build --release \ 
mv ./target/release/data-interchange-server ./binary

FROM alpine:latest

WORKDIR /data-interchange

COPY --from=builder ./ ./

EXPOSE 80:8080/tcp

RUN ./binary/data-interchange-server serve
