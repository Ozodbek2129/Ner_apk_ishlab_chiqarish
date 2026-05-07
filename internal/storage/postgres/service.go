package postgres

import (
	"context"
	"errors"
	"fmt"
	pb1 "ishlab_chiqarish/genproto/services"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NerProductionRepo) CreateGroup(ctx context.Context, req *pb1.CreateGroupReq) (*pb1.CreateGroupRes, error) {
	query := `insert into groups (
						id, name, master, created_at, updated_at
				) values ($1, $2, $3, $4, $5)`

	id := uuid.NewString()
	newtime := time.Now()

	_, err := s.DB.Exec(query, id, req.Name, req.Master, newtime, newtime)
	if err != nil {
		s.Log.Error("Error inserting groups", "error", err.Error())
		return nil, err
	}

	return &pb1.CreateGroupRes{Message: "Successful"}, nil
}

func (s *NerProductionRepo) UpdateGroup(ctx context.Context, req *pb1.UpdateGroupReq) (*pb1.UpdateGroupRes, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1
	var updatedAt time.Time

	if req.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *req.Name)
		argPos++
	}

	if req.Master != nil {
		setParts = append(setParts, fmt.Sprintf("master = $%d", argPos))
		args = append(args, *req.Master)
		argPos++
	}

	if len(setParts) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argPos))
	args = append(args, time.Now())
	argPos++

	query := fmt.Sprintf(`
		UPDATE groups 
		SET %s
		WHERE id = $%d
		RETURNING id, name, master, updated_at
	`, strings.Join(setParts, ", "), argPos)

	args = append(args, req.Id)

	err := s.DB.QueryRowContext(ctx, query, args...).Scan(
		&req.Id,
		&req.Name,
		&req.Master,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &pb1.UpdateGroupRes{Message: "Successful"}, nil
}

func (s *NerProductionRepo) DeleteGroup(ctx context.Context, req *pb1.DeleteGroupReq) (*pb1.DeleteGroupRes, error) {
	query := `update groups set deleted_at = $1 where id = $2 and deleted_at = 0`

	res, err := s.DB.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		s.Log.Error("Error Delete services", "error", err.Error())
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb1.DeleteGroupRes{
			Message: "No service delete updated",
		}, nil
	}

	return &pb1.DeleteGroupRes{
		Message: "Service Delete successfully",
	}, nil
}

func (s *NerProductionRepo) GetAllGroup(ctx context.Context, req *pb1.GetAllGroupReq) (*pb1.GetAllGroupRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `
	SELECT id, name, master 
	FROM groups 
	WHERE deleted_at = 0
	ORDER BY created_at DESC 
	LIMIT $1 OFFSET $2
`
	rows, err := s.DB.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		s.Log.Error("Error fetching services", "err", err)
		return nil, errors.New("failed to fetch services")
	}
	defer rows.Close()

	var groups []*pb1.Group
	for rows.Next() {
		var c pb1.Group
		if err := rows.Scan(&c.Id, &c.Name, &c.Master); err != nil {
			s.Log.Error("Error scanning services row", "error", err)
			return nil, err
		}
		groups = append(groups, &c)
	}

	if err := rows.Err(); err != nil {
		s.Log.Error("Row iteration error", "err", err)
		return nil, errors.New("error iterating rows")
	}

	return &pb1.GetAllGroupRes{Groups: groups}, nil
}

func (s *NerProductionRepo) CreateWorker(ctx context.Context, req *pb1.CreateWorkerReq) (*pb1.CreateWorkerRes, error) {
	query := `insert into workers (
					id, group_id, name, image, created_at, updated_at
				) values ($1, $2, $3, $4, $5, $6)`

	id := uuid.NewString()
	newtime := time.Now()

	_, err := s.DB.Exec(query, id, req.GroupId, req.Name, req.Image, newtime, newtime)
	if err != nil {
		s.Log.Error("Error inserting groups", "error", err.Error())
		return nil, err
	}

	return &pb1.CreateWorkerRes{Message: "Successfull"}, nil
}

func (s *NerProductionRepo) UpdateWorker(ctx context.Context, req *pb1.UpdateWorkerReq) (*pb1.UpdateWorkerRes, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1
	var updatedAt time.Time

	if req.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *req.Name)
		argPos++
	}

	if req.Image != nil {
		setParts = append(setParts, fmt.Sprintf("image = $%d", argPos))
		args = append(args, *req.Image)
		argPos++
	}

	if req.GroupId != nil {
		setParts = append(setParts, fmt.Sprintf("group_id = $%d", argPos))
		args = append(args, *req.GroupId)
		argPos++
	}

	if len(setParts) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argPos))
	args = append(args, time.Now())
	argPos++

	query := fmt.Sprintf(`
		UPDATE workers 
		SET %s
		WHERE id = $%d
		RETURNING id, name, image, group_id, updated_at
	`, strings.Join(setParts, ", "), argPos)

	args = append(args, req.Id)

	err := s.DB.QueryRowContext(ctx, query, args...).Scan(
		&req.Id,
		&req.Name,
		&req.Image,
		&req.GroupId,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &pb1.UpdateWorkerRes{Message: "Successfull"}, nil
}

func (s *NerProductionRepo) DeleteWorker(ctx context.Context, req *pb1.DeleteWorkerReq) (*pb1.DeleteWorkerRes, error) {
	query := `update workers set deleted_at = $1 where id = $2 and deleted_at = 0`

	res, err := s.DB.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		s.Log.Error("Error Delete services", "error", err.Error())
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb1.DeleteWorkerRes{
			Message: "No service delete",
		}, nil
	}

	return &pb1.DeleteWorkerRes{
		Message: "Service Delete successfully",
	}, nil
}

func (s *NerProductionRepo) GetAllWorker(ctx context.Context, req *pb1.GetAllWorkerReq) (*pb1.GetAllWorkerRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `
	SELECT id, name, image, group_id 
	FROM workers 
	WHERE deleted_at = 0 and group_id = $1
	ORDER BY created_at DESC 
	LIMIT $2 OFFSET $3
`
	rows, err := s.DB.QueryContext(ctx, query, req.GroupId, req.Limit, offset)
	if err != nil {
		s.Log.Error("Error fetching services", "err", err)
		return nil, errors.New("failed to fetch services")
	}
	defer rows.Close()

	var workers []*pb1.Worker
	for rows.Next() {
		var c pb1.Worker
		if err := rows.Scan(&c.Id, &c.Name, &c.Image, &c.GroupId); err != nil {
			s.Log.Error("Error scanning services row", "error", err)
			return nil, err
		}
		workers = append(workers, &c)
	}

	if err := rows.Err(); err != nil {
		s.Log.Error("Row iteration error", "err", err)
		return nil, errors.New("error iterating rows")
	}

	return &pb1.GetAllWorkerRes{Workers: workers}, nil
}

func (s *NerProductionRepo) CreateAttendance(ctx context.Context, req *pb1.CreateAttendanceReq) (*pb1.CreateAttendanceRes, error) {
	if len(req.Attendance) == 0 {
		return nil, fmt.Errorf("attendance list is empty")
	}

	query := `
		INSERT INTO attendance (
			id, worker_id, situation, reason, is_present, work_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	now := time.Now()

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx error: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, att := range req.Attendance {

		if att.WorkerId == "" {
			err = fmt.Errorf("worker_id is required")
			return nil, err
		}

		id := uuid.NewString()

		workDate := att.WorkDate.AsTime()

		_, err = tx.Exec(query,
			id,
			att.WorkerId,
			att.Situation,
			att.Reason,
			att.IsPresent,
			workDate,
			now,
			now,
		)

		if err != nil {
			return nil, fmt.Errorf("insert attendance failed: %w", err)
		}
	}

	return &pb1.CreateAttendanceRes{
		Message: "attendance created successfully",
	}, nil
}

func (s *NerProductionRepo) UpdateAttendance(ctx context.Context, req *pb1.UpdateAttendanceReq) (*pb1.UpdateAttendanceRes, error) {
	if len(req.Attendance) == 0 {
		return nil, fmt.Errorf("attendance list is empty")
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx error: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, att := range req.Attendance {

		if att.Id == "" || att.WorkerId == "" {
			err = fmt.Errorf("id and worker_id are required")
			return nil, err
		}

		setParts := []string{}
		args := []interface{}{}
		argPos := 1

		if att.Situation != nil {
			setParts = append(setParts, fmt.Sprintf("situation = $%d", argPos))
			args = append(args, *att.Situation)
			argPos++
		}

		if att.Reason != nil {
			setParts = append(setParts, fmt.Sprintf("reason = $%d", argPos))
			args = append(args, *att.Reason)
			argPos++
		}

		if att.IsPresent != nil {
			setParts = append(setParts, fmt.Sprintf("is_present = $%d", argPos))
			args = append(args, *att.IsPresent)
			argPos++
		}

		if att.WorkDate != nil {
			setParts = append(setParts, fmt.Sprintf("work_date = $%d", argPos))
			args = append(args, att.WorkDate.AsTime())
			argPos++
		}

		if len(setParts) == 0 {
			continue
		}

		setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argPos))
		args = append(args, time.Now())
		argPos++

		query := fmt.Sprintf(`
			UPDATE attendance
			SET %s
			WHERE id = $%d AND worker_id = $%d
		`, strings.Join(setParts, ", "), argPos, argPos+1)

		args = append(args, att.Id, att.WorkerId)

		res, execErr := tx.ExecContext(ctx, query, args...)
		if execErr != nil {
			err = fmt.Errorf("update attendance failed: %w", execErr)
			return nil, err
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			err = fmt.Errorf("attendance not found id=%s", att.Id)
			return nil, err
		}
	}

	return &pb1.UpdateAttendanceRes{
		Message: "Successfully updated",
	}, nil
}

func (s *NerProductionRepo) DeleteAttendance(ctx context.Context, req *pb1.DeleteAttendanceReq) (*pb1.DeleteAttendanceRes, error) {
	query := `update attendance set deleted_at = $1 where work_date = $2 and deleted_at = 0`

	res, err := s.DB.Exec(query, req.TodayDate, req.DeleteDate)
	if err != nil {
		s.Log.Error("Error Delete services", "error", err.Error())
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb1.DeleteAttendanceRes{
			Message: "No service delete",
		}, nil
	}

	return &pb1.DeleteAttendanceRes{Message: "Successfull"}, nil
}

func (s *NerProductionRepo) GetDailyAttendance(ctx context.Context, req *pb1.GetDailyAttendanceReq) (*pb1.GetDailyAttendanceRes, error) {
	query := `select id, worker_id, situation, reason, is_present, work_date from attendance where work_date = $1 and deleted_at = 0`

	rows, err := s.DB.Query(query, req.WorkDate)
	if err != nil {
		s.Log.Error("Error fetching services", "err", err)
		return nil, errors.New("failed to fetch services")
	}
	defer rows.Close()

	var attendance  []*pb1.Attendance
	for rows.Next() {
		var a pb1.Attendance
		if err := rows.Scan(&a.Id, &a.WorkerId, &a.Situation, &a.Reason, &a.IsPresent, &a.WorkDate); err != nil {
			s.Log.Error("Error scanning service row", "error", err)
			return  nil, err
		}
		attendance = append(attendance, &a)
	}

	if err := rows.Err(); err != nil {
		s.Log.Error("Row iteration error", "error", err)
		return nil, err
	}

	return &pb1.GetDailyAttendanceRes{Attendance: attendance}, nil
}

func (s *NerProductionRepo) GetAllAttendance(ctx context.Context, req *pb1.GetAllAttendanceReq) (*pb1.GetAllAttendanceRes, error) {
	if req.StartDate == nil || req.EndDate == nil {
		return nil, fmt.Errorf("start_date and end_date are required")
	}

	startDate := req.StartDate.AsTime()
	endDate := req.EndDate.AsTime()

	query := `
		SELECT id, worker_id, situation, reason, is_present, work_date
		FROM attendance
		WHERE work_date BETWEEN $1 AND $2
		AND deleted_at = 0
		ORDER BY work_date
	`

	rows, err := s.DB.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var attendances []*pb1.Attendance

	for rows.Next() {
		var (
			id         string
			workerID   string
			situation  string
			reason     string
			isPresent  bool
			workDate   time.Time
		)

		if err := rows.Scan(
			&id,
			&workerID,
			&situation,
			&reason,
			&isPresent,
			&workDate,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		attendances = append(attendances, &pb1.Attendance{
			Id:         id,
			WorkerId:   workerID,
			Situation:  situation,
			Reason:     reason,
			IsPresent:  isPresent,
			WorkDate:   timestamppb.New(workDate),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb1.GetAllAttendanceRes{
		Attendance: attendances,
	}, nil
}

func (s *NerProductionRepo) CreateTask(ctx context.Context, req *pb1.CreateTaskReq) (*pb1.CreateTaskRes, error) {
	query := `insert into tasks (id, worker_id, task, deadline, start_time, end_time, created_at, updated_at
					) values ($1, $2, $3, $4, $5, $6, $7, $8)`

	id := uuid.NewString()
	newtime := time.Now()

	_, err := s.DB.Exec(query, id, req.WorkerId, req.Task, req.Deadline, req.StartTime, req.EndTime, newtime, newtime)
	if err != nil {
		return nil, fmt.Errorf("insert tasks failed: %w", err)
	}

	return &pb1.CreateTaskRes{Message: "Successfull"}, nil
}

func (s *NerProductionRepo) UpdateTask(ctx context.Context, req *pb1.UpdateTaskReq) (*pb1.UpdateTaskRes, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if req.Deadline != nil {
		setParts = append(setParts, fmt.Sprintf("deadline = $%d", argPos))
		args = append(args, req.Deadline)
		argPos++
	}

	if req.EndTime != nil {
		setParts = append(setParts, fmt.Sprintf("end_time = $%d", argPos))
		args = append(args, req.EndTime.AsTime())
		argPos++
	}

	if len(setParts) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argPos))
	args = append(args, time.Now())
	argPos++

	query := fmt.Sprintf(`
		UPDATE tasks 
		SET %s
		WHERE id = $%d
		RETURNING id, worker_id, task, deadline, start_time, end_time
	`, strings.Join(setParts, ", "), argPos)

	args = append(args, req.Id)

	var (
		id         string
		workerID   string
		task       string
		deadline   string
		startTime  time.Time
		endTime    time.Time
	)

	err := s.DB.QueryRowContext(ctx, query, args...).Scan(
		&id,
		&workerID,
		&task,
		&deadline,
		&startTime,
		&endTime,
	)
	if err != nil {
		return nil, err
	}

	return &pb1.UpdateTaskRes{
		Message: "Successfully updated",
	}, nil
}

func (s *NerProductionRepo) DeleteTask(ctx context.Context, req *pb1.DeleteTaskReq) (*pb1.DeleteTaskRes, error) {
	query := `update tasks set deleted_at = $1 where id = $2 and deleted_at = 0`

	res, err := s.DB.Exec(query, time.Now().Unix(), req.Id)
	if err != nil {
		s.Log.Error("Error tasks delete", "error", err)
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return &pb1.DeleteTaskRes{
			Message: "No service delete",
		}, nil
	}

	return &pb1.DeleteTaskRes{Message: "Successfull"}, nil
}

func (s *NerProductionRepo) GetAllTask(ctx context.Context, req *pb1.GetAllTaskReq) (*pb1.GetAllTaskRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `
	SELECT id, worker_id, task, deadline, start_time, end_time 
	FROM tasks 
	WHERE deleted_at = 0 and worker_id = $1
	ORDER BY created_at DESC 
	LIMIT $2 OFFSET $3
`
	rows, err := s.DB.QueryContext(ctx, query, req.WorkerId, req.Limit, offset)
	if err != nil {
		s.Log.Error("Error fetching services", "err", err)
		return nil, errors.New("failed to fetch services")
	}
	defer rows.Close()

	var tasks []*pb1.Task
	for rows.Next() {
		var c pb1.Task
		if err := rows.Scan(&c.Id, &c.WorkerId, &c.Task, &c.Deadline, &c.StartTime, &c.EndTime); err != nil {
			s.Log.Error("Error scanning services row", "error", err)
			return nil, err
		}
		tasks = append(tasks, &c)
	}

	if err := rows.Err(); err != nil {
		s.Log.Error("Row iteration error", "err", err)
		return nil, errors.New("error iterating rows")
	}

	return &pb1.GetAllTaskRes{Tasks: tasks}, nil
}

func (s *NerProductionRepo) GetTaskById(ctx context.Context, req *pb1.GetTaskByIdReq) (*pb1.GetTaskByIdRes, error) {
	return nil, nil
}
