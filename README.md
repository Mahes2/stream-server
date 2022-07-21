# stream-server

## ⚡️ Prerequisite
-   [Go](https://go.dev/dl/) (1.18 or higher)
-   [Postgres](https://www.postgresql.org/download/)
-   [Java](https://www.oracle.com/java/technologies/downloads/) (8)
-   [Maven](https://maven.apache.org/download.cgi)

## ⚙️ Build
Execute query `script/test.sql` to create table and populate data

Build protobuf file used by go server:
```bash
make pb-go
```
Build proto jar file used by java client:
```bash
make pb-java
```

## 🤖 Run

To run grpc server A and B:
```bash
go run cmd/main.go grpc 9998
go run cmd/main.go grpc 9997
```

To run http server A and B:
```bash
go run cmd/main.go http 9998
go run cmd/main.go http 9997
```

## 🎯 Test
See [stream-client](https://github.com/Mahes2/stream-client/) for testing step
