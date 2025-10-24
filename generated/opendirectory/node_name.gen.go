// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class nodeName */


/* debug [class_header]: Header for nodeName */
// The class instance for the [nodeName] class.
var (
	NodeNameClass     _nodeNameClass
	NodeNameClassOnce sync.Once
)

func getnodeNameClass() _nodeNameClass {
	NodeNameClassOnce.Do(func() {
		NodeNameClass = _nodeNameClass{objc.GetClass("nodeName")}
	})
	return NodeNameClass
}

type _nodeNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for nodeName */
// An interface definition for the [nodeName] class.
type InodeName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for nodeName */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for nodeName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for nodeName */
// Alloc allocates a new instance without initialization.
func (nc _nodeNameClass) Alloc() nodeName {
	rv := objc.Send[nodeName](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _nodeNameClass) New() nodeName {
	rv := objc.Send[nodeName](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ nodeName) Init() nodeName {
	rv := objc.Send[nodeName](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ nodeName) Autorelease() nodeName {
	rv := objc.Send[nodeName](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewnodeName creates a new nodeName instance.
func NewnodeName() nodeName {
	return getnodeNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for nodeName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-c.ivar
type nodeName struct {
	objectivec.Object
}

// nodeNameFrom constructs a [nodeName] from an unsafe.Pointer.
func nodeNameFrom(ptr unsafe.Pointer) nodeName {
	return nodeName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for nodeName *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for nodeName */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for nodeName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for nodeName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for nodeName */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class nodeName */



