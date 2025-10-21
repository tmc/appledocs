// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterLoginParams] class.
var (
	MTRAccountLoginClusterLoginParamsClass     _MTRAccountLoginClusterLoginParamsClass
	MTRAccountLoginClusterLoginParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterLoginParamsClass() _MTRAccountLoginClusterLoginParamsClass {
	MTRAccountLoginClusterLoginParamsClassOnce.Do(func() {
		MTRAccountLoginClusterLoginParamsClass = _MTRAccountLoginClusterLoginParamsClass{objc.GetClass("MTRAccountLoginClusterLoginParams")}
	})
	return MTRAccountLoginClusterLoginParamsClass
}

type _MTRAccountLoginClusterLoginParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterLoginParams] class.
type IMTRAccountLoginClusterLoginParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoginParams
type MTRAccountLoginClusterLoginParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterLoginParamsFrom constructs a [MTRAccountLoginClusterLoginParams] from an unsafe.Pointer.
func MTRAccountLoginClusterLoginParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLoginParams {
	return MTRAccountLoginClusterLoginParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLoginParamsClass) Alloc() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterLoginParamsClass) New() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLoginParams) Init() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLoginParams) Autorelease() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLoginParams creates a new MTRAccountLoginClusterLoginParams instance.
func NewMTRAccountLoginClusterLoginParams() MTRAccountLoginClusterLoginParams {
	return getMTRAccountLoginClusterLoginParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/node
func (m_ MTRAccountLoginClusterLoginParams) Node() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("node"))
	return rv
}


// SetNode sets the value of the node property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/node
func (m_ MTRAccountLoginClusterLoginParams) SetNode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLoginParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLoginParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/setuppin
func (m_ MTRAccountLoginClusterLoginParams) SetupPIN() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("setupPIN"))
	return rv
}


// SetSetupPIN sets the value of the setupPIN property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/setuppin
func (m_ MTRAccountLoginClusterLoginParams) SetSetupPIN(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSetupPIN:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterLoginParams) TempAccountIdentifier() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("tempAccountIdentifier"))
	return rv
}


// SetTempAccountIdentifier sets the value of the tempAccountIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/tempaccountidentifier
func (m_ MTRAccountLoginClusterLoginParams) SetTempAccountIdentifier(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTempAccountIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLoginParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterloginparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLoginParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



