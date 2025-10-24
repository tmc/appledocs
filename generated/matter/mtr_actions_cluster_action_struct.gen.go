// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRActionsClusterActionStruct] class.
type IMTRActionsClusterActionStruct interface {
	objectivec.IObject
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionStruct
type MTRActionsClusterActionStruct struct {
	objectivec.Object
}

// MTRActionsClusterActionStructFrom constructs a [MTRActionsClusterActionStruct] from an unsafe.Pointer.
func MTRActionsClusterActionStructFrom(ptr unsafe.Pointer) MTRActionsClusterActionStruct {
	return MTRActionsClusterActionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterActionStructClass) Alloc() MTRActionsClusterActionStruct {
	rv := objc.Send[MTRActionsClusterActionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/actionid
func (m_ MTRActionsClusterActionStruct) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/actionid
func (m_ MTRActionsClusterActionStruct) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/endpointlistid
func (m_ MTRActionsClusterActionStruct) EndpointListID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpointListID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/endpointlistid
func (m_ MTRActionsClusterActionStruct) SetEndpointListID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpointListID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/name
func (m_ MTRActionsClusterActionStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/name
func (m_ MTRActionsClusterActionStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/state
func (m_ MTRActionsClusterActionStruct) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/state
func (m_ MTRActionsClusterActionStruct) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/supportedcommands
func (m_ MTRActionsClusterActionStruct) SupportedCommands() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("supportedCommands"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/supportedcommands
func (m_ MTRActionsClusterActionStruct) SetSupportedCommands(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportedCommands:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/type
func (m_ MTRActionsClusterActionStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/type
func (m_ MTRActionsClusterActionStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



