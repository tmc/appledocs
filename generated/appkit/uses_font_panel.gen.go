
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [usesFontPanel] class.
var usesFontPanelClass _usesFontPanelClass

func init() {
	usesFontPanelClass = _usesFontPanelClass{objc.GetClass("usesFontPanel")}
}

type _usesFontPanelClass struct {
	objc.Class
}

// An interface definition for the [usesFontPanel] class.
type IusesFontPanel interface {
	ID() objc.ID
}

type usesFontPanel struct {
	id objc.ID
}

func usesFontPanelFrom(ptr unsafe.Pointer) usesFontPanel {
	return usesFontPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ usesFontPanel) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _usesFontPanelClass) Alloc() usesFontPanel {
	rv := objc.Send[usesFontPanel](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _usesFontPanelClass) New() usesFontPanel {
	rv := objc.Send[usesFontPanel](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewusesFontPanel creates and returns a new initialized instance.
func NewusesFontPanel() usesFontPanel {
	return usesFontPanelClass.New()
}

// Init initializes the instance.
func (u_ usesFontPanel) Init() usesFontPanel {
	rv := objc.Send[usesFontPanel](u_.ID(), selInit)
	return rv
}
