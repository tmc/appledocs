// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPasskeyIndex] class.
var (
	MPasskeyIndexClass     _mPasskeyIndexClass
	MPasskeyIndexClassOnce sync.Once
)

func getmPasskeyIndexClass() _mPasskeyIndexClass {
	MPasskeyIndexClassOnce.Do(func() {
		MPasskeyIndexClass = _mPasskeyIndexClass{objc.GetClass("mPasskeyIndex")}
	})
	return MPasskeyIndexClass
}

type _mPasskeyIndexClass struct {
	class objc.Class
}

// An interface definition for the [mPasskeyIndex] class.
type ImPasskeyIndex interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/mPasskeyIndex
type mPasskeyIndex struct {
	objectivec.Object
}

// mPasskeyIndexFrom constructs a [mPasskeyIndex] from an unsafe.Pointer.
func mPasskeyIndexFrom(ptr unsafe.Pointer) mPasskeyIndex {
	return mPasskeyIndex{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPasskeyIndexClass) Alloc() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPasskeyIndexClass) New() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPasskeyIndex) Init() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPasskeyIndex) Autorelease() mPasskeyIndex {
	rv := objc.Send[mPasskeyIndex](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPasskeyIndex creates a new mPasskeyIndex instance.
func NewmPasskeyIndex() mPasskeyIndex {
	return getmPasskeyIndexClass().New()
}




