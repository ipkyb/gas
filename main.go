package main

import "github.com/ipkyb/gas/internal"

var handles = []internal.Cmdline{
	internal.CmdlineBuild,
	internal.CmdlineClean,
	internal.CmdlineInit,
	internal.CmdlineRun,
	internal.CmdlineVersion,
}

func main() {
	internal.CmdlineHandle(handles)
}
