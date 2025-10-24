// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterLaundryWasherMode */


/* debug [class_header]: Header for MTRBaseClusterLaundryWasherMode */
// The class instance for the [MTRBaseClusterLaundryWasherMode] class.
var (
	MTRBaseClusterLaundryWasherModeClass     _MTRBaseClusterLaundryWasherModeClass
	MTRBaseClusterLaundryWasherModeClassOnce sync.Once
)

func getMTRBaseClusterLaundryWasherModeClass() _MTRBaseClusterLaundryWasherModeClass {
	MTRBaseClusterLaundryWasherModeClassOnce.Do(func() {
		MTRBaseClusterLaundryWasherModeClass = _MTRBaseClusterLaundryWasherModeClass{objc.GetClass("MTRBaseClusterLaundryWasherMode")}
	})
	return MTRBaseClusterLaundryWasherModeClass
}

type _MTRBaseClusterLaundryWasherModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterLaundryWasherMode */
// An interface definition for the [MTRBaseClusterLaundryWasherMode] class.
type IMTRBaseClusterLaundryWasherMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterLaundryWasherMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterLaundryWasherMode */
	// methods:
	ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterLaundryWasherMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLaundryWasherModeClass) Alloc() MTRBaseClusterLaundryWasherMode {
	rv := objc.Send[MTRBaseClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterLaundryWasherModeClass) New() MTRBaseClusterLaundryWasherMode {
	rv := objc.Send[MTRBaseClusterLaundryWasherMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLaundryWasherMode) Init() MTRBaseClusterLaundryWasherMode {
	rv := objc.Send[MTRBaseClusterLaundryWasherMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLaundryWasherMode) Autorelease() MTRBaseClusterLaundryWasherMode {
	rv := objc.Send[MTRBaseClusterLaundryWasherMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLaundryWasherMode creates a new MTRBaseClusterLaundryWasherMode instance.
func NewMTRBaseClusterLaundryWasherMode() MTRBaseClusterLaundryWasherMode {
	return getMTRBaseClusterLaundryWasherModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterLaundryWasherMode */
// Cluster Laundry Washer Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Laundry Washer Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherMode
type MTRBaseClusterLaundryWasherMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLaundryWasherModeFrom constructs a [MTRBaseClusterLaundryWasherMode] from an unsafe.Pointer.
//
// Cluster Laundry Washer Mode
func MTRBaseClusterLaundryWasherModeFrom(ptr unsafe.Pointer) MTRBaseClusterLaundryWasherMode {
	return MTRBaseClusterLaundryWasherMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterLaundryWasherMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterLaundryWasherMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterLaundryWasherMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterLaundryWasherMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherMode/readAttributeCurrentMode(completion:)
func (m_ MTRBaseClusterLaundryWasherMode) ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentModeWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterLaundryWasherMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterLaundryWasherMode */



