// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterEndpointListStruct */


/* debug [class_header]: Header for MTRActionsClusterEndpointListStruct */
// The class instance for the [MTRActionsClusterEndpointListStruct] class.
var (
	MTRActionsClusterEndpointListStructClass     _MTRActionsClusterEndpointListStructClass
	MTRActionsClusterEndpointListStructClassOnce sync.Once
)

func getMTRActionsClusterEndpointListStructClass() _MTRActionsClusterEndpointListStructClass {
	MTRActionsClusterEndpointListStructClassOnce.Do(func() {
		MTRActionsClusterEndpointListStructClass = _MTRActionsClusterEndpointListStructClass{objc.GetClass("MTRActionsClusterEndpointListStruct")}
	})
	return MTRActionsClusterEndpointListStructClass
}

type _MTRActionsClusterEndpointListStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterEndpointListStruct */
// An interface definition for the [MTRActionsClusterEndpointListStruct] class.
type IMTRActionsClusterEndpointListStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterEndpointListStruct */
	// properties:
	EndpointListID() objc.IObject /* cross-framework: NSNumber */
	SetEndpointListID(value objc.IObject /* cross-framework: NSNumber */)
	Endpoints() objc.IObject /* cross-framework: NSArray */
	SetEndpoints(value objc.IObject /* cross-framework: NSArray */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterEndpointListStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterEndpointListStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEndpointListStructClass) Alloc() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterEndpointListStructClass) New() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEndpointListStruct) Init() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEndpointListStruct) Autorelease() MTRActionsClusterEndpointListStruct {
	rv := objc.Send[MTRActionsClusterEndpointListStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEndpointListStruct creates a new MTRActionsClusterEndpointListStruct instance.
func NewMTRActionsClusterEndpointListStruct() MTRActionsClusterEndpointListStruct {
	return getMTRActionsClusterEndpointListStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterEndpointListStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct
type MTRActionsClusterEndpointListStruct struct {
	objectivec.Object
}

// MTRActionsClusterEndpointListStructFrom constructs a [MTRActionsClusterEndpointListStruct] from an unsafe.Pointer.
func MTRActionsClusterEndpointListStructFrom(ptr unsafe.Pointer) MTRActionsClusterEndpointListStruct {
	return MTRActionsClusterEndpointListStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterEndpointListStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterEndpointListStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterEndpointListStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterEndpointListStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterEndpointListStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/endpointListID
func (m_ MTRActionsClusterEndpointListStruct) EndpointListID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpointListID"))
	return rv
}/* debug [instance_properties/getter]: endpointListID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/endpointListID
func (m_ MTRActionsClusterEndpointListStruct) SetEndpointListID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpointListID:"), value)
}/* debug [instance_properties/setter]: endpointListID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/endpoints
func (m_ MTRActionsClusterEndpointListStruct) Endpoints() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("endpoints"))
	return rv
}/* debug [instance_properties/getter]: endpoints */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/endpoints
func (m_ MTRActionsClusterEndpointListStruct) SetEndpoints(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoints:"), value)
}/* debug [instance_properties/setter]: endpoints */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/name
func (m_ MTRActionsClusterEndpointListStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/name
func (m_ MTRActionsClusterEndpointListStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/type
func (m_ MTRActionsClusterEndpointListStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEndpointListStruct/type
func (m_ MTRActionsClusterEndpointListStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterEndpointListStruct */



