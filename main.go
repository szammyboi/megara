package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
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

	file.Close()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "megara",
		Width:  1000,
		Height: 500,
		DisableResize: true,
		Frameless: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:        app.startup,
		SingleInstanceLock: &options.SingleInstanceLock{
            UniqueId:               "e3984e08-28dc-4e3d-b70a-45e961589cdc",
            OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
        },
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings {
				DarkModeTitleBar:   windows.RGB(34, 26, 2),
                DarkModeTitleText:  windows.RGB(34, 26, 2),
                DarkModeBorder:     windows.RGB(34, 26, 2),
				DarkModeTitleBarInactive: windows.RGB(54, 39, 18),
				DarkModeTitleTextInactive: windows.RGB(54, 39, 18),
                LightModeTitleBar:  windows.RGB(34, 26, 2),
                LightModeTitleText: windows.RGB(34, 26, 2),
                LightModeBorder:    windows.RGB(34, 26, 2),	
				LightModeTitleBarInactive: windows.RGB(54, 39, 18),
				LightModeTitleTextInactive: windows.RGB(54, 39, 18),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
