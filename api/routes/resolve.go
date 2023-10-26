package routes

import (
	"github.com/Thibesan/ShortUrlService/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url") //Fetch Request URL
	r := database.CreateClient(0)
	defer r.Close() //After Function is executed, close rdb client

	//URL Not Found
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "shortURL not found in database off specified key",
		})
	
		//Datbase Failed to Initialize
	} else if err !=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	//Increment DB Counter for Next Request
	_ = rInr.Incr(database.Ctx, "counter")

	//Return Value of KV Pair in rDB to redirect and resolve shortened URL
	return c.Redirect(value, 301)
}