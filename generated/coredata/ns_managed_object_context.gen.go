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
	// properties:
	DeletedObjects() unsafe.Pointer
	InsertedObjects() unsafe.Pointer
	ShouldDeleteInaccessibleFaults() bool /* primitive/slice/pointer. */
	SetShouldDeleteInaccessibleFaults(value bool /* primitive/slice/pointer. */)
	StalenessInterval() foundation.TimeInterval /* not a class type */
	SetStalenessInterval(value foundation.TimeInterval /* not a class type */)
	UndoManager() UndoManager /* not a class type */
	SetUndoManager(value UndoManager /* not a class type */)
	UpdatedObjects() unsafe.Pointer
	NSDeletedObjectsKey() string /* primitive/slice/pointer. */
	NSInsertedObjectsKey() string /* primitive/slice/pointer. */
	NSInvalidatedAllObjectsKey() string /* primitive/slice/pointer. */
	NSInvalidatedObjectsKey() string /* primitive/slice/pointer. */
	AutomaticallyMergesChangesFromParent() bool /* primitive/slice/pointer. */
	SetAutomaticallyMergesChangesFromParent(value bool /* primitive/slice/pointer. */)
	ConcurrencyType() ManagedObjectContextConcurrencyType /* not a class type */
	SetConcurrencyType(value ManagedObjectContextConcurrencyType /* not a class type */)
	HasChanges() bool /* primitive/slice/pointer. */
	SetHasChanges(value bool /* primitive/slice/pointer. */)
	MergePolicy() unsafe.Pointer
	SetMergePolicy(value unsafe.Pointer)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	Parent() IManagedObjectContext
	SetParent(value IManagedObjectContext)
	PersistentStoreCoordinator() IPersistentStoreCoordinator
	SetPersistentStoreCoordinator(value IPersistentStoreCoordinator)
	PropagatesDeletesAtEndOfEvent() bool /* primitive/slice/pointer. */
	SetPropagatesDeletesAtEndOfEvent(value bool /* primitive/slice/pointer. */)
	QueryGenerationToken() IQueryGenerationToken
	SetQueryGenerationToken(value IQueryGenerationToken)
	RegisteredObjects() IManagedObject
	SetRegisteredObjects(value IManagedObject)
	RetainsRegisteredObjects() bool /* primitive/slice/pointer. */
	SetRetainsRegisteredObjects(value bool /* primitive/slice/pointer. */)
	TransactionAuthor() string /* primitive/slice/pointer. */
	SetTransactionAuthor(value string /* primitive/slice/pointer. */)
	UserInfo() MutableDictionary /* not a class type */
	SetUserInfo(value MutableDictionary /* not a class type */)
	NSManagedObjectContextQueryGenerationKey() string /* primitive/slice/pointer. */
	NSRefreshedObjectsKey() string /* primitive/slice/pointer. */
	NSUpdatedObjectsKey() string /* primitive/slice/pointer. */
	// methods:
	AssignObjectToPersistentStore(object objectivec.IObject, store IPersistentStore)
	DeleteObject(object IManagedObject)
	DetectConflictsForObject(object IManagedObject)
	ExecuteFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) objc.IObject /* cross-framework: Array */
	InsertObject(object IManagedObject)
	MergeChangesFromContextDidSaveNotification(notification Notification /* not a class type */)
	ObjectWithID(objectID IManagedObjectID) IManagedObject
	ObtainPermanentIDsForObjectsError(objects []ManagedObject /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	PerformBlock(block unsafe.Pointer)
	ProcessPendingChanges()
	Redo()
	RefreshObjectMergeChanges(object IManagedObject, flag bool /* primitive/slice/pointer. */)
	Reset()
	Rollback()
	Save(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool /* primitive/slice/pointer. */
	Undo()
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



// Specifies the store in which a newly inserted object will be saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/assign(_:to:)
func (m_ ManagedObjectContext) AssignObjectToPersistentStore(object objectivec.IObject, store IPersistentStore) {
	objc.Send[objc.ID](m_.ID, objc.Sel("assignObject:toPersistentStore:"), object, store)
}


// Specifies an object that should be removed from its persistent store when changes are committed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/delete(_:)
func (m_ ManagedObjectContext) DeleteObject(object IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deleteObject:"), object)
}


// Marks an object for conflict detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/detectConflicts(for:)
func (m_ ManagedObjectContext) DetectConflictsForObject(object IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("detectConflictsForObject:"), object)
}


// Returns an array of objects that meet the criteria of the specified fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/executeFetchRequest:error:
func (m_ ManagedObjectContext) ExecuteFetchRequestError(request IFetchRequest, error_ unsafe.Pointer) objc.IObject /* cross-framework: Array */ {
	rv := objc.Send[Array](m_.ID, objc.Sel("executeFetchRequest:error:"), request, error_)
	return rv
}


// Registers an object to be inserted in the context’s persistent store the next time changes are saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insert(_:)
func (m_ ManagedObjectContext) InsertObject(object IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:"), object)
}


// Merges the changes specified in a given notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromContextDidSave:)
func (m_ ManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification Notification /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("mergeChangesFromContextDidSaveNotification:"), notification)
}


// Returns either an existing object from the context or a fault that represents that object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/object(with:)
func (m_ ManagedObjectContext) ObjectWithID(objectID IManagedObjectID) IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("objectWithID:"), objectID)
	return rv
}


// Allows a context that has registered as an observer of a value to be notified of a change to that value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/observeValue(forKeyPath:of:change:context:)
func (m_ ManagedObjectContext) ObserveValueForKeyPathOfObjectChangeContext(keyPath string /* primitive/slice/pointer. */, object objectivec.IObject, change foundation.IDictionary /* already interface */, context unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), objc.String(keyPath), object, change, context)
}


// Converts to permanent IDs the object IDs of the objects in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/obtainPermanentIDs(for:)
func (m_ ManagedObjectContext) ObtainPermanentIDsForObjectsError(objects []ManagedObject /* primitive/slice/pointer. */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
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
func (m_ ManagedObjectContext) RefreshObjectMergeChanges(object IManagedObject, flag bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshObject:mergeChanges:"), object, flag)
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
func (m_ ManagedObjectContext) Save(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("save:"), error_)
	return rv
}


// Creates a log of the inaccessible fault.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldHandleInaccessibleFault(_:for:triggeredByProperty:)
func (m_ ManagedObjectContext) ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldHandleInaccessibleFault:forObjectID:triggeredByProperty:"), fault, oid, property)
	return rv
}


// Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undo()
func (m_ ManagedObjectContext) Undo() {
	objc.Send[objc.ID](m_.ID, objc.Sel("undo"))
}


// The set of objects that will be removed from their persistent store during the next save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/deletedObjects
func (m_ ManagedObjectContext) DeletedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deletedObjects"))
	return rv
}


// The set of objects that have been inserted into the context but not yet saved in a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/insertedObjects
func (m_ ManagedObjectContext) InsertedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("insertedObjects"))
	return rv
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults
func (m_ ManagedObjectContext) ShouldDeleteInaccessibleFaults() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/shouldDeleteInaccessibleFaults
func (m_ ManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval
func (m_ ManagedObjectContext) StalenessInterval() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("stalenessInterval"))
	return rv
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/stalenessInterval
func (m_ ManagedObjectContext) SetStalenessInterval(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStalenessInterval:"), value)
}


// The object that provides undo support for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undoManager
func (m_ ManagedObjectContext) UndoManager() UndoManager /* not a class type */ {
	rv := objc.Send[UndoManager](m_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that provides undo support for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/undoManager
func (m_ ManagedObjectContext) SetUndoManager(value UndoManager /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUndoManager:"), value)
}


// The set of objects registered with the context that have uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/updatedObjects
func (m_ ManagedObjectContext) UpdatedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("updatedObjects"))
	return rv
}


// A key for the set of objects that were marked for deletion during the previous event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsdeletedobjectskey
func (m_ ManagedObjectContext) NSDeletedObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSDeletedObjectsKey"))
	return rv
}


// A key for the set of objects that were inserted into the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinsertedobjectskey
func (m_ ManagedObjectContext) NSInsertedObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInsertedObjectsKey"))
	return rv
}


// A key that specifies that all objects in the context have been invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinvalidatedallobjectskey
func (m_ ManagedObjectContext) NSInvalidatedAllObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInvalidatedAllObjectsKey"))
	return rv
}


// A key for the set of objects that were invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsinvalidatedobjectskey
func (m_ ManagedObjectContext) NSInvalidatedObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSInvalidatedObjectsKey"))
	return rv
}


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/automaticallymergeschangesfromparent
func (m_ ManagedObjectContext) AutomaticallyMergesChangesFromParent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/automaticallymergeschangesfromparent
func (m_ ManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}


// The concurrency type for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/concurrencytype-swift.property
func (m_ ManagedObjectContext) ConcurrencyType() ManagedObjectContextConcurrencyType /* not a class type */ {
	rv := objc.Send[ManagedObjectContextConcurrencyType](m_.ID, objc.Sel("concurrencyType"))
	return rv
}


// The concurrency type for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/concurrencytype-swift.property
func (m_ ManagedObjectContext) SetConcurrencyType(value ManagedObjectContextConcurrencyType /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConcurrencyType:"), value)
}


// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/haschanges
func (m_ ManagedObjectContext) HasChanges() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasChanges"))
	return rv
}


// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/haschanges
func (m_ ManagedObjectContext) SetHasChanges(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasChanges:"), value)
}


// The merge policy of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/mergepolicy
func (m_ ManagedObjectContext) MergePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mergePolicy"))
	return rv
}


// The merge policy of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/mergepolicy
func (m_ ManagedObjectContext) SetMergePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMergePolicy:"), value)
}


// The developer-provided name of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/name
func (m_ ManagedObjectContext) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The developer-provided name of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/name
func (m_ ManagedObjectContext) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}


// The parent of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/parent
func (m_ ManagedObjectContext) Parent() IManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](m_.ID, objc.Sel("parent"))
	return rv
}


// The parent of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/parent
func (m_ ManagedObjectContext) SetParent(value IManagedObjectContext) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParent:"), value)
}


// The persistent store coordinator of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/persistentstorecoordinator
func (m_ ManagedObjectContext) PersistentStoreCoordinator() IPersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](m_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}


// The persistent store coordinator of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/persistentstorecoordinator
func (m_ ManagedObjectContext) SetPersistentStoreCoordinator(value IPersistentStoreCoordinator) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPersistentStoreCoordinator:"), value)
}


// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/propagatesdeletesatendofevent
func (m_ ManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}


// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/propagatesdeletesatendofevent
func (m_ ManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPropagatesDeletesAtEndOfEvent:"), value)
}


// Returns the token associated with the query generation currently in use by this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/querygenerationtoken
func (m_ ManagedObjectContext) QueryGenerationToken() IQueryGenerationToken {
	rv := objc.Send[QueryGenerationToken](m_.ID, objc.Sel("queryGenerationToken"))
	return rv
}


// Returns the token associated with the query generation currently in use by this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/querygenerationtoken
func (m_ ManagedObjectContext) SetQueryGenerationToken(value IQueryGenerationToken) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQueryGenerationToken:"), value)
}


// The set of registered managed objects in the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/registeredobjects
func (m_ ManagedObjectContext) RegisteredObjects() IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("registeredObjects"))
	return rv
}


// The set of registered managed objects in the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/registeredobjects
func (m_ ManagedObjectContext) SetRegisteredObjects(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegisteredObjects:"), value)
}


// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/retainsregisteredobjects
func (m_ ManagedObjectContext) RetainsRegisteredObjects() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("retainsRegisteredObjects"))
	return rv
}


// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/retainsregisteredobjects
func (m_ ManagedObjectContext) SetRetainsRegisteredObjects(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRetainsRegisteredObjects:"), value)
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/transactionauthor
func (m_ ManagedObjectContext) TransactionAuthor() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/transactionauthor
func (m_ ManagedObjectContext) SetTransactionAuthor(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}


// The user information for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/userinfo
func (m_ ManagedObjectContext) UserInfo() MutableDictionary /* not a class type */ {
	rv := objc.Send[MutableDictionary](m_.ID, objc.Sel("userInfo"))
	return rv
}


// The user information for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/userinfo
func (m_ ManagedObjectContext) SetUserInfo(value MutableDictionary /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInfo:"), value)
}


// Constant used to reference the query generation token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontextquerygenerationkey
func (m_ ManagedObjectContext) NSManagedObjectContextQueryGenerationKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSManagedObjectContextQueryGenerationKey"))
	return rv
}


// A key for the set of objects that were refreshed but were not dirtied in the scope of this context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsrefreshedobjectskey
func (m_ ManagedObjectContext) NSRefreshedObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSRefreshedObjectsKey"))
	return rv
}


// A key for the set of objects that were updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsupdatedobjectskey
func (m_ ManagedObjectContext) NSUpdatedObjectsKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("NSUpdatedObjectsKey"))
	return rv
}



