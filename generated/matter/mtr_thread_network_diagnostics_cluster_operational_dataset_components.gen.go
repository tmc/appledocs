// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] class.
var (
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass     _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass() _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass {
	MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass = _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents")}
	})
	return MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass
}

type _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] class.
type IMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents
type MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom constructs a [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	return MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass) Alloc() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass) New() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) Init() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) Autorelease() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents creates a new MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents instance.
func NewMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	return getMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/delaypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) DelayPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("delayPresent"))
	return rv
}


// SetDelayPresent sets the value of the delayPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/delaypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetDelayPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pendingtimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PendingTimestampPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("pendingTimestampPresent"))
	return rv
}


// SetPendingTimestampPresent sets the value of the pendingTimestampPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pendingtimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPendingTimestampPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingTimestampPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/panidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PanIdPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("panIdPresent"))
	return rv
}


// SetPanIdPresent sets the value of the panIdPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/panidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPanIdPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanIdPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/masterkeypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MasterKeyPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("masterKeyPresent"))
	return rv
}


// SetMasterKeyPresent sets the value of the masterKeyPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/masterkeypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMasterKeyPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMasterKeyPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channelPresent"))
	return rv
}


// SetChannelPresent sets the value of the channelPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/extendedpanidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ExtendedPanIdPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extendedPanIdPresent"))
	return rv
}


// SetExtendedPanIdPresent sets the value of the extendedPanIdPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/extendedpanidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetExtendedPanIdPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanIdPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelmaskpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelMaskPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("channelMaskPresent"))
	return rv
}


// SetChannelMaskPresent sets the value of the channelMaskPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelmaskpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelMaskPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelMaskPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pskcpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PskcPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("pskcPresent"))
	return rv
}


// SetPskcPresent sets the value of the pskcPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pskcpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPskcPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPskcPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/securitypolicypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SecurityPolicyPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("securityPolicyPresent"))
	return rv
}


// SetSecurityPolicyPresent sets the value of the securityPolicyPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/securitypolicypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetSecurityPolicyPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurityPolicyPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/meshlocalprefixpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MeshLocalPrefixPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("meshLocalPrefixPresent"))
	return rv
}


// SetMeshLocalPrefixPresent sets the value of the meshLocalPrefixPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/meshlocalprefixpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMeshLocalPrefixPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshLocalPrefixPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/activetimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ActiveTimestampPresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("activeTimestampPresent"))
	return rv
}


// SetActiveTimestampPresent sets the value of the activeTimestampPresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/activetimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetActiveTimestampPresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestampPresent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/networknamepresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) NetworkNamePresent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("networkNamePresent"))
	return rv
}


// SetNetworkNamePresent sets the value of the networkNamePresent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/networknamepresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetNetworkNamePresent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkNamePresent:"), value)
}



