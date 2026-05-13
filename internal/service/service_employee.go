package service

import (
	"context"
	"database/sql"
	pb "ishlab_chiqarish/genproto/services"
	"ishlab_chiqarish/internal/storage"
	"ishlab_chiqarish/internal/storage/postgres"
	"log/slog"
)

type EmployeeServiceRepo struct {
	pb.UnimplementedServicesServiceServer
	Employee storage.IStorage
	Logger   *slog.Logger
}

func NewEmployeeServiceRepo(db *sql.DB, logger *slog.Logger) *EmployeeServiceRepo {
	return &EmployeeServiceRepo{
		Employee: postgres.NewPostgresStorage(db, logger),
		Logger:   logger,
	}
}

func (e *EmployeeServiceRepo) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.CreateGroupRes, error) {
	res, err := e.Employee.Production().CreateGroup(ctx, req)
	if err != nil {
		e.Logger.Error("Error CreateGroup service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) UpdateGroup(ctx context.Context, req *pb.UpdateGroupReq) (*pb.UpdateGroupRes, error) {
	res, err := e.Employee.Production().UpdateGroup(ctx, req)
	if err != nil {
		e.Logger.Error("Error UpdateGroup service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) DeleteGroup(ctx context.Context, req *pb.DeleteGroupReq) (*pb.DeleteGroupRes, error) {
	res, err := e.Employee.Production().DeleteGroup(ctx, req)
	if err != nil {
		e.Logger.Error("Error DeleteGroup service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetAllGroup(ctx context.Context, req *pb.GetAllGroupReq) (*pb.GetAllGroupRes, error) {
	res, err := e.Employee.Production().GetAllGroup(ctx, req)
	if err != nil {
		e.Logger.Error("Error GetAllGroup service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) CreateWorker(ctx context.Context, req *pb.CreateWorkerReq) (*pb.CreateWorkerRes, error) {
	res, err := e.Employee.Production().CreateWorker(ctx, req)
	if err != nil {
		e.Logger.Error("Error CreateWorker service file", "error", err.Error())
		return nil, nil
	}

	return res, nil
}

func (e *EmployeeServiceRepo) UpdateWorker(ctx context.Context, req *pb.UpdateWorkerReq) (*pb.UpdateWorkerRes, error) {
	res, err := e.Employee.Production().UpdateWorker(ctx, req)
	if err != nil {
		e.Logger.Error("Error UpdateWorker service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) DeleteWorker(ctx context.Context, req *pb.DeleteWorkerReq) (*pb.DeleteWorkerRes, error) {
	res, err := e.Employee.Production().DeleteWorker(ctx, req)
	if err != nil {
		e.Logger.Error("Error DeleteWorker service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetAllWorker(ctx context.Context, req *pb.GetAllWorkerReq) (*pb.GetAllWorkerRes, error) {
	res, err := e.Employee.Production().GetAllWorker(ctx, req)
	if err != nil {
		e.Logger.Error("Error GetAllWorker service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) CreateAttendance(ctx context.Context, req *pb.CreateAttendanceReq) (*pb.CreateAttendanceRes, error) {
	res, err := e.Employee.Production().CreateAttendance(ctx, req)
	if err != nil {
		e.Logger.Error("Error CreateAttendance service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) UpdateAttendance(ctx context.Context, req *pb.UpdateAttendanceReq) (*pb.UpdateAttendanceRes, error) {
	res, err := e.Employee.Production().UpdateAttendance(ctx, req)
	if err != nil {
		e.Logger.Error("Error UpdateAttendance service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) DeleteAttendance(ctx context.Context, req *pb.DeleteAttendanceReq) (*pb.DeleteAttendanceRes, error) {
	res, err := e.Employee.Production().DeleteAttendance(ctx, req)
	if err != nil {
		e.Logger.Error("Error DeleteAttendance service file")
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetDailyAttendance(ctx context.Context, req *pb.GetDailyAttendanceReq) (*pb.GetDailyAttendanceRes, error) {
	res, err := e.Employee.Production().GetDailyAttendance(ctx, req)
	if err != nil {
		e.Logger.Error("Error GetDailyAttendance service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetAllAttendance(ctx context.Context, req *pb.GetAllAttendanceReq) (*pb.GetAllAttendanceRes, error) {
	res, err := e.Employee.Production().GetAllAttendance(ctx, req)
	if err != nil {
		e.Logger.Error("Error GetAllAttendance service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	res, err := e.Employee.Production().CreateTask(ctx, req)
	if err != nil {
		e.Logger.Error("Error CreateTask service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) UpdateTask(ctx context.Context, req *pb.UpdateTaskReq) (*pb.UpdateTaskRes, error) {
	res, err := e.Employee.Production().UpdateTask(ctx, req)
	if err != nil {
		e.Logger.Error("Error UpdateTask service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) DeleteTask(ctx context.Context, req *pb.DeleteTaskReq) (*pb.DeleteTaskRes, error) {
	res, err := e.Employee.Production().DeleteTask(ctx, req)
	if err != nil {
		e.Logger.Error("Error DeleteTask service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetAllTask(ctx context.Context, req *pb.GetAllTaskReq) (*pb.GetAllTaskRes, error) {
	res, err := e.Employee.Production().GetAllTask(ctx, req)
	if err != nil {
		e.Logger.Error("Error GetAllTask service file", "error", err.Error())
		return nil, err
	}

	return res, nil
}

func (e *EmployeeServiceRepo) GetTaskById(ctx context.Context, req *pb.GetTaskByIdReq) (*pb.GetTaskByIdRes, error) {
	return nil, nil
}
