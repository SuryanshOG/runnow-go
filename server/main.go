package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	PistonAPIURL     string
	RateLimitRPS     float64
	RateLimitBurst   int
	MaxCodeLength    int
	AllowedLanguages []string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func loadConfig() Config {
	_ = godotenv.Load()

	port := getEnv("PORT", "8080")
	rps, _ := strconv.ParseFloat(getEnv("RATE_LIMIT_RPS", "5"), 64)
	burst, _ := strconv.Atoi(getEnv("RATE_LIMIT_BURST", "10"))
	maxLen, _ := strconv.Atoi(getEnv("MAX_CODE_LENGTH", "10000"))
	langs := strings.Split(getEnv("ALLOWED_LANGUAGES", "python,javascript,typescript,cpp,c,go,rust,java,ruby"), ",")

	trimmedLangs := make([]string, 0, len(langs))
	for _, lang := range langs {
		if t := strings.TrimSpace(lang); t != "" {
			trimmedLangs = append(trimmedLangs, t)
		}
	}

	return Config{
		Port:             port,
		PistonAPIURL:     getEnv("PISTON_API_URL", "http://localhost:2000"),
		RateLimitRPS:     rps,
		RateLimitBurst:   burst,
		MaxCodeLength:    maxLen,
		AllowedLanguages: trimmedLangs,
	}
}

func setupRouter(cfg Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	api := r.Group("/api")
	{
		api.POST("/execute", func(c *gin.Context) {
			c.JSON(501, gin.H{"error": "not implemented yet"})
		})
	}
	return r
}

func main() {
	cfg := loadConfig()

	log.Printf("Config: Port=%s, PistonURL=%s, RateLimitRPS=%.1f, RateLimitBurst=%d, MaxCodeLength=%d, Languages=%v",
		cfg.Port, cfg.PistonAPIURL, cfg.RateLimitRPS, cfg.RateLimitBurst, cfg.MaxCodeLength, cfg.AllowedLanguages)

	router := setupRouter(cfg)

	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
