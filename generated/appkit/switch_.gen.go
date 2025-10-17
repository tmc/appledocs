
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Switch_] class.
var Switch_Class _Switch_Class

func init() {
	Switch_Class = _Switch_Class{objc.GetClass("NSSwitch")}
}

type _Switch_Class struct {
	objc.Class
}

// An interface definition for the [Switch_] class.
type ISwitch_ interface {
	ID() objc.ID
}

type Switch_ struct {
	id objc.ID
}

func Switch_From(ptr unsafe.Pointer) Switch_ {
	return Switch_{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Switch_) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _Switch_Class) Alloc() Switch_ {
	rv := objc.Send[Switch_](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _Switch_Class) New() Switch_ {
	rv := objc.Send[Switch_](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSwitch_ creates and returns a new initialized instance.
func NewSwitch_() Switch_ {
	return Switch_Class.New()
}

// Init initializes the instance.
func (s_ Switch_) Init() Switch_ {
	rv := objc.Send[Switch_](s_.ID(), selInit)
	return rv
}
