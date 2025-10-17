
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintPanel] class.
var PrintPanelClass _PrintPanelClass

func init() {
	PrintPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}
}

type _PrintPanelClass struct {
	objc.Class
}

// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	ID() objc.ID
}

type PrintPanel struct {
	id objc.ID
}

func PrintPanelFrom(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PrintPanel) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.Send[PrintPanel](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPrintPanel creates and returns a new initialized instance.
func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

// Init initializes the instance.
func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.Send[PrintPanel](p_.ID(), selInit)
	return rv
}
