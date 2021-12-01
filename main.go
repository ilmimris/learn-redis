package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	redisKey := fmt.Sprintf("%s:%s", "user", id)
	val, err := cache.Get(ctx, redisKey).Bytes()
	if err != nil {
		return c.Next()
	}

	data := toJson(val)
	return c.JSON(fiber.Map{"Cached": data})
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is working 👊")
	})

	// get by id
	app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/users/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		redisKey := fmt.Sprintf("%s:%s", "user", id)
		cacheErr := cache.Set(ctx, redisKey, body, 10*time.Second).Err()
		if cacheErr != nil {
			return cacheErr
		}

		data := toJson(body)
		return c.JSON(fiber.Map{"Data": data})
	})

	app.Listen(":3000")
}
