// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */


/* debug [class_header]: Header for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
// An interface definition for the [MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams] class.
type IMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
	// properties:
	Passphrase() foundation.Data
	SetPassphrase(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass) Alloc() MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams
type MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams struct {
	objectivec.Object
}

// MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsFrom constructs a [MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams] from an unsafe.Pointer.
func MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsFrom(ptr unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	return MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */

// Initialize an MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams/init(responseValue:)
func NewMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams {
	instance := getMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsClass().Alloc()
	rv := objc.Send[MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRWiFiNetworkManagementClusterNetworkPassphraseResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkmanagementclusternetworkpassphraseresponseparams/passphrase
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) Passphrase() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("passphrase"))
	return rv
}/* debug [instance_properties/getter]: passphrase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkmanagementclusternetworkpassphraseresponseparams/passphrase
func (m_ MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams) SetPassphrase(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPassphrase:"), value)
}/* debug [instance_properties/setter]: passphrase */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkManagementClusterNetworkPassphraseResponseParams */


