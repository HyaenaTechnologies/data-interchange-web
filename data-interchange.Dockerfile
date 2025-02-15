FROM amd64/rust:alpine

WORKDIR /data-interchange

COPY ./ ./

RUN cargo build

FROM amd64/alpine:latest

WORKDIR /data-interchange

COPY --from=builder ./ ./ 

RUN  ./target/debug/htdinet serve
