package cmd

import (
	"context"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/pipego/exporter/config"
)

var (
	app = kingpin.New("exporter", "pipego exporter").Version(config.Version + "-build-" + config.Build)
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	return nil
}
