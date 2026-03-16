package postgres

import (
	"context"
	"database/sql"
	"errors"
	pb "ishlab_chiqarish/genproto/ishlab_chiqarish"
	"ishlab_chiqarish/internal/pkg/logger"
	"ishlab_chiqarish/internal/storage"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type NerProductionRepo struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewNerProductionRepo(db *sql.DB) storage.IProductionStorage {
	return &NerProductionRepo{
		DB:  db,
		Log: logger.NewLogger(),
	}
}

func (p *NerProductionRepo) NewContract(ctx context.Context, req *pb.NewContractReq) (*pb.NewContractRes, error) {
	query := `insert into production_contract (
						id, contract_name, contract_number, 
						contract_deadline, contract_file, 
						responsible_person, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	id := uuid.NewString()
	newtime := time.Now()
	_, err := p.DB.ExecContext(ctx, query, id, req.ContractName, req.ContractNumber, req.ContractDeadline, req.ContractFileUrl, req.ResponsiblePerson, newtime, newtime)
	if err != nil {
		p.Log.Error("Error inserting production", "error", err.Error())
		return nil, err
	}

	return &pb.NewContractRes{Message: "Successful"}, nil
}

func (p *NerProductionRepo) NewContractUpdate(ctx context.Context, req *pb.NewContractUpdateReq) (*pb.NewContractUpdateRes, error) {

	queryUpdate := `
		UPDATE production_contract 
		SET contract_name = $1, 
		    contract_number = $2, 
		    contract_deadline = $3, 
		    contract_file = $4, 
		    responsible_person = $5 
		WHERE id = $6 AND deleted_at = 0
	`

	result, err := p.DB.ExecContext(
		ctx,
		queryUpdate,
		req.ContractName,
		req.ContractNumber,
		req.ContractDeadline,
		req.ContractFileUrl,
		req.ResponsiblePerson,
		req.Id,
	)
	if err != nil {
		p.Log.Error("Error updating contract", "err", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb.NewContractUpdateRes{
			Message: "No contract updated",
		}, nil
	}

	return &pb.NewContractUpdateRes{
		Message: "Contract updated successfully",
	}, nil
}

func (p *NerProductionRepo) NewContractDelete(ctx context.Context, req *pb.NewContractDeleteReq) (*pb.NewContractDeleteRes, error) {
	query := `update production_contract set deleted_at = $1 where id = $2`
	res, err := p.DB.ExecContext(ctx, query, time.Now(), req.Id)
	if err != nil {
		p.Log.Error("Error Delete contract", "error", err.Error())
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb.NewContractDeleteRes{
			Message: "No contract delete updated",
		}, nil
	}

	return &pb.NewContractDeleteRes{
		Message: "Contract Deleteupdated successfully",
	}, nil
}

func (p *NerProductionRepo) NewContractGetName(ctx context.Context, req *pb.NewContractGetNameReq) (*pb.NewContractGetNameRes, error) {
	var contract pb.Contract
	query := ` select id, contract_name, contract_number, contract_deadline, contract_file, responsible_person from production_contract where contract_name = $1`
	err := p.DB.QueryRow(query, req.Name).Scan(&contract.Id, &contract.ContractName, &contract.ContractNumber, &contract.ContractDeadline, &contract.ContractFileUrl, &contract.ResponsiblePerson)

	if err == sql.ErrNoRows {
		p.Log.Error("No contract found", "name", req.Name)
		return nil, errors.New("no contract found")
	} else if err != nil {
		p.Log.Error("Error getting contract by name", "err", err)
		return nil, err
	}

	return &pb.NewContractGetNameRes{
		Id:                contract.Id,
		ContractName:      contract.ContractName,
		ContractNumber:    contract.ContractNumber,
		ContractDeadline:  contract.ContractDeadline,
		ContractFileUrl:   contract.ContractFileUrl,
		ResponsiblePerson: contract.ResponsiblePerson,
	}, nil
}

func (p *NerProductionRepo) NewContractGetAll(ctx context.Context, req *pb.NewContractGetAllReq) (*pb.NewContractGetAllRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `SELECT id, contract_name, contract_number, contract_deadline, contract_file, responsible_person FROM production_contract ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := p.DB.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		p.Log.Error("Error fetching production_contract", "err", err)
		return nil, errors.New("failed to fetch production_contract")
	}
	defer rows.Close()

	var contract []*pb.Contract
	for rows.Next() {
		var c pb.Contract
		if err := rows.Scan(&c.Id, &c.ContractName, &c.ContractNumber, &c.ContractDeadline, &c.ContractFileUrl, &c.ResponsiblePerson); err != nil {
			p.Log.Error("Error scanning production_contract row", "error", err)
			return nil, err
		}
		contract = append(contract, &c)
	}

	if err := rows.Err(); err != nil {
		p.Log.Error("Row iteration error", "err", err)
		return nil, errors.New("error iterating rows")
	}

	return &pb.NewContractGetAllRes{Contract: contract}, nil
}

func (p *NerProductionRepo) NewInsideTheContract(ctx context.Context, req *pb.NewInsideTheContractReq) (*pb.NewInsideTheContractRes, error) {
	query := `insert into inside_contract (
						id, contract_id, inside_contract_name, inside_contract_price, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6)`

	id := uuid.NewString()
	newtime := time.Now()
	_, err := p.DB.ExecContext(ctx, query, id, req.ContractId, req.Name, req.Price, newtime, newtime)
	if err != nil {
		p.Log.Error("Error inserting inside_contract", "error", err.Error())
		return nil, err
	}

	return &pb.NewInsideTheContractRes{Message: "Successful"}, nil
}
func (p *NerProductionRepo) NewInsideTheContractUpdate(ctx context.Context, req *pb.NewInsideTheContractUpdateReq) (*pb.NewInsideTheContractUpdateRes, error) {
	queryUpdate := `
		UPDATE inside_contract 
		SET inside_contract_name = $1, inside_contract_price = $2, updated_at = $3
		WHERE contract_id = $4 AND deleted_at = 0
	`
	result, err := p.DB.ExecContext(
		ctx,
		queryUpdate,
		req.Name,
		req.Price,
		time.Now(),
		req.ContractId,
	)
	if err != nil {
		p.Log.Error("Error updating inside_contract", "err", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb.NewInsideTheContractUpdateRes{
			Message: "No inside_contract updated",
		}, nil
	}

	return &pb.NewInsideTheContractUpdateRes{
		Message: "Inside_contract updated successfully",
	}, nil
}

func (p *NerProductionRepo) NewInsideTheContractDelete(ctx context.Context, req *pb.NewInsideTheContractDeleteReq) (*pb.NewInsideTheContractDeleteRes, error) {
	query := `update inside_contract set deleted_at = $1 where id = $2`
	res, err := p.DB.ExecContext(ctx, query, time.Now(), req.Id)
	if err != nil {
		p.Log.Error("Error Delete inside_contract", "error", err.Error())
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb.NewInsideTheContractDeleteRes{
			Message: "No inside_contract delete updated",
		}, nil
	}

	return &pb.NewInsideTheContractDeleteRes{
		Message: "Inside_contract Deleteupdated successfully",
	}, nil
}

func (p *NerProductionRepo) NewInsideTheContractGetAll(ctx context.Context, req *pb.NewInsideTheContractGetAllReq) (*pb.NewInsideTheContractGetAllRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `SELECT id, contract_id, inside_contract_name, inside_contract_price, created_at, updated_at FROM inside_contract ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := p.DB.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		p.Log.Error("Error fetching inside_contract", "err", err)
		return nil, errors.New("failed to fetch inside_contract")
	}
	defer rows.Close()

	var inside_contract []*pb.NewInsideTheContractt
	for rows.Next() {
		var c pb.NewInsideTheContractt
		if err := rows.Scan(&c.Id, &c.ContractId, &c.Name, &c.Price, &c.CreatedAt, &c.UpdatedAt); err != nil {
			p.Log.Error("Error scanning inside_contract row", "error", err)
			return nil, err
		}
		inside_contract = append(inside_contract, &c)
	}

	if err := rows.Err(); err != nil {
		p.Log.Error("Row iteration error", "err", err)
		return nil, errors.New("error iterating rows")
	}

	return &pb.NewInsideTheContractGetAllRes{All: inside_contract}, nil
}
