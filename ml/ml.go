package ml

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

var (
	mlAddress            string
	ErrMladdressNotFound = errors.New("ml address not found")
)

type Config struct {
	address string
}

func GetConfig() (*Config, error) {

	cfg := &Config{}

	mladdrEnv := os.Getenv("ML_ADDRESS")
	if mladdrEnv != "" {
		cfg.address = mladdrEnv
	} else {
		return nil, ErrMladdressNotFound
	}

	return cfg, nil
}

func SetUp(cfg *Config) error {
	slog.Debug("Starting ml...")

	mlAddress = "http://" + cfg.address

	// _, err := http.Get(mlAddress)
	// times := retry
	// for err != nil && times > 0 {
	// 	_, err = http.Get(mlAddress)
	// 	times--
	// 	time.After(time.Second)
	// }
	// if err != nil {
	// 	return err
	// }

	slog.Debug("ml started...")

	return nil
}

func GetProvidersTop(rawJson []byte) ([]byte, error) {
	slog.Debug("Getting providers top...")

	recom, err := GetProvidersTopFromModel(rawJson)
	if err != nil {
		slog.Error("err getting providers top from model", "err", err.Error())
		return nil, err
	}

	slog.Debug("recomendation", "recom", recom)

	return recom, nil
}

func GetProvidersTopFromModel(rawJson []byte) ([]byte, error) {
	slog.Debug("Getting providers top from model...")

	buf := bytes.NewReader(rawJson)

	slog.Debug("request for model", "mlAddress", mlAddress, "buf", string(rawJson))
	res, err := http.Post(mlAddress, "application/json", buf)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	recom := make([]byte, res.ContentLength)
	_, err = res.Body.Read(recom)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	slog.Debug("json from model", "json", string(recom))

	return recom, nil
}
