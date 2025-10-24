// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterPageTokenStruct */


/* debug [class_header]: Header for MTRChannelClusterPageTokenStruct */
// The class instance for the [MTRChannelClusterPageTokenStruct] class.
var (
	MTRChannelClusterPageTokenStructClass     _MTRChannelClusterPageTokenStructClass
	MTRChannelClusterPageTokenStructClassOnce sync.Once
)

func getMTRChannelClusterPageTokenStructClass() _MTRChannelClusterPageTokenStructClass {
	MTRChannelClusterPageTokenStructClassOnce.Do(func() {
		MTRChannelClusterPageTokenStructClass = _MTRChannelClusterPageTokenStructClass{objc.GetClass("MTRChannelClusterPageTokenStruct")}
	})
	return MTRChannelClusterPageTokenStructClass
}

type _MTRChannelClusterPageTokenStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterPageTokenStruct */
// An interface definition for the [MTRChannelClusterPageTokenStruct] class.
type IMTRChannelClusterPageTokenStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterPageTokenStruct */
	// properties:
	After() objc.IObject /* cross-framework: NSString */
	SetAfter(value objc.IObject /* cross-framework: NSString */)
	Before() objc.IObject /* cross-framework: NSString */
	SetBefore(value objc.IObject /* cross-framework: NSString */)
	Limit() objc.IObject /* cross-framework: NSNumber */
	SetLimit(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterPageTokenStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterPageTokenStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterPageTokenStructClass) Alloc() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterPageTokenStructClass) New() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterPageTokenStruct) Init() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterPageTokenStruct) Autorelease() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterPageTokenStruct creates a new MTRChannelClusterPageTokenStruct instance.
func NewMTRChannelClusterPageTokenStruct() MTRChannelClusterPageTokenStruct {
	return getMTRChannelClusterPageTokenStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterPageTokenStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct
type MTRChannelClusterPageTokenStruct struct {
	objectivec.Object
}

// MTRChannelClusterPageTokenStructFrom constructs a [MTRChannelClusterPageTokenStruct] from an unsafe.Pointer.
func MTRChannelClusterPageTokenStructFrom(ptr unsafe.Pointer) MTRChannelClusterPageTokenStruct {
	return MTRChannelClusterPageTokenStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterPageTokenStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterPageTokenStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterPageTokenStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterPageTokenStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterPageTokenStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) After() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("after"))
	return rv
}/* debug [instance_properties/getter]: after */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) SetAfter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAfter:"), value)
}/* debug [instance_properties/setter]: after */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterpagetokenstruct/before
func (m_ MTRChannelClusterPageTokenStruct) Before() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("before"))
	return rv
}/* debug [instance_properties/getter]: before */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterpagetokenstruct/before
func (m_ MTRChannelClusterPageTokenStruct) SetBefore(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBefore:"), value)
}/* debug [instance_properties/setter]: before */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterpagetokenstruct/limit
func (m_ MTRChannelClusterPageTokenStruct) Limit() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("limit"))
	return rv
}/* debug [instance_properties/getter]: limit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterpagetokenstruct/limit
func (m_ MTRChannelClusterPageTokenStruct) SetLimit(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLimit:"), value)
}/* debug [instance_properties/setter]: limit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterPageTokenStruct */



