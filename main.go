package main

import (
	"log"
	"os"

	"github.com/Sho372/emote-clone/emotions"
	"github.com/spf13/cobra"
)

func main() {
	app, err := emotions.New()
	if err != nil {
		log.Fatal(err)
	}
	cmd := buildEmoteCommand(app)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildEmoteCommand(app *emotions.App) *cobra.Command {
	var dest string

	emote := &cobra.Command {
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

	emote.Flags().StringVar(&dest, "dest", "clipboard", "Where to send your emotion")

	return emote
}
