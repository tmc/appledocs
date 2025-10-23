// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mWeAreAborting] class.
var (
	MWeAreAbortingClass     _mWeAreAbortingClass
	MWeAreAbortingClassOnce sync.Once
)

func getmWeAreAbortingClass() _mWeAreAbortingClass {
	MWeAreAbortingClassOnce.Do(func() {
		MWeAreAbortingClass = _mWeAreAbortingClass{objc.GetClass("mWeAreAborting")}
	})
	return MWeAreAbortingClass
}

type _mWeAreAbortingClass struct {
	class objc.Class
}

// An interface definition for the [mWeAreAborting] class.
type ImWeAreAborting interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mWeAreAborting
type mWeAreAborting struct {
	objectivec.Object
}

// mWeAreAbortingFrom constructs a [mWeAreAborting] from an unsafe.Pointer.
func mWeAreAbortingFrom(ptr unsafe.Pointer) mWeAreAborting {
	return mWeAreAborting{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mWeAreAbortingClass) Alloc() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mWeAreAbortingClass) New() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mWeAreAborting) Init() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mWeAreAborting) Autorelease() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmWeAreAborting creates a new mWeAreAborting instance.
func NewmWeAreAborting() mWeAreAborting {
	return getmWeAreAbortingClass().New()
}




