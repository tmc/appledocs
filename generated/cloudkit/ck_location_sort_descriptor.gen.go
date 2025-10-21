// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKLocationSortDescriptor] class.
var (
	CKLocationSortDescriptorClass     _CKLocationSortDescriptorClass
	CKLocationSortDescriptorClassOnce sync.Once
)

func getCKLocationSortDescriptorClass() _CKLocationSortDescriptorClass {
	CKLocationSortDescriptorClassOnce.Do(func() {
		CKLocationSortDescriptorClass = _CKLocationSortDescriptorClass{objc.GetClass("CKLocationSortDescriptor")}
	})
	return CKLocationSortDescriptorClass
}

type _CKLocationSortDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [CKLocationSortDescriptor] class.
type ICKLocationSortDescriptor interface {
	foundation.ISortDescriptor
}

// An object for sorting records that contain location data.
//
// You can add a location sort descriptor to your queries when searching for records. At creation time, you must provide the sort descriptor with a key that has a object as its value. The sort descriptor uses the value of that key to perform the sort. CloudKit computes distance by drawing a direct line between the two locations that follows the curvature of the Earth. Distances don’t account for altitude changes between the two locations.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor
type CKLocationSortDescriptor struct {
	foundation.SortDescriptor
}

// CKLocationSortDescriptorFrom constructs a [CKLocationSortDescriptor] from an unsafe.Pointer.
//
// An object for sorting records that contain location data.
func CKLocationSortDescriptorFrom(ptr unsafe.Pointer) CKLocationSortDescriptor {
	return CKLocationSortDescriptor{
		SortDescriptor: foundation.SortDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKLocationSortDescriptorClass) Alloc() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKLocationSortDescriptorClass) New() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKLocationSortDescriptor) Init() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKLocationSortDescriptor) Autorelease() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKLocationSortDescriptor creates a new CKLocationSortDescriptor instance.
func NewCKLocationSortDescriptor() CKLocationSortDescriptor {
	return getCKLocationSortDescriptorClass().New()
}




// Creates a location sort descriptor from a serialized instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(coder:)
func NewCKLocationSortDescriptorWithCoder(aDecoder foundation.ICoder) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[CKLocationSortDescriptor](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}



// Creates a location sort descriptor using the specified key and relative location.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(key:relativeLocation:)
func NewCKLocationSortDescriptorWithKeyRelativeLocation(key string, relativeLocation corelocation.ILocation) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[CKLocationSortDescriptor](instance.ID, objc.Sel("initWithKey:relativeLocation:"), objc.String(key), relativeLocation)
	rv.Autorelease()
	return rv
}


// The reference location for sorting records.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/relativeLocation
func (c_ CKLocationSortDescriptor) RelativeLocation() corelocation.Location {
	rv := objc.Send[corelocation.Location](c_.ID, objc.Sel("relativeLocation"))
	return rv
}


