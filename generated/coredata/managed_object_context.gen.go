// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectContext] class.
var (
	managedObjectContextClass     _ManagedObjectContextClass
	managedObjectContextClassOnce sync.Once
)

func getManagedObjectContextClass() _ManagedObjectContextClass {
	managedObjectContextClassOnce.Do(func() {
		managedObjectContextClass = _ManagedObjectContextClass{objc.GetClass("NSManagedObjectContext")}
	})
	return managedObjectContextClass
}

type _ManagedObjectContextClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectContext] class.
type IManagedObjectContext interface {
	objectivec.IObject
	AssignObjectToPersistentStore(object objc.ID, store unsafe.Pointer)
	MergeChangesFromContextDidSaveNotification(notification unsafe.Pointer)
	PerformBlock(block unsafe.Pointer)
	PerformBlockAndWait(block unsafe.Pointer)
	ProcessPendingChanges()
	RefreshAllObjects()
	Reset()
	Save(error unsafe.Pointer) bool
	SetQueryGenerationFromTokenError(generation unsafe.Pointer, error unsafe.Pointer) bool
}

// An object space to manipulate and track changes to managed objects.
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
func (mc _ManagedObjectContextClass) MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData unsafe.Pointer, contexts unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mergeChangesFromRemoteContextSave:intoContexts:"), changeNotificationData, contexts)
}
// Specifies the store in which a newly inserted object will be saved.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/assign(_:to:)
func (m_ ManagedObjectContext) AssignObjectToPersistentStore(object objc.ID, store unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("assignObject:toPersistentStore:"), object, store)
}
// Merges the changes specified in a given notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/mergeChanges(fromContextDidSave:)
func (m_ ManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("mergeChangesFromContextDidSaveNotification:"), notification)
}
// Allows a context that has registered as an observer of a value to be notified of a change to that value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/observeValue(forKeyPath:of:change:context:)
func (m_ ManagedObjectContext) ObserveValueForKeyPathOfObjectChangeContext(keyPath string, object objc.ID, change unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), objc.String(keyPath), object, change, context)
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
// Refreshes all of the registered managed objects in the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/refreshAllObjects()
func (m_ ManagedObjectContext) RefreshAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("refreshAllObjects"))
}
// Returns the context to its base state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/reset()
func (m_ ManagedObjectContext) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}
// Attempts to commit unsaved changes to registered objects to the context’s parent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/save()
func (m_ ManagedObjectContext) Save(error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("save:"), error)
	return rv
}
// Sets the query generation this context should use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectContext/setQueryGenerationFrom(_:)
func (m_ ManagedObjectContext) SetQueryGenerationFromTokenError(generation unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setQueryGenerationFromToken:error:"), generation, error)
	return rv
}

