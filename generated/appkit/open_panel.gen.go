
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenPanel] class.
var OpenPanelClass _OpenPanelClass

func init() {
	OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}
}

type _OpenPanelClass struct {
	objc.Class
}

// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ID() objc.ID
}

type OpenPanel struct {
	id objc.ID
}

func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenPanel) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenPanel creates and returns a new initialized instance.
func NewOpenPanel() OpenPanel {
	return OpenPanelClass.New()
}

// Init initializes the instance.
func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.Send[OpenPanel](o_.ID(), selInit)
	return rv
}
// A Boolean that indicates whether the user can choose files in the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.Send[bool](o_.ID(), objc.RegisterName("canChooseFiles"))
	return rv
}
// SetCanChooseFiles sets the value of the canChooseFiles property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("setCanChooseFiles:"), value)
}
