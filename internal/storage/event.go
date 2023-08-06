package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
)

const (
	EventTableName            = "collector.t_event"
	EventTableTimestampColumn = "ts"
	EventTableSourceColumn    = "source"
	EventTableDataColumn      = "data"
)

type EventStorage interface {
	InsertEvents(ctx context.Context, events []Event) error
}

type EventStorageImpl struct {
	Storage
}

var _ EventStorage = (*EventStorageImpl)(nil)

func NewStorage(config Config) (*EventStorageImpl, error) {
	password, ok := os.LookupEnv(config.PasswordEnv)
	if !ok {
		return nil, fmt.Errorf("no DB password found at env %s", config.PasswordEnv)
	}
	config.Password = password

	s := Storage{
		Config: config,
	}

	return &EventStorageImpl{
		Storage: s,
	}, nil
}

func (s *EventStorageImpl) InsertEvents(ctx context.Context, events []Event) (err error) {
	if err := s.Connect(); err != nil {
		return err
	}
	defer func() {
		err = s.Disconnect()
	}()

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, event := range events {
		sqQuery := squirrel.Insert(EventTableName).
			Columns(EventTableTimestampColumn, EventTableSourceColumn, EventTableDataColumn).
			Values(event.Timestamp, event.Source, event.Data).
			Suffix(
				fmt.Sprintf(
					"ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s, %s = EXCLUDED.%s",
					EventTableTimestampColumn,
					EventTableSourceColumn,
					EventTableSourceColumn,
					EventTableDataColumn,
					EventTableDataColumn,
				),
			).
			PlaceholderFormat(squirrel.Dollar)

		query, args, err := sqQuery.ToSql()
		if err != nil {
			return err
		}

		rows, err := tx.Query(query, args...)
		if err != nil {
			return err
		}
		if err := rows.Err(); err != nil {
			return err
		}
		if err := rows.Close(); err != nil {
			return err
		}
	}

	return tx.Commit()
}
