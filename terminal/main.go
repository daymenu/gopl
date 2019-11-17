package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Terminal")
	window.Connect("destroy", gtk.MainQuit)

	window.Show()

	gtk.Main()
}
