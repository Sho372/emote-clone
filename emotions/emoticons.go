package emotions

import (
	"github.com/Sho372/emote-clone/config"
	"io"
)

type App struct {
	Out    io.Writer
	Config *config.Config
}
