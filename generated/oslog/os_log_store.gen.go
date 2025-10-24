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
	// properties:
	// methods:
}

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





