// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSLogStore */


/* debug [class_header]: Header for OSLogStore */
// The class instance for the [OSLogStore] class.
var (
	OSLogStoreClass     _OSLogStoreClass
	OSLogStoreClassOnce sync.Once
)

func getOSLogStoreClass() _OSLogStoreClass {
	OSLogStoreClassOnce.Do(func() {
		OSLogStoreClass = _OSLogStoreClass{objc.GetClass("OSLogStore")}
	})
	return OSLogStoreClass
}

type _OSLogStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogStore */
// An interface definition for the [OSLogStore] class.
type IOSLogStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OSLogStore */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogStore */
	// methods:
	EntriesEnumeratorAndReturnError(error_ unsafe.Pointer) IOSLogEnumerator
	EntriesEnumeratorWithOptionsPositionPredicateError(options OSLogEnumeratorOptions, position IOSLogPosition, predicate foundation.Predicate, error_ unsafe.Pointer) IOSLogEnumerator
	PositionWithDate(date objc.IObject /* cross-framework: NSDate */) IOSLogPosition
	PositionWithTimeIntervalSinceEnd(seconds float64) IOSLogPosition
	PositionWithTimeIntervalSinceLatestBoot(seconds float64) IOSLogPosition
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogStore */
// Alloc allocates a new instance without initialization.
func (oc _OSLogStoreClass) Alloc() OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogStoreClass) New() OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogStore) Init() OSLogStore {
	rv := objc.Send[OSLogStore](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogStore) Autorelease() OSLogStore {
	rv := objc.Send[OSLogStore](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogStore creates a new OSLogStore instance.
func NewOSLogStore() OSLogStore {
	return getOSLogStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogStore */
// A set of entries from the unified logging system.
//
// Instances of this class represent a fixed range of entries and may be backed by a or your Mac’s local store. In Swift, Use the function to retrieve a filtered array of log entries. In Objective-C, use instances of this class to create objects. One store can support multiple instances concurrently.


// A set of entries from the unified logging system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore
type OSLogStore struct {
	objectivec.Object
}

// OSLogStoreFrom constructs a [OSLogStore] from an unsafe.Pointer.
//
// A set of entries from the unified logging system.
func OSLogStoreFrom(ptr unsafe.Pointer) OSLogStore {
	return OSLogStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(scope:)
func NewOSLogStoreWithScopeError(scope OSLogStoreScope, error_ unsafe.Pointer) OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithScope:error:"), scope, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewOSLogStoreWithScopeError */


// Creates a log store based on a log archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(url:)
func NewOSLogStoreWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewOSLogStoreWithURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogStore */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(scope:)
func (oc _OSLogStoreClass) StoreWithScopeError(scope OSLogStoreScope, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("storeWithScope:error:"), scope, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StoreWithScopeError) */


// Creates a log store based on a log archive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(url:)
func (oc _OSLogStoreClass) StoreWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("storeWithURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StoreWithURLError) */


// Creates a log store representing the Mac’s local store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/local()
func (oc _OSLogStoreClass) LocalStoreAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("localStoreAndReturnError:"), error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalStoreAndReturnError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogStore */

// Returns a log enumerator with default options for viewing the entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorAndReturnError:
func (o_ OSLogStore) EntriesEnumeratorAndReturnError(error_ unsafe.Pointer) IOSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("entriesEnumeratorAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: EntriesEnumeratorAndReturnError */


// Returns a log enumerator based on an underlying store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorWithOptions:position:predicate:error:
func (o_ OSLogStore) EntriesEnumeratorWithOptionsPositionPredicateError(options OSLogEnumeratorOptions, position IOSLogPosition, predicate foundation.Predicate, error_ unsafe.Pointer) IOSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("entriesEnumeratorWithOptions:position:predicate:error:"), options, position, predicate, error_)
	return rv
}/* debug [instance_methods/method]: EntriesEnumeratorWithOptionsPositionPredicateError */


// Returns a position representing the time specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(date:)
func (o_ OSLogStore) PositionWithDate(date objc.IObject /* cross-framework: NSDate */) IOSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("positionWithDate:"), date)
	return rv
}/* debug [instance_methods/method]: PositionWithDate */


// Returns a position representing time since the end of the time range that the entries span.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceEnd:)
func (o_ OSLogStore) PositionWithTimeIntervalSinceEnd(seconds float64) IOSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("positionWithTimeIntervalSinceEnd:"), seconds)
	return rv
}/* debug [instance_methods/method]: PositionWithTimeIntervalSinceEnd */


// Returns a position representing time since the last boot in the series of entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceLatestBoot:)
func (o_ OSLogStore) PositionWithTimeIntervalSinceLatestBoot(seconds float64) IOSLogPosition {
	rv := objc.Send[OSLogPosition](o_.ID, objc.Sel("positionWithTimeIntervalSinceLatestBoot:"), seconds)
	return rv
}/* debug [instance_methods/method]: PositionWithTimeIntervalSinceLatestBoot */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogStore */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogStore */


