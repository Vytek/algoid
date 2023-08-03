package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/slog"
)

// Constants
const ServiceName = "ALGOID"
const ServiceVersion = "v0.1 Alfa"

/*
Check: https://github.com/goccy/go-json
https://github.com/w3c-ccg/community/blob/main/work_items.md
*/

func main() {
	myFigure := figure.NewFigure(ServiceName+" "+ServiceVersion, "", true)
	myFigure.Print()
	slog.Infof("Service: %s Version: %s", ServiceName, ServiceVersion)
	//slog.Warn("warning log message")
	//slog.Infof("info log %s", "message")
	//slog.Debugf("debug %s", "message")
}
