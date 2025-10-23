// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mOBEXSession] class.
var (
	MOBEXSessionClass     _mOBEXSessionClass
	MOBEXSessionClassOnce sync.Once
)

func getmOBEXSessionClass() _mOBEXSessionClass {
	MOBEXSessionClassOnce.Do(func() {
		MOBEXSessionClass = _mOBEXSessionClass{objc.GetClass("mOBEXSession")}
	})
	return MOBEXSessionClass
}

type _mOBEXSessionClass struct {
	class objc.Class
}

// An interface definition for the [mOBEXSession] class.
type ImOBEXSession interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mOBEXSession
type mOBEXSession struct {
	objectivec.Object
}

// mOBEXSessionFrom constructs a [mOBEXSession] from an unsafe.Pointer.
func mOBEXSessionFrom(ptr unsafe.Pointer) mOBEXSession {
	return mOBEXSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mOBEXSessionClass) Alloc() mOBEXSession {
	rv := objc.Send[mOBEXSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mOBEXSessionClass) New() mOBEXSession {
	rv := objc.Send[mOBEXSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mOBEXSession) Init() mOBEXSession {
	rv := objc.Send[mOBEXSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mOBEXSession) Autorelease() mOBEXSession {
	rv := objc.Send[mOBEXSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmOBEXSession creates a new mOBEXSession instance.
func NewmOBEXSession() mOBEXSession {
	return getmOBEXSessionClass().New()
}




