proto: 
	protoc --go_out=. --go_opt=module=example.com/route_guide --go-grpc_out=. --go-grpc_opt=module=example.com/route_guide route_guide.proto
	go mod tidy

clean:
	rm -rf route_guide_grpc.pb.go route_guide.pb.go