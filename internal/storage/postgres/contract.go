package postgres

import (
	"context"
	"database/sql"
	"ishlab_chiqarish/internal/pkg/logger"
	"ishlab_chiqarish/internal/storage"
	"log/slog"
	pb "ishlab_chiqarish/genproto/ishlab_chiqarish"
)

type NerProductionRepo struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewNerProductionRepo(db *sql.DB) storage.IProductionStorage {
	return &NerProductionRepo{
		DB:    db,
		Log:   logger.NewLogger(),
	}
}

func (p *NerProductionRepo) NewContract(context.Context, *pb.NewContractReq) (*pb.NewContractRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewContractUpdate(context.Context, *pb.NewContractUpdateReq) (*pb.NewContractUpdateRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewContractDelete(context.Context, *pb.NewContractDeleteReq) (*pb.NewContractDeleteRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewContractGetName(context.Context, *pb.NewContractGetNameReq) (*pb.NewContractGetNameRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewContractGetAll(context.Context, *pb.NewContractGetAllReq) (*pb.NewContractGetAllRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewInsideTheContract(context.Context, *pb.NewInsideTheContractReq) (*pb.NewInsideTheContractRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewInsideTheContractUpdate(context.Context, *pb.NewInsideTheContractUpdateReq) (*pb.NewInsideTheContractUpdateRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewInsideTheContractDelete(context.Context, *pb.NewInsideTheContractDeleteReq) (*pb.NewInsideTheContractDeleteRes, error) {
	return nil, nil
}
func (p *NerProductionRepo) NewInsideTheContractGetAll(context.Context, *pb.NewInsideTheContractGetAllReq) (*pb.NewInsideTheContractGetAllRes, error) {
	return nil, nil
}
