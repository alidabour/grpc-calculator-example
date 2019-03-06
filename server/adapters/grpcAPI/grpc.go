package grpcAPI

import (
	"context"
	"log"
	"server/adapters/grpcAPI/entity"
	"server/usecase"
	"strconv"

	"google.golang.org/grpc"
)

//NewServer register calc server
func NewServer(gserver *grpc.Server, evalUC usecase.Eval) {
	evalServer := &server{
		usecase: evalUC,
	}
	entity.RegisterCalcServer(gserver, evalServer)
}

type server struct {
	usecase usecase.Eval
}

//Eval endpoint to evaluate expression
func (s server) Eval(ctx context.Context, req *entity.Request) (*entity.Result, error) {
	log.Printf("Received a request: %+v\n", req)
	result := &entity.Result{}
	res, err := s.usecase.Eval(req.Value)
	if err != nil {
		return result, err
	}
	result.Value = strconv.FormatFloat(res, 'G', -1, 64)
	return result, nil
}
