// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchRequest] class.
var fetchRequestClass = _FetchRequestClass{objc.GetClass("NSFetchRequest")}

type _FetchRequestClass struct {
	class objc.Class
}

// An interface definition for the [FetchRequest] class.
type IFetchRequest interface {
	IPersistentStoreRequest
	Execute(error unsafe.Pointer) unsafe.Pointer
}

// A description of search criteria used to retrieve data from a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest

type FetchRequest struct {
	PersistentStoreRequest
}

// FetchRequestFrom constructs a [FetchRequest] from an unsafe.Pointer.
//
// A description of search criteria used to retrieve data from a persistent store.
func FetchRequestFrom(ptr unsafe.Pointer) FetchRequest {
	return FetchRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (fc _FetchRequestClass) Alloc() FetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (fc _FetchRequestClass) New() FetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FetchRequest) Init() FetchRequest {
	rv := objc.Send[FetchRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FetchRequest) Autorelease() FetchRequest {
	rv := objc.Send[FetchRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchRequest creates a new FetchRequest instance.
func NewFetchRequest() FetchRequest {
	return fetchRequestClass.New()
}


// Initializes a fetch request configured with a given entity name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/init(entityName:)
func NewFetchRequestWithEntityName(entityName string) FetchRequest {
	instance := fetchRequestClass.Alloc()
	rv := objc.Send[FetchRequest](instance.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	rv.Autorelease()
	return rv
}


// Returns a fetch request configured with a given entity name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/fetchRequestWithEntityName:
func (fc _FetchRequestClass) FetchRequestWithEntityName(entityName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fetchRequestWithEntityName:"), objc.String(entityName))
	return rv
}
// Executes the fetch request against the managed object context that is associated with the current queue. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/execute()
func (f_ FetchRequest) Execute(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("execute:"), error)
	return rv
}

