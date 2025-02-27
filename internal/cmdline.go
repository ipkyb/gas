package internal

import (
	"fmt"
	cmdline_handler "gas/internal/cmdline-handler"
	"os"
)

var CmdlineBuild = Cmdline{
	Key:     "build",
	Options: nil,
	Handler: cmdline_handler.Build,
}

var CmdlineClean = Cmdline{
	Key:     "clean",
	Options: nil,
	Handler: cmdline_handler.Clean,
}

var CmdlineInit = Cmdline{
	Key:     "init",
	Options: nil,
	Handler: cmdline_handler.Init,
}

var CmdlineRun = Cmdline{
	Key:     "run",
	Options: nil,
	Handler: cmdline_handler.Run,
}

var CmdlineVersion = Cmdline{
	Key:     "version",
	Options: nil,
	Handler: cmdline_handler.Version(Version),
}

type Cmdline struct {
	Key     string
	Options []string
	Handler func()
}

func CmdlineHandle(handles []Cmdline) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: gas <command> [arguments]")
		return
	}

	key := args[1]
	for _, h := range handles {
		if h.Key == key {
			h.Handler()
			break
		}
	}
}
