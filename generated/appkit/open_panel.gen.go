// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OpenPanel] class.
var (
	openPanelClass     _OpenPanelClass
	openPanelClassOnce sync.Once
)

func getOpenPanelClass() _OpenPanelClass {
	openPanelClassOnce.Do(func() {
		openPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}
	})
	return openPanelClass
}

type _OpenPanelClass struct {
	class objc.Class
}

// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ISavePanel
}

// A panel that prompts the user to select a file to open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel

type OpenPanel struct {
	SavePanel
}

// OpenPanelFrom constructs a [OpenPanel] from an unsafe.Pointer.
//
// A panel that prompts the user to select a file to open.
func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: SavePanelFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.Send[OpenPanel](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenPanel) Autorelease() OpenPanel {
	rv := objc.Send[OpenPanel](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenPanel creates a new OpenPanel instance.
func NewOpenPanel() OpenPanel {
	return getOpenPanelClass().New()
}




