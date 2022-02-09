package cmd

import (
	"fmt"
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"github.com/spf13/cobra"
	"github.com/tyrm/gemini-forum/config"
	"github.com/tyrm/gemini-forum/db/postgres"
	"github.com/tyrm/gemini-forum/gemini"
	"github.com/tyrm/gemini-forum/kv/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	rootCmd.AddCommand(geminiCmd)
}

var geminiCmd = &cobra.Command{
	Use:   "gemini",
	Short: "Run the gemini server",
	//TODO Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		requiredVars := []string{
			"POSTGRES_DSN",
			"REDIS_ADDRESS",
		}
		c, err := config.CollectConfig(requiredVars)
		if err != nil {
			log.Fatalf("error gathering configuration: %s", err.Error())
			return
		}

		// start logger
		err = loggo.ConfigureLoggers(c.LoggerConfig)
		if err != nil {
			log.Fatalf("error configuring logger: %s", err.Error())
			return
		}
		_, err = loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
		if err != nil {
			log.Fatalf("error configuring color logger: %s", err.Error())
			return
		}

		logger := loggo.GetLogger("main")
		logger.Infof("starting main process")

		logger.Debugf("creating db client")
		dbClient, err := postgres.NewClient(c)
		if err != nil {
			logger.Errorf("db client: %s", err.Error())
			return
		}

		logger.Debugf("creating kv client")
		kvClient, err := redis.NewClient(c.RedisAddress, c.RedisDB, c.RedisPassword)
		if err != nil {
			logger.Errorf("db client: %s", err.Error())
			return
		}

		logger.Debugf("creating gemini server")
		server, err := gemini.NewServer(c, dbClient, kvClient)
		if err != nil {
			logger.Errorf("gemini server: %s", err.Error())
			return
		}

		// ** start application **
		errChan := make(chan error)

		// start web server
		logger.Infof("starting gemini server")
		go func(errChan chan error) {
			err := server.Run()
			if err != nil {
				errChan <- fmt.Errorf("gemini: %s", err.Error())
			}
		}(errChan)

		// Wait for SIGINT and SIGTERM (HIT CTRL-C)
		nch := make(chan os.Signal)
		signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-nch:
			logger.Infof("got sig: %s", sig)
		case err := <-errChan:
			logger.Criticalf(err.Error())
		}

		logger.Infof("main process done")
	},
}
