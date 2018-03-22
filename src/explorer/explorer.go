package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func read(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fi.IsDir() {
		return filename
	}
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return "EXIF data missing or incomplete"
	}

	builder := strings.Builder{}

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Sprintf("%s\n", camModel.String())
	builder.WriteString(fmt.Sprintf("Camera model %s\n", camModel.String()))

	lensModel, _ := x.Get(exif.LensModel)
	builder.WriteString(fmt.Sprintf("Lens model %s\n", lensModel.String()))

	fmt.Println(lensModel.Type)

	focal, _ := x.Get(exif.FocalLength)
	focalLength, _ := focal.Int(1)
	builder.WriteString(fmt.Sprintf("Focal length %d mm\n", focalLength))

	apperture, _ := x.Get(exif.ApertureValue)
	numer, denom, _ := apperture.Rat2(0) // retrieve first (only) rat. value

	fmt.Println(numer, denom)

	//	builder.WriteString(fmt.Sprintf("Apperture F/%f\n", app))

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	builder.WriteString(fmt.Sprintf("Taken: %s\n", tm))

	// lat, long, _ := x.LatLong()
	// builder.WriteString(fmt.Sprintf("lat %f long %f\n", lat, long))

	return builder.String()
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Unable to create treeView:", err)
	}

	cellRenderer, err := gtk.CellRendererTextNew()
	column, err := gtk.TreeViewColumnNewWithAttribute("Tree", cellRenderer, "text", 0)

	treeView.ExpandAll()

	treeView.AppendColumn(column)

	treeStore, err := gtk.TreeStoreNew(glib.TYPE_STRING)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dir := user.HomeDir + "/gocode/src/explorer"
	listDir(dir, 0, treeStore, treeStore.Append(nil))

	treeView.SetModel(treeStore)
	treeView.ExpandAll()
	selection, err := treeView.GetSelection()
	selection.SetMode(gtk.SELECTION_SINGLE)

	label, err := gtk.LabelNew("$$")

	treeView.Connect("cursor-changed", func() {

		selection, _ := treeView.GetSelection()
		imodel, iter, _ := selection.GetSelected()
		model := imodel.(*gtk.TreeModel)
		value, _ := model.GetValue(iter, 0)
		path, _ := value.GetString()
		label.SetText(read(path))

	})

	paneView, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)

	paneView.Add1(treeView)
	paneView.Add2(label)

	// Add the iconView to the window.
	win.Add(paneView)

	// Set the default window size.
	win.SetDefaultSize(800, 200)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func listDir(path string, level int, treeStore *gtk.TreeStore, parent *gtk.TreeIter) {

	files, _ := ioutil.ReadDir(path)
	level++

	treeStore.SetValue(parent, 0, path)

	for _, file := range files {

		switch {

		case file.IsDir():
			//fmt.Printf("D %s %d\n", file.Name(), level)

			child := treeStore.Append(parent)

			listDir(file.Name(), level, treeStore, child)
		default:
			//fmt.Printf("F %s %d\n", file.Name(), level)

			child := treeStore.Append(parent)
			treeStore.SetValue(child, 0, file.Name())
		}
	}
}

//go build -v -tags gtk_3_18 -gcflags "-N -l"
