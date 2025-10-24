// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEModeClusterModeTagStruct */


/* debug [class_header]: Header for MTREnergyEVSEModeClusterModeTagStruct */
// The class instance for the [MTREnergyEVSEModeClusterModeTagStruct] class.
var (
	MTREnergyEVSEModeClusterModeTagStructClass     _MTREnergyEVSEModeClusterModeTagStructClass
	MTREnergyEVSEModeClusterModeTagStructClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterModeTagStructClass() _MTREnergyEVSEModeClusterModeTagStructClass {
	MTREnergyEVSEModeClusterModeTagStructClassOnce.Do(func() {
		MTREnergyEVSEModeClusterModeTagStructClass = _MTREnergyEVSEModeClusterModeTagStructClass{objc.GetClass("MTREnergyEVSEModeClusterModeTagStruct")}
	})
	return MTREnergyEVSEModeClusterModeTagStructClass
}

type _MTREnergyEVSEModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEModeClusterModeTagStruct */
// An interface definition for the [MTREnergyEVSEModeClusterModeTagStruct] class.
type IMTREnergyEVSEModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEModeClusterModeTagStruct */
	// properties:
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterModeTagStructClass) Alloc() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEModeClusterModeTagStructClass) New() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Init() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Autorelease() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterModeTagStruct creates a new MTREnergyEVSEModeClusterModeTagStruct instance.
func NewMTREnergyEVSEModeClusterModeTagStruct() MTREnergyEVSEModeClusterModeTagStruct {
	return getMTREnergyEVSEModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct
type MTREnergyEVSEModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterModeTagStructFrom constructs a [MTREnergyEVSEModeClusterModeTagStruct] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterModeTagStruct {
	return MTREnergyEVSEModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/value
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/value
func (m_ MTREnergyEVSEModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclustermodetagstruct/mfgcode
func (m_ MTREnergyEVSEModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevsemodeclustermodetagstruct/mfgcode
func (m_ MTREnergyEVSEModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEModeClusterModeTagStruct */



