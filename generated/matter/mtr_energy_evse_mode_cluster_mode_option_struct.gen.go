// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTREnergyEVSEModeClusterModeOptionStruct */
// The class instance for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
var (
	MTREnergyEVSEModeClusterModeOptionStructClass     _MTREnergyEVSEModeClusterModeOptionStructClass
	MTREnergyEVSEModeClusterModeOptionStructClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterModeOptionStructClass() _MTREnergyEVSEModeClusterModeOptionStructClass {
	MTREnergyEVSEModeClusterModeOptionStructClassOnce.Do(func() {
		MTREnergyEVSEModeClusterModeOptionStructClass = _MTREnergyEVSEModeClusterModeOptionStructClass{objc.GetClass("MTREnergyEVSEModeClusterModeOptionStruct")}
	})
	return MTREnergyEVSEModeClusterModeOptionStructClass
}

type _MTREnergyEVSEModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEModeClusterModeOptionStruct */
// An interface definition for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
type IMTREnergyEVSEModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) Alloc() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) New() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Init() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Autorelease() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterModeOptionStruct creates a new MTREnergyEVSEModeClusterModeOptionStruct instance.
func NewMTREnergyEVSEModeClusterModeOptionStruct() MTREnergyEVSEModeClusterModeOptionStruct {
	return getMTREnergyEVSEModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct
type MTREnergyEVSEModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterModeOptionStructFrom constructs a [MTREnergyEVSEModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterModeOptionStruct {
	return MTREnergyEVSEModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclustermodeoptionstruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclustermodeoptionstruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEModeClusterModeOptionStruct */



