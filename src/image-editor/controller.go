package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)

	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	windowScroll, err := gtk.ScrolledWindowNew(nil, nil)
	bailOnError(err)

	image, err := gtk.ImageNew()

	windowScroll.Add(image)
	windowMain.SetTitle("Simple Example")
	windowMain.SetDefaultSize(1200, 600)
	windowMain.SetPosition(gtk.WIN_POS_CENTER)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	bailOnError(err)
	windowMain.Container.Add(box)

	menuBar, err := gtk.MenuBarNew()
	bailOnError(err)
	menuFile, err := gtk.MenuNew()
	bailOnError(err)
	menuEdit, err := gtk.MenuNew()
	bailOnError(err)

	menuItemFile, err := gtk.MenuItemNewWithLabel("File")
	bailOnError(err)
	menuItemOpen, err := gtk.MenuItemNewWithLabel("Open")
	bailOnError(err)
	menuItemSave, err := gtk.MenuItemNewWithLabel("Save")
	bailOnError(err)
	menuItemQuit, err := gtk.MenuItemNewWithLabel("Quit")
	bailOnError(err)
	menuItemAbout, err := gtk.MenuItemNewWithLabel("About")
	bailOnError(err)
	menuItemEdit, err := gtk.MenuItemNewWithLabel("Edit")
	bailOnError(err)
	menuItemCopy, err := gtk.MenuItemNewWithLabel("Copy")
	bailOnError(err)
	menuItemPaste, err := gtk.MenuItemNewWithLabel("Paste")
	bailOnError(err)

	// menu bar
	menuBar.MenuShell.Append(menuItemFile)
	menuBar.MenuShell.Append(menuItemEdit)
	menuBar.MenuShell.Append(menuItemAbout)

	// file menu
	menuItemFile.SetSubmenu(menuFile)
	menuFile.MenuShell.Append(menuItemOpen)
	menuFile.MenuShell.Append(menuItemSave)
	menuFile.MenuShell.Append(menuItemQuit)

	// edit menu
	menuItemEdit.SetSubmenu(menuEdit)
	menuEdit.MenuShell.Append(menuItemCopy)
	menuEdit.MenuShell.Append(menuItemPaste)

	box.PackStart(menuBar, false, false, 0)
	box.PackStart(windowScroll, true, true, 0)

	menuItemQuit.Connect("activate", gtk.MainQuit)
	menuItemOpen.Connect("activate", func() {

		fcd, err := gtk.FileChooserDialogNewWith2Buttons("Open file", windowMain,
			gtk.FILE_CHOOSER_ACTION_OPEN,
			"Open", gtk.RESPONSE_OK,
			"Cancel", gtk.RESPONSE_CANCEL)
		bailOnError(err)

		if gtk.ResponseType(fcd.Run()) == gtk.RESPONSE_OK {
			fmt.Printf("%s\n", fcd.GetFilename())

			image.SetFromFile(fcd.GetFilename())

			if err != nil {
				md := gtk.MessageDialogNew(windowMain, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error")
				md.FormatSecondaryText(fmt.Sprintf("%v", err))
				md.Run()
			}
		}

		fcd.Close()
	})

	menuItemAbout.Connect("activate", func() {

		ad, _ := gtk.AboutDialogNew()
		ad.SetTransientFor(windowMain)
		bailOnError(err)
		ad.Run()
		ad.Close()
	})

	//paneView.Add1(treeView)
	//paneView.Add2(label)

	//win.Add(toolbar)
	// Add the iconView to the window.

	//_ = paneView

	windowMain.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

//go build -v -tags gtk_3_18 -gcflags "-N -l"
