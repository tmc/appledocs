// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRControllerFactory */


/* debug [class_header]: Header for MTRControllerFactory */
// The class instance for the [MTRControllerFactory] class.
var (
	MTRControllerFactoryClass     _MTRControllerFactoryClass
	MTRControllerFactoryClassOnce sync.Once
)

func getMTRControllerFactoryClass() _MTRControllerFactoryClass {
	MTRControllerFactoryClassOnce.Do(func() {
		MTRControllerFactoryClass = _MTRControllerFactoryClass{objc.GetClass("MTRControllerFactory")}
	})
	return MTRControllerFactoryClass
}

type _MTRControllerFactoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRControllerFactory */
// An interface definition for the [MTRControllerFactory] class.
type IMTRControllerFactory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRControllerFactory */
	// properties:
	IsRunning() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRControllerFactory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRControllerFactory */
// Alloc allocates a new instance without initialization.
func (mc _MTRControllerFactoryClass) Alloc() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRControllerFactoryClass) New() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRControllerFactory) Init() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRControllerFactory) Autorelease() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRControllerFactory creates a new MTRControllerFactory instance.
func NewMTRControllerFactory() MTRControllerFactory {
	return getMTRControllerFactoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRControllerFactory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactory
type MTRControllerFactory struct {
	objectivec.Object
}

// MTRControllerFactoryFrom constructs a [MTRControllerFactory] from an unsafe.Pointer.
func MTRControllerFactoryFrom(ptr unsafe.Pointer) MTRControllerFactory {
	return MTRControllerFactory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRControllerFactory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRControllerFactory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactory/sharedInstance()
func (mc _MTRControllerFactoryClass) SharedInstance() MTRControllerFactory {
	rv := objc.Send[MTRControllerFactory](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRControllerFactory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRControllerFactory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRControllerFactory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRControllerFactory/isRunning
func (m_ MTRControllerFactory) IsRunning() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRControllerFactory */



