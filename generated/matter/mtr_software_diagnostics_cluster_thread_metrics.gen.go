// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRSoftwareDiagnosticsClusterThreadMetrics */


/* debug [class_header]: Header for MTRSoftwareDiagnosticsClusterThreadMetrics */
// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsClass     _MTRSoftwareDiagnosticsClusterThreadMetricsClass
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsClass() _MTRSoftwareDiagnosticsClusterThreadMetricsClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsClass = _MTRSoftwareDiagnosticsClusterThreadMetricsClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetrics")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSoftwareDiagnosticsClusterThreadMetrics */
// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
type IMTRSoftwareDiagnosticsClusterThreadMetrics interface {
	IMTRSoftwareDiagnosticsClusterThreadMetricsStruct
	
/* debug [class_interface_properties]: Properties for MTRSoftwareDiagnosticsClusterThreadMetrics */
	// properties:
	Id() objc.IObject /* cross-framework: NSNumber */
	SetId(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */
	SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */)
	StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */
	SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */)
	StackSize() objc.IObject /* cross-framework: NSNumber */
	SetStackSize(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSoftwareDiagnosticsClusterThreadMetrics */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSoftwareDiagnosticsClusterThreadMetrics */
// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) New() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Init() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetrics creates a new MTRSoftwareDiagnosticsClusterThreadMetrics instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetrics() MTRSoftwareDiagnosticsClusterThreadMetrics {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSoftwareDiagnosticsClusterThreadMetrics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics
type MTRSoftwareDiagnosticsClusterThreadMetrics struct {
	MTRSoftwareDiagnosticsClusterThreadMetricsStruct
}

// MTRSoftwareDiagnosticsClusterThreadMetricsFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetrics] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetrics {
	return MTRSoftwareDiagnosticsClusterThreadMetrics{
		MTRSoftwareDiagnosticsClusterThreadMetricsStruct: MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSoftwareDiagnosticsClusterThreadMetrics *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSoftwareDiagnosticsClusterThreadMetrics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSoftwareDiagnosticsClusterThreadMetrics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSoftwareDiagnosticsClusterThreadMetrics */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSoftwareDiagnosticsClusterThreadMetrics */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}/* debug [instance_properties/getter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}/* debug [instance_properties/setter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackFreeCurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeCurrent"))
	return rv
}/* debug [instance_properties/getter]: stackFreeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackFreeCurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeCurrent:"), value)
}/* debug [instance_properties/setter]: stackFreeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackFreeMinimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeMinimum"))
	return rv
}/* debug [instance_properties/getter]: stackFreeMinimum */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackFreeMinimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeMinimum:"), value)
}/* debug [instance_properties/setter]: stackFreeMinimum */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackSize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackSize"))
	return rv
}/* debug [instance_properties/getter]: stackSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics/stackSize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackSize:"), value)
}/* debug [instance_properties/setter]: stackSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSoftwareDiagnosticsClusterThreadMetrics */



