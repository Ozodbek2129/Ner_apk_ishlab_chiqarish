package storage

import (
	pb "ishlab_chiqarish/genproto/ishlab_chiqarish"
	"context"
)

type IStorage interface {
	Production() IProductionStorage
	Close()
}

type IProductionStorage interface {
	NewContract(context.Context ,*pb.NewContractReq) (*pb.NewContractRes, error)
	NewContractUpdate(context.Context ,*pb.NewContractUpdateReq) (*pb.NewContractUpdateRes, error)
	NewContractDelete(context.Context ,*pb.NewContractDeleteReq) (*pb.NewContractDeleteRes, error)
	NewContractGetName(context.Context ,*pb.NewContractGetNameReq) (*pb.NewContractGetNameRes, error)
	NewContractGetAll(context.Context ,*pb.NewContractGetAllReq) (*pb.NewContractGetAllRes, error)
	NewInsideTheContract(context.Context ,*pb.NewInsideTheContractReq) (*pb.NewInsideTheContractRes, error)
	NewInsideTheContractUpdate(context.Context ,*pb.NewInsideTheContractUpdateReq) (*pb.NewInsideTheContractUpdateRes, error)
	NewInsideTheContractDelete(context.Context ,*pb.NewInsideTheContractDeleteReq) (*pb.NewInsideTheContractDeleteRes, error)
	NewInsideTheContractGetAll(context.Context ,*pb.NewInsideTheContractGetAllReq) (*pb.NewInsideTheContractGetAllRes, error)
}