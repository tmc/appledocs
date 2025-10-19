// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Nib] class.
var nibClass = _NibClass{objc.GetClass("NSNib")}

type _NibClass struct {
	class objc.Class
}

// An interface definition for the [Nib] class.
type INib interface {
	objectivec.IObject
}

// An object wrapper, or container, for an Interface Builder nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib

type Nib struct {
	objectivec.Object
}

// NibFrom constructs a [Nib] from an unsafe.Pointer.
//
// An object wrapper, or container, for an Interface Builder nib file.
func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NibClass) Alloc() Nib {
	rv := objc.Send[Nib](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NibClass) New() Nib {
	rv := objc.Send[Nib](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Nib) Init() Nib {
	rv := objc.Send[Nib](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Nib) Autorelease() Nib {
	rv := objc.Send[Nib](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNib creates a new Nib instance.
func NewNib() Nib {
	return nibClass.New()
}




