// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAttributeCacheContainer */


/* debug [class_header]: Header for MTRAttributeCacheContainer */
// The class instance for the [MTRAttributeCacheContainer] class.
var (
	MTRAttributeCacheContainerClass     _MTRAttributeCacheContainerClass
	MTRAttributeCacheContainerClassOnce sync.Once
)

func getMTRAttributeCacheContainerClass() _MTRAttributeCacheContainerClass {
	MTRAttributeCacheContainerClassOnce.Do(func() {
		MTRAttributeCacheContainerClass = _MTRAttributeCacheContainerClass{objc.GetClass("MTRAttributeCacheContainer")}
	})
	return MTRAttributeCacheContainerClass
}

type _MTRAttributeCacheContainerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAttributeCacheContainer */
// An interface definition for the [MTRAttributeCacheContainer] class.
type IMTRAttributeCacheContainer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAttributeCacheContainer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAttributeCacheContainer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAttributeCacheContainer */
// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeCacheContainerClass) Alloc() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAttributeCacheContainerClass) New() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeCacheContainer) Init() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeCacheContainer) Autorelease() MTRAttributeCacheContainer {
	rv := objc.Send[MTRAttributeCacheContainer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeCacheContainer creates a new MTRAttributeCacheContainer instance.
func NewMTRAttributeCacheContainer() MTRAttributeCacheContainer {
	return getMTRAttributeCacheContainerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAttributeCacheContainer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeCacheContainer
type MTRAttributeCacheContainer struct {
	objectivec.Object
}

// MTRAttributeCacheContainerFrom constructs a [MTRAttributeCacheContainer] from an unsafe.Pointer.
func MTRAttributeCacheContainerFrom(ptr unsafe.Pointer) MTRAttributeCacheContainer {
	return MTRAttributeCacheContainer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAttributeCacheContainer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAttributeCacheContainer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAttributeCacheContainer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAttributeCacheContainer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAttributeCacheContainer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAttributeCacheContainer */



