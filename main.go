package main

// To compile: CGO_CFLAGS=-Wno-deprecated-declarations go build -v

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatalf("Unable to create window: %v", err)
	}

	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	label, err := gtk.LabelNew("Drag using middle mouse button")
	if err != nil {
		log.Fatalf("Unable to create label: %v", err)
	}

	eventBox, err := gtk.EventBoxNew()
	if err != nil {
		log.Fatalf("Unable to create event box: %v", err)
	}

	eventBox.Connect("button-press-event", func(tree *gtk.EventBox, event *gdk.Event) bool {
		button := gdk.EventButtonNewFromEvent(event)
		switch button.Button() {
		case gdk.BUTTON_PRIMARY:
			return true
		case gdk.BUTTON_MIDDLE:
			window.BeginMoveDrag(
				gdk.BUTTON_MIDDLE,
				int(button.XRoot()),
				int(button.YRoot()),
				button.Time(),
			)
			return true
		case gdk.BUTTON_SECONDARY:
			window.Close()
			return true
		default:
			return false
		}
	})

	eventBox.Add(label)

	window.Add(eventBox)
	window.SetTitle("Undecorated")
	window.SetDefaultSize(800, 600)
	window.SetDecorated(false)
	window.SetKeepAbove(true)
	window.SetSkipPagerHint(true)
	window.SetSkipTaskbarHint(true)
	window.SetTypeHint(gdk.WINDOW_TYPE_HINT_DND)
	window.Stick()

	window.ShowAll()

	gtk.Main()
}
