// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKDeletedObject] class.
var (
	HKDeletedObjectClass     _HKDeletedObjectClass
	HKDeletedObjectClassOnce sync.Once
)

func getHKDeletedObjectClass() _HKDeletedObjectClass {
	HKDeletedObjectClassOnce.Do(func() {
		HKDeletedObjectClass = _HKDeletedObjectClass{objc.GetClass("HKDeletedObject")}
	})
	return HKDeletedObjectClass
}

type _HKDeletedObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKDeletedObject] class.
type IHKDeletedObject interface {
	objectivec.IObject
}

// An object that represents a sample that has been deleted from the HealthKit store.
//
// Use queries to generate a list of recently deleted objects. Create a query using the method. When the system calls the result handler, it passes the parameter an array of instances matching the query. Deleted objects are temporary; the system may remove them from the HealthKit store at any time to free up space. To guarantee that you receive notifications for all deleted objects, create an and register it for background delivery. The system then wakes your app and calls the observer query’s update handler whenever the matching objects change—including deletions. However, the query does not provide a list of deleted objects. To determine which objects were deleted, use the observer query’s update handler to create an anchored object query for the newly deleted objects.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDeletedObject
type HKDeletedObject struct {
	objectivec.Object
}

// HKDeletedObjectFrom constructs a [HKDeletedObject] from an unsafe.Pointer.
//
// An object that represents a sample that has been deleted from the HealthKit store.
func HKDeletedObjectFrom(ptr unsafe.Pointer) HKDeletedObject {
	return HKDeletedObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDeletedObjectClass) Alloc() HKDeletedObject {
	rv := objc.Send[HKDeletedObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDeletedObjectClass) New() HKDeletedObject {
	rv := objc.Send[HKDeletedObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDeletedObject) Init() HKDeletedObject {
	rv := objc.Send[HKDeletedObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDeletedObject) Autorelease() HKDeletedObject {
	rv := objc.Send[HKDeletedObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDeletedObject creates a new HKDeletedObject instance.
func NewHKDeletedObject() HKDeletedObject {
	return getHKDeletedObjectClass().New()
}


// The metadata associated with the deleted object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDeletedObject/metadata
func (h_ HKDeletedObject) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("metadata"))
	return rv
}

// The universally unique identifier (UUID) for the HealthKit object that was deleted from the store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDeletedObject/uuid
func (h_ HKDeletedObject) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("UUID"))
	return rv
}



