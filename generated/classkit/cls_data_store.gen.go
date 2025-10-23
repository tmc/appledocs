// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SDataStore] class.
var (
	SDataStoreClass     _SDataStoreClass
	SDataStoreClassOnce sync.Once
)

func getSDataStoreClass() _SDataStoreClass {
	SDataStoreClassOnce.Do(func() {
		SDataStoreClass = _SDataStoreClass{objc.GetClass("CLSDataStore")}
	})
	return SDataStoreClass
}

type _SDataStoreClass struct {
	class objc.Class
}

// An interface definition for the [SDataStore] class.
type ISDataStore interface {
	objectivec.IObject
	// properties:
	ActiveContext() ICLSContext
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	MainAppContext() ICLSContext
	RunningActivity() ICLSActivity
	// methods:
	CompleteAllAssignedActivitiesMatching(contextPath []string /* primitive/slice/pointer. */)
	ContextsMatchingPredicateCompletion(predicate objc.IObject /* cross-framework Predicate */, completion unsafe.Pointer)
	ContextsMatchingIdentifierPathCompletion(identifierPath []string /* primitive/slice/pointer. */, completion unsafe.Pointer)
	FetchActivityForURLCompletion(url foundation.objc.IObject /* cross-framework URL */, completion unsafe.Pointer)
	RemoveContext(context ICLSContext)
	SaveWithCompletion(completion unsafe.Pointer)
}

// A container for all the ClassKit data in your app.
//
// Use the ClassKit data store to build and access contexts ( instances) that you use to advertise your app’s assignable content. Contexts in turn provide access to activities ( instances) and activity items ( , , and instances) that you use to record progress through assignments. You don’t instantiate a data store yourself. Instead, use the single data store instance throughout your app. The data store provides access to the app’s one and only main context through the property. This property acts as the root context in your context hierarchy that you can use as a starting point when searching for descendant contexts. To build contexts, you adopt the protocol in one of your classes, typically one that exists for the lifetime of your app, and assign an instance of that class as the shared data store’s property. Then, when the data store needs a context that it’s never seen before, it asks your delegate to build it. After you make changes to any context, activity, or activity item, call the data store’s method to commit the changes, and propagate them through the network.


// A container for all the ClassKit data in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore
type SDataStore struct {
	objectivec.Object
}

// SDataStoreFrom constructs a [SDataStore] from an unsafe.Pointer.
//
// A container for all the ClassKit data in your app.
func SDataStoreFrom(ptr unsafe.Pointer) SDataStore {
	return SDataStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SDataStoreClass) Alloc() SDataStore {
	rv := objc.Send[SDataStore](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SDataStoreClass) New() SDataStore {
	rv := objc.Send[SDataStore](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SDataStore) Init() SDataStore {
	rv := objc.Send[SDataStore](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SDataStore) Autorelease() SDataStore {
	rv := objc.Send[SDataStore](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSDataStore creates a new SDataStore instance.
func NewSDataStore() SDataStore {
	return getSDataStoreClass().New()
}



// The shared data store object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/shared
func (sc _SDataStoreClass) Shared() SDataStore {
	rv := objc.Send[SDataStore](objc.ID(sc.class), objc.Sel("shared"))
	return rv
}

// Marks all of the assigned and active activities for the given context path as complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/completeAllAssignedActivities(matching:)
func (s_ SDataStore) CompleteAllAssignedActivitiesMatching(contextPath []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("completeAllAssignedActivitiesMatching:"), contextPath)
}


// Fetches all the contexts matching a predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/contexts(matching:completion:)
func (s_ SDataStore) ContextsMatchingPredicateCompletion(predicate objc.IObject /* cross-framework Predicate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("contextsMatchingPredicate:completion:"), predicate, completion)
}


// Fetches all the contexts along a given identifier path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/contexts(matchingIdentifierPath:completion:)
func (s_ SDataStore) ContextsMatchingIdentifierPathCompletion(identifierPath []string /* primitive/slice/pointer. */, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("contextsMatchingIdentifierPath:completion:"), identifierPath, completion)
}


// Fetches an activity for a given document so you can record progress on the associated task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/fetchActivity(for:completion:)
func (s_ SDataStore) FetchActivityForURLCompletion(url foundation.objc.IObject /* cross-framework URL */, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fetchActivityForURL:completion:"), url, completion)
}


// Marks a context for removal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/remove(_:)
func (s_ SDataStore) RemoveContext(context ICLSContext) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeContext:"), context)
}


// Saves any changes you’ve made in the data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/save(completion:)
func (s_ SDataStore) SaveWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("saveWithCompletion:"), completion)
}


// The currently active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/activeContext
func (s_ SDataStore) ActiveContext() ICLSContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("activeContext"))
	return rv
}


// The data store delegate instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/delegate
func (s_ SDataStore) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The data store delegate instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/delegate
func (s_ SDataStore) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// The app’s top-level context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/mainAppContext
func (s_ SDataStore) MainAppContext() ICLSContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("mainAppContext"))
	return rv
}


// The currently running activity within the currently active context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/runningActivity
func (s_ SDataStore) RunningActivity() ICLSActivity {
	rv := objc.Send[SActivity](s_.ID, objc.Sel("runningActivity"))
	return rv
}


// The shared data store object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSDataStore/shared
func (s_ SDataStore) Shared() ICLSDataStore {
	rv := objc.Send[SDataStore](s_.ID, objc.Sel("shared"))
	return rv
}



