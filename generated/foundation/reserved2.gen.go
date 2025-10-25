// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class reserved2 */


/* debug [class_header]: Header for reserved2 */
// The class instance for the [reserved2] class.
var (
	Reserved2Class     _reserved2Class
	Reserved2ClassOnce sync.Once
)

func getreserved2Class() _reserved2Class {
	Reserved2ClassOnce.Do(func() {
		Reserved2Class = _reserved2Class{objc.GetClass("reserved2")}
	})
	return Reserved2Class
}

type _reserved2Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for reserved2 */
// An interface definition for the [reserved2] class.
type Ireserved2 interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for reserved2 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for reserved2 */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for reserved2 */
// Alloc allocates a new instance without initialization.
func (rc _reserved2Class) Alloc() reserved2 {
	rv := objc.Send[reserved2](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _reserved2Class) New() reserved2 {
	rv := objc.Send[reserved2](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reserved2) Init() reserved2 {
	rv := objc.Send[reserved2](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reserved2) Autorelease() reserved2 {
	rv := objc.Send[reserved2](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreserved2 creates a new reserved2 instance.
func Newreserved2() reserved2 {
	return getreserved2Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for reserved2 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/reserved2
type reserved2 struct {
	objectivec.Object
}

// reserved2From constructs a [reserved2] from an unsafe.Pointer.
func reserved2From(ptr unsafe.Pointer) reserved2 {
	return reserved2{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for reserved2 *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for reserved2 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for reserved2 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for reserved2 */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for reserved2 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class reserved2 */



