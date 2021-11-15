package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "gorm.io/driver/postgres"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"service_catalog/models"
	"service_catalog/web"
	"sync"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, echo.HeaderOrigin, echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	// dsn := "host=localhost port=5432 user=postgres dbname=service_catalog password=postgres sslmode=disable "
	db, err := gorm.Open(sqlite.Open("service_catalog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	db.AutoMigrate(models.Service{}, models.Version{})
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	sc := web.NewServiceCatalogSvc(db)
	e.GET("/services/:id", sc.FindById(ctx))
	e.GET("/services/:id/versions", sc.FindVersions(ctx))
	e.GET("/services", sc.List(ctx))
	go func() {
		defer wg.Done()
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	signal.Notify(quit, os.Interrupt)
	<-quit
	cancel()
	sctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	log.Info("Shutting down gateway")
	if err := e.Shutdown(sctx); err != nil {
		log.Fatal(err)
	}
}
