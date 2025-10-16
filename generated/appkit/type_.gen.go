
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [type_] class.
var type_Class _type_Class

func init() {
	type_Class = _type_Class{objc.GetClass("type")}
}

type _type_Class struct {
	objc.Class
}

// An interface definition for the [type_] class.
type Itype_ interface {
	ID() objc.ID
}

type type_ struct {
	id objc.ID
}

func type_From(ptr unsafe.Pointer) type_ {
	return type_{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ type_) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _type_Class) Alloc() type_ {
	rv := objc.Send[type_](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _type_Class) New() type_ {
	rv := objc.Send[type_](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtype_ creates and returns a new initialized instance.
func Newtype_() type_ {
	return type_Class.New()
}

// Init initializes the instance.
func (t_ type_) Init() type_ {
	rv := objc.Send[type_](t_.ID(), selInit)
	return rv
}
