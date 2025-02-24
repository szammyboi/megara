package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ColorScheme struct {
	Name     string `json:"name"`
	Base     string `json:"base"`
	Overlay1 string `json:"overlay1"`
	Overlay2 string `json:"overlay2"`
	Overlay3 string `json:"overlay3"`
	Text     string `json:"text"`
	Primary1 string `json:"primary1"`
	Primary2 string `json:"primary2"`
	Primary3 string `json:"primary3"`
	Primary4 string `json:"primary4"`
}

type ColorSchemeLayout struct {
	Active string
	ColorSchemes map[string]toml.Primitive `toml:"colorschemes"`
}

type ColorDefaults string

const (
    Name      ColorDefaults = "Megara"
    Base      ColorDefaults = "#201a01"
    Overlay1  ColorDefaults = "#342711"
    Overlay2  ColorDefaults = "#342711"
    Overlay3  ColorDefaults = "#342711"
    Text      ColorDefaults = "#d9ae80"
    Primary1  ColorDefaults = "#ef540a"
    Primary2  ColorDefaults = "#cc094d"
    Primary3  ColorDefaults = "#578b99"
    Primary4  ColorDefaults = "#e58c00"
)

func GetDefaultColors() ColorScheme {
	return ColorScheme {
		Name:     string(Name) , 
		Base:     string(Base),
		Overlay1: string(Overlay1),
		Overlay2: string(Overlay2),
		Overlay3: string(Overlay3),
		Text:     string(Text),
		Primary1: string(Primary1),
		Primary2: string(Primary2),
		Primary3: string(Primary3),
		Primary4: string(Primary4),
	}
}

func ReadColors() (ColorScheme, []ColorScheme) {
	var active ColorScheme
	var schemes []ColorScheme

	schemes = append(schemes, GetDefaultColors())
	active = schemes[0]

	file, err := os.ReadFile(MEGARA_DIR + "colors.toml")
	if (err != nil) {
		return active, schemes
	}

	data := string(file)
	var layout ColorSchemeLayout 
	tml, err := toml.Decode(data, &layout)

	if err != nil {
		return active, schemes
	}

	for name, table := range layout.ColorSchemes {
		var scheme ColorScheme
		tml.PrimitiveDecode(table, &scheme)
		scheme.Name = name 
		if (name == layout.Active) {
			active = scheme
		}
		schemes = append(schemes, scheme)
	}

	return active, schemes
}

func (a *App) LoadColors() {
	a.active, a.color_schemes = ReadColors()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
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
					a.active, a.color_schemes = ReadColors()
					runtime.EventsEmit(a.ctx, "updateColors", a.color_schemes)
					runtime.EventsEmit(a.ctx, "updateActiveColor", a.active)
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


