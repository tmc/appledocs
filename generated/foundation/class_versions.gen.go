// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class classVersions */


/* debug [class_header]: Header for classVersions */
// The class instance for the [classVersions] class.
var (
	ClassVersionsClass     _classVersionsClass
	ClassVersionsClassOnce sync.Once
)

func getclassVersionsClass() _classVersionsClass {
	ClassVersionsClassOnce.Do(func() {
		ClassVersionsClass = _classVersionsClass{objc.GetClass("classVersions")}
	})
	return ClassVersionsClass
}

type _classVersionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for classVersions */
// An interface definition for the [classVersions] class.
type IclassVersions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for classVersions */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for classVersions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for classVersions */
// Alloc allocates a new instance without initialization.
func (cc _classVersionsClass) Alloc() classVersions {
	rv := objc.Send[classVersions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _classVersionsClass) New() classVersions {
	rv := objc.Send[classVersions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ classVersions) Init() classVersions {
	rv := objc.Send[classVersions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ classVersions) Autorelease() classVersions {
	rv := objc.Send[classVersions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewclassVersions creates a new classVersions instance.
func NewclassVersions() classVersions {
	return getclassVersionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for classVersions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/classVersions
type classVersions struct {
	objectivec.Object
}

// classVersionsFrom constructs a [classVersions] from an unsafe.Pointer.
func classVersionsFrom(ptr unsafe.Pointer) classVersions {
	return classVersions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for classVersions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for classVersions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for classVersions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for classVersions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for classVersions */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class classVersions */



