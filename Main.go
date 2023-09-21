package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", settings)
	app.Post("/beeload/compare/release", PostCompareRelease)
	app.Post("/beeload/create/bucket", PostCreateBucket)
	app.Get("/beeload/get/tabledatacurrenttests", GetTableDataCurrentTests)
	app.Get("/beeload/get/bucketList", GetBucketList)
	app.Get("/beeload/get/projectList", GetProjectList)
	app.Get("/beeload/get/tableDataReports", GetTableDataReports)
	app.Get("/beeload/get/tableDataTests", GetTableDataTests)
	app.Get("/beeload/get/tableDataStatus", GetTableDataStatus)
	app.Get("/beeload/get/versionList", GetVersionsList)
}

func main() {
	// Чтение конфигурации из файла
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Error reading config file:", err)
		return
	} else {
		logrus.Debug("Config file: readed sucessfully.")
		//fmt.Println("Config file: readed sucessfully.")
	}
	fmt.Println("Debug mode: ", viper.GetBool("server.debug"))
	fmt.Println("Log level: ", viper.GetString("server.log_level"))
	InitLogger(viper.GetBool("server.debug"), viper.GetString("server.log_level"))
	// Настройка Fiber
	app := fiber.New()
	setupRoutes(app)

	// Запуск сервера в горутине
	go func() {
		port := viper.GetInt("server.port")
		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			logrus.Error("Error starting server:", err)
			fmt.Println("Error starting server:", err)
		}
	}()

	// Обработка сигналов для завершения программы
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	logrus.Error("Shutting down...")
	fmt.Println("Shutting down...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		logrus.Error(err)
		fmt.Println("Error shutting down server:", err)
	}
	//TODO: АНТОН НЕ ТРОГАЙ КОД!
}
