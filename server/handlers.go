package server

import (
	"backend/datamodel"
	"backend/ml"
	"backend/storage"
	"encoding/json"
	"log/slog"
	"net/http"
)

func setupHandlers() {
	http.HandleFunc("POST /providers/{$}", getProviders)
}

func getProviders(w http.ResponseWriter, r *http.Request) {

	tr := &datamodel.Transaction{}

	rawBody := make([]byte, r.ContentLength)

	_, err := r.Body.Read(rawBody)
	if err != nil {
		slog.Error("err reading request body", "err", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(rawBody, tr)
	if err != nil {
		slog.Error("err unmarshaling request body", "err", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	providers, err := storage.GetSuitableProviders(tr)
	if err != nil {
		slog.Error("err getting suitable providers", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//No suitable providers
	if len(providers) == 0 {
		slog.Debug("no suitable providers found")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	providersJson, err := json.Marshal(providers)
	if err != nil {
		slog.Error("err marshaling response body", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	providersTop, err := ml.GetProvidersTop(providersJson)
	if err != nil {
		slog.Error("err getting providers top", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(providersTop)
}
