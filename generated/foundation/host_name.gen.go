// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hostName] class.
var (
	HostNameClass     _hostNameClass
	HostNameClassOnce sync.Once
)

func gethostNameClass() _hostNameClass {
	HostNameClassOnce.Do(func() {
		HostNameClass = _hostNameClass{objc.GetClass("hostName")}
	})
	return HostNameClass
}

type _hostNameClass struct {
	class objc.Class
}

// An interface definition for the [hostName] class.
type IhostName interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/hostName-c.ivar

type hostName struct {
	objectivec.Object
}

// hostNameFrom constructs a [hostName] from an unsafe.Pointer.
func hostNameFrom(ptr unsafe.Pointer) hostName {
	return hostName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _hostNameClass) Alloc() hostName {
	rv := objc.Send[hostName](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _hostNameClass) New() hostName {
	rv := objc.Send[hostName](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hostName) Init() hostName {
	rv := objc.Send[hostName](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hostName) Autorelease() hostName {
	rv := objc.Send[hostName](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhostName creates a new hostName instance.
func NewhostName() hostName {
	return gethostNameClass().New()
}




