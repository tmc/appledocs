// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterMessageResponseOptionStruct */


/* debug [class_header]: Header for MTRMessagesClusterMessageResponseOptionStruct */
// The class instance for the [MTRMessagesClusterMessageResponseOptionStruct] class.
var (
	MTRMessagesClusterMessageResponseOptionStructClass     _MTRMessagesClusterMessageResponseOptionStructClass
	MTRMessagesClusterMessageResponseOptionStructClassOnce sync.Once
)

func getMTRMessagesClusterMessageResponseOptionStructClass() _MTRMessagesClusterMessageResponseOptionStructClass {
	MTRMessagesClusterMessageResponseOptionStructClassOnce.Do(func() {
		MTRMessagesClusterMessageResponseOptionStructClass = _MTRMessagesClusterMessageResponseOptionStructClass{objc.GetClass("MTRMessagesClusterMessageResponseOptionStruct")}
	})
	return MTRMessagesClusterMessageResponseOptionStructClass
}

type _MTRMessagesClusterMessageResponseOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterMessageResponseOptionStruct */
// An interface definition for the [MTRMessagesClusterMessageResponseOptionStruct] class.
type IMTRMessagesClusterMessageResponseOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterMessageResponseOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MessageResponseID() objc.IObject /* cross-framework: NSNumber */
	SetMessageResponseID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterMessageResponseOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterMessageResponseOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageResponseOptionStructClass) Alloc() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterMessageResponseOptionStructClass) New() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Init() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Autorelease() MTRMessagesClusterMessageResponseOptionStruct {
	rv := objc.Send[MTRMessagesClusterMessageResponseOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageResponseOptionStruct creates a new MTRMessagesClusterMessageResponseOptionStruct instance.
func NewMTRMessagesClusterMessageResponseOptionStruct() MTRMessagesClusterMessageResponseOptionStruct {
	return getMTRMessagesClusterMessageResponseOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterMessageResponseOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct
type MTRMessagesClusterMessageResponseOptionStruct struct {
	objectivec.Object
}

// MTRMessagesClusterMessageResponseOptionStructFrom constructs a [MTRMessagesClusterMessageResponseOptionStruct] from an unsafe.Pointer.
func MTRMessagesClusterMessageResponseOptionStructFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageResponseOptionStruct {
	return MTRMessagesClusterMessageResponseOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterMessageResponseOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterMessageResponseOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterMessageResponseOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterMessageResponseOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterMessageResponseOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/label
func (m_ MTRMessagesClusterMessageResponseOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageResponseOptionStruct/label
func (m_ MTRMessagesClusterMessageResponseOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessageresponseoptionstruct/messageresponseid
func (m_ MTRMessagesClusterMessageResponseOptionStruct) MessageResponseID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageResponseID"))
	return rv
}/* debug [instance_properties/getter]: messageResponseID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessageresponseoptionstruct/messageresponseid
func (m_ MTRMessagesClusterMessageResponseOptionStruct) SetMessageResponseID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageResponseID:"), value)
}/* debug [instance_properties/setter]: messageResponseID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterMessageResponseOptionStruct */



