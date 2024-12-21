package storage

import (
	"backend/datamodel"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var (
	db             *sql.DB
	ErrEmptyConfig = errors.New("empty config")
	ErrTimeout     = errors.New("db ping timeout")
)

type Config struct {
	creds      string
	retryCount int
}

func GetConfig() (*Config, error) {

	cfg := &Config{}

	credsEnv := os.Getenv("DB_CREDENTIALS")
	if credsEnv != "" {
		return nil, ErrEmptyConfig
	}
	cfg.creds = credsEnv

	retryCountEnv := os.Getenv("DB_RETRIES_COUNT")
	if retryCountEnv != "" {
		retryCount, err := strconv.ParseInt(retryCountEnv, 10, 0)
		if err != nil {
			return nil, err
		}
		cfg.retryCount = int(retryCount)
	} else {
		cfg.retryCount = 5
	}

	return cfg, nil
}

func Connect(cfg *Config) error {
	if db != nil {
		return nil
	}

	var err error
	db, err = sql.Open("postgres", cfg.creds)
	if err != nil {
		return err
	}

	err = db.Ping()
	retry := cfg.retryCount
	for ; retry > 0 && err != nil; retry-- {
		err = db.Ping()
		time.Sleep(time.Second)
	}

	if err != nil {
		return errors.Join(ErrTimeout, err)
	}

	return nil
}

func GetSuitableProviders(tr *datamodel.Transaction) ([]datamodel.Provider, error) {

	rows, err := db.Query(
		`SELECT * FROM providers p
		WHERE p.min_sum <= $1 AND p.max_sum >= $1 AND p.currency = $2
		`,
		tr.Amount, tr.Cur)

	if err != nil {
		slog.Error("err getting suitable providers", "err", err.Error())
		return nil, err
	}

	providers := make([]datamodel.Provider, 0, 20)

	for rows.Next() {
		var p datamodel.Provider
		err := rows.Scan(&p.Id, &p.Conversion, &p.AvgTime, &p.MinSum, &p.MaxSum, &p.LimitMax, &p.LimitMin, &p.LimitByCard, &p.Commission, &p.Currency)
		if err != nil {
			slog.Error("err scanning provider", "err", err.Error())
			return nil, err
		}
		providers = append(providers, p)
	}
	slog.Debug("got suitable providers", "count", len(providers))

	return providers, nil
}
