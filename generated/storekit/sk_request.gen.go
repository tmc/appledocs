// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Request] class.
var (
	RequestClass     _RequestClass
	RequestClassOnce sync.Once
)

func getRequestClass() _RequestClass {
	RequestClassOnce.Do(func() {
		RequestClass = _RequestClass{objc.GetClass("SKRequest")}
	})
	return RequestClass
}

type _RequestClass struct {
	class objc.Class
}

// An interface definition for the [Request] class.
type IRequest interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other StoreKit classes.


// A parent class referenced by other StoreKit classes. [Full Topic]
type Request struct {
	objectivec.Object
}

// RequestFrom constructs a [Request] from an unsafe.Pointer.
//
// A parent class referenced by other StoreKit classes.
func RequestFrom(ptr unsafe.Pointer) Request {
	return Request{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RequestClass) Alloc() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RequestClass) New() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Request) Init() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Request) Autorelease() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRequest creates a new Request instance.
func NewRequest() Request {
	return getRequestClass().New()
}




