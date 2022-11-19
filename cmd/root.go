package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	text = `Alien Invasion`
)

// Command is the structure of cli
type Command struct {
	rootCmd *cobra.Command
}

// NewCommand the command line boot-loader
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "alien-invasion",
		Short: "Alien Invasion command line",
		Long:  "Alien Invasion command line",
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

// Start the all command line
func (c *Command) Start() {
	var rootCommands = []*cobra.Command{
		{
			Use:   "start",
			Short: "Start Alien Invasion",
			Long:  "Start Alien Invasion",
			PreRun: func(cmd *cobra.Command, args []string) {
				fmt.Println(text) // Show display text
			},
			Run: func(cmd *cobra.Command, args []string) {

			},
			PostRun: func(cmd *cobra.Command, args []string) {

			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.AddCommand(command)
	}

	err := c.rootCmd.Execute()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
