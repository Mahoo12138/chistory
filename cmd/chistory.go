package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	greetingBanner = `
	▄████████    ▄█    █▄     ▄█     ▄████████     ███      ▄██████▄     ▄████████ ▄██   ▄   
	███    ███   ███    ███   ███    ███    ███ ▀█████████▄ ███    ███   ███    ███ ███   ██▄ 
	███    █▀    ███    ███   ███▌   ███    █▀     ▀███▀▀██ ███    ███   ███    ███ ███▄▄▄███ 
	███         ▄███▄▄▄▄███▄▄ ███▌   ███            ███   ▀ ███    ███  ▄███▄▄▄▄██▀ ▀▀▀▀▀▀███ 
	███        ▀▀███▀▀▀▀███▀  ███▌ ▀███████████     ███     ███    ███ ▀▀███▀▀▀▀▀   ▄██   ███ 
	███    █▄    ███    ███   ███           ███     ███     ███    ███ ▀███████████ ███   ███ 
	███    ███   ███    ███   ███     ▄█    ███     ███     ███    ███   ███    ███ ███   ███ 
	████████▀    ███    █▀    █▀    ▄████████▀     ▄████▀    ▀██████▀    ███    ███  ▀█████▀         
`
)

var (
	mode   string
	addr   string
	port   int
	data   string
	driver string
	dsn    string

	rootCmd = &cobra.Command{
		Use:   "chistory",
		Short: `An open-source, self-hosted chat history manager.`,
		Run: func(_cmd *cobra.Command, _args []string) {
			ctx, cancel := context.WithCancel(context.Background())

			c := make(chan os.Signal, 1)

			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				sig := <-c
				fmt.Sprintf("%s received.\n", sig.String())
				cancel()
			}()
			print(mode, addr, port, driver)
			printGreetings()
			<-ctx.Done()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "demo", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().StringVarP(&addr, "addr", "a", "", "address of server")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "port of server")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data directory")
	rootCmd.PersistentFlags().StringVarP(&driver, "driver", "", "", "database driver")
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "", "", "database source name(aka. DSN)")

	err := viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	if err != nil {
		panic(err)
	}

	viper.SetDefault("mode", "demo")
	viper.SetDefault("driver", "sqlite")
	viper.SetDefault("addr", "")
	viper.SetDefault("port", 8081)
	viper.SetEnvPrefix("chistory")
}

func initConfig() {
	viper.AutomaticEnv()
	print("initial over")
	var err error
	// profile, err = _profile.GetProfile()
	if err != nil {
		fmt.Printf("failed to get profile, error: %+v\n", err)
		return
	}
}
func printGreetings() {
	print(greetingBanner)
}

func Execute() error {
	return rootCmd.Execute()
}
