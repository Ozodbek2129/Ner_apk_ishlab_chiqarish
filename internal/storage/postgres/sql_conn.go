package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"ishlab_chiqarish/internal/config"
	"ishlab_chiqarish/internal/storage"

	_ "github.com/lib/pq"
)

type postgresStorage struct {
	db    *sql.DB
	log   *slog.Logger
}

func NewPostgresStorage(db *sql.DB, log *slog.Logger) storage.IStorage {
	return &postgresStorage{
		db:  db,
		log: log,
	}
}

func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (p *postgresStorage) Close() {
	p.db.Close()
}

func (p *postgresStorage) Production() storage.IProductionStorage {
	return NewNerProductionRepo(p.db)
}
