// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCCleanModeClusterChangeToModeResponseParams] class.
var (
	MTRRVCCleanModeClusterChangeToModeResponseParamsClass     _MTRRVCCleanModeClusterChangeToModeResponseParamsClass
	MTRRVCCleanModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRRVCCleanModeClusterChangeToModeResponseParamsClass() _MTRRVCCleanModeClusterChangeToModeResponseParamsClass {
	MTRRVCCleanModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRRVCCleanModeClusterChangeToModeResponseParamsClass = _MTRRVCCleanModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRRVCCleanModeClusterChangeToModeResponseParams")}
	})
	return MTRRVCCleanModeClusterChangeToModeResponseParamsClass
}

type _MTRRVCCleanModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCCleanModeClusterChangeToModeResponseParams] class.
type IMTRRVCCleanModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	Status() foundation.Number
	SetStatus(value foundation.INumber)
	StatusText() string
	SetStatusText(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCCleanModeClusterChangeToModeResponseParams
type MTRRVCCleanModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRRVCCleanModeClusterChangeToModeResponseParamsFrom constructs a [MTRRVCCleanModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRRVCCleanModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRRVCCleanModeClusterChangeToModeResponseParams {
	return MTRRVCCleanModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCCleanModeClusterChangeToModeResponseParamsClass) Alloc() MTRRVCCleanModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCCleanModeClusterChangeToModeResponseParamsClass) New() MTRRVCCleanModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) Init() MTRRVCCleanModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) Autorelease() MTRRVCCleanModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCCleanModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCCleanModeClusterChangeToModeResponseParams creates a new MTRRVCCleanModeClusterChangeToModeResponseParams instance.
func NewMTRRVCCleanModeClusterChangeToModeResponseParams() MTRRVCCleanModeClusterChangeToModeResponseParams {
	return getMTRRVCCleanModeClusterChangeToModeResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomoderesponseparams/status
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomoderesponseparams/status
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) StatusText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("statusText"))
	return rv
}


// SetStatusText sets the value of the statusText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclusterchangetomoderesponseparams/statustext
func (m_ MTRRVCCleanModeClusterChangeToModeResponseParams) SetStatusText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), objc.String(value))
}



