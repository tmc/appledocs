// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRWaterHeaterModeClusterModeOptionStruct */
// The class instance for the [MTRWaterHeaterModeClusterModeOptionStruct] class.
var (
	MTRWaterHeaterModeClusterModeOptionStructClass     _MTRWaterHeaterModeClusterModeOptionStructClass
	MTRWaterHeaterModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterModeOptionStructClass() _MTRWaterHeaterModeClusterModeOptionStructClass {
	MTRWaterHeaterModeClusterModeOptionStructClassOnce.Do(func() {
		MTRWaterHeaterModeClusterModeOptionStructClass = _MTRWaterHeaterModeClusterModeOptionStructClass{objc.GetClass("MTRWaterHeaterModeClusterModeOptionStruct")}
	})
	return MTRWaterHeaterModeClusterModeOptionStructClass
}

type _MTRWaterHeaterModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterModeClusterModeOptionStruct */
// An interface definition for the [MTRWaterHeaterModeClusterModeOptionStruct] class.
type IMTRWaterHeaterModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterModeClusterModeOptionStruct */
	// properties:
	ModeTags() objc.IObject /* cross-framework: NSArray */
	SetModeTags(value objc.IObject /* cross-framework: NSArray */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterModeOptionStructClass) Alloc() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterModeClusterModeOptionStructClass) New() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Init() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Autorelease() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterModeOptionStruct creates a new MTRWaterHeaterModeClusterModeOptionStruct instance.
func NewMTRWaterHeaterModeClusterModeOptionStruct() MTRWaterHeaterModeClusterModeOptionStruct {
	return getMTRWaterHeaterModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct
type MTRWaterHeaterModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterModeOptionStructFrom constructs a [MTRWaterHeaterModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterModeOptionStruct {
	return MTRWaterHeaterModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/modeTags
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}/* debug [instance_properties/getter]: modeTags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/modeTags
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}/* debug [instance_properties/setter]: modeTags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodeoptionstruct/label
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodeoptionstruct/label
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodeoptionstruct/mode
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwaterheatermodeclustermodeoptionstruct/mode
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterModeClusterModeOptionStruct */



