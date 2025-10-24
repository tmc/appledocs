// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterServiceArea */


/* debug [class_header]: Header for MTRClusterServiceArea */
// The class instance for the [MTRClusterServiceArea] class.
var (
	MTRClusterServiceAreaClass     _MTRClusterServiceAreaClass
	MTRClusterServiceAreaClassOnce sync.Once
)

func getMTRClusterServiceAreaClass() _MTRClusterServiceAreaClass {
	MTRClusterServiceAreaClassOnce.Do(func() {
		MTRClusterServiceAreaClass = _MTRClusterServiceAreaClass{objc.GetClass("MTRClusterServiceArea")}
	})
	return MTRClusterServiceAreaClass
}

type _MTRClusterServiceAreaClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterServiceArea */
// An interface definition for the [MTRClusterServiceArea] class.
type IMTRClusterServiceArea interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterServiceArea */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterServiceArea */
	// methods:
	SelectAreasWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSelectAreasParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterServiceArea */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterServiceAreaClass) Alloc() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterServiceAreaClass) New() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterServiceArea) Init() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterServiceArea) Autorelease() MTRClusterServiceArea {
	rv := objc.Send[MTRClusterServiceArea](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterServiceArea creates a new MTRClusterServiceArea instance.
func NewMTRClusterServiceArea() MTRClusterServiceArea {
	return getMTRClusterServiceAreaClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterServiceArea */
// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.


// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea
type MTRClusterServiceArea struct {
	MTRGenericCluster
}

// MTRClusterServiceAreaFrom constructs a [MTRClusterServiceArea] from an unsafe.Pointer.
//
// Cluster Service Area The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.
func MTRClusterServiceAreaFrom(ptr unsafe.Pointer) MTRClusterServiceArea {
	return MTRClusterServiceArea{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterServiceArea *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterServiceArea */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterServiceArea */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterServiceArea */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterServiceArea/selectAreas(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterServiceArea) SelectAreasWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRServiceAreaClusterSelectAreasParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAreasWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}/* debug [instance_methods/method]: SelectAreasWithParamsExpectedValuesExpectedValueIntervalCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterServiceArea */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterServiceArea */



