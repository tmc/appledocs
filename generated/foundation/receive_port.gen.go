// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class receivePort */


/* debug [class_header]: Header for receivePort */
// The class instance for the [receivePort] class.
var (
	ReceivePortClass     _receivePortClass
	ReceivePortClassOnce sync.Once
)

func getreceivePortClass() _receivePortClass {
	ReceivePortClassOnce.Do(func() {
		ReceivePortClass = _receivePortClass{objc.GetClass("receivePort")}
	})
	return ReceivePortClass
}

type _receivePortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for receivePort */
// An interface definition for the [receivePort] class.
type IreceivePort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for receivePort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for receivePort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for receivePort */
// Alloc allocates a new instance without initialization.
func (rc _receivePortClass) Alloc() receivePort {
	rv := objc.Send[receivePort](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _receivePortClass) New() receivePort {
	rv := objc.Send[receivePort](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ receivePort) Init() receivePort {
	rv := objc.Send[receivePort](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ receivePort) Autorelease() receivePort {
	rv := objc.Send[receivePort](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreceivePort creates a new receivePort instance.
func NewreceivePort() receivePort {
	return getreceivePortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for receivePort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/receivePort-c.ivar
type receivePort struct {
	objectivec.Object
}

// receivePortFrom constructs a [receivePort] from an unsafe.Pointer.
func receivePortFrom(ptr unsafe.Pointer) receivePort {
	return receivePort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for receivePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for receivePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for receivePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for receivePort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for receivePort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class receivePort */



