
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mnemonicLocation] class.
var mnemonicLocationClass _mnemonicLocationClass

func init() {
	mnemonicLocationClass = _mnemonicLocationClass{objc.GetClass("mnemonicLocation")}
}

type _mnemonicLocationClass struct {
	objc.Class
}

// An interface definition for the [mnemonicLocation] class.
type ImnemonicLocation interface {
	ID() objc.ID
}

type mnemonicLocation struct {
	id objc.ID
}

func mnemonicLocationFrom(ptr unsafe.Pointer) mnemonicLocation {
	return mnemonicLocation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mnemonicLocation) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mnemonicLocationClass) Alloc() mnemonicLocation {
	rv := objc.Send[mnemonicLocation](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mnemonicLocationClass) New() mnemonicLocation {
	rv := objc.Send[mnemonicLocation](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmnemonicLocation creates and returns a new initialized instance.
func NewmnemonicLocation() mnemonicLocation {
	return mnemonicLocationClass.New()
}

// Init initializes the instance.
func (m_ mnemonicLocation) Init() mnemonicLocation {
	rv := objc.Send[mnemonicLocation](m_.ID(), selInit)
	return rv
}
