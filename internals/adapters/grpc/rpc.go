package grpc

// import (
// 	"context"

// 	"github.com/somT-oss/urlshortner/internals/adapters/grpc/pb"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )



// func (adapter Adapter) ShortenUrl(ctx context.Context, req *pb.MainUrl) (*pb.ShortenedUrl, error) {
// 	var err error
// 	ans := &pb.ShortenedUrl{}

// 	if len(req.MainUrl) == 0{
// 		return ans, status.Error(codes.InvalidArgument, "missing required argument")
// 	}
// 	answer, err := adapter.api.GenerateShortUrl(req.MainUrl)
// 	if err != nil {
// 		return ans, status.Error(codes.Internal,  "unexpected error")
// 	}
// 	ans = &pb.ShortenedUrl{
// 		ShortUrl: answer.,
// 	}
// }