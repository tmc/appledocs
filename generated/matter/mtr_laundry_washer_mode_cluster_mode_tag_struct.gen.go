// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLaundryWasherModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRLaundryWasherModeClusterModeTagStruct */
// The class instance for the [MTRLaundryWasherModeClusterModeTagStruct] class.
var (
	MTRLaundryWasherModeClusterModeTagStructClass     _MTRLaundryWasherModeClusterModeTagStructClass
	MTRLaundryWasherModeClusterModeTagStructClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterModeTagStructClass() _MTRLaundryWasherModeClusterModeTagStructClass {
	MTRLaundryWasherModeClusterModeTagStructClassOnce.Do(func() {
		MTRLaundryWasherModeClusterModeTagStructClass = _MTRLaundryWasherModeClusterModeTagStructClass{objc.GetClass("MTRLaundryWasherModeClusterModeTagStruct")}
	})
	return MTRLaundryWasherModeClusterModeTagStructClass
}

type _MTRLaundryWasherModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLaundryWasherModeClusterModeTagStruct */
// An interface definition for the [MTRLaundryWasherModeClusterModeTagStruct] class.
type IMTRLaundryWasherModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLaundryWasherModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLaundryWasherModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLaundryWasherModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) Alloc() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) New() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Init() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Autorelease() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterModeTagStruct creates a new MTRLaundryWasherModeClusterModeTagStruct instance.
func NewMTRLaundryWasherModeClusterModeTagStruct() MTRLaundryWasherModeClusterModeTagStruct {
	return getMTRLaundryWasherModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLaundryWasherModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct
type MTRLaundryWasherModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterModeTagStructFrom constructs a [MTRLaundryWasherModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterModeTagStruct {
	return MTRLaundryWasherModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLaundryWasherModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLaundryWasherModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLaundryWasherModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLaundryWasherModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLaundryWasherModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclustermodetagstruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclustermodetagstruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLaundryWasherModeClusterModeTagStruct */



