// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchDeleteRequest] class.
var (
	batchDeleteRequestClass     _BatchDeleteRequestClass
	batchDeleteRequestClassOnce sync.Once
)

func getBatchDeleteRequestClass() _BatchDeleteRequestClass {
	batchDeleteRequestClassOnce.Do(func() {
		batchDeleteRequestClass = _BatchDeleteRequestClass{objc.GetClass("NSBatchDeleteRequest")}
	})
	return batchDeleteRequestClass
}

type _BatchDeleteRequestClass struct {
	class objc.Class
}

// An interface definition for the [BatchDeleteRequest] class.
type IBatchDeleteRequest interface {
	IPersistentStoreRequest
}

// A request that deletes objects in the SQLite persistent store without loading them into memory.
//
// — available only when using a SQLite persistent store — deletes managed objects at the SQL level of the persistent store. This request is quicker and more efficient than using a context to fetch a large number of objects into memory, delete them, and then save those deletions back to the store. You create a request using an instance of that identifies the objects to delete. Alternatively, you can provide an array of identifiers from specific objects of the same entity type; mixing entity types results in an error when you execute the request. doesn’t automatically merge a request’s deletions because they happen at the SQL level. Subsequently, you must remove any deleted objects from memory after the request finishes. To determine the objects a request deletes, configure it to return the of each deleted object and use those identifiers to update your contexts, as the following example shows: Alternatively, you can use persistent history tracking to make your contexts aware of changes that happen at the persistent store level. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest
type BatchDeleteRequest struct {
	PersistentStoreRequest
}

// BatchDeleteRequestFrom constructs a [BatchDeleteRequest] from an unsafe.Pointer.
//
// A request that deletes objects in the SQLite persistent store without loading them into memory.
func BatchDeleteRequestFrom(ptr unsafe.Pointer) BatchDeleteRequest {
	return BatchDeleteRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchDeleteRequestClass) Alloc() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchDeleteRequestClass) New() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchDeleteRequest) Init() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchDeleteRequest) Autorelease() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchDeleteRequest creates a new BatchDeleteRequest instance.
func NewBatchDeleteRequest() BatchDeleteRequest {
	return getBatchDeleteRequestClass().New()
}


// The type of result the request provides when it executes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/resultType
func (b_ BatchDeleteRequest) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resultType"))
	return rv
}


// SetResultType sets the value of the resultType property.
// The type of result the request provides when it executes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest/resultType
func (b_ BatchDeleteRequest) SetResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultType:"), value)
}


