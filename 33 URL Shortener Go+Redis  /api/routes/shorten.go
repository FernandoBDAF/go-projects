package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fbdaf/shorthnerURL-fiber-redis/database"
	"github.com/fbdaf/shorthnerURL-fiber-redis/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"customShort"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL                string        `json:"url"`
	CustomShort        string        `json:"customShort"`
	Expiry             time.Duration `json:"expiry"`
	RateLimitRemaining int           `json:"rateLimitRemaining"`
	XRateLimitReset    time.Time     `json:"xRateLimitReset"`
}

func ShortenURL(c *fiber.Ctx) error {
	// parse body
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()

	value, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valueInt, _ := strconv.Atoi(value)
		if valueInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "rate limit exceeded", "rateLimitReset": limit / time.Nanosecond / time.Minute})
		}
	}

	// check if the url is valid
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	// check if the url is already in the database
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	// create short url
	r := database.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "URL custom short is already in use"})
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to create short URL"})
	}

	// return response
	resp := response{
		URL:                body.URL,
		CustomShort:        id,
		Expiry:             body.Expiry,
		RateLimitRemaining: 10,
		XRateLimitReset:    time.Now().Add(time.Minute * 30),
	}
	r2.Decr(database.Ctx, c.IP())
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.RateLimitRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = time.Now().Add(ttl)
	resp.URL = os.Getenv("DOMAIN") + "/" + resp.CustomShort
	return c.Status(fiber.StatusOK).JSON(resp)
}
