// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNHumanBodyRecognizedPoint3D */


/* debug [class_header]: Header for VNHumanBodyRecognizedPoint3D */
// The class instance for the [HumanBodyRecognizedPoint3D] class.
var (
	HumanBodyRecognizedPoint3DClass     _HumanBodyRecognizedPoint3DClass
	HumanBodyRecognizedPoint3DClassOnce sync.Once
)

func getHumanBodyRecognizedPoint3DClass() _HumanBodyRecognizedPoint3DClass {
	HumanBodyRecognizedPoint3DClassOnce.Do(func() {
		HumanBodyRecognizedPoint3DClass = _HumanBodyRecognizedPoint3DClass{objc.GetClass("VNHumanBodyRecognizedPoint3D")}
	})
	return HumanBodyRecognizedPoint3DClass
}

type _HumanBodyRecognizedPoint3DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HumanBodyRecognizedPoint3D */
// An interface definition for the [HumanBodyRecognizedPoint3D] class.
type IHumanBodyRecognizedPoint3D interface {
	IRecognizedPoint3D
	
/* debug [class_interface_properties]: Properties for HumanBodyRecognizedPoint3D */
	// properties:
	LocalPosition() objectivec.IObject
	ParentJoint() HumanBodyPose3DObservationJointName /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HumanBodyRecognizedPoint3D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HumanBodyRecognizedPoint3D */
// Alloc allocates a new instance without initialization.
func (hc _HumanBodyRecognizedPoint3DClass) Alloc() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HumanBodyRecognizedPoint3DClass) New() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanBodyRecognizedPoint3D) Init() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanBodyRecognizedPoint3D) Autorelease() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanBodyRecognizedPoint3D creates a new HumanBodyRecognizedPoint3D instance.
func NewHumanBodyRecognizedPoint3D() HumanBodyRecognizedPoint3D {
	return getHumanBodyRecognizedPoint3DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HumanBodyRecognizedPoint3D */
// A recognized 3D point that includes a parent joint.


// A recognized 3D point that includes a parent joint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D
type HumanBodyRecognizedPoint3D struct {
	RecognizedPoint3D
}

// HumanBodyRecognizedPoint3DFrom constructs a [HumanBodyRecognizedPoint3D] from an unsafe.Pointer.
//
// A recognized 3D point that includes a parent joint.
func HumanBodyRecognizedPoint3DFrom(ptr unsafe.Pointer) HumanBodyRecognizedPoint3D {
	return HumanBodyRecognizedPoint3D{
		RecognizedPoint3D: RecognizedPoint3DFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HumanBodyRecognizedPoint3D *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HumanBodyRecognizedPoint3D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HumanBodyRecognizedPoint3D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HumanBodyRecognizedPoint3D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HumanBodyRecognizedPoint3D */

// The three-dimensional position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D/localPosition
func (h_ HumanBodyRecognizedPoint3D) LocalPosition() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("localPosition"))
	return rv
}/* debug [instance_properties/getter]: localPosition */


// The parent joint in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D/parentJoint
func (h_ HumanBodyRecognizedPoint3D) ParentJoint() HumanBodyPose3DObservationJointName /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("parentJoint"))
	return rv
}/* debug [instance_properties/getter]: parentJoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNHumanBodyRecognizedPoint3D */



