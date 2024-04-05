# Schedule

- Ngày học lí thuyết : 29/02/2024 -> 08/03/2024
- Ngày viết code example : 29/02/2024 -> 06/03/2024
- [Báo cáo](https://docs.google.com/document/d/1ZE5Xary64_o2ZE4YkqoBnIHIzPPbIYkLmBIb6tNF5LA/edit)

# Structure

```
C:\USERS\ADMIN\DOCUMENTS\HIRO\GRPC-COURSE
└───example
    ├───proto
    ├───proto_file
    └───udemy
        └───ex
            ├───client
            └───server
```

- /example: Chứa code triển khai grpc
- /example/proto: Chứa code dược gen từ /proto_file sử dụng protoc

# Description

<p sytle="color='#22d3ee'">Source code cơ bản triển khai gRPC sử dụng protoc. Trong đó gRPC server là Caculator Server - cung cấp các service tính tổng, trung bình, tìm max ..</p>

Định nghĩa 1 Calulater Server như sau:

```
service CaculateSum {
    rpc SumTowInteger(SumRequest) returns (SumResponse);
    rpc SumArr(SumArrRequest) returns (SumResponse);
    rpc Primes (PrimesRequest) returns (stream PrimesResponse);
    rpc Avg (stream PrimesRequest) returns (AvgResponse);
    rpc Max (stream PrimesRequest) returns (stream PrimesResponse);
}
```

# Running

```
Run server: go run ./example/udemy/ex/server
Run client: go run ./example/udemy/ex/client
```
