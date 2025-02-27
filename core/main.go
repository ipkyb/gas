package main

import "gas/internal"

var handles = []internal.Cmdline{
	internal.CmdlineRun,
	internal.CmdlineVersion,
}

func main() {
	internal.CmdlineHandle(handles)
}
