package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/database_actions"
	"github.com/japhy-tech/backend-test/internal/daemon"
	database "github.com/japhy-tech/backend-test/internal/outbound/database/breeds"
	"github.com/japhy-tech/backend-test/internal/usecase/breeds"
	"github.com/japhy-tech/backend-test/internal/utils/converter"
)

const (
	MysqlDSN = "root:root@(mysql-test:3306)/core?parseTime=true"
	ApiPort  = "5000"
)

func main() {
	logger := charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	err := database_actions.InitMigrator(MysqlDSN)
	if err != nil {
		logger.Fatal(err.Error())
	}

	msg, err := database_actions.RunMigrate("up", 0)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(msg)
	}

	db, err := sql.Open("mysql", MysqlDSN)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	db.SetMaxIdleConns(0)

	err = db.Ping()
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	logger.Info("Database connected")

	breedsRepo := database.NewBreeds(db)

	breedsCRUD := breeds.NewBreeds(breedsRepo)

	converter.NewConverter(breedsRepo).ConvertCSVToSQL(context.Background(), "breeds.csv")

	app := daemon.NewApp(logger, breedsCRUD)

	r := mux.NewRouter()
	app.RegisterRoutes(r.PathPrefix("/v1").Subrouter(), logger)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	err = http.ListenAndServe(
		net.JoinHostPort("", ApiPort),
		r,
	)

	// =============================== Starting Msg ===============================
	logger.Info(fmt.Sprintf("Service started and listen on port %s", ApiPort))
}
