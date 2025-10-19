// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchUpdateRequest] class.
var batchUpdateRequestClass = _BatchUpdateRequestClass{objc.GetClass("NSBatchUpdateRequest")}

type _BatchUpdateRequestClass struct {
	class objc.Class
}

// A request to Core Data to do a batch update of data in a persistent store without loading any data into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest

type BatchUpdateRequest struct {
	PersistentStoreRequest
}

// BatchUpdateRequestFrom constructs a [BatchUpdateRequest] from an unsafe.Pointer.
//
// A request to Core Data to do a batch update of data in a persistent store without loading any data into memory.
func BatchUpdateRequestFrom(ptr unsafe.Pointer) BatchUpdateRequest {
	return BatchUpdateRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (bc _BatchUpdateRequestClass) Alloc() BatchUpdateRequest {
	rv := objc.Send[BatchUpdateRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (bc _BatchUpdateRequestClass) New() BatchUpdateRequest {
	rv := objc.Send[BatchUpdateRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchUpdateRequest) Init() BatchUpdateRequest {
	rv := objc.Send[BatchUpdateRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchUpdateRequest) Autorelease() BatchUpdateRequest {
	rv := objc.Send[BatchUpdateRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchUpdateRequest creates a new BatchUpdateRequest instance.
func NewBatchUpdateRequest() BatchUpdateRequest {
	return batchUpdateRequestClass.New()
}
// Creates a batch-update request for a managed entity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entity:)
func NewBatchUpdateRequestWithEntity(entity unsafe.Pointer) BatchUpdateRequest {
	instance := batchUpdateRequestClass.Alloc()
	rv := objc.Send[BatchUpdateRequest](instance.ID, objc.Sel("initWithEntity:"), entity)
	rv.Autorelease()
	return rv
}
// Creates a batch-update request for a named managed entity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entityName:)
func NewBatchUpdateRequestWithEntityName(entityName string) BatchUpdateRequest {
	instance := batchUpdateRequestClass.Alloc()
	rv := objc.Send[BatchUpdateRequest](instance.ID, objc.Sel("initWithEntityName:"), entityName)
	rv.Autorelease()
	return rv
}


// Creates a batch-update request for a named managed entity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/batchUpdateRequestWithEntityName:
func (bc _BatchUpdateRequestClass) BatchUpdateRequestWithEntityName(entityName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("batchUpdateRequestWithEntityName:"), entityName)
	return rv
}

