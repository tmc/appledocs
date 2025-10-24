// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class customAttributes */


/* debug [class_header]: Header for customAttributes */
// The class instance for the [customAttributes] class.
var (
	CustomAttributesClass     _customAttributesClass
	CustomAttributesClassOnce sync.Once
)

func getcustomAttributesClass() _customAttributesClass {
	CustomAttributesClassOnce.Do(func() {
		CustomAttributesClass = _customAttributesClass{objc.GetClass("customAttributes")}
	})
	return CustomAttributesClass
}

type _customAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for customAttributes */
// An interface definition for the [customAttributes] class.
type IcustomAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for customAttributes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for customAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for customAttributes */
// Alloc allocates a new instance without initialization.
func (cc _customAttributesClass) Alloc() customAttributes {
	rv := objc.Send[customAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _customAttributesClass) New() customAttributes {
	rv := objc.Send[customAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ customAttributes) Init() customAttributes {
	rv := objc.Send[customAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ customAttributes) Autorelease() customAttributes {
	rv := objc.Send[customAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcustomAttributes creates a new customAttributes instance.
func NewcustomAttributes() customAttributes {
	return getcustomAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for customAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-c.ivar
type customAttributes struct {
	objectivec.Object
}

// customAttributesFrom constructs a [customAttributes] from an unsafe.Pointer.
func customAttributesFrom(ptr unsafe.Pointer) customAttributes {
	return customAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for customAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for customAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for customAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for customAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for customAttributes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class customAttributes */



