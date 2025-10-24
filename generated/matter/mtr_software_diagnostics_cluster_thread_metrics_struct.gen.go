// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSoftwareDiagnosticsClusterThreadMetricsStruct */


/* debug [class_header]: Header for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClass     _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass() _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsStructClass = _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetricsStruct")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
type IMTRSoftwareDiagnosticsClusterThreadMetricsStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
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

	
/* debug [class_interface_methods]: Methods for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) New() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Init() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct creates a new MTRSoftwareDiagnosticsClusterThreadMetricsStruct instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct
type MTRSoftwareDiagnosticsClusterThreadMetricsStruct struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return MTRSoftwareDiagnosticsClusterThreadMetricsStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSoftwareDiagnosticsClusterThreadMetricsStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSoftwareDiagnosticsClusterThreadMetricsStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}/* debug [instance_properties/getter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}/* debug [instance_properties/setter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackFreeCurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeCurrent"))
	return rv
}/* debug [instance_properties/getter]: stackFreeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackFreeCurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeCurrent:"), value)
}/* debug [instance_properties/setter]: stackFreeCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackFreeMinimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeMinimum"))
	return rv
}/* debug [instance_properties/getter]: stackFreeMinimum */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackFreeMinimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeMinimum:"), value)
}/* debug [instance_properties/setter]: stackFreeMinimum */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackSize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackSize"))
	return rv
}/* debug [instance_properties/getter]: stackSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct/stackSize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackSize:"), value)
}/* debug [instance_properties/setter]: stackSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSoftwareDiagnosticsClusterThreadMetricsStruct */



