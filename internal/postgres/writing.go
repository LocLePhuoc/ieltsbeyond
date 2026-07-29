package postgres

import (
	"context"

	"ieltsbeyond/internal/writing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ writing.TaskRepository = (*WritingTaskRepository)(nil)

type WritingTaskRepository struct {
	db *pgxpool.Pool
}

func NewWritingTaskRepository(db *pgxpool.Pool) *WritingTaskRepository {
	return &WritingTaskRepository{db: db}
}

func (s *WritingTaskRepository) GetAllTasks1(ctx context.Context, cursor int, limit int) ([]writing.Task1, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, question, type, image_path FROM writing_tasks1
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`, limit, cursor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]writing.Task1, 0)
	for rows.Next() {
		task, err := scanTask1(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *WritingTaskRepository) GetAllTasks2(ctx context.Context, cursor int, limit int) ([]writing.Task2, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, question, category FROM writing_tasks2
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`, limit, cursor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]writing.Task2, 0)
	for rows.Next() {
		task, err := scanTask2(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *WritingTaskRepository) GetTask1(ctx context.Context, taskId string) (writing.Task1, error) {
	row := s.db.QueryRow(ctx, `SELECT id, question, type, image_path FROM writing_tasks1 WHERE id = $1`, taskId)
	return scanTask1(row)
}

func (s *WritingTaskRepository) GetTask2(ctx context.Context, taskId string) (writing.Task2, error) {
	row := s.db.QueryRow(ctx, `SELECT id, question, category FROM writing_tasks2 WHERE id = $1`, taskId)
	return scanTask2(row)
}

func (s *WritingTaskRepository) UpsertTask1(ctx context.Context, task writing.Task1) (writing.Task1, error) {
	if task.Id == "" {
		task.Id = uuid.NewString()
	}
	row := s.db.QueryRow(ctx, `
		INSERT INTO writing_tasks1 (id, question, type, image_path)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET question = $2, type = $3, image_path = $4, updated_at = now()
		RETURNING id, question, type, image_path
	`, task.Id, task.Question, task.Type, task.ImagePath)
	return scanTask1(row)
}

func (s *WritingTaskRepository) UpsertTask2(ctx context.Context, task writing.Task2) (writing.Task2, error) {
	if task.Id == "" {
		task.Id = uuid.NewString()
	}
	row := s.db.QueryRow(ctx, `
		INSERT INTO writing_tasks2 (id, question, category)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET question = $2, category = $3, updated_at = now()
		RETURNING id, question, category
	`, task.Id, task.Question, task.Category)
	return scanTask2(row)
}

func scanTask1(row rowScanner) (writing.Task1, error) {
	var task writing.Task1
	if err := row.Scan(&task.Id, &task.Question, &task.Type, &task.ImagePath); err != nil {
		return writing.Task1{}, err
	}
	return task, nil
}

func scanTask2(row rowScanner) (writing.Task2, error) {
	var task writing.Task2
	if err := row.Scan(&task.Id, &task.Question, &task.Category); err != nil {
		return writing.Task2{}, err
	}
	return task, nil
}
