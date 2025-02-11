package main

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
)

type ColorScheme struct {
	Name string
	Background string
	Overlay string
	Text string
	Color1 string
	Color2 string
	Color3 string
	Color4 string
}

type ColorSchemeLayout struct {
	ColorSchemes map[string]toml.Primitive `toml:"colorschemes"`
}

func GetDefaultColors() ColorScheme {
	return ColorScheme {
		Name: "Megara",
		Background: "#201a01",
		Overlay: "#342711",
		Text: "#d9ae80",
		Color1: "#ef540a",
		Color2: "#cc094d",
		Color3: "#578b99",
		Color4: "#e58c00",
	}
}

func ReadColors() []ColorScheme {
	var schemes []ColorScheme

	schemes = append(schemes, GetDefaultColors())
	file, err := os.ReadFile(MEGARA_DIR + "colors.toml")
	if (err != nil) {
		return schemes
	}

	data := string(file)
	var layout ColorSchemeLayout 
	tml, err := toml.Decode(data, &layout)

	if err != nil {
		return schemes
	}

	for name, table := range layout.ColorSchemes {
		var scheme ColorScheme
		tml.PrimitiveDecode(table, &scheme)
		scheme.Name = name 
		schemes = append(schemes, scheme)
	}

	return schemes
}

func (a *App) LoadColors() {
	a.color_schemes = ReadColors()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
    }
	defer watcher.Close()

	go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
				if event.Op == fsnotify.Write {
					a.color_schemes = ReadColors()
				}
            case _, ok := <-watcher.Errors:
                if !ok {
                    return
                }
            }
        }
    }()

	err = watcher.Add(MEGARA_DIR + "colors.toml")
	if (err != nil) {
		return
	}

	<-make(chan struct{})
}


