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
	Cancel()
	Start()
}

// An abstract class that represents a request to the App Store.
//
// To make a request, initialize a subclass of —such as or —set the property, and call the method.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest
type Request struct {
	objectivec.Object
}

// RequestFrom constructs a [Request] from an unsafe.Pointer.
//
// An abstract class that represents a request to the App Store.
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


// Cancels a previously started request.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/cancel()
func (r_ Request) Cancel() {
	objc.Send[objc.ID](r_.ID, objc.Sel("cancel"))
}

// Sends the request to the Apple App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/start()
func (r_ Request) Start() {
	objc.Send[objc.ID](r_.ID, objc.Sel("start"))
}

// The delegate of the request object.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/delegate
func (r_ Request) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the request object.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/delegate
func (r_ Request) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}



