// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AsynchronousFetchRequest] class.
var (
	asynchronousFetchRequestClass     _AsynchronousFetchRequestClass
	asynchronousFetchRequestClassOnce sync.Once
)

func getAsynchronousFetchRequestClass() _AsynchronousFetchRequestClass {
	asynchronousFetchRequestClassOnce.Do(func() {
		asynchronousFetchRequestClass = _AsynchronousFetchRequestClass{objc.GetClass("NSAsynchronousFetchRequest")}
	})
	return asynchronousFetchRequestClass
}

type _AsynchronousFetchRequestClass struct {
	class objc.Class
}

// An interface definition for the [AsynchronousFetchRequest] class.
type IAsynchronousFetchRequest interface {
	IPersistentStoreRequest
}

// A fetch request that retrieves results asynchronously and supports progress notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest
type AsynchronousFetchRequest struct {
	PersistentStoreRequest
}

// AsynchronousFetchRequestFrom constructs a [AsynchronousFetchRequest] from an unsafe.Pointer.
//
// A fetch request that retrieves results asynchronously and supports progress notification.
func AsynchronousFetchRequestFrom(ptr unsafe.Pointer) AsynchronousFetchRequest {
	return AsynchronousFetchRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AsynchronousFetchRequestClass) Alloc() AsynchronousFetchRequest {
	rv := objc.Send[AsynchronousFetchRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AsynchronousFetchRequestClass) New() AsynchronousFetchRequest {
	rv := objc.Send[AsynchronousFetchRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AsynchronousFetchRequest) Init() AsynchronousFetchRequest {
	rv := objc.Send[AsynchronousFetchRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AsynchronousFetchRequest) Autorelease() AsynchronousFetchRequest {
	rv := objc.Send[AsynchronousFetchRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAsynchronousFetchRequest creates a new AsynchronousFetchRequest instance.
func NewAsynchronousFetchRequest() AsynchronousFetchRequest {
	return getAsynchronousFetchRequestClass().New()
}


// Initializes a new asynchronous fetch request configured with the provided fetch request and completion block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAsynchronousFetchRequest/init(fetchRequest:completionBlock:)
func NewAsynchronousFetchRequestWithFetchRequestCompletionBlock(request unsafe.Pointer, blk unsafe.Pointer) AsynchronousFetchRequest {
	instance := getAsynchronousFetchRequestClass().Alloc()
	rv := objc.Send[AsynchronousFetchRequest](instance.ID, objc.Sel("initWithFetchRequest:completionBlock:"), request, blk)
	rv.Autorelease()
	return rv
}



