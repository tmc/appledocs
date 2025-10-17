
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Nib] class.
var NibClass _NibClass

func init() {
	NibClass = _NibClass{objc.GetClass("NSNib")}
}

type _NibClass struct {
	objc.Class
}

// An interface definition for the [Nib] class.
type INib interface {
	ID() objc.ID
}

type Nib struct {
	id objc.ID
}

func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ Nib) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _NibClass) Alloc() Nib {
	rv := objc.Send[Nib](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _NibClass) New() Nib {
	rv := objc.Send[Nib](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewNib creates and returns a new initialized instance.
func NewNib() Nib {
	return NibClass.New()
}

// Init initializes the instance.
func (n_ Nib) Init() Nib {
	rv := objc.Send[Nib](n_.ID(), selInit)
	return rv
}
