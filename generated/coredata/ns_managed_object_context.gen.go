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
	NSDeletedObjectsKey() string
	NSInsertedObjectsKey() string
	NSInvalidatedAllObjectsKey() string
	NSInvalidatedObjectsKey() string
	AutomaticallyMergesChangesFromParent() bool
	SetAutomaticallyMergesChangesFromParent(value bool)
	ConcurrencyType() unsafe.Pointer
	SetConcurrencyType(value unsafe.Pointer)
	DeletedObjects() IManagedObject
	SetDeletedObjects(value IManagedObject)
	HasChanges() bool
	SetHasChanges(value bool)
	InsertedObjects() IManagedObject
	SetInsertedObjects(value IManagedObject)
	MergePolicy() unsafe.Pointer
	SetMergePolicy(value unsafe.Pointer)
	Name() string
	SetName(value string)
	Parent() IManagedObjectContext
	SetParent(value IManagedObjectContext)
	PersistentStoreCoordinator() IPersistentStoreCoordinator
	SetPersistentStoreCoordinator(value IPersistentStoreCoordinator)
	PropagatesDeletesAtEndOfEvent() bool
	SetPropagatesDeletesAtEndOfEvent(value bool)
	QueryGenerationToken() IQueryGenerationToken
	SetQueryGenerationToken(value IQueryGenerationToken)
	RegisteredObjects() IManagedObject
	SetRegisteredObjects(value IManagedObject)
	RetainsRegisteredObjects() bool
	SetRetainsRegisteredObjects(value bool)
	ShouldDeleteInaccessibleFaults() bool
	SetShouldDeleteInaccessibleFaults(value bool)
	StalenessInterval() unsafe.Pointer
	SetStalenessInterval(value unsafe.Pointer)
	TransactionAuthor() string
	SetTransactionAuthor(value string)
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	UpdatedObjects() IManagedObject
	SetUpdatedObjects(value IManagedObject)
	UserInfo() foundation.MutableDictionary
	SetUserInfo(value foundation.MutableDictionary)
	NSManagedObjectContextQueryGenerationKey() string
	NSRefreshedObjectsKey() string
	NSUpdatedObjectsKey() string
	Save(error_ unsafe.Pointer) bool
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



// Attempts to commit unsaved changes to registered objects to the context’s parent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()
func (m_ ManagedObjectContext) Save(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("save:"), error_)
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


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/automaticallymergeschangesfromparent
func (m_ ManagedObjectContext) AutomaticallyMergesChangesFromParent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}


// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/automaticallymergeschangesfromparent
func (m_ ManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}


// The concurrency type for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/concurrencytype-swift.property
func (m_ ManagedObjectContext) ConcurrencyType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("concurrencyType"))
	return rv
}


// The concurrency type for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/concurrencytype-swift.property
func (m_ ManagedObjectContext) SetConcurrencyType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConcurrencyType:"), value)
}


// The set of objects that will be removed from their persistent store during the next save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/deletedobjects
func (m_ ManagedObjectContext) DeletedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("deletedObjects"))
	return rv
}


// The set of objects that will be removed from their persistent store during the next save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/deletedobjects
func (m_ ManagedObjectContext) SetDeletedObjects(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeletedObjects:"), value)
}


// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/haschanges
func (m_ ManagedObjectContext) HasChanges() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasChanges"))
	return rv
}


// A Boolean value that indicates whether the context has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/haschanges
func (m_ ManagedObjectContext) SetHasChanges(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasChanges:"), value)
}


// The set of objects that have been inserted into the context but not yet saved in a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/insertedobjects
func (m_ ManagedObjectContext) InsertedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("insertedObjects"))
	return rv
}


// The set of objects that have been inserted into the context but not yet saved in a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/insertedobjects
func (m_ ManagedObjectContext) SetInsertedObjects(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInsertedObjects:"), value)
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
func (m_ ManagedObjectContext) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The developer-provided name of the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/name
func (m_ ManagedObjectContext) SetName(value string) {
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
func (m_ ManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}


// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/propagatesdeletesatendofevent
func (m_ ManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool) {
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
func (m_ ManagedObjectContext) RetainsRegisteredObjects() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("retainsRegisteredObjects"))
	return rv
}


// A Boolean value that indicates whether the context keeps strong references to all registered managed objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/retainsregisteredobjects
func (m_ ManagedObjectContext) SetRetainsRegisteredObjects(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRetainsRegisteredObjects:"), value)
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/shoulddeleteinaccessiblefaults
func (m_ ManagedObjectContext) ShouldDeleteInaccessibleFaults() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}


// A Boolean value that determines whether the context turns inaccessible faults into deleted objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/shoulddeleteinaccessiblefaults
func (m_ ManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/stalenessinterval
func (m_ ManagedObjectContext) StalenessInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("stalenessInterval"))
	return rv
}


// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/stalenessinterval
func (m_ ManagedObjectContext) SetStalenessInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStalenessInterval:"), value)
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/transactionauthor
func (m_ ManagedObjectContext) TransactionAuthor() string {
	rv := objc.Send[string](m_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// The author for the context that is used as an identifier in persistent history transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/transactionauthor
func (m_ ManagedObjectContext) SetTransactionAuthor(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
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
func (m_ ManagedObjectContext) SetUndoManager(value foundation.UndoManager) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUndoManager:"), value)
}


// The set of objects registered with the context that have uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/updatedobjects
func (m_ ManagedObjectContext) UpdatedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](m_.ID, objc.Sel("updatedObjects"))
	return rv
}


// The set of objects registered with the context that have uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/updatedobjects
func (m_ ManagedObjectContext) SetUpdatedObjects(value IManagedObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdatedObjects:"), value)
}


// The user information for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/userinfo
func (m_ ManagedObjectContext) UserInfo() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](m_.ID, objc.Sel("userInfo"))
	return rv
}


// The user information for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/userinfo
func (m_ ManagedObjectContext) SetUserInfo(value foundation.MutableDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserInfo:"), value)
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



