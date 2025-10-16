
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [keyViewSelectionDirection] class.
var keyViewSelectionDirectionClass _keyViewSelectionDirectionClass

func init() {
	keyViewSelectionDirectionClass = _keyViewSelectionDirectionClass{objc.GetClass("keyViewSelectionDirection")}
}

type _keyViewSelectionDirectionClass struct {
	objc.Class
}

// An interface definition for the [keyViewSelectionDirection] class.
type IkeyViewSelectionDirection interface {
	ID() objc.ID
}

type keyViewSelectionDirection struct {
	id objc.ID
}

func keyViewSelectionDirectionFrom(ptr unsafe.Pointer) keyViewSelectionDirection {
	return keyViewSelectionDirection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (k_ keyViewSelectionDirection) ID() objc.ID {
	return k_.id
}

// Alloc allocates a new instance without initialization.
func (kc _keyViewSelectionDirectionClass) Alloc() keyViewSelectionDirection {
	rv := objc.Send[keyViewSelectionDirection](objc.ID(kc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (kc _keyViewSelectionDirectionClass) New() keyViewSelectionDirection {
	rv := objc.Send[keyViewSelectionDirection](objc.ID(kc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewkeyViewSelectionDirection creates and returns a new initialized instance.
func NewkeyViewSelectionDirection() keyViewSelectionDirection {
	return keyViewSelectionDirectionClass.New()
}

// Init initializes the instance.
func (k_ keyViewSelectionDirection) Init() keyViewSelectionDirection {
	rv := objc.Send[keyViewSelectionDirection](k_.ID(), selInit)
	return rv
}
