package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type Databases struct {
	PrimaryDB *pg.DB
	LKDB      *pg.DB
}

type loggerHook struct {
	instance string
	logger   *zap.Logger
}

func NewPostgresDBs(primaryURL, lkURL string, schema string, logger *zap.Logger) (*Databases, error) {
	primaryDB, err := newDB(primaryURL, "")
	if err != nil {
		return nil, err
	}

	primaryDB.AddQueryHook(loggerHook{"primary DB", logger})

	lkDB, err := newDB(lkURL, schema)
	if err != nil {
		return nil, err
	}

	lkDB.AddQueryHook(loggerHook{"lk DB", logger})

	return &Databases{
		PrimaryDB: primaryDB,
		LKDB:      lkDB,
	}, nil
}

func (dbs *Databases) Close() error {
	if err := dbs.PrimaryDB.Close(); err != nil {
		return err
	}

	if err := dbs.LKDB.Close(); err != nil {
		return err
	}

	return nil
}

func newDB(url string, schema string) (*pg.DB, error) {
	opts, err := pg.ParseURL(url)
	if err != nil {
		return nil, err
	}

	if schema != "" {
		opts.OnConnect = func(ctx context.Context, cn *pg.Conn) error {
			_, _ = cn.Exec("set search_path=?", schema)
			return nil
		}
	}

	db := pg.Connect(opts)

	var n int
	_, err = db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return nil, fmt.Errorf("cant connect database %v, %v", url, err.Error())
	}

	return db, nil
}

func (h loggerHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	query, err := q.FormattedQuery()

	if err == nil {
		// h.logger.Debug("PostgresDB", zap.ByteString("query", query))
		fmt.Println(string(query))
	}

	return c, nil
}

func (h loggerHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}
