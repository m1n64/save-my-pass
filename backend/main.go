package main

import (
	actions3 "backend/modules/categories/actions"
	actions2 "backend/modules/users/actions"
	"backend/modules/users/middlewares"
	"backend/services"
	"backend/system/actions"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"runtime"
	"time"
)

func main() {
	services.InitDBConnection()
	services.Migrations()

	r := gin.Default()
	routes(r)

	fmt.Println("Server started on 80 port")
	r.Run(":80")
}

// @title Save My Pass - API
// @version 1.0
// @description This is a Save My Pass API.

// @host localhost
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func routes(r *gin.Engine) {
	r.GET("/ping", actions.Ping)

	user := r.Group("/user")
	{
		user.GET("", middlewares.AuthMiddleware(), actions2.GetUser)
		user.POST("/register", actions2.UserRegister)
		user.POST("/login", actions2.UserLogin)
	}

	authEndpoints := r.Group("/", middlewares.AuthMiddleware())
	category := authEndpoints.Group("/category")
	{
		category.GET("/all", actions3.GetCategories)
		category.POST("/create", actions3.CreateCategory)
		category.PUT("/update/:id", actions3.UpdateCategory)
		category.DELETE("/delete/:id", actions3.DeleteCategory)
	}

	swaggerURL := ginSwagger.URL("http://localhost/api/docs/swagger.json")
	r.GET("/api/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	prometheusInit(r)
}

func prometheusInit(r *gin.Engine) {
	myMetric := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "backend_api",
		Help: "API Metrics",
	})

	ramUsage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "backend_api_ram_usage_bytes",
		Help: "RAM usage of the API services",
	})

	cpuUsage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "backend_api_cpu_usage_percent",
		Help: "Current CPU usage percent of API Service",
	})

	prometheus.MustRegister(myMetric, ramUsage, cpuUsage)

	go func() {
		for {
			percent, err := cpu.Percent(time.Second, false)
			if err != nil {
				log.Println("Ошибка при получении процента использования ЦПУ:", err)
				continue
			}
			cpuUsage.Set(percent[0])
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			ramUsage.Set(float64(m.Alloc) / 1024 / 1024)
			time.Sleep(10 * time.Second)
		}
	}()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
