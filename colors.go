package main

import (
	"os"
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ColorScheme struct {
	Name string `json:"name"`
	Background string `json:"background"`
	Overlay string `json:"overlay"`
	Text string `json:"text"`
	Color1 string `json:"color1"`
	Color2 string `json:"color2"`
	Color3 string `json:"color3"`
	Color4 string `json:"color4"`
}

type ColorSchemeLayout struct {
	Active string
	ColorSchemes map[string]toml.Primitive `toml:"colorschemes"`
}

type ColorDefaults string

const (
    Name          ColorDefaults = "Megara"
    Background    ColorDefaults = "#201a01"
    Overlay       ColorDefaults = "#342711"
    Text          ColorDefaults = "#d9ae80"
    Color1        ColorDefaults = "#ef540a"
    Color2        ColorDefaults = "#cc094d"
    Color3        ColorDefaults = "#578b99"
    Color4        ColorDefaults = "#e58c00"
)

var DefaultColors  = []struct {
    Value  ColorDefaults
    TSName string
}{
    {Name, "NAME"},
    {Background, "BACKGROUND"},
    {Overlay, "OVERLAY"},
    {Text, "TEXT"},
    {Color1, "COLOR1"},
    {Color2, "COLOR2"},
    {Color3, "COLOR3"},
    {Color4, "COLOR4"},
}

func GetDefaultColors() ColorScheme {
	return ColorScheme {
		Name: string(Name) , 
		Background: string(Background),
		Overlay: string(Overlay),
		Text: string(Text),
		Color1: string(Color1),
		Color2: string(Color2),
		Color3: string(Color3),
		Color4: string(Color4),
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


