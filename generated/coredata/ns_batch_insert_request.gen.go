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
	// properties:
	DictionaryHandler() bool /* primitive/slice/pointer. */
	SetDictionaryHandler(value bool /* primitive/slice/pointer. */)
	Entity() IEntityDescription
	SetEntity(value IEntityDescription)
	EntityName() string /* primitive/slice/pointer. */
	SetEntityName(value string /* primitive/slice/pointer. */)
	ManagedObjectHandler() bool /* primitive/slice/pointer. */
	SetManagedObjectHandler(value bool /* primitive/slice/pointer. */)
	ObjectsToInsert() string /* primitive/slice/pointer. */
	SetObjectsToInsert(value string /* primitive/slice/pointer. */)
	ResultType() BatchInsertRequestResultType
	SetResultType(value BatchInsertRequestResultType)
	// methods:
}

// A request to insert a batch of data in a persistent store.


// A request to insert a batch of data in a persistent store.
//
// [Full Topic]
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



// Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entity:managedObjectHandler:)
func NewBatchInsertRequestWithEntityManagedObjectHandler(entity IEntityDescription, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntity:managedObjectHandler:"), entity, handler)
	rv.Autorelease()
	return rv
}


// Creates a batch-insertion request for a named managed entity, and specifies a closure that provides data dictionaries for insertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:dictionaryHandler:)
func NewBatchInsertRequestWithEntityNameDictionaryHandler(entityName string /* primitive/slice/pointer. */, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntityName:dictionaryHandler:"), objc.String(entityName), handler)
	rv.Autorelease()
	return rv
}


// Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/init(entityName:managedObjectHandler:)
func NewBatchInsertRequestWithEntityNameManagedObjectHandler(entityName string /* primitive/slice/pointer. */, handler unsafe.Pointer) BatchInsertRequest {
	instance := getBatchInsertRequestClass().Alloc()
	rv := objc.Send[BatchInsertRequest](instance.ID, objc.Sel("initWithEntityName:managedObjectHandler:"), objc.String(entityName), handler)
	rv.Autorelease()
	return rv
}



// Creates a batch-insertion request for a named managed entity, and specifies a closure that provides data dictionaries for insertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequest/batchInsertRequestWithEntityName:dictionaryHandler:
func (bc _BatchInsertRequestClass) BatchInsertRequestWithEntityNameDictionaryHandler(entityName string /* primitive/slice/pointer. */, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("batchInsertRequestWithEntityName:dictionaryHandler:"), objc.String(entityName), handler)
	return rv
}


// A closure that provides a dictionary for your app to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/dictionaryhandler
func (b_ BatchInsertRequest) DictionaryHandler() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("dictionaryHandler"))
	return rv
}


// A closure that provides a dictionary for your app to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/dictionaryhandler
func (b_ BatchInsertRequest) SetDictionaryHandler(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDictionaryHandler:"), value)
}


// The managed entity to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/entity
func (b_ BatchInsertRequest) Entity() IEntityDescription {
	rv := objc.Send[EntityDescription](b_.ID, objc.Sel("entity"))
	return rv
}


// The managed entity to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/entity
func (b_ BatchInsertRequest) SetEntity(value IEntityDescription) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEntity:"), value)
}


// The name of the managed entity to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/entityname
func (b_ BatchInsertRequest) EntityName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("entityName"))
	return rv
}


// The name of the managed entity to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/entityname
func (b_ BatchInsertRequest) SetEntityName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEntityName:"), objc.String(value))
}


// A closure that provides a managed object for your app to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/managedobjecthandler
func (b_ BatchInsertRequest) ManagedObjectHandler() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("managedObjectHandler"))
	return rv
}


// A closure that provides a managed object for your app to insert data into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/managedobjecthandler
func (b_ BatchInsertRequest) SetManagedObjectHandler(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setManagedObjectHandler:"), value)
}


// An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/objectstoinsert
func (b_ BatchInsertRequest) ObjectsToInsert() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("objectsToInsert"))
	return rv
}


// An array of dictionaries that represents the objects to insert with the keys as attribute names and their assigned values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/objectstoinsert
func (b_ BatchInsertRequest) SetObjectsToInsert(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObjectsToInsert:"), objc.String(value))
}


// The type of result that Core Data returns from this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/resulttype
func (b_ BatchInsertRequest) ResultType() BatchInsertRequestResultType {
	rv := objc.Send[BatchInsertRequestResultType](b_.ID, objc.Sel("resultType"))
	return rv
}


// The type of result that Core Data returns from this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchinsertrequest/resulttype
func (b_ BatchInsertRequest) SetResultType(value BatchInsertRequestResultType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultType:"), value)
}


