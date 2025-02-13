FROM amd64/golang:alpine

WORKDIR /data-interchange

COPY ./ ./

RUN go env  \ 
go list ./source \ 
go vet ./source \ 
go fix ./source \ 
go fmt ./source \ 
go vet ./source \ 
GOOS=linux GOARCH=amd64 go build -o ./binary/htdinet ./source/main.go

FROM amd64/alpine:latest

COPY --from=builder ./ ./ 

RUN ./binary/htdinet serve
