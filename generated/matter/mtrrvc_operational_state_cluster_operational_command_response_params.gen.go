// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterOperationalCommandResponseParams] class.
var (
	MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass     _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass
	MTRRVCOperationalStateClusterOperationalCommandResponseParamsClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterOperationalCommandResponseParamsClass() _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass {
	MTRRVCOperationalStateClusterOperationalCommandResponseParamsClassOnce.Do(func() {
		MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass = _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass{objc.GetClass("MTRRVCOperationalStateClusterOperationalCommandResponseParams")}
	})
	return MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass
}

type _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterOperationalCommandResponseParams] class.
type IMTRRVCOperationalStateClusterOperationalCommandResponseParams interface {
	objectivec.IObject
	CommandResponseState() MTRRVCOperationalStateClusterErrorStateStruct
	SetCommandResponseState(value IMTRRVCOperationalStateClusterErrorStateStruct)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterOperationalCommandResponseParams
type MTRRVCOperationalStateClusterOperationalCommandResponseParams struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterOperationalCommandResponseParamsFrom constructs a [MTRRVCOperationalStateClusterOperationalCommandResponseParams] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterOperationalCommandResponseParamsFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	return MTRRVCOperationalStateClusterOperationalCommandResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass) Alloc() MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterOperationalCommandResponseParamsClass) New() MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterOperationalCommandResponseParams) Init() MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterOperationalCommandResponseParams) Autorelease() MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterOperationalCommandResponseParams creates a new MTRRVCOperationalStateClusterOperationalCommandResponseParams instance.
func NewMTRRVCOperationalStateClusterOperationalCommandResponseParams() MTRRVCOperationalStateClusterOperationalCommandResponseParams {
	return getMTRRVCOperationalStateClusterOperationalCommandResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalcommandresponseparams/commandresponsestate
func (m_ MTRRVCOperationalStateClusterOperationalCommandResponseParams) CommandResponseState() MTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("commandResponseState"))
	return rv
}


// SetCommandResponseState sets the value of the commandResponseState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalcommandresponseparams/commandresponsestate
func (m_ MTRRVCOperationalStateClusterOperationalCommandResponseParams) SetCommandResponseState(value IMTRRVCOperationalStateClusterErrorStateStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandResponseState:"), value)
}



