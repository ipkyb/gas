package cmdline_handler

const (
	Startup_Filepath_Config string = "gas.config.json"
	Startup_Dirpath_Plugins string = "plugins"
)

func Run() {
	panic("Unimplemented")
}

// func foo() {
// 	cfg := internal.ConfigDefault()

// 	if err := configLoad(Startup_Filepath_Config, &cfg); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	if err := pluginsLoad(Startup_Dirpath_Plugins); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	if err := fiberRun(cfg.Http); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
