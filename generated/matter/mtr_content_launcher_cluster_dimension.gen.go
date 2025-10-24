// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterDimension */


/* debug [class_header]: Header for MTRContentLauncherClusterDimension */
// The class instance for the [MTRContentLauncherClusterDimension] class.
var (
	MTRContentLauncherClusterDimensionClass     _MTRContentLauncherClusterDimensionClass
	MTRContentLauncherClusterDimensionClassOnce sync.Once
)

func getMTRContentLauncherClusterDimensionClass() _MTRContentLauncherClusterDimensionClass {
	MTRContentLauncherClusterDimensionClassOnce.Do(func() {
		MTRContentLauncherClusterDimensionClass = _MTRContentLauncherClusterDimensionClass{objc.GetClass("MTRContentLauncherClusterDimension")}
	})
	return MTRContentLauncherClusterDimensionClass
}

type _MTRContentLauncherClusterDimensionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterDimension */
// An interface definition for the [MTRContentLauncherClusterDimension] class.
type IMTRContentLauncherClusterDimension interface {
	IMTRContentLauncherClusterDimensionStruct
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterDimension */
	// properties:
	Height() objc.IObject /* cross-framework: NSNumber */
	SetHeight(value objc.IObject /* cross-framework: NSNumber */)
	Metric() objc.IObject /* cross-framework: NSNumber */
	SetMetric(value objc.IObject /* cross-framework: NSNumber */)
	Width() objc.IObject /* cross-framework: NSNumber */
	SetWidth(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterDimension */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterDimension */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterDimensionClass) Alloc() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterDimensionClass) New() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterDimension) Init() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterDimension) Autorelease() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterDimension creates a new MTRContentLauncherClusterDimension instance.
func NewMTRContentLauncherClusterDimension() MTRContentLauncherClusterDimension {
	return getMTRContentLauncherClusterDimensionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension
type MTRContentLauncherClusterDimension struct {
	MTRContentLauncherClusterDimensionStruct
}

// MTRContentLauncherClusterDimensionFrom constructs a [MTRContentLauncherClusterDimension] from an unsafe.Pointer.
func MTRContentLauncherClusterDimensionFrom(ptr unsafe.Pointer) MTRContentLauncherClusterDimension {
	return MTRContentLauncherClusterDimension{
		MTRContentLauncherClusterDimensionStruct: MTRContentLauncherClusterDimensionStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterDimension *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterDimension */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterDimension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterDimension */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterDimension */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/height
func (m_ MTRContentLauncherClusterDimension) Height() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/height
func (m_ MTRContentLauncherClusterDimension) SetHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/metric
func (m_ MTRContentLauncherClusterDimension) Metric() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("metric"))
	return rv
}/* debug [instance_properties/getter]: metric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/metric
func (m_ MTRContentLauncherClusterDimension) SetMetric(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetric:"), value)
}/* debug [instance_properties/setter]: metric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/width
func (m_ MTRContentLauncherClusterDimension) Width() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension/width
func (m_ MTRContentLauncherClusterDimension) SetWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterDimension */



