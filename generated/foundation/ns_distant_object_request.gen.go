// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DistantObjectRequest] class.
var (
	DistantObjectRequestClass     _DistantObjectRequestClass
	DistantObjectRequestClassOnce sync.Once
)

func getDistantObjectRequestClass() _DistantObjectRequestClass {
	DistantObjectRequestClassOnce.Do(func() {
		DistantObjectRequestClass = _DistantObjectRequestClass{objc.GetClass("NSDistantObjectRequest")}
	})
	return DistantObjectRequestClass
}

type _DistantObjectRequestClass struct {
	class objc.Class
}

// An interface definition for the [DistantObjectRequest] class.
type IDistantObjectRequest interface {
	objectivec.IObject
	Connection() IConnection
	Conversation() objc.ID
	Invocation() IInvocation
}

// An object used by the distributed objects system to help handle invocations between different processes.
//
// Do not create objects directly. Unless you are getting involved with the low-level details of distributed objects, there should never be a need to access an . To intercept and possibly process requests yourself, implement the delegate method .


// An object used by the distributed objects system to help handle invocations between different processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest
type DistantObjectRequest struct {
	objectivec.Object
}

// DistantObjectRequestFrom constructs a [DistantObjectRequest] from an unsafe.Pointer.
//
// An object used by the distributed objects system to help handle invocations between different processes.
func DistantObjectRequestFrom(ptr unsafe.Pointer) DistantObjectRequest {
	return DistantObjectRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DistantObjectRequestClass) Alloc() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DistantObjectRequestClass) New() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistantObjectRequest) Init() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistantObjectRequest) Autorelease() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistantObjectRequest creates a new DistantObjectRequest instance.
func NewDistantObjectRequest() DistantObjectRequest {
	return getDistantObjectRequestClass().New()
}



// Returns the object involved in the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/connection
func (d_ DistantObjectRequest) Connection() IConnection {
	rv := objc.Send[NSConnection](d_.ID, objc.Sel("connection"))
	return rv
}


// Returns the token object representing the conversation in which the receiver was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/conversation
func (d_ DistantObjectRequest) Conversation() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("conversation"))
	return rv
}


// Returns the object for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/invocation
func (d_ DistantObjectRequest) Invocation() IInvocation {
	rv := objc.Send[NSInvocation](d_.ID, objc.Sel("invocation"))
	return rv
}



