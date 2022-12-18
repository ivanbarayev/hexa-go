package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"main/config"
	"time"
)

var (
	uri string
)

func NewEngineServer(cfg *config.Config) (*nats.EncodedConn, error) {
	uri = fmt.Sprintf("%s:%s", cfg.Nats.SERVER_HOST, cfg.Nats.SERVER_PORT)
	if cfg.Nats.SERVER_HOST == "" || cfg.Nats.SERVER_PORT == "" {
		uri = nats.DefaultURL
	}
	// Connect Options.
	opts := []nats.Option{nats.Name("Data Subscriber")}
	opts = setupConnOptions(opts)

	nc, _ := nats.Connect(uri)

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	//defer ec.Close()

	return ec, err
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
