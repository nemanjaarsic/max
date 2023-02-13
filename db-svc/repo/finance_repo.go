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

type financeRepo struct {
	db        *pgx.Conn
	tableName string
}

// Check interface complience
var _ FinanceRepository = (*financeRepo)(nil)

func NewFinanceRepo() *financeRepo {
	return &financeRepo{
		db:        NewDatabaseConnection(),
		tableName: "\"Financies\"",
	}
}

func (r *financeRepo) GetBalanceByUserID(ctx context.Context, userID string) (float32, util.Error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	uid, uErr := uuid.Parse(userID)
	if uErr != nil {
		return -1, util.Error{
			Origin: "[FinanceRepo GetBalanceByUserID]",
			Err:    uErr,
			Code:   http.StatusBadRequest,
		}
	}
	query, args, _ := psql.Select("*").From(r.tableName).Where(sq.Eq{"\"UserId\"": uid}).ToSql()
	rows, qErr := r.db.Query(ctx, query, args...)
	if qErr != nil {
		return -1, util.Error{
			Origin: "[FinanceRepo GetBalanceByUserID]",
			Err:    uErr,
			Code:   http.StatusBadRequest,
		}
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return -1, util.Error{
			Origin: "[FinanceRepo GetBalanceByUserID]",
			Err:    uErr,
			Code:   http.StatusBadRequest,
		}
	}

	if !rows.Next() {
		return -1, util.Error{
			Origin: "[FinanceRepo GetBalanceByUserID]",
			Err:    uErr,
			Code:   http.StatusMethodNotAllowed,
		}
	}

	var f model.Finance
	if err := rows.Scan(&f.ID, &f.UserID, &f.Balance); err != nil {
		return -1, util.Error{
			Origin: "[FinanceRepo GetBalanceByUserID]",
			Err:    uErr,
			Code:   http.StatusInternalServerError,
		}
	}
	return f.Balance, util.Error{}
}

func (r *financeRepo) Insert(ctx context.Context, balance model.Finance) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	idUID, uErr := uuid.NewRandom()
	if uErr != nil {
		return fmt.Errorf("[FinanciesRepo Insert] %vv", uErr)
	}
	userUID, uErr := uuid.Parse(balance.UserID)
	if uErr != nil {
		return fmt.Errorf("[FinanciesRepo Insert] %vv", uErr)
	}
	query, args, _ := psql.Insert(r.tableName).Columns("\"Id\", \"UserId\", \"Balance\"").Values(idUID, userUID, balance.Balance).ToSql()
	command, err := r.db.Exec(ctx, query, args...)
	if command.RowsAffected() == 0 {
		return fmt.Errorf("[FinanciesRepo Insert] Invalid argument")
	}
	return err
}

func (r *financeRepo) Update(ctx context.Context, balance model.Finance) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	userUID, uErr := uuid.Parse(balance.UserID)
	if uErr != nil {
		return fmt.Errorf("[FinanciesRepo Update] %vv", uErr)
	}
	query, args, _ := psql.Update(r.tableName).Set("\"Balance\"", balance.Balance).Where(sq.Eq{"\"UserId\"": userUID}).ToSql()
	command, err := r.db.Exec(ctx, query, args...)
	if command.RowsAffected() == 0 {
		return fmt.Errorf("[FinanciesRepo Update] Invalid argument")
	}

	return err
}
