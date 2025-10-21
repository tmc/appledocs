// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams] class.
var (
	MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass     _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass
	MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClassOnce sync.Once
)

func getMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass() _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass {
	MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClassOnce.Do(func() {
		MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass = _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass{objc.GetClass("MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams")}
	})
	return MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass
}

type _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams] class.
type IMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams
type MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams struct {
	objectivec.Object
}

// MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsFrom constructs a [MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams] from an unsafe.Pointer.
func MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	return MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass) Alloc() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass) New() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) Init() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) Autorelease() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams creates a new MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams instance.
func NewMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	return getMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass().New()
}


// Initialize an MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams/init(responseValue:)
func NewMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	instance := getMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass().Alloc()
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams/passphrase
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) Passphrase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("passphrase"))
	return rv
}


// SetPassphrase sets the value of the passphrase property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams/passphrase
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) SetPassphrase(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPassphrase:"), value)
}

