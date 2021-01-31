package main

import (
	"log"
	"os"

	commands "github.com/Sho372/emote-clone/commands"
	"github.com/spf13/cobra"
)

func main() {
	app, err := commands.New()
	if err != nil {
		log.Fatal(err)
	}
	cmdEmote := buildEmoteCommand(app)     // cmdEmote as root command
	cmdVersion := buildVersionCommand(app) // cmdVersion as sub command
	cmdEmote.AddCommand(cmdVersion)
	if err := cmdEmote.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildEmoteCommand(app *commands.App) *cobra.Command {
	var dest string

	cmdEmote := &cobra.Command{
		Use: "emote",
		PreRun: func(cmd *cobra.Command, args []string) {
			app.Out = cmd.OutOrStdout()
		},
		Run: func(cmd *cobra.Command, args []string) {
			emotionName := args[0]
			app.Emote(emotionName, dest)
		},
		Args: cobra.ExactArgs(1),
	}

	cmdEmote.Flags().StringVar(&dest, "dest", "clipboard", "Where to send your emotion")

	return cmdEmote
}

func buildVersionCommand(app *commands.App) *cobra.Command {
	cmdVersion := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			app.Version()
		},
	}
	return cmdVersion
}
