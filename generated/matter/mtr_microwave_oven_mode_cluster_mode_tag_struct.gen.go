// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMicrowaveOvenModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRMicrowaveOvenModeClusterModeTagStruct */
// The class instance for the [MTRMicrowaveOvenModeClusterModeTagStruct] class.
var (
	MTRMicrowaveOvenModeClusterModeTagStructClass     _MTRMicrowaveOvenModeClusterModeTagStructClass
	MTRMicrowaveOvenModeClusterModeTagStructClassOnce sync.Once
)

func getMTRMicrowaveOvenModeClusterModeTagStructClass() _MTRMicrowaveOvenModeClusterModeTagStructClass {
	MTRMicrowaveOvenModeClusterModeTagStructClassOnce.Do(func() {
		MTRMicrowaveOvenModeClusterModeTagStructClass = _MTRMicrowaveOvenModeClusterModeTagStructClass{objc.GetClass("MTRMicrowaveOvenModeClusterModeTagStruct")}
	})
	return MTRMicrowaveOvenModeClusterModeTagStructClass
}

type _MTRMicrowaveOvenModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMicrowaveOvenModeClusterModeTagStruct */
// An interface definition for the [MTRMicrowaveOvenModeClusterModeTagStruct] class.
type IMTRMicrowaveOvenModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMicrowaveOvenModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMicrowaveOvenModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMicrowaveOvenModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenModeClusterModeTagStructClass) Alloc() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMicrowaveOvenModeClusterModeTagStructClass) New() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Init() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Autorelease() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenModeClusterModeTagStruct creates a new MTRMicrowaveOvenModeClusterModeTagStruct instance.
func NewMTRMicrowaveOvenModeClusterModeTagStruct() MTRMicrowaveOvenModeClusterModeTagStruct {
	return getMTRMicrowaveOvenModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMicrowaveOvenModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct
type MTRMicrowaveOvenModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRMicrowaveOvenModeClusterModeTagStructFrom constructs a [MTRMicrowaveOvenModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRMicrowaveOvenModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRMicrowaveOvenModeClusterModeTagStruct {
	return MTRMicrowaveOvenModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMicrowaveOvenModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMicrowaveOvenModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMicrowaveOvenModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMicrowaveOvenModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMicrowaveOvenModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/mfgCode
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/mfgCode
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovenmodeclustermodetagstruct/value
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovenmodeclustermodetagstruct/value
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMicrowaveOvenModeClusterModeTagStruct */



