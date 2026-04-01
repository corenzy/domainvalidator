package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/corenzy/domainvalidator/models"
	"github.com/corenzy/domainvalidator/services"
)

// HandleLookup processes POST /lookup requests.
// It expects a JSON body with a "domain" field.
func HandleLookup(c *fiber.Ctx) error {
	var req models.LookupRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid request body. Expected JSON with a 'domain' field.",
		})
	}

	if req.Domain == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Success: false,
			Message: "The 'domain' field is required.",
		})
	}

	start := time.Now()
	result, err := services.Lookup(req.Domain)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Success:      false,
			Message:      err.Error(),
			LookupTimeMs: elapsed,
		})
	}

	return c.JSON(models.APIResponse{
		Success:      true,
		Message:      fmt.Sprintf("DNS lookup completed for %s", result.Domain),
		LookupTimeMs: elapsed,
		Result:       result,
	})
}
