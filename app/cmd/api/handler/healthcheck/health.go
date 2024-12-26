package healthcheck

import "github.com/yaken-org/hakushi/pkg/server"

func Health(c server.Context) error {
	return c.JSON(200, "OK")
}
