package routes

import (
	"time"
)


type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_remaining"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	
	if err := c.BodyParser(&body); err!= nil {
		return c.Status(fiber.StatusBadRequest)
						.JSON(fiber.map{"error":"cannot parse JSON"})
	}

	//Rate Limiting
	//Determine User Session by Unique Identifier (IP Address)
	//Decrement RateRemaining
	//Output Reset Timer for Refreshes
	//Proceed with remainder of function call for returning response w/ URL

	//Validate User Input

	//Check if URL is valid
	if !govalidator.IsURL(body.URL){
		return c.status(fiber.StatusBadRequest)
						.JSON(fiber.map{"error":"Invalid URL"})
	}

	//Check LocalHost domain error (inf loop if source URL is modified)
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable)
						.JSON(fiber.map{"error":"Requested Resource Cannot Be Accessed :)"})
	}

	//Enforce HTTP, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}
