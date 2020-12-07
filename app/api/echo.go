package api

import (
	"context"
	"fmt"
	"github.com/gogf/katyusha-demos/protocol/pb"
)

// 回显服务
var Echo = &apiEcho{}

type apiEcho struct{}

func (s *apiEcho) Say(ctx context.Context, r *pb.SayReq) (*pb.SayRes, error) {
	text := fmt.Sprintf(`> %s`, r.Content)
	return &pb.SayRes{Content: text}, nil
}
