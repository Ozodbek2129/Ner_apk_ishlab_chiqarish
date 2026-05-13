package storage

import (
	pb "ishlab_chiqarish/genproto/contract"
	pb1 "ishlab_chiqarish/genproto/services"
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

	CreateGroup(ctx context.Context, req *pb1.CreateGroupReq) (*pb1.CreateGroupRes, error)
	UpdateGroup(ctx context.Context, req *pb1.UpdateGroupReq) (*pb1.UpdateGroupRes, error)
	DeleteGroup(ctx context.Context, req *pb1.DeleteGroupReq) (*pb1.DeleteGroupRes, error)
	GetAllGroup(ctx context.Context, req *pb1.GetAllGroupReq) (*pb1.GetAllGroupRes, error)
	CreateWorker(ctx context.Context, req *pb1.CreateWorkerReq) (*pb1.CreateWorkerRes, error)
	UpdateWorker(ctx context.Context, req *pb1.UpdateWorkerReq) (*pb1.UpdateWorkerRes, error)
	DeleteWorker(ctx context.Context, req *pb1.DeleteWorkerReq) (*pb1.DeleteWorkerRes, error)
	GetAllWorker(ctx context.Context, req *pb1.GetAllWorkerReq) (*pb1.GetAllWorkerRes, error)
	CreateAttendance(ctx context.Context, req *pb1.CreateAttendanceReq) (*pb1.CreateAttendanceRes, error)
	UpdateAttendance(ctx context.Context, req *pb1.UpdateAttendanceReq) (*pb1.UpdateAttendanceRes, error)
	DeleteAttendance(ctx context.Context, req *pb1.DeleteAttendanceReq) (*pb1.DeleteAttendanceRes, error)
	GetDailyAttendance(ctx context.Context, req *pb1.GetDailyAttendanceReq) (*pb1.GetDailyAttendanceRes, error)
	GetAllAttendance(ctx context.Context, req *pb1.GetAllAttendanceReq) (*pb1.GetAllAttendanceRes, error)
	CreateTask(ctx context.Context, req *pb1.CreateTaskReq) (*pb1.CreateTaskRes, error)
	UpdateTask(ctx context.Context, req *pb1.UpdateTaskReq) (*pb1.UpdateTaskRes, error)
	DeleteTask(ctx context.Context, req *pb1.DeleteTaskReq) (*pb1.DeleteTaskRes, error)
	GetAllTask(ctx context.Context, req *pb1.GetAllTaskReq) (*pb1.GetAllTaskRes, error)
	GetTaskById(ctx context.Context, req *pb1.GetTaskByIdReq) (*pb1.GetTaskByIdRes, error)
}