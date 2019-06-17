package dependent

import "google.golang.org/grpc"

var GrpcClient *grpc.ClientConn
func init() {
	var e error
	GrpcClient, e = grpc.Dial("localhost:6001", grpc.WithInsecure())
	if e!=nil {
		panic(e)
	}
}
