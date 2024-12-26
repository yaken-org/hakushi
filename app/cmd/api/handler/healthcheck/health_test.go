package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/yaken-org/hakushi/pkg/server"
)

func TestHealth(t *testing.T) {
	s := server.New()
	e := s.Echo()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Health(server.NewContext(c)); err != nil {
		t.Fatalf("Health() failed: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("Health() failed: %v", rec.Code)
	}

	if rec.Body.String() != "\"OK\"\n" {
		t.Fatalf("Health() failed: %v", rec.Body.String())
	}
}
