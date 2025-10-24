// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRDishwasherModeClusterModeTagStruct */
// The class instance for the [MTRDishwasherModeClusterModeTagStruct] class.
var (
	MTRDishwasherModeClusterModeTagStructClass     _MTRDishwasherModeClusterModeTagStructClass
	MTRDishwasherModeClusterModeTagStructClassOnce sync.Once
)

func getMTRDishwasherModeClusterModeTagStructClass() _MTRDishwasherModeClusterModeTagStructClass {
	MTRDishwasherModeClusterModeTagStructClassOnce.Do(func() {
		MTRDishwasherModeClusterModeTagStructClass = _MTRDishwasherModeClusterModeTagStructClass{objc.GetClass("MTRDishwasherModeClusterModeTagStruct")}
	})
	return MTRDishwasherModeClusterModeTagStructClass
}

type _MTRDishwasherModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherModeClusterModeTagStruct */
// An interface definition for the [MTRDishwasherModeClusterModeTagStruct] class.
type IMTRDishwasherModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterModeTagStructClass) Alloc() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherModeClusterModeTagStructClass) New() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterModeTagStruct) Init() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterModeTagStruct) Autorelease() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterModeTagStruct creates a new MTRDishwasherModeClusterModeTagStruct instance.
func NewMTRDishwasherModeClusterModeTagStruct() MTRDishwasherModeClusterModeTagStruct {
	return getMTRDishwasherModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct
type MTRDishwasherModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRDishwasherModeClusterModeTagStructFrom constructs a [MTRDishwasherModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRDishwasherModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterModeTagStruct {
	return MTRDishwasherModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/mfgCode
func (m_ MTRDishwasherModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/mfgCode
func (m_ MTRDishwasherModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclustermodetagstruct/value
func (m_ MTRDishwasherModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclustermodetagstruct/value
func (m_ MTRDishwasherModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherModeClusterModeTagStruct */



