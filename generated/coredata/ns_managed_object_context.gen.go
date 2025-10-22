// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AssignObjectToPersistentStore(object objectivec.IObject, store IPersistentStore)
	CountForFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) uint
	DetectConflictsForObject(object IManagedObject)
	ExecuteRequestError(request IPersistentStoreRequest, error_ unsafe.Pointer) PersistentStoreResult
	ExecuteFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) foundation.Array
	InsertObject(object IManagedObject)
	Lock()
	MergeChangesFromContextDidSaveNotification(notification foundation.INotification)
	ObjectWithID(objectID IManagedObjectID) ManagedObject
	ObtainPermanentIDsForObjectsError(objects []ManagedObject, error_ unsafe.Pointer) bool
	PerformBlock(block unsafe.Pointer)
	PerformBlockAndWait(block unsafe.Pointer)
	ProcessPendingChanges()
	Redo()
	RefreshObjectMergeChanges(object IManagedObject, flag bool)
	RefreshAllObjects()
	ObjectRegisteredForID(objectID IManagedObjectID) ManagedObject
	Reset()
	Rollback()
	Save(error_ unsafe.Pointer) bool
	SetQueryGenerationFromTokenError(generation IQueryGenerationToken, error_ unsafe.Pointer) bool
	ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool
	TryLock() bool
	Undo()
	Unlock()
	AutomaticallyMergesChangesFromParent() bool
	SetAutomaticallyMergesChangesFromParent(value bool)
	ConcurrencyType() ManagedObjectContextConcurrencyType
	DeletedObjects() unsafe.Pointer
	HasChanges() bool
	MergePolicy() objc.ID
	SetMergePolicy(value objc.ID)
	Name() string
	SetName(value string)
	PersistentStoreCoordinator() NSPersistentStoreCoordinator
	SetPersistentStoreCoordinator(value IPersistentStoreCoordinator)
	PropagatesDeletesAtEndOfEvent() bool
	SetPropagatesDeletesAtEndOfEvent(value bool)
	QueryGenerationToken() NSQueryGenerationToken
	RegisteredObjects() unsafe.Pointer
	RetainsRegisteredObjects() bool
	SetRetainsRegisteredObjects(value bool)
	ShouldDeleteInaccessibleFaults() bool
	SetShouldDeleteInaccessibleFaults(value bool)
	StalenessInterval() foundation.TimeInterval
	SetStalenessInterval(value foundation.TimeInterval)
	TransactionAuthor() string
	SetTransactionAuthor(value string)
	UpdatedObjects() unsafe.Pointer
	UserInfo() foundation.MutableDictionary
	NSDeletedObjectsKey() string
	NSInsertedObjectsKey() string
	NSInvalidatedAllObjectsKey() string
	NSInvalidatedObjectsKey() string
	InsertedObjects() NSManagedObject
	SetInsertedObjects(value IManagedObject)
	Parent() NSManagedObjectContext
	SetParent(value IManagedObjectContext)
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.IUndoManager)
	NSManagedObjectContextQueryGenerationKey() string
	NSRefreshedObjectsKey() string
	NSUpdatedObjectsKey() string
}

// An object space to manipulate and track changes to managed objects.
//
// A context consists of a group of related model objects that represent an internally consistent view of one or more persistent stores. Changes to managed objects remain in memory in the associated context until Core Data saves that context to one or more persistent stores. A single managed object instance exists in one and only one context, but multiple copies of an object can exist in different contexts. Therefore, an object is unique to a particular context.


// An object space to manipulate and track changes to managed objects.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/init(concurrencyType:)

func NewManagedObjectContextWithConcurrencyType(ct ManagedObjectContextConcurrencyType) ManagedObjectContext {
	instance := getManagedObjectContextClass().Alloc()
	rv := objc.Send[ManagedObjectContext](instance.ID, objc.Sel("initWithConcurrencyType:"), ct)
	rv.Autorelease()
	return rv
}



// Handles changes from other processes or from a serialized state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromRemoteContextSave:into:)

func (mc _ManagedObjectContextClass) MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData objectivec.IObject, contexts []ManagedObjectContext) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mergeChangesFromRemoteContextSave:intoContexts:"), changeNotificationData, contexts)
}



// Specifies the store in which a newly inserted object will be saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/assign(_:to:)

func (m_ ManagedObjectContext) AssignObjectToPersistentStore(object objectivec.IObject, store IPersistentStore) {
	objc.Send[objc.ID](m_.ID, objc.Sel("assignObject:toPersistentStore:"), object, store)
}



// Returns the number of objects the specified request fetches when it executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/count(for:)-93zbm

func (m_ ManagedObjectContext) CountForFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("countForFetchRequest:error:"), request, error_)
	return rv
}



// Marks an object for conflict detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/detectConflicts(for:)

func (m_ ManagedObjectContext) DetectConflictsForObject(object IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("detectConflictsForObject:"), object)
}



// Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/execute(_:)

func (m_ ManagedObjectContext) ExecuteRequestError(request IPersistentStoreRequest, error_ unsafe.Pointer) PersistentStoreResult {
	rv := objc.Send[PersistentStoreResult](m_.ID, objc.Sel("executeRequest:error:"), request, error_)
	return rv
}



// Returns an array of objects that meet the criteria of the specified fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/executeFetchRequest:error:

func (m_ ManagedObjectContext) ExecuteFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](m_.ID, objc.Sel("executeFetchRequest:error:"), request, error_)
	return rv
}



// Registers an object to be inserted in the context’s persistent store the next time changes are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insert(_:)

func (m_ ManagedObjectContext) InsertObject(object IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:"), object)
}



// Attempts to acquire a lock on the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/lock()

func (m_ ManagedObjectContext) Lock() {
	objc.Send[objc.ID](m_.ID, objc.Sel("lock"))
}



// Merges the changes specified in a given notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromContextDidSave:)

func (m_ ManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification foundation.INotification) {
	objc.Send[objc.ID](m_.ID, objc.Sel("mergeChangesFromContextDidSaveNotification:"), notification)
}



// Returns either an existing object from the context or a fault that represents that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/object(with:)

func (m_ ManagedObjectContext) ObjectWithID(objectID IManagedObjectID) ManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("objectWithID:"), objectID)
	return rv
}



// Allows a context that has registered as an observer of a value to be notified of a change to that value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/observeValue(forKeyPath:of:change:context:)

func (m_ ManagedObjectContext) ObserveValueForKeyPathOfObjectChangeContext(keyPath string, object objectivec.IObject, change unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), objc.String(keyPath), object, change, context)
}



// Converts to permanent IDs the object IDs of the objects in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/obtainPermanentIDs(for:)

func (m_ ManagedObjectContext) ObtainPermanentIDsForObjectsError(objects []ManagedObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("obtainPermanentIDsForObjects:error:"), objects, error_)
	return rv
}



// Asynchronously performs the specified closure on the context’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/perform(_:)

func (m_ ManagedObjectContext) PerformBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performBlock:"), block)
}



// Synchronously performs the specified closure on the context’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/performAndWait(_:)-ypye

func (m_ ManagedObjectContext) PerformBlockAndWait(block unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performBlockAndWait:"), block)
}



// Forces the context to process changes to the object graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/processPendingChanges()

func (m_ ManagedObjectContext) ProcessPendingChanges() {
	objc.Send[objc.ID](m_.ID, objc.Sel("processPendingChanges"))
}



// Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/redo()

func (m_ ManagedObjectContext) Redo() {
	objc.Send[objc.ID](m_.ID, objc.Sel("redo"))
}



// Updates the persistent properties of a managed object to use the latest values from the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refresh(_:mergeChanges:)

func (m_ ManagedObjectContext) RefreshObjectMergeChanges(object IManagedObject, flag bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshObject:mergeChanges:"), object, flag)
}



// Refreshes all of the registered managed objects in the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refreshAllObjects()

func (m_ ManagedObjectContext) RefreshAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshAllObjects"))
}



// Returns an object that exists in the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObject(for:)

func (m_ ManagedObjectContext) ObjectRegisteredForID(objectID IManagedObjectID) ManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("objectRegisteredForID:"), objectID)
	return rv
}



// Returns the context to its base state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/reset()

func (m_ ManagedObjectContext) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}



// Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/rollback()

func (m_ ManagedObjectContext) Rollback() {
	objc.Send[objc.ID](m_.ID, objc.Sel("rollback"))
}



// Attempts to commit unsaved changes to registered objects to the context’s parent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()

func (m_ ManagedObjectContext) Save(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("save:"), error_)
	return rv
}



// Sets the query generation this context should use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/setQueryGenerationFrom(_:)

func (m_ ManagedObjectContext) SetQueryGenerationFromTokenError(generation IQueryGenerationToken, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setQueryGenerationFromToken:error:"), generation, error_)
	return rv
}



// Creates a log of the inaccessible fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldHandleInaccessibleFault(_:for:triggeredByProperty:)

func (m_ ManagedObjectContext) ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldHandleInaccessibleFault:forObjectID:triggeredByProperty:"), fault, oid, property)
	return rv
}



// Attempts to acquire a lock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/tryLock()

func (m_ ManagedObjectContext) TryLock() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("tryLock"))
	return rv
}



// Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undo()

func (m_ ManagedObjectContext) Undo() {
	objc.Send[objc.ID](m_.ID, objc.Sel("undo"))
}



// Relinquishes a previously acquired lock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/unlock()

func (m_ ManagedObjectContext) Unlock() {
	objc.Send[objc.ID](m_.ID, objc.Sel("unlock"))
}


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/automaticallyMergesChangesFromParent

func (m_ ManagedObjectContext) AutomaticallyMergesChangesFromParent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/automaticallyMergesChangesFromParent

func (m_ ManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}


// The concurrency type for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/concurrencyType-swift.property

func (m_ ManagedObjectContext) ConcurrencyType() ManagedObjectContextConcurrencyType {
	rv := objc.Send[ManagedObjectContextConcurrencyType](m_.ID, objc.Sel("concurrencyType"))
	return rv
}


// The set of objects that will be removed from their persistent store during the next save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/deletedObjects

func (m_ ManagedObjectContext) DeletedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deletedObjects"))
	return rv
}


// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/hasChanges

func (m_ ManagedObjectContext) HasChanges() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasChanges"))
	return rv
}


// The merge policy of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergePolicy

func (m_ ManagedObjectContext) MergePolicy() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mergePolicy"))
	return rv
}


// The merge policy of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergePolicy

func (m_ ManagedObjectContext) SetMergePolicy(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMergePolicy:"), value)
}


// The developer-provided name of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/name

func (m_ ManagedObjectContext) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The developer-provided name of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/name

func (m_ ManagedObjectContext) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}


// The persistent store coordinator of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/persistentStoreCoordinator

func (m_ ManagedObjectContext) PersistentStoreCoordinator() NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](m_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}


// The persistent store coordinator of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/persistentStoreCoordinator

func (m_ ManagedObjectContext) SetPersistentStoreCoordinator(value IPersistentStoreCoordinator) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPersistentStoreCoordinator:"), value)
}


// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/propagatesDeletesAtEndOfEvent

func (m_ ManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}


// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/propagatesDeletesAtEndOfEvent

func (m_ ManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropagatesDeletesAtEndOfEvent:"), value)
}


// Returns the token associated with the query generation currently in use by this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/queryGenerationToken

func (m_ ManagedObjectContext) QueryGenerationToken() NSQueryGenerationToken {
	rv := objc.Send[NSQueryGenerationToken](m_.ID, objc.Sel("queryGenerationToken"))
	return rv
}


// The set of registered managed objects in the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/registeredObjects

func (m_ ManagedObjectContext) RegisteredObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("registeredObjects"))
	return rv
}


// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/retainsRegisteredObjects

func (m_ ManagedObjectContext) RetainsRegisteredObjects() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("retainsRegisteredObjects"))
	return rv
}


// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/retainsRegisteredObjects

func (m_ ManagedObjectContext) SetRetainsRegisteredObjects(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRetainsRegisteredObjects:"), value)
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults

func (m_ ManagedObjectContext) ShouldDeleteInaccessibleFaults() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults

func (m_ ManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval

func (m_ ManagedObjectContext) StalenessInterval() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("stalenessInterval"))
	return rv
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval

func (m_ ManagedObjectContext) SetStalenessInterval(value foundation.TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStalenessInterval:"), value)
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/transactionAuthor

func (m_ ManagedObjectContext) TransactionAuthor() string {
	rv := objc.Send[string](m_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/transactionAuthor

func (m_ ManagedObjectContext) SetTransactionAuthor(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}


// The set of objects registered with the context that have uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/updatedObjects

func (m_ ManagedObjectContext) UpdatedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("updatedObjects"))
	return rv
}


// The user information for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/userInfo

func (m_ ManagedObjectContext) UserInfo() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](m_.ID, objc.Sel("userInfo"))
	return rv
}


// A key for the set of objects that were marked for deletion during the previous event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsdeletedobjectskey

func (m_ ManagedObjectContext) NSDeletedObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSDeletedObjectsKey"))
	return rv
}


// A key for the set of objects that were inserted into the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinsertedobjectskey

func (m_ ManagedObjectContext) NSInsertedObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInsertedObjectsKey"))
	return rv
}


// A key that specifies that all objects in the context have been invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinvalidatedallobjectskey

func (m_ ManagedObjectContext) NSInvalidatedAllObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInvalidatedAllObjectsKey"))
	return rv
}


// A key for the set of objects that were invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinvalidatedobjectskey

func (m_ ManagedObjectContext) NSInvalidatedObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInvalidatedObjectsKey"))
	return rv
}


// The set of objects that have been inserted into the context but not yet saved in a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/insertedobjects

func (m_ ManagedObjectContext) InsertedObjects() NSManagedObject {
	rv := objc.Send[NSManagedObject](m_.ID, objc.Sel("insertedObjects"))
	return rv
}


// The set of objects that have been inserted into the context but not yet saved in a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/insertedobjects

func (m_ ManagedObjectContext) SetInsertedObjects(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInsertedObjects:"), value)
}


// The parent of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/parent

func (m_ ManagedObjectContext) Parent() NSManagedObjectContext {
	rv := objc.Send[NSManagedObjectContext](m_.ID, objc.Sel("parent"))
	return rv
}


// The parent of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/parent

func (m_ ManagedObjectContext) SetParent(value IManagedObjectContext) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParent:"), value)
}


// The object that provides undo support for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/undomanager

func (m_ ManagedObjectContext) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](m_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that provides undo support for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/undomanager

func (m_ ManagedObjectContext) SetUndoManager(value foundation.IUndoManager) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUndoManager:"), value)
}


// Constant used to reference the query generation token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontextquerygenerationkey

func (m_ ManagedObjectContext) NSManagedObjectContextQueryGenerationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSManagedObjectContextQueryGenerationKey"))
	return rv
}


// A key for the set of objects that were refreshed but were not dirtied in the scope of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrefreshedobjectskey

func (m_ ManagedObjectContext) NSRefreshedObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSRefreshedObjectsKey"))
	return rv
}


// A key for the set of objects that were updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsupdatedobjectskey

func (m_ ManagedObjectContext) NSUpdatedObjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSUpdatedObjectsKey"))
	return rv
}


