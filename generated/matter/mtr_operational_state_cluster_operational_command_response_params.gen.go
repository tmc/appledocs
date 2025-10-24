// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalStateClusterOperationalCommandResponseParams] class.
var (
	MTROperationalStateClusterOperationalCommandResponseParamsClass     _MTROperationalStateClusterOperationalCommandResponseParamsClass
	MTROperationalStateClusterOperationalCommandResponseParamsClassOnce sync.Once
)

func getMTROperationalStateClusterOperationalCommandResponseParamsClass() _MTROperationalStateClusterOperationalCommandResponseParamsClass {
	MTROperationalStateClusterOperationalCommandResponseParamsClassOnce.Do(func() {
		MTROperationalStateClusterOperationalCommandResponseParamsClass = _MTROperationalStateClusterOperationalCommandResponseParamsClass{objc.GetClass("MTROperationalStateClusterOperationalCommandResponseParams")}
	})
	return MTROperationalStateClusterOperationalCommandResponseParamsClass
}

type _MTROperationalStateClusterOperationalCommandResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterOperationalCommandResponseParams] class.
type IMTROperationalStateClusterOperationalCommandResponseParams interface {
	objectivec.IObject
	// properties:
	CommandResponseState() IMTROperationalStateClusterErrorStateStruct
	SetCommandResponseState(value IMTROperationalStateClusterErrorStateStruct)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterOperationalCommandResponseParams
type MTROperationalStateClusterOperationalCommandResponseParams struct {
	objectivec.Object
}

// MTROperationalStateClusterOperationalCommandResponseParamsFrom constructs a [MTROperationalStateClusterOperationalCommandResponseParams] from an unsafe.Pointer.
func MTROperationalStateClusterOperationalCommandResponseParamsFrom(ptr unsafe.Pointer) MTROperationalStateClusterOperationalCommandResponseParams {
	return MTROperationalStateClusterOperationalCommandResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterOperationalCommandResponseParamsClass) Alloc() MTROperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterOperationalCommandResponseParamsClass) New() MTROperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterOperationalCommandResponseParams) Init() MTROperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterOperationalCommandResponseParams) Autorelease() MTROperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterOperationalCommandResponseParams creates a new MTROperationalStateClusterOperationalCommandResponseParams instance.
func NewMTROperationalStateClusterOperationalCommandResponseParams() MTROperationalStateClusterOperationalCommandResponseParams {
	return getMTROperationalStateClusterOperationalCommandResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationalcommandresponseparams/commandresponsestate
func (m_ MTROperationalStateClusterOperationalCommandResponseParams) CommandResponseState() IMTROperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("commandResponseState"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationalcommandresponseparams/commandresponsestate
func (m_ MTROperationalStateClusterOperationalCommandResponseParams) SetCommandResponseState(value IMTROperationalStateClusterErrorStateStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandResponseState:"), value)
}
