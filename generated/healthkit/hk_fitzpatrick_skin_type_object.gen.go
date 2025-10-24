// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKFitzpatrickSkinTypeObject */


/* debug [class_header]: Header for HKFitzpatrickSkinTypeObject */
// The class instance for the [HKFitzpatrickSkinTypeObject] class.
var (
	HKFitzpatrickSkinTypeObjectClass     _HKFitzpatrickSkinTypeObjectClass
	HKFitzpatrickSkinTypeObjectClassOnce sync.Once
)

func getHKFitzpatrickSkinTypeObjectClass() _HKFitzpatrickSkinTypeObjectClass {
	HKFitzpatrickSkinTypeObjectClassOnce.Do(func() {
		HKFitzpatrickSkinTypeObjectClass = _HKFitzpatrickSkinTypeObjectClass{objc.GetClass("HKFitzpatrickSkinTypeObject")}
	})
	return HKFitzpatrickSkinTypeObjectClass
}

type _HKFitzpatrickSkinTypeObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKFitzpatrickSkinTypeObject */
// An interface definition for the [HKFitzpatrickSkinTypeObject] class.
type IHKFitzpatrickSkinTypeObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKFitzpatrickSkinTypeObject */
	// properties:
	SkinType() HKFitzpatrickSkinType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKFitzpatrickSkinTypeObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKFitzpatrickSkinTypeObject */
// Alloc allocates a new instance without initialization.
func (hc _HKFitzpatrickSkinTypeObjectClass) Alloc() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKFitzpatrickSkinTypeObjectClass) New() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKFitzpatrickSkinTypeObject) Init() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKFitzpatrickSkinTypeObject) Autorelease() HKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKFitzpatrickSkinTypeObject creates a new HKFitzpatrickSkinTypeObject instance.
func NewHKFitzpatrickSkinTypeObject() HKFitzpatrickSkinTypeObject {
	return getHKFitzpatrickSkinTypeObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKFitzpatrickSkinTypeObject */
// This class acts as a wrapper for the enumeration.


// This class acts as a wrapper for the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinTypeObject
type HKFitzpatrickSkinTypeObject struct {
	objectivec.Object
}

// HKFitzpatrickSkinTypeObjectFrom constructs a [HKFitzpatrickSkinTypeObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKFitzpatrickSkinTypeObjectFrom(ptr unsafe.Pointer) HKFitzpatrickSkinTypeObject {
	return HKFitzpatrickSkinTypeObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKFitzpatrickSkinTypeObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKFitzpatrickSkinTypeObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKFitzpatrickSkinTypeObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKFitzpatrickSkinTypeObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKFitzpatrickSkinTypeObject */

// The user’s skin type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinTypeObject/skinType
func (h_ HKFitzpatrickSkinTypeObject) SkinType() HKFitzpatrickSkinType {
	rv := objc.Send[HKFitzpatrickSkinType](h_.ID, objc.Sel("skinType"))
	return rv
}/* debug [instance_properties/getter]: skinType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKFitzpatrickSkinTypeObject */



