package repo

import (
	"context"
	"fmt"
	"max-db-svc/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	sq "github.com/Masterminds/squirrel"
)

type userRepo struct {
	db        *pgx.Conn
	tableName string
}

func NewUserRepo() *userRepo {
	dbc := NewDatabaseConnection()
	return &userRepo{
		db:        dbc,
		tableName: "\"Users\"",
	}
}

// Check interface complience
var _ UserRepository = (*userRepo)(nil)

func (r *userRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, _ := psql.Select("*").From(r.tableName).Where(sq.Eq{"\"Username\"": username}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetByUsername] DB query error: %v", qErr)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetByUsername] %v", err)
	}

	if !rows.Next() {
		return model.User{}, fmt.Errorf("[UserRepo GetByUsername] Invalid username")
	}

	var u model.User
	rows.Scan(&u.ID, &u.Username, &u.Name, &u.Password)
	return u, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, ID string) (model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	uid, uErr := uuid.Parse(ID)
	if uErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByID] Invalid user id")
	}
	query, args, _ := psql.Select("*").From(r.tableName).Where(sq.Eq{"\"Id\"": uid}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByID] DB query error: %v", qErr)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByID] %v", err)
	}

	if !rows.Next() {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByID] Invalid user id")
	}
	var u model.User
	rows.Scan(&u.ID, &u.Username, &u.Name, &u.Password)
	return u, nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]model.User, error) {
	query, args, _ := sq.Select("*").From(r.tableName).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return []model.User{}, fmt.Errorf("[UserRepo GetAllUsers] DB query error: %v", qErr)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Name, &u.Password); err != nil {
			return users, fmt.Errorf("[UserRepo GetAllUsers]  %v", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return users, fmt.Errorf("[UserRepo GetAllUsers]  %v", err)
	}

	return users, nil
}

func (r *userRepo) GetUserByToken(ctx context.Context, token string) (model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, _ := psql.Select("*").From("\"Auth\"").Where(sq.Eq{"\"Token\"": token}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] DB query error: %v", qErr)
	}
	defer rows.Close()
	if !rows.Next() {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] Invalid user token")
	}
	var t model.Token
	rows.Scan(&t.ID, &t.UserID, &t.Token, &t.Active, &t.CreatedAt)

	uid, uErr := uuid.Parse(t.UserID)
	if uErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] Invalid user id")
	}

	query, args, _ = psql.Select("*").From(r.tableName).Where(sq.Eq{"\"Id\"": uid}).ToSql()
	rows, qErr = r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] DB query error: %v", qErr)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] %v", err)
	}

	if !rows.Next() {
		return model.User{}, fmt.Errorf("[UserRepo GetUserByToken] Invalid user id")
	}
	var u model.User
	rows.Scan(&u.ID, &u.Username, &u.Name, &u.Password)
	return u, nil
}
