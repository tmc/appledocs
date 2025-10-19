// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DistantObject] class.
var (
	distantObjectClass     _DistantObjectClass
	distantObjectClassOnce sync.Once
)

func getDistantObjectClass() _DistantObjectClass {
	distantObjectClassOnce.Do(func() {
		distantObjectClass = _DistantObjectClass{objc.GetClass("NSDistantObject")}
	})
	return distantObjectClass
}

type _DistantObjectClass struct {
	class objc.Class
}

// An interface definition for the [DistantObject] class.
type IDistantObject interface {
	IProxy
}

// A proxy for objects in other applications or threads.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObject
type DistantObject struct {
	Proxy
}

// DistantObjectFrom constructs a [DistantObject] from an unsafe.Pointer.
//
// A proxy for objects in other applications or threads.
func DistantObjectFrom(ptr unsafe.Pointer) DistantObject {
	return DistantObject{
		Proxy: ProxyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DistantObjectClass) Alloc() DistantObject {
	rv := objc.Send[DistantObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DistantObjectClass) New() DistantObject {
	rv := objc.Send[DistantObject](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistantObject) Init() DistantObject {
	rv := objc.Send[DistantObject](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistantObject) Autorelease() DistantObject {
	rv := objc.Send[DistantObject](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistantObject creates a new DistantObject instance.
func NewDistantObject() DistantObject {
	return getDistantObjectClass().New()
}




