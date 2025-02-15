FROM amd64/rust:alpine

WORKDIR /data-interchange

COPY ./ ./

RUN cargo update \ 
cargo check \  
cargo build --release \ 
mv ./target/release/data-interchange-server ./binary

FROM amd64/alpine:latest

WORKDIR /data-interchange

COPY --from=builder ./ ./

RUN ./binary/data-interchange-server serve
