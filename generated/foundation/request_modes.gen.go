// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [requestModes] class.
var (
	RequestModesClass     _requestModesClass
	RequestModesClassOnce sync.Once
)

func getrequestModesClass() _requestModesClass {
	RequestModesClassOnce.Do(func() {
		RequestModesClass = _requestModesClass{objc.GetClass("requestModes")}
	})
	return RequestModesClass
}

type _requestModesClass struct {
	class objc.Class
}

// An interface definition for the [requestModes] class.
type IrequestModes interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/requestModes-c.ivar

type requestModes struct {
	objectivec.Object
}

// requestModesFrom constructs a [requestModes] from an unsafe.Pointer.
func requestModesFrom(ptr unsafe.Pointer) requestModes {
	return requestModes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _requestModesClass) Alloc() requestModes {
	rv := objc.Send[requestModes](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _requestModesClass) New() requestModes {
	rv := objc.Send[requestModes](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ requestModes) Init() requestModes {
	rv := objc.Send[requestModes](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ requestModes) Autorelease() requestModes {
	rv := objc.Send[requestModes](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrequestModes creates a new requestModes instance.
func NewrequestModes() requestModes {
	return getrequestModesClass().New()
}




