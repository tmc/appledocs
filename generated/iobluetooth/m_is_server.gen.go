// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIsServer] class.
var (
	MIsServerClass     _mIsServerClass
	MIsServerClassOnce sync.Once
)

func getmIsServerClass() _mIsServerClass {
	MIsServerClassOnce.Do(func() {
		MIsServerClass = _mIsServerClass{objc.GetClass("mIsServer")}
	})
	return MIsServerClass
}

type _mIsServerClass struct {
	class objc.Class
}

// An interface definition for the [mIsServer] class.
type ImIsServer interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mIsServer
type mIsServer struct {
	objectivec.Object
}

// mIsServerFrom constructs a [mIsServer] from an unsafe.Pointer.
func mIsServerFrom(ptr unsafe.Pointer) mIsServer {
	return mIsServer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIsServerClass) Alloc() mIsServer {
	rv := objc.Send[mIsServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIsServerClass) New() mIsServer {
	rv := objc.Send[mIsServer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIsServer) Init() mIsServer {
	rv := objc.Send[mIsServer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIsServer) Autorelease() mIsServer {
	rv := objc.Send[mIsServer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIsServer creates a new mIsServer instance.
func NewmIsServer() mIsServer {
	return getmIsServerClass().New()
}




