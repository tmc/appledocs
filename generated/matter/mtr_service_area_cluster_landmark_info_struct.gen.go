// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterLandmarkInfoStruct */


/* debug [class_header]: Header for MTRServiceAreaClusterLandmarkInfoStruct */
// The class instance for the [MTRServiceAreaClusterLandmarkInfoStruct] class.
var (
	MTRServiceAreaClusterLandmarkInfoStructClass     _MTRServiceAreaClusterLandmarkInfoStructClass
	MTRServiceAreaClusterLandmarkInfoStructClassOnce sync.Once
)

func getMTRServiceAreaClusterLandmarkInfoStructClass() _MTRServiceAreaClusterLandmarkInfoStructClass {
	MTRServiceAreaClusterLandmarkInfoStructClassOnce.Do(func() {
		MTRServiceAreaClusterLandmarkInfoStructClass = _MTRServiceAreaClusterLandmarkInfoStructClass{objc.GetClass("MTRServiceAreaClusterLandmarkInfoStruct")}
	})
	return MTRServiceAreaClusterLandmarkInfoStructClass
}

type _MTRServiceAreaClusterLandmarkInfoStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterLandmarkInfoStruct */
// An interface definition for the [MTRServiceAreaClusterLandmarkInfoStruct] class.
type IMTRServiceAreaClusterLandmarkInfoStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterLandmarkInfoStruct */
	// properties:
	LandmarkTag() objc.IObject /* cross-framework: NSNumber */
	SetLandmarkTag(value objc.IObject /* cross-framework: NSNumber */)
	RelativePositionTag() objc.IObject /* cross-framework: NSNumber */
	SetRelativePositionTag(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterLandmarkInfoStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterLandmarkInfoStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterLandmarkInfoStructClass) Alloc() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterLandmarkInfoStructClass) New() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) Init() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) Autorelease() MTRServiceAreaClusterLandmarkInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterLandmarkInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterLandmarkInfoStruct creates a new MTRServiceAreaClusterLandmarkInfoStruct instance.
func NewMTRServiceAreaClusterLandmarkInfoStruct() MTRServiceAreaClusterLandmarkInfoStruct {
	return getMTRServiceAreaClusterLandmarkInfoStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterLandmarkInfoStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct
type MTRServiceAreaClusterLandmarkInfoStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterLandmarkInfoStructFrom constructs a [MTRServiceAreaClusterLandmarkInfoStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterLandmarkInfoStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterLandmarkInfoStruct {
	return MTRServiceAreaClusterLandmarkInfoStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterLandmarkInfoStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterLandmarkInfoStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterLandmarkInfoStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterLandmarkInfoStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterLandmarkInfoStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/landmarkTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) LandmarkTag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("landmarkTag"))
	return rv
}/* debug [instance_properties/getter]: landmarkTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterLandmarkInfoStruct/landmarkTag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) SetLandmarkTag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLandmarkTag:"), value)
}/* debug [instance_properties/setter]: landmarkTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterlandmarkinfostruct/relativepositiontag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) RelativePositionTag() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("relativePositionTag"))
	return rv
}/* debug [instance_properties/getter]: relativePositionTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterlandmarkinfostruct/relativepositiontag
func (m_ MTRServiceAreaClusterLandmarkInfoStruct) SetRelativePositionTag(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRelativePositionTag:"), value)
}/* debug [instance_properties/setter]: relativePositionTag */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterLandmarkInfoStruct */



