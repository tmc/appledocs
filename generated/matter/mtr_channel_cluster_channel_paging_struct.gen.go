// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterChannelPagingStruct */


/* debug [class_header]: Header for MTRChannelClusterChannelPagingStruct */
// The class instance for the [MTRChannelClusterChannelPagingStruct] class.
var (
	MTRChannelClusterChannelPagingStructClass     _MTRChannelClusterChannelPagingStructClass
	MTRChannelClusterChannelPagingStructClassOnce sync.Once
)

func getMTRChannelClusterChannelPagingStructClass() _MTRChannelClusterChannelPagingStructClass {
	MTRChannelClusterChannelPagingStructClassOnce.Do(func() {
		MTRChannelClusterChannelPagingStructClass = _MTRChannelClusterChannelPagingStructClass{objc.GetClass("MTRChannelClusterChannelPagingStruct")}
	})
	return MTRChannelClusterChannelPagingStructClass
}

type _MTRChannelClusterChannelPagingStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterChannelPagingStruct */
// An interface definition for the [MTRChannelClusterChannelPagingStruct] class.
type IMTRChannelClusterChannelPagingStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterChannelPagingStruct */
	// properties:
	NextToken() IMTRChannelClusterPageTokenStruct
	SetNextToken(value IMTRChannelClusterPageTokenStruct)
	PreviousToken() IMTRChannelClusterPageTokenStruct
	SetPreviousToken(value IMTRChannelClusterPageTokenStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterChannelPagingStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterChannelPagingStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelPagingStructClass) Alloc() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterChannelPagingStructClass) New() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelPagingStruct) Init() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelPagingStruct) Autorelease() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelPagingStruct creates a new MTRChannelClusterChannelPagingStruct instance.
func NewMTRChannelClusterChannelPagingStruct() MTRChannelClusterChannelPagingStruct {
	return getMTRChannelClusterChannelPagingStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterChannelPagingStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct
type MTRChannelClusterChannelPagingStruct struct {
	objectivec.Object
}

// MTRChannelClusterChannelPagingStructFrom constructs a [MTRChannelClusterChannelPagingStruct] from an unsafe.Pointer.
func MTRChannelClusterChannelPagingStructFrom(ptr unsafe.Pointer) MTRChannelClusterChannelPagingStruct {
	return MTRChannelClusterChannelPagingStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterChannelPagingStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterChannelPagingStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterChannelPagingStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterChannelPagingStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterChannelPagingStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/nextToken
func (m_ MTRChannelClusterChannelPagingStruct) NextToken() IMTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("nextToken"))
	return rv
}/* debug [instance_properties/getter]: nextToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/nextToken
func (m_ MTRChannelClusterChannelPagingStruct) SetNextToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextToken:"), value)
}/* debug [instance_properties/setter]: nextToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelpagingstruct/previoustoken
func (m_ MTRChannelClusterChannelPagingStruct) PreviousToken() IMTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("previousToken"))
	return rv
}/* debug [instance_properties/getter]: previousToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelpagingstruct/previoustoken
func (m_ MTRChannelClusterChannelPagingStruct) SetPreviousToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousToken:"), value)
}/* debug [instance_properties/setter]: previousToken */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterChannelPagingStruct */



