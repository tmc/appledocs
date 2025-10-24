// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKLocationSortDescriptor */


/* debug [class_header]: Header for CKLocationSortDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKLocationSortDescriptor */
// An interface definition for the [CKLocationSortDescriptor] class.
type ICKLocationSortDescriptor interface {
	ISortDescriptor
	
/* debug [class_interface_properties]: Properties for CKLocationSortDescriptor */
	// properties:
	RelativeLocation() corelocation.Location
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKLocationSortDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKLocationSortDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CKLocationSortDescriptorClass) Alloc() CKLocationSortDescriptor {
	rv := objc.Send[CKLocationSortDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKLocationSortDescriptor */
// An object for sorting records that contain location data.
//
// You can add a location sort descriptor to your queries when searching for records. At creation time, you must provide the sort descriptor with a key that has a object as its value. The sort descriptor uses the value of that key to perform the sort. CloudKit computes distance by drawing a direct line between the two locations that follows the curvature of the Earth. Distances don’t account for altitude changes between the two locations.


// An object for sorting records that contain location data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor
type CKLocationSortDescriptor struct {
	SortDescriptor
}

// CKLocationSortDescriptorFrom constructs a [CKLocationSortDescriptor] from an unsafe.Pointer.
//
// An object for sorting records that contain location data.
func CKLocationSortDescriptorFrom(ptr unsafe.Pointer) CKLocationSortDescriptor {
	return CKLocationSortDescriptor{
		SortDescriptor: SortDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKLocationSortDescriptor */

// Creates a location sort descriptor from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(coder:)
func NewCKLocationSortDescriptorWithCoder(aDecoder foundation.Coder) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[CKLocationSortDescriptor](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKLocationSortDescriptorWithCoder */


// Creates a location sort descriptor using the specified key and relative location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/init(key:relativeLocation:)
func NewCKLocationSortDescriptorWithKeyRelativeLocation(key objc.IObject /* cross-framework: NSString */, relativeLocation corelocation.Location) CKLocationSortDescriptor {
	instance := getCKLocationSortDescriptorClass().Alloc()
	rv := objc.Send[CKLocationSortDescriptor](instance.ID, objc.Sel("initWithKey:relativeLocation:"), key, relativeLocation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKLocationSortDescriptorWithKeyRelativeLocation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKLocationSortDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKLocationSortDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKLocationSortDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKLocationSortDescriptor */

// The reference location for sorting records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKLocationSortDescriptor/relativeLocation
func (c_ CKLocationSortDescriptor) RelativeLocation() corelocation.Location {
	rv := objc.Send[corelocation.Location](c_.ID, objc.Sel("relativeLocation"))
	return rv
}/* debug [instance_properties/getter]: relativeLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKLocationSortDescriptor */


