package data

import (
	"context"
	"fmt"
	"time"

	baseentities "github.com/barisakdas/Framework/Abstraction/baseentities"
	"github.com/jackc/pgx/v4"
)

type BaseRepository struct {
	db *pgx.Conn
}

func NewBaseRepository(db *pgx.Conn) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) GetAll(ctx context.Context, tableName string) ([]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []interface{}

	for rows.Next() {
		var entity baseentities.BaseEntity

		err = rows.Scan(&entity.Id, &entity.CreatedBy, &entity.CreatedAt, &entity.UpdatedBy, &entity.UpdatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, entity)
	}

	return results, nil
}

func (r *BaseRepository) GetById(ctx context.Context, tableName string, id uint64) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	row := r.db.QueryRow(ctx, query, id)

	var entity baseentities.BaseEntity

	err := row.Scan(&entity.Id, &entity.CreatedBy, &entity.CreatedAt, &entity.UpdatedBy, &entity.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *BaseRepository) Create(ctx context.Context, tableName string, entity interface{}) error {
	baseEntity := entity.(baseentities.BaseEntity)

	query := fmt.Sprintf("INSERT INTO %s (created_by, created_at, updated_by, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", tableName)

	now := time.Now()

	err := r.db.QueryRow(ctx, query, baseEntity.CreatedBy, now, baseEntity.UpdatedBy, now).Scan(&baseEntity.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) Update(ctx context.Context, tableName string, entity interface{}) error {
	baseEntity := entity.(baseentities.BaseEntity)

	query := fmt.Sprintf("UPDATE %s SET updated_by = $1, updated_at = $2 WHERE id = $3", tableName)

	_, err := r.db.Exec(ctx, query, baseEntity.UpdatedBy, time.Now(), baseEntity.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) Delete(ctx context.Context, tableName string, id uint64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
