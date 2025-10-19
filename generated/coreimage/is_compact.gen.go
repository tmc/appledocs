// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [isCompact] class.
var isCompactClass = _isCompactClass{objc.GetClass("isCompact")}

type _isCompactClass struct {
	class objc.Class
}

// An interface definition for the [isCompact] class.
type IisCompact interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/isCompact-c.ivar

type isCompact struct {
	objectivec.Object
}

// isCompactFrom constructs a [isCompact] from an unsafe.Pointer.
func isCompactFrom(ptr unsafe.Pointer) isCompact {
	return isCompact{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ic _isCompactClass) Alloc() isCompact {
	rv := objc.Send[isCompact](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _isCompactClass) New() isCompact {
	rv := objc.Send[isCompact](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isCompact) Init() isCompact {
	rv := objc.Send[isCompact](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isCompact) Autorelease() isCompact {
	rv := objc.Send[isCompact](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisCompact creates a new isCompact instance.
func NewisCompact() isCompact {
	return isCompactClass.New()
}




