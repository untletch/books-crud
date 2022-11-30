package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spf13/viper"
	"github.com/untletch/gin-books/pkg/common/db"
	"github.com/untletch/gin-books/pkg/books"
	"github.com/untletch/gin-books/auth"
)

func main() {
	viper.SetConfigFile("./pkg/common/env/.env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	proxies := viper.GetString("PROXIES")
	gin.SetMode(viper.GetString("Gin_mode"))
	dbConn := db.Init(viper.GetString("DB_URI)"))

	r := gin.Default()
	r.setTrustedProxies(strings.Split(proxies, ","))

	auth.RegisterRoutes(r, dbConn)
	books.RegisterRoutes(r, dbConn)

	r.Run(viper.GetString("PORT"))
}
