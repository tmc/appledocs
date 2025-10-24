// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterProgramCastStruct */


/* debug [class_header]: Header for MTRChannelClusterProgramCastStruct */
// The class instance for the [MTRChannelClusterProgramCastStruct] class.
var (
	MTRChannelClusterProgramCastStructClass     _MTRChannelClusterProgramCastStructClass
	MTRChannelClusterProgramCastStructClassOnce sync.Once
)

func getMTRChannelClusterProgramCastStructClass() _MTRChannelClusterProgramCastStructClass {
	MTRChannelClusterProgramCastStructClassOnce.Do(func() {
		MTRChannelClusterProgramCastStructClass = _MTRChannelClusterProgramCastStructClass{objc.GetClass("MTRChannelClusterProgramCastStruct")}
	})
	return MTRChannelClusterProgramCastStructClass
}

type _MTRChannelClusterProgramCastStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterProgramCastStruct */
// An interface definition for the [MTRChannelClusterProgramCastStruct] class.
type IMTRChannelClusterProgramCastStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterProgramCastStruct */
	// properties:
	Role() objc.IObject /* cross-framework: NSString */
	SetRole(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterProgramCastStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterProgramCastStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramCastStructClass) Alloc() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterProgramCastStructClass) New() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramCastStruct) Init() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramCastStruct) Autorelease() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramCastStruct creates a new MTRChannelClusterProgramCastStruct instance.
func NewMTRChannelClusterProgramCastStruct() MTRChannelClusterProgramCastStruct {
	return getMTRChannelClusterProgramCastStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterProgramCastStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct
type MTRChannelClusterProgramCastStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramCastStructFrom constructs a [MTRChannelClusterProgramCastStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramCastStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramCastStruct {
	return MTRChannelClusterProgramCastStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterProgramCastStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterProgramCastStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterProgramCastStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterProgramCastStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterProgramCastStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/role
func (m_ MTRChannelClusterProgramCastStruct) Role() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("role"))
	return rv
}/* debug [instance_properties/getter]: role */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/role
func (m_ MTRChannelClusterProgramCastStruct) SetRole(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRole:"), value)
}/* debug [instance_properties/setter]: role */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramcaststruct/name
func (m_ MTRChannelClusterProgramCastStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterprogramcaststruct/name
func (m_ MTRChannelClusterProgramCastStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterProgramCastStruct */



