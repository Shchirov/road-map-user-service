package userrepository

import (
	"gorm.io/gorm"
	"log/slog"
	"road-map-user-server/internal/model"
	"road-map-user-server/internal/storage"
)

type Repository struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewRepository(db *gorm.DB, log *slog.Logger) *Repository {
	return &Repository{db: db, log: log}
}

func (r *Repository) GetUserById(userId string) (*model.User, error) {
	user := &model.User{}
	err := r.db.First(user, userId).Error
	if err != nil {
		return nil, storage.ErrUserNotFound
	}
	return user, nil
}
