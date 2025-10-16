
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Typesetter] class.
var TypesetterClass _TypesetterClass

func init() {
	TypesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}
}

type _TypesetterClass struct {
	objc.Class
}

// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	ID() objc.ID
}

type Typesetter struct {
	id objc.ID
}

func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ Typesetter) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TypesetterClass) Alloc() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TypesetterClass) New() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTypesetter creates and returns a new initialized instance.
func NewTypesetter() Typesetter {
	return TypesetterClass.New()
}

// Init initializes the instance.
func (t_ Typesetter) Init() Typesetter {
	rv := objc.Send[Typesetter](t_.ID(), selInit)
	return rv
}
