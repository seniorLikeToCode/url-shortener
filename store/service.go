package store

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server to ensure the connection is established
	ping, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error init Redis: %v", err)
	}

	log.Printf("\n Redis started successfully: ping message = {%s}", ping)
	storeService.redisClient = redisClient
	return storeService
}

/* 	We want to be able to save the mapping between the originalUrl
and the generated shorturl */
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		log.Fatalf("Failed saving key url | Error : %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
	}
}

/*	We should be able to retrieve the initial long URL once the short
	is Provided. This is when users will be calling the shortLink in the
	url, so what we need to do here is to retrieve the long url and think
	about redirect. */
func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		log.Printf("URL not found | shortUrl: %s\n", shortUrl)
		return "", nil
	} else if err != nil {
		log.Fatalf("Failed to retrieve initial URL | Error: %v - shortUrl: %s\n", err, shortUrl)
		return "", err
	}
	return result, nil
}
