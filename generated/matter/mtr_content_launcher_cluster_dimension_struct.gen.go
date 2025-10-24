// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRContentLauncherClusterDimensionStruct */


/* debug [class_header]: Header for MTRContentLauncherClusterDimensionStruct */
// The class instance for the [MTRContentLauncherClusterDimensionStruct] class.
var (
	MTRContentLauncherClusterDimensionStructClass     _MTRContentLauncherClusterDimensionStructClass
	MTRContentLauncherClusterDimensionStructClassOnce sync.Once
)

func getMTRContentLauncherClusterDimensionStructClass() _MTRContentLauncherClusterDimensionStructClass {
	MTRContentLauncherClusterDimensionStructClassOnce.Do(func() {
		MTRContentLauncherClusterDimensionStructClass = _MTRContentLauncherClusterDimensionStructClass{objc.GetClass("MTRContentLauncherClusterDimensionStruct")}
	})
	return MTRContentLauncherClusterDimensionStructClass
}

type _MTRContentLauncherClusterDimensionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRContentLauncherClusterDimensionStruct */
// An interface definition for the [MTRContentLauncherClusterDimensionStruct] class.
type IMTRContentLauncherClusterDimensionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRContentLauncherClusterDimensionStruct */
	// properties:
	Height() objc.IObject /* cross-framework: NSNumber */
	SetHeight(value objc.IObject /* cross-framework: NSNumber */)
	Metric() objc.IObject /* cross-framework: NSNumber */
	SetMetric(value objc.IObject /* cross-framework: NSNumber */)
	Width() objc.IObject /* cross-framework: NSNumber */
	SetWidth(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRContentLauncherClusterDimensionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRContentLauncherClusterDimensionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterDimensionStructClass) Alloc() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRContentLauncherClusterDimensionStructClass) New() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterDimensionStruct) Init() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterDimensionStruct) Autorelease() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterDimensionStruct creates a new MTRContentLauncherClusterDimensionStruct instance.
func NewMTRContentLauncherClusterDimensionStruct() MTRContentLauncherClusterDimensionStruct {
	return getMTRContentLauncherClusterDimensionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRContentLauncherClusterDimensionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct
type MTRContentLauncherClusterDimensionStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterDimensionStructFrom constructs a [MTRContentLauncherClusterDimensionStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterDimensionStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterDimensionStruct {
	return MTRContentLauncherClusterDimensionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRContentLauncherClusterDimensionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRContentLauncherClusterDimensionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRContentLauncherClusterDimensionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRContentLauncherClusterDimensionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRContentLauncherClusterDimensionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/height
func (m_ MTRContentLauncherClusterDimensionStruct) Height() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/height
func (m_ MTRContentLauncherClusterDimensionStruct) SetHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/metric
func (m_ MTRContentLauncherClusterDimensionStruct) Metric() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("metric"))
	return rv
}/* debug [instance_properties/getter]: metric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/metric
func (m_ MTRContentLauncherClusterDimensionStruct) SetMetric(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetric:"), value)
}/* debug [instance_properties/setter]: metric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/width
func (m_ MTRContentLauncherClusterDimensionStruct) Width() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct/width
func (m_ MTRContentLauncherClusterDimensionStruct) SetWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRContentLauncherClusterDimensionStruct */



