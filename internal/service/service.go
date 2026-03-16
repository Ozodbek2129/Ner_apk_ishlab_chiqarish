package service

import (
	"context"
	"database/sql"
	"fmt"
	pb "ishlab_chiqarish/genproto/ishlab_chiqarish"
	"ishlab_chiqarish/internal/storage"
	"ishlab_chiqarish/internal/storage/postgres"
	"log/slog"
)

type ProdutionContract struct {
	pb.UnimplementedIshlabChiqarishServiceServer
	Production storage.IStorage

	Logger *slog.Logger
}

func NewProductionContractService(db *sql.DB, logger *slog.Logger) *ProdutionContract {
	return &ProdutionContract{
		Production:   postgres.NewPostgresStorage(db, logger),
		Logger: logger,
	}
}

func (p *ProdutionContract) NewContract(ctx context.Context, req *pb.NewContractReq) (*pb.NewContractRes, error) {
	res, err := p.Production.Production().NewContract(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error creating production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewContractUpdate(ctx context.Context , req *pb.NewContractUpdateReq) (*pb.NewContractUpdateRes, error) {
	res, err := p.Production.Production().NewContractUpdate(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error updating production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewContractDelete(ctx context.Context , req *pb.NewContractDeleteReq) (*pb.NewContractDeleteRes, error) {
	res, err := p.Production.Production().NewContractDelete(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error deleting production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewContractGetName(ctx context.Context , req *pb.NewContractGetNameReq) (*pb.NewContractGetNameRes, error) {
	res, err := p.Production.Production().NewContractGetName(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error getname production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewContractGetAll(ctx context.Context , req *pb.NewContractGetAllReq) (*pb.NewContractGetAllRes, error) {
	res, err := p.Production.Production().NewContractGetAll(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error NewContractGetAll production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewInsideTheContract(ctx context.Context , req *pb.NewInsideTheContractReq) (*pb.NewInsideTheContractRes, error) {
	res, err := p.Production.Production().NewInsideTheContract(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error inside_the_contract service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewInsideTheContractUpdate(ctx context.Context , req *pb.NewInsideTheContractUpdateReq) (*pb.NewInsideTheContractUpdateRes, error) {
	res, err := p.Production.Production().NewInsideTheContractUpdate(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error Inside_the_contract_Update service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewInsideTheContractDelete(ctx context.Context , req *pb.NewInsideTheContractDeleteReq) (*pb.NewInsideTheContractDeleteRes, error) {
	res, err := p.Production.Production().NewInsideTheContractDelete(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error NewInsideContractDelete service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}

func (p *ProdutionContract) NewInsideTheContractGetAll(ctx context.Context , req *pb.NewInsideTheContractGetAllReq) (*pb.NewInsideTheContractGetAllRes, error) {
	res, err := p.Production.Production().NewInsideTheContractGetAll(ctx, req)
	if err != nil {
		p.Logger.Error(fmt.Sprintf("Error updating production service: %v", err.Error()))
		return nil, err
	}

	return res, nil
}
