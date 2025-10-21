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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/actionid
func (m_ MTRActionsClusterActionStruct) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/actionid
func (m_ MTRActionsClusterActionStruct) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/endpointlistid
func (m_ MTRActionsClusterActionStruct) EndpointListID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpointListID"))
	return rv
}


// SetEndpointListID sets the value of the endpointListID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/endpointlistid
func (m_ MTRActionsClusterActionStruct) SetEndpointListID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpointListID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/name
func (m_ MTRActionsClusterActionStruct) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/name
func (m_ MTRActionsClusterActionStruct) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/state
func (m_ MTRActionsClusterActionStruct) State() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/state
func (m_ MTRActionsClusterActionStruct) SetState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/supportedcommands
func (m_ MTRActionsClusterActionStruct) SupportedCommands() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("supportedCommands"))
	return rv
}


// SetSupportedCommands sets the value of the supportedCommands property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/supportedcommands
func (m_ MTRActionsClusterActionStruct) SetSupportedCommands(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportedCommands:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/type
func (m_ MTRActionsClusterActionStruct) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionstruct/type
func (m_ MTRActionsClusterActionStruct) SetType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



