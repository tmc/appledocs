
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [init_] class.
var init_Class _init_Class

func init() {
	init_Class = _init_Class{objc.GetClass("init")}
}

type _init_Class struct {
	objc.Class
}

// An interface definition for the [init_] class.
type Iinit_ interface {
	ID() objc.ID
}

type init_ struct {
	id objc.ID
}

func init_From(ptr unsafe.Pointer) init_ {
	return init_{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ init_) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _init_Class) Alloc() init_ {
	rv := objc.Send[init_](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _init_Class) New() init_ {
	rv := objc.Send[init_](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newinit_ creates and returns a new initialized instance.
func Newinit_() init_ {
	return init_Class.New()
}

// Init initializes the instance.
func (i_ init_) Init() init_ {
	rv := objc.Send[init_](i_.ID(), selInit)
	return rv
}
