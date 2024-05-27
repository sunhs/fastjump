package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	configFile   string
	fetchBrokers string
	fetchTopic   string
	groupid      string
	sendBrokers  string
	sendTopic    string
	targetQps    int
	appid        string
	businessType string
)

var rootCmd = &cobra.Command{
	Use:       "<executable> " + fmt.Sprintf("%s|%s|%s", kafka_utils.NormaliztionAuditMessage, kafka_utils.LikeMessage, kafka_utils.AuditDispatcherMessage),
	Short:     "Send kafka messages at certain qps.",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{string(kafka_utils.NormaliztionAuditMessage), string(kafka_utils.LikeMessage), string(kafka_utils.AuditDispatcherMessage)},
	Run: func(cmd *cobra.Command, args []string) {
		sender := kafka_utils.MessageSender{}

		if len(sendBrokers) == 0 {
			sendBrokers = fetchBrokers
		}
		if len(sendTopic) == 0 {
			sendTopic = fetchTopic
		}

		sender.Init(fetchBrokers, fetchTopic, groupid, sendBrokers, sendTopic, targetQps, kafka_utils.MessageType(args[0]), appid, businessType)

		wg := new(sync.WaitGroup)
		wg.Add(1)
		sender.Start(wg)
		wg.Wait()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to config file.")
	rootCmd.MarkFlagRequired("config")
	cobra.OnInitialize(initViper)
}

func initViper() {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(color.Red.Sprint(err))
	}

	fetchBrokers = viper.GetString("fetchBrokers")
	fetchTopic = viper.GetString("fetchTopic")
	groupid = viper.GetString("groupid")
	sendBrokers = viper.GetString("sendBrokers")
	sendTopic = viper.GetString("sendTopic")
	targetQps = viper.GetInt("targetQps")
	appid = viper.GetString("appid")
	businessType = viper.GetString("businessType")

	log.Println("Parsed configs:")
	log.Println(">>>>>>>> fetchBrokers:", fetchBrokers)
	log.Println(">>>>>>>> fetchTopic:", fetchTopic)
	log.Println(">>>>>>>> groupid:", groupid)
	log.Println(">>>>>>>> sendBrokers:", sendBrokers)
	log.Println(">>>>>>>> sendTopic:", sendTopic)
	log.Println(">>>>>>>> targetQps:", targetQps)
	log.Println(">>>>>>>> appid:", appid)
	log.Println(">>>>>>>> businessType:", businessType)
}
