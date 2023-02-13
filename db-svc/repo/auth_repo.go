package repo

import (
	"context"
	"fmt"
	"max-db-svc/model"
	"max-db-svc/util"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type authRepo struct {
	db        *pgx.Conn
	tableName string
}

func NewAuthRepo() *authRepo {
	dbc := NewDatabaseConnection()
	return &authRepo{
		db:        dbc,
		tableName: "\"Auth\"",
	}
}

// Check interface complience
var _ AuthRepository = (*authRepo)(nil)

func (r *authRepo) Insert(ctx context.Context, userID, token string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	idUID, uErr := uuid.NewRandom()
	if uErr != nil {
		return fmt.Errorf("[AuthRepo Insert] %vv", uErr)
	}
	userUID, uErr := uuid.Parse(userID)
	if uErr != nil {
		return fmt.Errorf("[AuthRepo Insert] %vv", uErr)
	}
	query, args, _ := psql.Insert(r.tableName).
		Columns("\"Id\", \"UserId\", \"Token\", \"Active\", \"Created_at\"").
		Values(idUID, userUID, token, true, sq.Expr("NOW()")).ToSql()
	command, err := r.db.Exec(ctx, query, args...)
	if command.RowsAffected() == 0 {
		return fmt.Errorf("[AuthRepo Insert] invalid argument")
	}
	return err
}

func (r *authRepo) FindByUserID(ctx context.Context, userID string) (model.Token, util.Error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, _ := psql.Select("*").From(r.tableName).Where(sq.Eq{"\"UserId\"": userID}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindByUserID] DB query error",
			Err:    qErr,
			Code:   http.StatusBadRequest,
		}
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindByUserID]",
			Err:    qErr,
			Code:   http.StatusBadRequest,
		}
	}

	if !rows.Next() {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindByUserID]",
			Code:   http.StatusNotFound,
		}
	}

	var t model.Token
	rows.Scan(&t.Token, &t.Active, &t.CreatedAt)
	return t, util.Error{}
}

func (r *authRepo) FindToken(ctx context.Context, token string) (model.Token, util.Error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, args, _ := psql.Select("*").From(r.tableName).Where(sq.Eq{"\"Token\"": token}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindToken] DB query error",
			Err:    qErr,
			Code:   http.StatusBadRequest,
		}
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindToken]",
			Err:    qErr,
			Code:   http.StatusBadRequest,
		}
	}

	if !rows.Next() {
		return model.Token{}, util.Error{
			Origin: "[AuthRepo FindToken]",
			Code:   http.StatusNotFound,
		}
	}

	var t model.Token
	rows.Scan(&t.ID, &t.UserID, &t.Token, &t.Active, &t.CreatedAt)
	return t, util.Error{}
}
