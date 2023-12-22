package cockroach

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthRepository struct {
	db *sql.DB
	contracts.AuthRepositoryContract
}

func GenerateAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FirstAdmin(ctx context.Context) (bool, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO lv.users (user_id, name, email, password, birthdate, type, doc, phone, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NULL, NULL)")

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	password := os.Getenv("ADMIN_PASSWORD")
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return false, err
	}

	date, err := time.Parse(time.RFC3339Nano, "1990-01-01T00:00:00.000Z-03:00")

	email := os.Getenv("ADMIN_EMAIL")
	_, err = stmt.ExecContext(ctx, uuid.New().String(), "Admin", email, passwordEncrypted, date, 1, "00000000000", "00000000000")

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	err = tx.Commit()

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	return true, nil
}

func (r *AuthRepository) SignInADMIN(ctx context.Context, email string, password string) (*contracts.UserContract, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM lv.users WHERE email = $1 AND deleted_at IS NULL AND type = 1")

	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, email)

	user := contracts.UserContract{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Type, &user.Doc, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) SignInCLIENT(ctx context.Context, email string, password string) (*contracts.UserContract, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM lv.users WHERE email = $1 AND deleted_at IS NULL AND type = 0")

	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, email)

	user := contracts.UserContract{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Type, &user.Doc, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) SignUpCLIENT(ctx context.Context, user contracts.UserContract) (bool, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO lv.users (user_id, name, email, password, birthdate, type, doc, phone, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NULL, NULL)")

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	_, err = stmt.ExecContext(ctx, uuid.New().String(), user.Name, user.Email, passwordEncrypted, user.BirthDate, 0, user.Doc, user.Phone)

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	err = tx.Commit()

	if err != nil {
		rb := tx.Rollback()
		if rb != nil {
			return false, rb
		}
		return false, err
	}

	return true, nil
}

func (r *AuthRepository) SignInWORKER(ctx context.Context, email string, password string) (*contracts.UserContract, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM lv.users WHERE email = $1 AND deleted_at IS NULL AND type = 2")

	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, email)

	user := contracts.UserContract{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Type, &user.Doc, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserByID(ctx context.Context, id string) (*contracts.UserContract, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM lv.users WHERE user_id = $1 AND deleted_at IS NULL")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)

	user := contracts.UserContract{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Type, &user.Doc, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
