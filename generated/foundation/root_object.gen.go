// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class rootObject */


/* debug [class_header]: Header for rootObject */
// The class instance for the [rootObject] class.
var (
	RootObjectClass     _rootObjectClass
	RootObjectClassOnce sync.Once
)

func getrootObjectClass() _rootObjectClass {
	RootObjectClassOnce.Do(func() {
		RootObjectClass = _rootObjectClass{objc.GetClass("rootObject")}
	})
	return RootObjectClass
}

type _rootObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for rootObject */
// An interface definition for the [rootObject] class.
type IrootObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for rootObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for rootObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for rootObject */
// Alloc allocates a new instance without initialization.
func (rc _rootObjectClass) Alloc() rootObject {
	rv := objc.Send[rootObject](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _rootObjectClass) New() rootObject {
	rv := objc.Send[rootObject](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rootObject) Init() rootObject {
	rv := objc.Send[rootObject](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rootObject) Autorelease() rootObject {
	rv := objc.Send[rootObject](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrootObject creates a new rootObject instance.
func NewrootObject() rootObject {
	return getrootObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for rootObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootObject-c.ivar
type rootObject struct {
	objectivec.Object
}

// rootObjectFrom constructs a [rootObject] from an unsafe.Pointer.
func rootObjectFrom(ptr unsafe.Pointer) rootObject {
	return rootObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for rootObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for rootObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for rootObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for rootObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for rootObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class rootObject */



