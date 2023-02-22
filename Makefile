proto-greet:
	protoc -Igreet/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial greet/proto/greet.proto

build-greet:
	protoc -Igreet/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial greet/proto/greet.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client

clean-greet:
	rm -rf greet/proto/*.pb.go
	rm -rf bin/greet/client
	rm -rf bin/greet/server

clean:
	rm -rf greet/proto/*.pb.go
	rm -rf bin/greet/client
	rm -rf bin/greet/server