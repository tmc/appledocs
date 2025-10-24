// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenModeClusterModeTagStruct */


/* debug [class_header]: Header for MTROvenModeClusterModeTagStruct */
// The class instance for the [MTROvenModeClusterModeTagStruct] class.
var (
	MTROvenModeClusterModeTagStructClass     _MTROvenModeClusterModeTagStructClass
	MTROvenModeClusterModeTagStructClassOnce sync.Once
)

func getMTROvenModeClusterModeTagStructClass() _MTROvenModeClusterModeTagStructClass {
	MTROvenModeClusterModeTagStructClassOnce.Do(func() {
		MTROvenModeClusterModeTagStructClass = _MTROvenModeClusterModeTagStructClass{objc.GetClass("MTROvenModeClusterModeTagStruct")}
	})
	return MTROvenModeClusterModeTagStructClass
}

type _MTROvenModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenModeClusterModeTagStruct */
// An interface definition for the [MTROvenModeClusterModeTagStruct] class.
type IMTROvenModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterModeTagStructClass) Alloc() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenModeClusterModeTagStructClass) New() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterModeTagStruct) Init() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterModeTagStruct) Autorelease() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterModeTagStruct creates a new MTROvenModeClusterModeTagStruct instance.
func NewMTROvenModeClusterModeTagStruct() MTROvenModeClusterModeTagStruct {
	return getMTROvenModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct
type MTROvenModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTROvenModeClusterModeTagStructFrom constructs a [MTROvenModeClusterModeTagStruct] from an unsafe.Pointer.
func MTROvenModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTROvenModeClusterModeTagStruct {
	return MTROvenModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/mfgCode
func (m_ MTROvenModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/mfgCode
func (m_ MTROvenModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclustermodetagstruct/value
func (m_ MTROvenModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclustermodetagstruct/value
func (m_ MTROvenModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenModeClusterModeTagStruct */



