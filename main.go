package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type Reaction struct {
	Reaction string `json:"reaction"`
	URL      string `json:"url"`
}

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // use environment variable
		Password: os.Getenv("REDIS_PASSWORD"), // use environment variable
		DB:       0,                           // use default DB
	})

	allowedUrls := strings.Split(os.Getenv("ALLOWED_URLS"), ",") // parse allowed urls

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		url := c.Query("url")
		log.Println(url)
		result, _ := rdb.HGetAll(ctx, url).Result()

		reactions := make(map[string]int)
		for k, v := range result {
			count, _ := strconv.Atoi(v)
			reactions[k] = count
		}

		c.JSON(200, reactions)
	})

	r.POST("/", func(c *gin.Context) {
		var reaction Reaction
		if err := c.ShouldBindJSON(&reaction); err == nil {
			if reaction.Reaction == "" || reaction.URL == "" {
				c.JSON(400, gin.H{"status": "Bad request. Missing `Reaction` or `URL` parameters."})
				return
			}

			isAllowed := false
			for _, allowedUrl := range allowedUrls {
				if strings.Contains(reaction.URL, allowedUrl) {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				c.JSON(400, gin.H{"status": "Bad request. Invalid URL."})
				return
			}

			rdb.HIncrBy(ctx, reaction.URL, reaction.Reaction, 1)
			c.JSON(201, gin.H{"status": "created"})
		} else {
			c.JSON(400, gin.H{"status": "Bad request"})
		}
	})

	r.Run(":8080")
}
