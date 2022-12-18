package server

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	mw_logger "github.com/gofiber/fiber/v2/middleware/logger"
	cm "main/pkg/utils/common"
	"os"
)

const (
	ctxTimeout = 3
)

func (s *server) NewHttpServer() (server *fiber.App, err error) {
	server = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ServerHeader:  os.Getenv("SERVER_HEADER"),
		AppName:       os.Getenv("APP_TITLE") + " " + os.Getenv("APP_VERSION"),
	})

	server.Use(cors.New())
	server.Use(mw_logger.New())

	server.Get("/doc/*", swagger.HandlerDefault)

	server.Get("/", func(c *fiber.Ctx) error {
		s.logger.Infof("Health check RequestID: %d", cm.GenNum())
		return c.SendString("Everything is OK ! ;)")
	})

	go func() {
		URI := fmt.Sprintf("%s:%s", "", s.cfg.Http.PORT)
		if s.cfg.Server.APP_ENV == "production" {
			if err := server.ListenTLS(URI, s.cfg.Http.SSL_CERT_PATH, s.cfg.Http.SSL_CERT_KEY); err != nil {
				s.logger.Fatalf("Error starting Server with SSL : ", err)
			}
		} else {
			if err = server.Listen(URI); err != nil {
				s.logger.Fatalf("Error starting Server : ", err)
			}
		}

		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Http.PORT)
	}()

	s.logger.Infof("Server is listening on PORT: %s", s.cfg.Http.PORT)

	return
}
