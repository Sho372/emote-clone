package emotions

import (
	"fmt"
	"io"
	"os"

	"github.com/Sho372/emote-clone/config"
	"github.com/atotto/clipboard"
)

type App struct {
	Out    io.Writer
	Config *config.Config
}

func New() (*App, error) {
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	a := &App{
		Out:    os.Stdout,
		Config: c,
	}
	return a, nil
}

func (a *App) Emote(name string, dest string) {
	emoticon := a.Config.Emoticon[name]

	switch dest {
	case "clipboard":
		clipboard.WriteAll(emoticon)
		// put out emoticon to clipboard
		fmt.Fprintf(a.Out, emoticon, "was copied to the clipboard")
	default:
		// put out emoticon to standard output
		fmt.Fprintf(a.Out, emoticon)
	}

}
