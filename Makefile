all:
	protoc -Igreet/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial greet/proto/greet.proto
	go build -o bin/greet/server ./greet/server
	go build -o bin/greet/client ./greet/client
	protoc -Isum/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial sum/proto/sum.proto
	go build -o bin/sum/server ./sum/server
	go build -o bin/sum/client ./sum/client

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

proto-sum:
	protoc -Isum/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial sum/proto/sum.proto

build-sum:
	protoc -Isum/proto --go_out=. --go_opt=module=github.com/araujo88/grpc-go-tutorial --go-grpc_out=. --go-grpc_opt=module=github.com/araujo88/grpc-go-tutorial sum/proto/sum.proto
	go build -o bin/sum/server ./sum/server
	go build -o bin/sum/client ./sum/client

clean-sum:
	rm -rf sum/proto/*.pb.go
	rm -rf bin/sum/client
	rm -rf bin/sum/server

clean:
	rm -rf greet/proto/*.pb.go
	rm -rf bin/greet/client
	rm -rf bin/greet/server
	rm -rf sum/proto/*.pb.go
	rm -rf bin/sum/client
	rm -rf bin/sum/server