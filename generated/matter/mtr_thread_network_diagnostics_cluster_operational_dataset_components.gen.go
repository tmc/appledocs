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
	// properties:
	ActiveTimestampPresent() objc.IObject /* cross-framework: NSNumber */
	SetActiveTimestampPresent(value objc.IObject /* cross-framework: NSNumber */)
	ChannelMaskPresent() objc.IObject /* cross-framework: NSNumber */
	SetChannelMaskPresent(value objc.IObject /* cross-framework: NSNumber */)
	ChannelPresent() objc.IObject /* cross-framework: NSNumber */
	SetChannelPresent(value objc.IObject /* cross-framework: NSNumber */)
	DelayPresent() objc.IObject /* cross-framework: NSNumber */
	SetDelayPresent(value objc.IObject /* cross-framework: NSNumber */)
	ExtendedPanIdPresent() objc.IObject /* cross-framework: NSNumber */
	SetExtendedPanIdPresent(value objc.IObject /* cross-framework: NSNumber */)
	MasterKeyPresent() objc.IObject /* cross-framework: NSNumber */
	SetMasterKeyPresent(value objc.IObject /* cross-framework: NSNumber */)
	MeshLocalPrefixPresent() objc.IObject /* cross-framework: NSNumber */
	SetMeshLocalPrefixPresent(value objc.IObject /* cross-framework: NSNumber */)
	NetworkNamePresent() objc.IObject /* cross-framework: NSNumber */
	SetNetworkNamePresent(value objc.IObject /* cross-framework: NSNumber */)
	PanIdPresent() objc.IObject /* cross-framework: NSNumber */
	SetPanIdPresent(value objc.IObject /* cross-framework: NSNumber */)
	PendingTimestampPresent() objc.IObject /* cross-framework: NSNumber */
	SetPendingTimestampPresent(value objc.IObject /* cross-framework: NSNumber */)
	PskcPresent() objc.IObject /* cross-framework: NSNumber */
	SetPskcPresent(value objc.IObject /* cross-framework: NSNumber */)
	SecurityPolicyPresent() objc.IObject /* cross-framework: NSNumber */
	SetSecurityPolicyPresent(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/activetimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ActiveTimestampPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeTimestampPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/activetimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetActiveTimestampPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestampPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelmaskpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelMaskPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelMaskPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelmaskpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelMaskPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelMaskPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/channelpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/delaypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) DelayPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/delaypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetDelayPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/extendedpanidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ExtendedPanIdPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extendedPanIdPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/extendedpanidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetExtendedPanIdPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanIdPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/masterkeypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MasterKeyPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("masterKeyPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/masterkeypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMasterKeyPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMasterKeyPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/meshlocalprefixpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MeshLocalPrefixPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("meshLocalPrefixPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/meshlocalprefixpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMeshLocalPrefixPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshLocalPrefixPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/networknamepresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) NetworkNamePresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkNamePresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/networknamepresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetNetworkNamePresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkNamePresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/panidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PanIdPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("panIdPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/panidpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPanIdPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanIdPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pendingtimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PendingTimestampPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pendingTimestampPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pendingtimestamppresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPendingTimestampPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingTimestampPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pskcpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PskcPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pskcPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/pskcpresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPskcPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPskcPresent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/securitypolicypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SecurityPolicyPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("securityPolicyPresent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusteroperationaldatasetcomponents/securitypolicypresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetSecurityPolicyPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurityPolicyPresent:"), value)
}



