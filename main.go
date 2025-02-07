package main

import (
	"embed"
	"strings"
	"os"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/build
var assets embed.FS
var MEGARA_DIR string


func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	secondInstanceArgs := secondInstanceData.Args

    println("user opened second instance", strings.Join(secondInstanceData.Args, ","))
    println("user opened second from", secondInstanceData.WorkingDirectory)
    runtime.WindowUnminimise(a.ctx)
    runtime.Show(a.ctx)
    go runtime.EventsEmit(a.ctx, "launchArgs", secondInstanceArgs)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	config_dir, err := os.UserConfigDir()

	if err != nil {
		return
	}

	MEGARA_DIR = config_dir + "/megara/"


	err = os.Mkdir(MEGARA_DIR, 0755)
	if err != nil && !os.IsExist(err) {
        fmt.Println("Error:", err)
        return
    }

	file, err := os.Create(MEGARA_DIR + "test.toml")
	if err != nil && !os.IsExist(err) {
        fmt.Println("Error:", err)
        return
    }

	file.WriteString("test string")
	file.Close()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "megara",
		Width:  900,
		Height: 800,
		DisableResize: true,
		Frameless: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 34, G: 26, B: 2, A: 1},
		OnStartup:        app.startup,
		SingleInstanceLock: &options.SingleInstanceLock{
            UniqueId:               "e3984e08-28dc-4e3d-b70a-45e961589cdc",
            OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
        },
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
