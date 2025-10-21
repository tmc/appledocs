// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [OSLogStore] class.
type IOSLogStore interface {
	objectivec.IObject
	EntriesEnumeratorAndReturnError(error_ unsafe.Pointer) unsafe.Pointer
	EntriesEnumeratorWithOptionsPositionPredicateError(options unsafe.Pointer, position unsafe.Pointer, predicate unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	PositionWithDate(date unsafe.Pointer) unsafe.Pointer
	PositionWithTimeIntervalSinceEnd(seconds TimeInterval) unsafe.Pointer
	PositionWithTimeIntervalSinceLatestBoot(seconds TimeInterval) unsafe.Pointer
}

// A set of entries from the unified logging system.
//
// Instances of this class represent a fixed range of entries and may be backed by a or your Mac’s local store. In Swift, Use the function to retrieve a filtered array of log entries. In Objective-C, use instances of this class to create objects. One store can support multiple instances concurrently.
//
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

// Alloc allocates a new instance without initialization.
func (oc _OSLogStoreClass) Alloc() OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(scope:)
func NewOSLogStoreWithScopeError(scope unsafe.Pointer, error_ unsafe.Pointer) OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithScope:error:"), scope, error_)
	return rv
}

// Creates a log store based on a log archive.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(url:)
func NewOSLogStoreWithURLError(url unsafe.Pointer, error_ unsafe.Pointer) OSLogStore {
	rv := objc.Send[OSLogStore](objc.ID(getOSLogStoreClass().class), objc.Sel("storeWithURL:error:"), url, error_)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(scope:)
func (oc _OSLogStoreClass) StoreWithScopeError(scope unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("storeWithScope:error:"), scope, error_)
	return rv
}

// Creates a log store based on a log archive.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/init(url:)
func (oc _OSLogStoreClass) StoreWithURLError(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("storeWithURL:error:"), url, error_)
	return rv
}

// Creates a log store representing the Mac’s local store.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/local()
func (oc _OSLogStoreClass) LocalStoreAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("localStoreAndReturnError:"), error_)
	return rv
}

// Returns a log enumerator with default options for viewing the entries.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorAndReturnError:
func (o_ OSLogStore) EntriesEnumeratorAndReturnError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("entriesEnumeratorAndReturnError:"), error_)
	return rv
}

// Returns a log enumerator based on an underlying store.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/entriesEnumeratorWithOptions:position:predicate:error:
func (o_ OSLogStore) EntriesEnumeratorWithOptionsPositionPredicateError(options unsafe.Pointer, position unsafe.Pointer, predicate unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("entriesEnumeratorWithOptions:position:predicate:error:"), options, position, predicate, error_)
	return rv
}

// Returns a position representing the time specified.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(date:)
func (o_ OSLogStore) PositionWithDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("positionWithDate:"), date)
	return rv
}

// Returns a position representing time since the end of the time range that the entries span.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceEnd:)
func (o_ OSLogStore) PositionWithTimeIntervalSinceEnd(seconds TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("positionWithTimeIntervalSinceEnd:"), seconds)
	return rv
}

// Returns a position representing time since the last boot in the series of entries.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogStore/position(timeIntervalSinceLatestBoot:)
func (o_ OSLogStore) PositionWithTimeIntervalSinceLatestBoot(seconds TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("positionWithTimeIntervalSinceLatestBoot:"), seconds)
	return rv
}


