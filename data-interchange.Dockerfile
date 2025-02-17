FROM rust:alpine

WORKDIR /data-interchange-web

COPY ./ ./

RUN cargo check \  
cargo build --release \ 
mv ./target/release/data-interchange-server ./binary

FROM alpine:latest

WORKDIR /data-interchange-web

COPY --from=builder ./ ./

EXPOSE 80:8080/tcp

CMD ["./binary/data-interchange-server", "serve"]
