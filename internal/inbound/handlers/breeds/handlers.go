package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/internal/ports"
)

type genericFunc[T any] func(context.Context, T) error

func Handle[T any](fn genericFunc[T], logger *charmLog.Logger) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var input T

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Error when calling generic handler", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, something went wrong"))
			return
		}

		if err := fn(r.Context(), input); err != nil {
			logger.Error("Error when calling generic handler", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, something went wrong"))
			return
		}

		marshaled, err := json.Marshal(input)
		if err != nil {
			logger.Error("Error when calling generic handler", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, something went wrong"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(marshaled)

	}
}

func HandleGet(repo ports.BreedsCRUD, logger *charmLog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if len(vars) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, you didn't provided any information about your pet"))
			return
		}

		for key, value := range vars {

			resp, err := repo.Read(r.Context(), key, value)
			if err != nil {
				logger.Error("Error when calling get handler", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Sorry, something went wrong"))
				return
			}

			marshaled, err := json.Marshal(resp)
			if err != nil {
				logger.Error("Error when calling get handler", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Sorry, something went wrong"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(marshaled)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, we couldn't find any pet with the information you provided"))

	}
}

func HandleDelete(repo ports.BreedsCRUD, logger *charmLog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if len(vars) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, you didn't provided any information about your pet"))
			return
		}

		convertedId, err := strconv.Atoi(vars["id"])
		if err != nil {
			logger.Error("Error when calling get handler", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, something went wrong"))
			return
		}

		err = repo.Delete(r.Context(), convertedId)
		if err != nil {
			logger.Error("Error when calling get handler", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sorry, something went wrong"))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
