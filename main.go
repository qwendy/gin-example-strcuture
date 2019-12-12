package main

import (
	"gin-example-structure/config"
	"gin-example-structure/db"
	"gin-example-structure/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

var versionCmd = &cobra.Command{
	Use:   "dev",
	Short: "develop环境 使用config-dev.toml配置文件",
	Long:  `develop环境 使用config-dev.toml配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Init("config-dev.toml")
	},
}

func init() {
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.toml", "config file (default is config.toml)")
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
	config.Init(cfgFile)
}
func main() {
	db.Init()
	server.Init()
}
