// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OpenPanel] class.
var (
	OpenPanelClass     _OpenPanelClass
	OpenPanelClassOnce sync.Once
)

func getOpenPanelClass() _OpenPanelClass {
	OpenPanelClassOnce.Do(func() {
		OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}
	})
	return OpenPanelClass
}

type _OpenPanelClass struct {
	class objc.Class
}

// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ISavePanel
}

// A panel that prompts the user to select a file to open.
//
// Apps use the Open panel as a convenient way to query the user for the name of a file to open. In macOS 10.15 and later, the system always draws Open panels in a separate process, regardless of whether the app is sandboxed. When the user chooses a file to open, macOS adds that file to the app’s sandbox. Prior to macOS 10.15, the system drew the panels in a separate process only for sandboxed apps.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

// Creates a new Open panel and initializes it with a default configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/openPanel
func (oc _OpenPanelClass) OpenPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("openPanel"))
	return rv
}

// A Boolean that indicates whether the user can choose files in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canChooseFiles"))
	return rv
}

// SetCanChooseFiles sets the value of the canChooseFiles property.
// A Boolean that indicates whether the user can choose files in the panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanChooseFiles:"), value)
}
