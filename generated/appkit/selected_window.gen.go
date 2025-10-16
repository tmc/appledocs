
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedWindow] class.
var selectedWindowClass _selectedWindowClass

func init() {
	selectedWindowClass = _selectedWindowClass{objc.GetClass("selectedWindow")}
}

type _selectedWindowClass struct {
	objc.Class
}

// An interface definition for the [selectedWindow] class.
type IselectedWindow interface {
	ID() objc.ID
}

type selectedWindow struct {
	id objc.ID
}

func selectedWindowFrom(ptr unsafe.Pointer) selectedWindow {
	return selectedWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedWindow) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedWindowClass) Alloc() selectedWindow {
	rv := objc.Send[selectedWindow](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedWindowClass) New() selectedWindow {
	rv := objc.Send[selectedWindow](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedWindow creates and returns a new initialized instance.
func NewselectedWindow() selectedWindow {
	return selectedWindowClass.New()
}

// Init initializes the instance.
func (s_ selectedWindow) Init() selectedWindow {
	rv := objc.Send[selectedWindow](s_.ID(), selInit)
	return rv
}
