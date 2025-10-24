// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class attributes */


/* debug [class_header]: Header for attributes */
// The class instance for the [attributes] class.
var (
	AttributesClass     _attributesClass
	AttributesClassOnce sync.Once
)

func getattributesClass() _attributesClass {
	AttributesClassOnce.Do(func() {
		AttributesClass = _attributesClass{objc.GetClass("attributes")}
	})
	return AttributesClass
}

type _attributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for attributes */
// An interface definition for the [attributes] class.
type Iattributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for attributes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for attributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for attributes */
// Alloc allocates a new instance without initialization.
func (ac _attributesClass) Alloc() attributes {
	rv := objc.Send[attributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _attributesClass) New() attributes {
	rv := objc.Send[attributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ attributes) Init() attributes {
	rv := objc.Send[attributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ attributes) Autorelease() attributes {
	rv := objc.Send[attributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// Newattributes creates a new attributes instance.
func Newattributes() attributes {
	return getattributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for attributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-c.ivar
type attributes struct {
	objectivec.Object
}

// attributesFrom constructs a [attributes] from an unsafe.Pointer.
func attributesFrom(ptr unsafe.Pointer) attributes {
	return attributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for attributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for attributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for attributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for attributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for attributes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class attributes */



