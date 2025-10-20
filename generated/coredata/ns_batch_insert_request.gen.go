// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchInsertRequest] class.
var (
	BatchInsertRequestClass     _BatchInsertRequestClass
	BatchInsertRequestClassOnce sync.Once
)

func getBatchInsertRequestClass() _BatchInsertRequestClass {
	BatchInsertRequestClassOnce.Do(func() {
		BatchInsertRequestClass = _BatchInsertRequestClass{objc.GetClass("NSBatchInsertRequest")}
	})
	return BatchInsertRequestClass
}

type _BatchInsertRequestClass struct {
	class objc.Class
}

// An interface definition for the [BatchInsertRequest] class.
type IBatchInsertRequest interface {
	IPersistentStoreRequest
}

// A request to insert a batch of data in a persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest
type BatchInsertRequest struct {
	PersistentStoreRequest
}

// BatchInsertRequestFrom constructs a [BatchInsertRequest] from an unsafe.Pointer.
//
// A request to insert a batch of data in a persistent store.
func BatchInsertRequestFrom(ptr unsafe.Pointer) BatchInsertRequest {
	return BatchInsertRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchInsertRequestClass) Alloc() BatchInsertRequest {
	rv := objc.Send[BatchInsertRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchInsertRequestClass) New() BatchInsertRequest {
	rv := objc.Send[BatchInsertRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchInsertRequest) Init() BatchInsertRequest {
	rv := objc.Send[BatchInsertRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchInsertRequest) Autorelease() BatchInsertRequest {
	rv := objc.Send[BatchInsertRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchInsertRequest creates a new BatchInsertRequest instance.
func NewBatchInsertRequest() BatchInsertRequest {
	return getBatchInsertRequestClass().New()
}


// Creates a batch-insertion request for a named managed entity, and specifies a closure that provides data dictionaries for insertion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:dictionaryHandler:)
func NewBatchInsertRequestWithEntityNameDictionaryHandler(entityName string, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntityName:dictionaryHandler:"), objc.String(entityName), handler)
	rv.Autorelease()
	return rv
}

// Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:managedObjectHandler:)
func NewBatchInsertRequestWithEntityNameManagedObjectHandler(entityName string, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntityName:managedObjectHandler:"), objc.String(entityName), handler)
	rv.Autorelease()
	return rv
}

// Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entity:managedObjectHandler:)
func NewBatchInsertRequestWithEntityManagedObjectHandler(entity unsafe.Pointer, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntity:managedObjectHandler:"), entity, handler)
	rv.Autorelease()
	return rv
}


// Creates a batch-insertion request for a named managed entity, and specifies a closure that provides data dictionaries for insertion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/batchInsertRequestWithEntityName:dictionaryHandler:
func (bc _BatchInsertRequestClass) BatchInsertRequestWithEntityNameDictionaryHandler(entityName string, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("batchInsertRequestWithEntityName:dictionaryHandler:"), objc.String(entityName), handler)
	return rv
}


