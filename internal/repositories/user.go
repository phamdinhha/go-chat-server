package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-chat-server/internal/models"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, req *models.User) (*models.User, error) {
	createdUser := &models.User{}
	if err := r.db.QueryRowxContext(
		ctx,
		CREATE_USER_QUERY,
		req.ID,
		req.UserName,
		req.Email,
		req.Password,
	).StructScan(createdUser); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return createdUser, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.GetContext(ctx, user, GET_USER_BY_EMAIL, email); err != nil {
		return nil, err
	}
	return user, nil
}
