// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRWaterHeaterModeClusterModeTagStruct */
// The class instance for the [MTRWaterHeaterModeClusterModeTagStruct] class.
var (
	MTRWaterHeaterModeClusterModeTagStructClass     _MTRWaterHeaterModeClusterModeTagStructClass
	MTRWaterHeaterModeClusterModeTagStructClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterModeTagStructClass() _MTRWaterHeaterModeClusterModeTagStructClass {
	MTRWaterHeaterModeClusterModeTagStructClassOnce.Do(func() {
		MTRWaterHeaterModeClusterModeTagStructClass = _MTRWaterHeaterModeClusterModeTagStructClass{objc.GetClass("MTRWaterHeaterModeClusterModeTagStruct")}
	})
	return MTRWaterHeaterModeClusterModeTagStructClass
}

type _MTRWaterHeaterModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterModeClusterModeTagStruct */
// An interface definition for the [MTRWaterHeaterModeClusterModeTagStruct] class.
type IMTRWaterHeaterModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterModeClusterModeTagStruct */
	// properties:
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterModeTagStructClass) Alloc() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterModeClusterModeTagStructClass) New() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Init() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Autorelease() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterModeTagStruct creates a new MTRWaterHeaterModeClusterModeTagStruct instance.
func NewMTRWaterHeaterModeClusterModeTagStruct() MTRWaterHeaterModeClusterModeTagStruct {
	return getMTRWaterHeaterModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct
type MTRWaterHeaterModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterModeTagStructFrom constructs a [MTRWaterHeaterModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterModeTagStruct {
	return MTRWaterHeaterModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/value
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/value
func (m_ MTRWaterHeaterModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodetagstruct/mfgcode
func (m_ MTRWaterHeaterModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodetagstruct/mfgcode
func (m_ MTRWaterHeaterModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterModeClusterModeTagStruct */



