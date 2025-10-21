// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterUpdateFabricLabelParams] class.
var (
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClass     _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterUpdateFabricLabelParamsClass() _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass {
	MTROperationalCredentialsClusterUpdateFabricLabelParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterUpdateFabricLabelParamsClass = _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass{objc.GetClass("MTROperationalCredentialsClusterUpdateFabricLabelParams")}
	})
	return MTROperationalCredentialsClusterUpdateFabricLabelParamsClass
}

type _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterUpdateFabricLabelParams] class.
type IMTROperationalCredentialsClusterUpdateFabricLabelParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateFabricLabelParams
type MTROperationalCredentialsClusterUpdateFabricLabelParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterUpdateFabricLabelParamsFrom constructs a [MTROperationalCredentialsClusterUpdateFabricLabelParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterUpdateFabricLabelParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterUpdateFabricLabelParams {
	return MTROperationalCredentialsClusterUpdateFabricLabelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass) Alloc() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterUpdateFabricLabelParamsClass) New() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Init() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Autorelease() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateFabricLabelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterUpdateFabricLabelParams creates a new MTROperationalCredentialsClusterUpdateFabricLabelParams instance.
func NewMTROperationalCredentialsClusterUpdateFabricLabelParams() MTROperationalCredentialsClusterUpdateFabricLabelParams {
	return getMTROperationalCredentialsClusterUpdateFabricLabelParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/label
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/label
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/serversideprocessingtimeout
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterupdatefabriclabelparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterUpdateFabricLabelParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



