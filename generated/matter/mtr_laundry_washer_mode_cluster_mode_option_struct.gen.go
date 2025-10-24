// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLaundryWasherModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRLaundryWasherModeClusterModeOptionStruct */
// The class instance for the [MTRLaundryWasherModeClusterModeOptionStruct] class.
var (
	MTRLaundryWasherModeClusterModeOptionStructClass     _MTRLaundryWasherModeClusterModeOptionStructClass
	MTRLaundryWasherModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterModeOptionStructClass() _MTRLaundryWasherModeClusterModeOptionStructClass {
	MTRLaundryWasherModeClusterModeOptionStructClassOnce.Do(func() {
		MTRLaundryWasherModeClusterModeOptionStructClass = _MTRLaundryWasherModeClusterModeOptionStructClass{objc.GetClass("MTRLaundryWasherModeClusterModeOptionStruct")}
	})
	return MTRLaundryWasherModeClusterModeOptionStructClass
}

type _MTRLaundryWasherModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLaundryWasherModeClusterModeOptionStruct */
// An interface definition for the [MTRLaundryWasherModeClusterModeOptionStruct] class.
type IMTRLaundryWasherModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLaundryWasherModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLaundryWasherModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLaundryWasherModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterModeOptionStructClass) Alloc() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLaundryWasherModeClusterModeOptionStructClass) New() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Init() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Autorelease() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterModeOptionStruct creates a new MTRLaundryWasherModeClusterModeOptionStruct instance.
func NewMTRLaundryWasherModeClusterModeOptionStruct() MTRLaundryWasherModeClusterModeOptionStruct {
	return getMTRLaundryWasherModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLaundryWasherModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct
type MTRLaundryWasherModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterModeOptionStructFrom constructs a [MTRLaundryWasherModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterModeOptionStruct {
	return MTRLaundryWasherModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLaundryWasherModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLaundryWasherModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLaundryWasherModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLaundryWasherModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLaundryWasherModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/label
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/label
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclustermodeoptionstruct/mode
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclustermodeoptionstruct/mode
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLaundryWasherModeClusterModeOptionStruct */



