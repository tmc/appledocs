// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BatchUpdateRequest] class.
var (
	BatchUpdateRequestClass     _BatchUpdateRequestClass
	BatchUpdateRequestClassOnce sync.Once
)

func getBatchUpdateRequestClass() _BatchUpdateRequestClass {
	BatchUpdateRequestClassOnce.Do(func() {
		BatchUpdateRequestClass = _BatchUpdateRequestClass{objc.GetClass("NSBatchUpdateRequest")}
	})
	return BatchUpdateRequestClass
}

type _BatchUpdateRequestClass struct {
	class objc.Class
}

// An interface definition for the [BatchUpdateRequest] class.
type IBatchUpdateRequest interface {
	IPersistentStoreRequest
	Entity() IEntityDescription
	EntityName() string
	IncludesSubentities() bool
	SetIncludesSubentities(value bool)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.Predicate)
	PropertiesToUpdate() objc.ID
	SetPropertiesToUpdate(value objc.ID)
	ResultType() unsafe.Pointer
	SetResultType(value unsafe.Pointer)
}

// A request to Core Data to do a batch update of data in a persistent store without loading any data into memory.


// A request to Core Data to do a batch update of data in a persistent store without loading any data into memory.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getBatchUpdateRequestClass().New()
}



// Creates a batch-update request for a managed entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entity:)
func NewBatchUpdateRequestWithEntity(entity IEntityDescription) BatchUpdateRequest {
	instance := getBatchUpdateRequestClass().Alloc()
	rv := objc.Send[BatchUpdateRequest](instance.ID, objc.Sel("initWithEntity:"), entity)
	rv.Autorelease()
	return rv
}


// Creates a batch-update request for a named managed entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/init(entityName:)
func NewBatchUpdateRequestWithEntityName(entityName string) BatchUpdateRequest {
	instance := getBatchUpdateRequestClass().Alloc()
	rv := objc.Send[BatchUpdateRequest](instance.ID, objc.Sel("initWithEntityName:"), objc.String(entityName))
	rv.Autorelease()
	return rv
}



// Creates a batch-update request for a named managed entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/batchUpdateRequestWithEntityName:
func (bc _BatchUpdateRequestClass) BatchUpdateRequestWithEntityName(entityName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("batchUpdateRequestWithEntityName:"), objc.String(entityName))
	return rv
}


// The managed entity to update data for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/entity
func (b_ BatchUpdateRequest) Entity() IEntityDescription {
	rv := objc.Send[EntityDescription](b_.ID, objc.Sel("entity"))
	return rv
}


// The name of the managed entity to update data for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/entityName
func (b_ BatchUpdateRequest) EntityName() string {
	rv := objc.Send[string](b_.ID, objc.Sel("entityName"))
	return rv
}


// A Boolean value that indicates whether to update subentities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/includesSubentities
func (b_ BatchUpdateRequest) IncludesSubentities() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesSubentities"))
	return rv
}


// A Boolean value that indicates whether to update subentities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/includesSubentities
func (b_ BatchUpdateRequest) SetIncludesSubentities(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesSubentities:"), value)
}


// A predicate that identifies the objects to update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/predicate
func (b_ BatchUpdateRequest) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](b_.ID, objc.Sel("predicate"))
	return rv
}


// A predicate that identifies the objects to update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/predicate
func (b_ BatchUpdateRequest) SetPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPredicate:"), value)
}


// A dictionary of property description pairs that describe the updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/propertiesToUpdate
func (b_ BatchUpdateRequest) PropertiesToUpdate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("propertiesToUpdate"))
	return rv
}


// A dictionary of property description pairs that describe the updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/propertiesToUpdate
func (b_ BatchUpdateRequest) SetPropertiesToUpdate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPropertiesToUpdate:"), value)
}


// The type of result that Core Data returns from the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/resultType
func (b_ BatchUpdateRequest) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resultType"))
	return rv
}


// The type of result that Core Data returns from the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateRequest/resultType
func (b_ BatchUpdateRequest) SetResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultType:"), value)
}


