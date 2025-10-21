// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] class.
var (
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass     _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClassOnce sync.Once
)

func getMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass() _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass {
	MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClassOnce.Do(func() {
		MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass = _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass{objc.GetClass("MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams")}
	})
	return MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass
}

type _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] class.
type IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams
type MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams struct {
	objectivec.Object
}

// MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsFrom constructs a [MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams] from an unsafe.Pointer.
func MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	return MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass) Alloc() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass) New() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) Init() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) Autorelease() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams creates a new MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams instance.
func NewMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams() MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams {
	return getMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/serverSideProcessingTimeout
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams/timedInvokeTimeoutMs
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



