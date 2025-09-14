package storage

import (
	"fmt"

	"github.com/petrick-ribeiro/devops-pucpr/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(user, pass, host, port, dbname string) (*PostgresStorage, error) {
	dsn := fmt.Sprintf(`
		user=%s
		password=%s
		host=%s
		port=%s
		dbname=%s
		sslmode=disable`, user, pass, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	err := s.db.AutoMigrate(&types.Todo{})
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) GetAll() ([]*types.Todo, error) {
	todos := []*types.Todo{}

	result := s.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (s *PostgresStorage) Get(id uint64) (*types.Todo, error) {
	todo := types.Todo{}

	result := s.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (s *PostgresStorage) Insert(todo *types.Todo) error {
	result := s.db.Create(&todo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PostgresStorage) Update(ut *types.Todo, id uint64) (*types.Todo, error) {
	todo := types.Todo{}

	result := s.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if ut.Title != "" {
		todo.Title = ut.Title
	}

	if ut.Description != "" {
		todo.Description = ut.Description
	}

	todo.UpdatedAt = ut.UpdatedAt
	todo.Done = ut.Done

	result = s.db.Save(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (s *PostgresStorage) Delete(id uint64) (*types.Todo, error) {
	todo := types.Todo{}

	result := s.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result := s.db.Delete(&todo, id); result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}
