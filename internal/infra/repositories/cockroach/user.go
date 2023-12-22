package cockroach

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
	contracts.UserRepositoryContract
}

func GenerateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*contracts.UserContract, error) {
	var user contracts.UserContract
	err := r.db.QueryRowContext(ctx, "SELECT * FROM lv.users WHERE user_id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Type, &user.Doc, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
