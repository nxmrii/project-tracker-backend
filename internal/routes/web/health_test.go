package web

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthEndpoint(t *testing.T) {

	app := fiber.New()

	RegisterHealthRoute(app)

	req, err := http.NewRequest(
		http.MethodGet,
		"/health",
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

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.StatusCode,
		)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf(
			"failed to read response: %v",
			err,
		)
	}

	var result map[string]string

	if err := json.Unmarshal(
		body,
		&result,
	); err != nil {
		t.Fatalf(
			"invalid response JSON: %v",
			err,
		)
	}

	if result["status"] != "ok" {
		t.Fatalf(
			"expected status ok, got %s",
			result["status"],
		)
	}
}
