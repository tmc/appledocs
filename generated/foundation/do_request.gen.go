// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [doRequest] class.
var (
	DoRequestClass     _doRequestClass
	DoRequestClassOnce sync.Once
)

func getdoRequestClass() _doRequestClass {
	DoRequestClassOnce.Do(func() {
		DoRequestClass = _doRequestClass{objc.GetClass("doRequest")}
	})
	return DoRequestClass
}

type _doRequestClass struct {
	class objc.Class
}

// An interface definition for the [doRequest] class.
type IdoRequest interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/doRequest

type doRequest struct {
	objectivec.Object
}

// doRequestFrom constructs a [doRequest] from an unsafe.Pointer.
func doRequestFrom(ptr unsafe.Pointer) doRequest {
	return doRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _doRequestClass) Alloc() doRequest {
	rv := objc.Send[doRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _doRequestClass) New() doRequest {
	rv := objc.Send[doRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ doRequest) Init() doRequest {
	rv := objc.Send[doRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ doRequest) Autorelease() doRequest {
	rv := objc.Send[doRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdoRequest creates a new doRequest instance.
func NewdoRequest() doRequest {
	return getdoRequestClass().New()
}




