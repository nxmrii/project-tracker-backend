package web

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestUnknownRouteReturns404(t *testing.T) {

	app := fiber.New()

	RegisterHealthRoute(app)

	req, err := http.NewRequest(
		http.MethodGet,
		"/this-route-does-not-exist",
		nil,
	)

	if err != nil {
		t.Fatalf(
			"failed to create request: %v",
			err,
		)
	}

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf(
			"request failed: %v",
			err,
		)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			resp.StatusCode,
		)
	}
}
