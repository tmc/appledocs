// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterAreaInfoStruct */


/* debug [class_header]: Header for MTRServiceAreaClusterAreaInfoStruct */
// The class instance for the [MTRServiceAreaClusterAreaInfoStruct] class.
var (
	MTRServiceAreaClusterAreaInfoStructClass     _MTRServiceAreaClusterAreaInfoStructClass
	MTRServiceAreaClusterAreaInfoStructClassOnce sync.Once
)

func getMTRServiceAreaClusterAreaInfoStructClass() _MTRServiceAreaClusterAreaInfoStructClass {
	MTRServiceAreaClusterAreaInfoStructClassOnce.Do(func() {
		MTRServiceAreaClusterAreaInfoStructClass = _MTRServiceAreaClusterAreaInfoStructClass{objc.GetClass("MTRServiceAreaClusterAreaInfoStruct")}
	})
	return MTRServiceAreaClusterAreaInfoStructClass
}

type _MTRServiceAreaClusterAreaInfoStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterAreaInfoStruct */
// An interface definition for the [MTRServiceAreaClusterAreaInfoStruct] class.
type IMTRServiceAreaClusterAreaInfoStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterAreaInfoStruct */
	// properties:
	LocationInfo() IMTRDataTypeLocationDescriptorStruct
	SetLocationInfo(value IMTRDataTypeLocationDescriptorStruct)
	LandmarkInfo() IMTRServiceAreaClusterLandmarkInfoStruct
	SetLandmarkInfo(value IMTRServiceAreaClusterLandmarkInfoStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterAreaInfoStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterAreaInfoStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterAreaInfoStructClass) Alloc() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterAreaInfoStructClass) New() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterAreaInfoStruct) Init() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterAreaInfoStruct) Autorelease() MTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterAreaInfoStruct creates a new MTRServiceAreaClusterAreaInfoStruct instance.
func NewMTRServiceAreaClusterAreaInfoStruct() MTRServiceAreaClusterAreaInfoStruct {
	return getMTRServiceAreaClusterAreaInfoStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterAreaInfoStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct
type MTRServiceAreaClusterAreaInfoStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterAreaInfoStructFrom constructs a [MTRServiceAreaClusterAreaInfoStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterAreaInfoStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterAreaInfoStruct {
	return MTRServiceAreaClusterAreaInfoStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterAreaInfoStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterAreaInfoStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterAreaInfoStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterAreaInfoStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterAreaInfoStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/locationInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) LocationInfo() IMTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](m_.ID, objc.Sel("locationInfo"))
	return rv
}/* debug [instance_properties/getter]: locationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaInfoStruct/locationInfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) SetLocationInfo(value IMTRDataTypeLocationDescriptorStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocationInfo:"), value)
}/* debug [instance_properties/setter]: locationInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareainfostruct/landmarkinfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) LandmarkInfo() IMTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](m_.ID, objc.Sel("landmarkInfo"))
	return rv
}/* debug [instance_properties/getter]: landmarkInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareainfostruct/landmarkinfo
func (m_ MTRServiceAreaClusterAreaInfoStruct) SetLandmarkInfo(value IMTRServiceAreaClusterLandmarkInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLandmarkInfo:"), value)
}/* debug [instance_properties/setter]: landmarkInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterAreaInfoStruct */



