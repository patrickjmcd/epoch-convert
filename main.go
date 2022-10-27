package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	utc bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&utc, "utc", "u", false, "display time in UTC")
}

var rootCmd = &cobra.Command{
	Use:   "epoch-convert",
	Short: "Convert epoch time into human readable time",
	Long:  `Convert epoch time into human readable time, either in the current locale time zone or UTC`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var epoch int64
		var err error
		if len(args) == 0 {
			epoch = time.Now().Unix()
			fmt.Printf("Current epoch time: %d\n", epoch)
		} else {
			epoch, err = strconv.ParseInt(args[0], 10, 64)
			if epoch > 9999999999 {
				epoch = epoch / 1000
			}
		}
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse epoch time")
			return
		}
		t := time.Unix(epoch, 0)
		if utc {
			t = t.UTC()
		}
		fmt.Println(t.Format(time.RFC3339))

	},
}

// execute adds all child commands to the root command sets flags appropriately.
func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	execute()
}
