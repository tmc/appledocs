// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mReturnImage] class.
var (
	MReturnImageClass     _mReturnImageClass
	MReturnImageClassOnce sync.Once
)

func getmReturnImageClass() _mReturnImageClass {
	MReturnImageClassOnce.Do(func() {
		MReturnImageClass = _mReturnImageClass{objc.GetClass("mReturnImage")}
	})
	return MReturnImageClass
}

type _mReturnImageClass struct {
	class objc.Class
}

// An interface definition for the [mReturnImage] class.
type ImReturnImage interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mReturnImage
type mReturnImage struct {
	objectivec.Object
}

// mReturnImageFrom constructs a [mReturnImage] from an unsafe.Pointer.
func mReturnImageFrom(ptr unsafe.Pointer) mReturnImage {
	return mReturnImage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mReturnImageClass) Alloc() mReturnImage {
	rv := objc.Send[mReturnImage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mReturnImageClass) New() mReturnImage {
	rv := objc.Send[mReturnImage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mReturnImage) Init() mReturnImage {
	rv := objc.Send[mReturnImage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mReturnImage) Autorelease() mReturnImage {
	rv := objc.Send[mReturnImage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmReturnImage creates a new mReturnImage instance.
func NewmReturnImage() mReturnImage {
	return getmReturnImageClass().New()
}




