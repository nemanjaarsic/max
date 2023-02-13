package repo

import (
	"context"
	"fmt"
	"max-db-svc/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type transactionRepo struct {
	db        *pgx.Conn
	tableName string
}

var _ TransactionRepository = (*transactionRepo)(nil)

func NewTransactionRepo() *transactionRepo {
	return &transactionRepo{
		db:        NewDatabaseConnection(),
		tableName: "\"Transactions\"",
	}
}

func (r *transactionRepo) Insert(ctx context.Context, tx model.Transaction) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	idUID, uErr := uuid.NewRandom()
	if uErr != nil {
		return fmt.Errorf("[TransactionRepo Insert] %vv", uErr)
	}
	userUID, uErr := uuid.Parse(tx.UserID)
	if uErr != nil {
		return fmt.Errorf("[TransactionRepo Insert] %vv", uErr)
	}
	query, args, _ := psql.Insert(r.tableName).
		Columns("\"Id\", \"UserId\", \"OperationType\", \"Amount\", \"Timestamp\"").
		Values(idUID, userUID, tx.OperationType, tx.Amount, sq.Expr("NOW()")).ToSql()
	command, err := r.db.Exec(ctx, query, args...)
	if command.RowsAffected() == 0 {
		return fmt.Errorf("[TransactionRepo Insert] invalid argument")
	}
	return err
}
