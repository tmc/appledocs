// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterActionStruct */


/* debug [class_header]: Header for MTRActionsClusterActionStruct */
// The class instance for the [MTRActionsClusterActionStruct] class.
var (
	MTRActionsClusterActionStructClass     _MTRActionsClusterActionStructClass
	MTRActionsClusterActionStructClassOnce sync.Once
)

func getMTRActionsClusterActionStructClass() _MTRActionsClusterActionStructClass {
	MTRActionsClusterActionStructClassOnce.Do(func() {
		MTRActionsClusterActionStructClass = _MTRActionsClusterActionStructClass{objc.GetClass("MTRActionsClusterActionStruct")}
	})
	return MTRActionsClusterActionStructClass
}

type _MTRActionsClusterActionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterActionStruct */
// An interface definition for the [MTRActionsClusterActionStruct] class.
type IMTRActionsClusterActionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterActionStruct */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	EndpointListID() objc.IObject /* cross-framework: NSNumber */
	SetEndpointListID(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
	SupportedCommands() objc.IObject /* cross-framework: NSNumber */
	SetSupportedCommands(value objc.IObject /* cross-framework: NSNumber */)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterActionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterActionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterActionStructClass) Alloc() MTRActionsClusterActionStruct {
	rv := objc.Send[MTRActionsClusterActionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterActionStructClass) New() MTRActionsClusterActionStruct {
	rv := objc.Send[MTRActionsClusterActionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterActionStruct) Init() MTRActionsClusterActionStruct {
	rv := objc.Send[MTRActionsClusterActionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterActionStruct) Autorelease() MTRActionsClusterActionStruct {
	rv := objc.Send[MTRActionsClusterActionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterActionStruct creates a new MTRActionsClusterActionStruct instance.
func NewMTRActionsClusterActionStruct() MTRActionsClusterActionStruct {
	return getMTRActionsClusterActionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterActionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct
type MTRActionsClusterActionStruct struct {
	objectivec.Object
}

// MTRActionsClusterActionStructFrom constructs a [MTRActionsClusterActionStruct] from an unsafe.Pointer.
func MTRActionsClusterActionStructFrom(ptr unsafe.Pointer) MTRActionsClusterActionStruct {
	return MTRActionsClusterActionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterActionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterActionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterActionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterActionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterActionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/actionID
func (m_ MTRActionsClusterActionStruct) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/actionID
func (m_ MTRActionsClusterActionStruct) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/endpointListID
func (m_ MTRActionsClusterActionStruct) EndpointListID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpointListID"))
	return rv
}/* debug [instance_properties/getter]: endpointListID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/endpointListID
func (m_ MTRActionsClusterActionStruct) SetEndpointListID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpointListID:"), value)
}/* debug [instance_properties/setter]: endpointListID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/name
func (m_ MTRActionsClusterActionStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/name
func (m_ MTRActionsClusterActionStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/state
func (m_ MTRActionsClusterActionStruct) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/state
func (m_ MTRActionsClusterActionStruct) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/supportedCommands
func (m_ MTRActionsClusterActionStruct) SupportedCommands() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("supportedCommands"))
	return rv
}/* debug [instance_properties/getter]: supportedCommands */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/supportedCommands
func (m_ MTRActionsClusterActionStruct) SetSupportedCommands(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportedCommands:"), value)
}/* debug [instance_properties/setter]: supportedCommands */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/type
func (m_ MTRActionsClusterActionStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct/type
func (m_ MTRActionsClusterActionStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterActionStruct */



