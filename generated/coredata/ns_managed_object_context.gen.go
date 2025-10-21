// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ManagedObjectContext] class.
var (
	ManagedObjectContextClass     _ManagedObjectContextClass
	ManagedObjectContextClassOnce sync.Once
)

func getManagedObjectContextClass() _ManagedObjectContextClass {
	ManagedObjectContextClassOnce.Do(func() {
		ManagedObjectContextClass = _ManagedObjectContextClass{objc.GetClass("NSManagedObjectContext")}
	})
	return ManagedObjectContextClass
}

type _ManagedObjectContextClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectContext] class.
type IManagedObjectContext interface {
	objectivec.IObject
	AssignObjectToPersistentStore(object objc.ID, store unsafe.Pointer)
	CountForFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) uint
	DetectConflictsForObject(object unsafe.Pointer)
	ExecuteRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	ExecuteFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	InsertObject(object unsafe.Pointer)
	Lock()
	MergeChangesFromContextDidSaveNotification(notification unsafe.Pointer)
	ObjectWithID(objectID unsafe.Pointer) unsafe.Pointer
	ObtainPermanentIDsForObjectsError(objects unsafe.Pointer, error_ unsafe.Pointer) bool
	PerformBlock(block unsafe.Pointer)
	PerformBlockAndWait(block unsafe.Pointer)
	ProcessPendingChanges()
	Redo()
	RefreshObjectMergeChanges(object unsafe.Pointer, flag bool)
	RefreshAllObjects()
	ObjectRegisteredForID(objectID unsafe.Pointer) unsafe.Pointer
	Reset()
	Rollback()
	Save(error_ unsafe.Pointer) bool
	SetQueryGenerationFromTokenError(generation unsafe.Pointer, error_ unsafe.Pointer) bool
	ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault unsafe.Pointer, oid unsafe.Pointer, property unsafe.Pointer) bool
	TryLock() bool
	Undo()
	Unlock()
}

// An object space to manipulate and track changes to managed objects.
//
// A context consists of a group of related model objects that represent an internally consistent view of one or more persistent stores. Changes to managed objects remain in memory in the associated context until Core Data saves that context to one or more persistent stores. A single managed object instance exists in one and only one context, but multiple copies of an object can exist in different contexts. Therefore, an object is unique to a particular context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext
type ManagedObjectContext struct {
	objectivec.Object
}

// ManagedObjectContextFrom constructs a [ManagedObjectContext] from an unsafe.Pointer.
//
// An object space to manipulate and track changes to managed objects.
func ManagedObjectContextFrom(ptr unsafe.Pointer) ManagedObjectContext {
	return ManagedObjectContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectContextClass) Alloc() ManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectContextClass) New() ManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObjectContext) Init() ManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObjectContext) Autorelease() ManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObjectContext creates a new ManagedObjectContext instance.
func NewManagedObjectContext() ManagedObjectContext {
	return getManagedObjectContextClass().New()
}




// Creates a context that uses the specified concurrency type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/init(concurrencyType:)
func NewManagedObjectContextWithConcurrencyType(ct unsafe.Pointer) ManagedObjectContext {
	instance := getManagedObjectContextClass().Alloc()
	rv := objc.Send[ManagedObjectContext](instance.ID, objc.Sel("initWithConcurrencyType:"), ct)
	rv.Autorelease()
	return rv
}


// Handles changes from other processes or from a serialized state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromRemoteContextSave:into:)
func (mc _ManagedObjectContextClass) MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData objc.ID, contexts unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mergeChangesFromRemoteContextSave:intoContexts:"), changeNotificationData, contexts)
}

// Specifies the store in which a newly inserted object will be saved.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/assign(_:to:)
func (m_ ManagedObjectContext) AssignObjectToPersistentStore(object objc.ID, store unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("assignObject:toPersistentStore:"), object, store)
}

// Returns the number of objects the specified request fetches when it executes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/count(for:)-93zbm
func (m_ ManagedObjectContext) CountForFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("countForFetchRequest:error:"), request, error_)
	return rv
}

// Marks an object for conflict detection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/detectConflicts(for:)
func (m_ ManagedObjectContext) DetectConflictsForObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("detectConflictsForObject:"), object)
}

// Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/execute(_:)
func (m_ ManagedObjectContext) ExecuteRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("executeRequest:error:"), request, error_)
	return rv
}

// Returns an array of objects that meet the criteria of the specified fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/executeFetchRequest:error:
func (m_ ManagedObjectContext) ExecuteFetchRequestError(request unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("executeFetchRequest:error:"), request, error_)
	return rv
}

// Registers an object to be inserted in the context’s persistent store the next time changes are saved.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insert(_:)
func (m_ ManagedObjectContext) InsertObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:"), object)
}

// Attempts to acquire a lock on the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/lock()
func (m_ ManagedObjectContext) Lock() {
	objc.Send[objc.ID](m_.ID, objc.Sel("lock"))
}

// Merges the changes specified in a given notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromContextDidSave:)
func (m_ ManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("mergeChangesFromContextDidSaveNotification:"), notification)
}

// Returns either an existing object from the context or a fault that represents that object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/object(with:)
func (m_ ManagedObjectContext) ObjectWithID(objectID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectWithID:"), objectID)
	return rv
}

// Allows a context that has registered as an observer of a value to be notified of a change to that value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/observeValue(forKeyPath:of:change:context:)
func (m_ ManagedObjectContext) ObserveValueForKeyPathOfObjectChangeContext(keyPath string, object objc.ID, change unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), objc.String(keyPath), object, change, context)
}

// Converts to permanent IDs the object IDs of the objects in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/obtainPermanentIDs(for:)
func (m_ ManagedObjectContext) ObtainPermanentIDsForObjectsError(objects unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("obtainPermanentIDsForObjects:error:"), objects, error_)
	return rv
}

// Asynchronously performs the specified closure on the context’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/perform(_:)
func (m_ ManagedObjectContext) PerformBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performBlock:"), block)
}

// Synchronously performs the specified closure on the context’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/performAndWait(_:)-ypye
func (m_ ManagedObjectContext) PerformBlockAndWait(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performBlockAndWait:"), block)
}

// Forces the context to process changes to the object graph.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/processPendingChanges()
func (m_ ManagedObjectContext) ProcessPendingChanges() {
	objc.Send[objc.ID](m_.ID, objc.Sel("processPendingChanges"))
}

// Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/redo()
func (m_ ManagedObjectContext) Redo() {
	objc.Send[objc.ID](m_.ID, objc.Sel("redo"))
}

// Updates the persistent properties of a managed object to use the latest values from the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refresh(_:mergeChanges:)
func (m_ ManagedObjectContext) RefreshObjectMergeChanges(object unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshObject:mergeChanges:"), object, flag)
}

// Refreshes all of the registered managed objects in the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refreshAllObjects()
func (m_ ManagedObjectContext) RefreshAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshAllObjects"))
}

// Returns an object that exists in the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObject(for:)
func (m_ ManagedObjectContext) ObjectRegisteredForID(objectID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectRegisteredForID:"), objectID)
	return rv
}

// Returns the context to its base state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/reset()
func (m_ ManagedObjectContext) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/rollback()
func (m_ ManagedObjectContext) Rollback() {
	objc.Send[objc.ID](m_.ID, objc.Sel("rollback"))
}

// Attempts to commit unsaved changes to registered objects to the context’s parent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()
func (m_ ManagedObjectContext) Save(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("save:"), error_)
	return rv
}

// Sets the query generation this context should use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/setQueryGenerationFrom(_:)
func (m_ ManagedObjectContext) SetQueryGenerationFromTokenError(generation unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setQueryGenerationFromToken:error:"), generation, error_)
	return rv
}

// Creates a log of the inaccessible fault.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldHandleInaccessibleFault(_:for:triggeredByProperty:)
func (m_ ManagedObjectContext) ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault unsafe.Pointer, oid unsafe.Pointer, property unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldHandleInaccessibleFault:forObjectID:triggeredByProperty:"), fault, oid, property)
	return rv
}

// Attempts to acquire a lock.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/tryLock()
func (m_ ManagedObjectContext) TryLock() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("tryLock"))
	return rv
}

// Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undo()
func (m_ ManagedObjectContext) Undo() {
	objc.Send[objc.ID](m_.ID, objc.Sel("undo"))
}

// Relinquishes a previously acquired lock.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/unlock()
func (m_ ManagedObjectContext) Unlock() {
	objc.Send[objc.ID](m_.ID, objc.Sel("unlock"))
}

// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/automaticallyMergesChangesFromParent
func (m_ ManagedObjectContext) AutomaticallyMergesChangesFromParent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}


// SetAutomaticallyMergesChangesFromParent sets the value of the automaticallyMergesChangesFromParent property.
// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/automaticallyMergesChangesFromParent
func (m_ ManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}

// The concurrency type for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/concurrencyType-swift.property
func (m_ ManagedObjectContext) ConcurrencyType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("concurrencyType"))
	return rv
}

// The set of objects that will be removed from their persistent store during the next save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/deletedObjects
func (m_ ManagedObjectContext) DeletedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deletedObjects"))
	return rv
}

// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/hasChanges
func (m_ ManagedObjectContext) HasChanges() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasChanges"))
	return rv
}

// The merge policy of the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergePolicy
func (m_ ManagedObjectContext) MergePolicy() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mergePolicy"))
	return rv
}


// SetMergePolicy sets the value of the mergePolicy property.
// The merge policy of the context.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergePolicy
func (m_ ManagedObjectContext) SetMergePolicy(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMergePolicy:"), value)
}

// The developer-provided name of the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/name
func (m_ ManagedObjectContext) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The developer-provided name of the context.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/name
func (m_ ManagedObjectContext) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// The persistent store coordinator of the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/persistentStoreCoordinator
func (m_ ManagedObjectContext) PersistentStoreCoordinator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}


// SetPersistentStoreCoordinator sets the value of the persistentStoreCoordinator property.
// The persistent store coordinator of the context.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/persistentStoreCoordinator
func (m_ ManagedObjectContext) SetPersistentStoreCoordinator(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPersistentStoreCoordinator:"), value)
}

// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/propagatesDeletesAtEndOfEvent
func (m_ ManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}


// SetPropagatesDeletesAtEndOfEvent sets the value of the propagatesDeletesAtEndOfEvent property.
// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/propagatesDeletesAtEndOfEvent
func (m_ ManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropagatesDeletesAtEndOfEvent:"), value)
}

// Returns the token associated with the query generation currently in use by this context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/queryGenerationToken
func (m_ ManagedObjectContext) QueryGenerationToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("queryGenerationToken"))
	return rv
}

// The set of registered managed objects in the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObjects
func (m_ ManagedObjectContext) RegisteredObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("registeredObjects"))
	return rv
}

// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/retainsRegisteredObjects
func (m_ ManagedObjectContext) RetainsRegisteredObjects() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("retainsRegisteredObjects"))
	return rv
}


// SetRetainsRegisteredObjects sets the value of the retainsRegisteredObjects property.
// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/retainsRegisteredObjects
func (m_ ManagedObjectContext) SetRetainsRegisteredObjects(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRetainsRegisteredObjects:"), value)
}

// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults
func (m_ ManagedObjectContext) ShouldDeleteInaccessibleFaults() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}


// SetShouldDeleteInaccessibleFaults sets the value of the shouldDeleteInaccessibleFaults property.
// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults
func (m_ ManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}

// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval
func (m_ ManagedObjectContext) StalenessInterval() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("stalenessInterval"))
	return rv
}


// SetStalenessInterval sets the value of the stalenessInterval property.
// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval
func (m_ ManagedObjectContext) SetStalenessInterval(value TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStalenessInterval:"), value)
}

// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/transactionAuthor
func (m_ ManagedObjectContext) TransactionAuthor() string {
	rv := objc.Send[string](m_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// SetTransactionAuthor sets the value of the transactionAuthor property.
// The author for the context that is used as an identifier in persistent history transactions.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/transactionAuthor
func (m_ ManagedObjectContext) SetTransactionAuthor(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}

// The set of objects registered with the context that have uncommitted changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/updatedObjects
func (m_ ManagedObjectContext) UpdatedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("updatedObjects"))
	return rv
}

// The user information for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/userInfo
func (m_ ManagedObjectContext) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userInfo"))
	return rv
}


