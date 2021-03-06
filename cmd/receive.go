package cmd

import (
	"github.com/derricw/siggo/model"
	"github.com/derricw/siggo/signal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "receive all outstanding messages",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := model.GetConfig()
		if err != nil {
			log.Fatalf("failed to read config @ %s", model.DefaultConfigPath())
		}
		initLogging(cfg)

		if cfg.UserNumber == "" {
			log.Fatalf("no user phone number configured @ %s", model.DefaultConfigPath())
		}

		var signalAPI model.SignalAPI = signal.NewSignal(cfg.UserNumber)
		if mock != "" {
			signalAPI = setupMock(mock, cfg)
		}

		s := model.NewSiggo(signalAPI, cfg)

		s.NewInfo = func(conv *model.Conversation) {
			log.Printf("From: %v | Conv: \n%s", conv.Contact, conv.String())
		}
		s.ReceiveForever()
		<-make(chan struct{})
	},
}
