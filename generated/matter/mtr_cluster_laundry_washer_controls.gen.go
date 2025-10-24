// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterLaundryWasherControls */


/* debug [class_header]: Header for MTRClusterLaundryWasherControls */
// The class instance for the [MTRClusterLaundryWasherControls] class.
var (
	MTRClusterLaundryWasherControlsClass     _MTRClusterLaundryWasherControlsClass
	MTRClusterLaundryWasherControlsClassOnce sync.Once
)

func getMTRClusterLaundryWasherControlsClass() _MTRClusterLaundryWasherControlsClass {
	MTRClusterLaundryWasherControlsClassOnce.Do(func() {
		MTRClusterLaundryWasherControlsClass = _MTRClusterLaundryWasherControlsClass{objc.GetClass("MTRClusterLaundryWasherControls")}
	})
	return MTRClusterLaundryWasherControlsClass
}

type _MTRClusterLaundryWasherControlsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterLaundryWasherControls */
// An interface definition for the [MTRClusterLaundryWasherControls] class.
type IMTRClusterLaundryWasherControls interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterLaundryWasherControls */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterLaundryWasherControls */
	// methods:
	WriteAttributeSpinSpeedCurrentWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterLaundryWasherControls */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterLaundryWasherControlsClass) Alloc() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterLaundryWasherControlsClass) New() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterLaundryWasherControls) Init() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterLaundryWasherControls) Autorelease() MTRClusterLaundryWasherControls {
	rv := objc.Send[MTRClusterLaundryWasherControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterLaundryWasherControls creates a new MTRClusterLaundryWasherControls instance.
func NewMTRClusterLaundryWasherControls() MTRClusterLaundryWasherControls {
	return getMTRClusterLaundryWasherControlsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterLaundryWasherControls */
// Cluster Laundry Washer Controls This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.


// Cluster Laundry Washer Controls This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls
type MTRClusterLaundryWasherControls struct {
	MTRGenericCluster
}

// MTRClusterLaundryWasherControlsFrom constructs a [MTRClusterLaundryWasherControls] from an unsafe.Pointer.
//
// Cluster Laundry Washer Controls This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.
func MTRClusterLaundryWasherControlsFrom(ptr unsafe.Pointer) MTRClusterLaundryWasherControls {
	return MTRClusterLaundryWasherControls{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterLaundryWasherControls *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterLaundryWasherControls */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterLaundryWasherControls */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterLaundryWasherControls */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterLaundryWasherControls/writeAttributeSpinSpeedCurrent(withValue:expectedValueInterval:)
func (m_ MTRClusterLaundryWasherControls) WriteAttributeSpinSpeedCurrentWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSpinSpeedCurrentWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}/* debug [instance_methods/method]: WriteAttributeSpinSpeedCurrentWithValueExpectedValueInterval */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterLaundryWasherControls */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterLaundryWasherControls */



