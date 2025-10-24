// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [native] class.
var (
	NativeClass     _nativeClass
	NativeClassOnce sync.Once
)

func getnativeClass() _nativeClass {
	NativeClassOnce.Do(func() {
		NativeClass = _nativeClass{objc.GetClass("native")}
	})
	return NativeClass
}

type _nativeClass struct {
	class objc.Class
}

// An interface definition for the [native] class.
type Inative interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-c.ivar
type native struct {
	objectivec.Object
}

// nativeFrom constructs a [native] from an unsafe.Pointer.
func nativeFrom(ptr unsafe.Pointer) native {
	return native{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _nativeClass) Alloc() native {
	rv := objc.Send[native](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _nativeClass) New() native {
	rv := objc.Send[native](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ native) Init() native {
	rv := objc.Send[native](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ native) Autorelease() native {
	rv := objc.Send[native](n_.ID, objc.Sel("autorelease"))
	return rv
}

// Newnative creates a new native instance.
func Newnative() native {
	return getnativeClass().New()
}




