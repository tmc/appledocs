// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] class.
type IMTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsClass) Alloc() MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents
type MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom constructs a [MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponentsFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents {
	return MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/activeTimestampPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ActiveTimestampPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeTimestampPresent"))
	return rv
}/* debug [instance_properties/getter]: activeTimestampPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/activeTimestampPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetActiveTimestampPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestampPresent:"), value)
}/* debug [instance_properties/setter]: activeTimestampPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/channelMaskPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelMaskPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelMaskPresent"))
	return rv
}/* debug [instance_properties/getter]: channelMaskPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/channelMaskPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelMaskPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelMaskPresent:"), value)
}/* debug [instance_properties/setter]: channelMaskPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/channelPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ChannelPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelPresent"))
	return rv
}/* debug [instance_properties/getter]: channelPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/channelPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetChannelPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelPresent:"), value)
}/* debug [instance_properties/setter]: channelPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/delayPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) DelayPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayPresent"))
	return rv
}/* debug [instance_properties/getter]: delayPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/delayPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetDelayPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayPresent:"), value)
}/* debug [instance_properties/setter]: delayPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/extendedPanIdPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) ExtendedPanIdPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extendedPanIdPresent"))
	return rv
}/* debug [instance_properties/getter]: extendedPanIdPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/extendedPanIdPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetExtendedPanIdPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanIdPresent:"), value)
}/* debug [instance_properties/setter]: extendedPanIdPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/masterKeyPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MasterKeyPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("masterKeyPresent"))
	return rv
}/* debug [instance_properties/getter]: masterKeyPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/masterKeyPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMasterKeyPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMasterKeyPresent:"), value)
}/* debug [instance_properties/setter]: masterKeyPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/meshLocalPrefixPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) MeshLocalPrefixPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("meshLocalPrefixPresent"))
	return rv
}/* debug [instance_properties/getter]: meshLocalPrefixPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/meshLocalPrefixPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetMeshLocalPrefixPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshLocalPrefixPresent:"), value)
}/* debug [instance_properties/setter]: meshLocalPrefixPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/networkNamePresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) NetworkNamePresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("networkNamePresent"))
	return rv
}/* debug [instance_properties/getter]: networkNamePresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/networkNamePresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetNetworkNamePresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkNamePresent:"), value)
}/* debug [instance_properties/setter]: networkNamePresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/panIdPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PanIdPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("panIdPresent"))
	return rv
}/* debug [instance_properties/getter]: panIdPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/panIdPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPanIdPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanIdPresent:"), value)
}/* debug [instance_properties/setter]: panIdPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/pendingTimestampPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PendingTimestampPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pendingTimestampPresent"))
	return rv
}/* debug [instance_properties/getter]: pendingTimestampPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/pendingTimestampPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPendingTimestampPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingTimestampPresent:"), value)
}/* debug [instance_properties/setter]: pendingTimestampPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/pskcPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) PskcPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pskcPresent"))
	return rv
}/* debug [instance_properties/getter]: pskcPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/pskcPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetPskcPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPskcPresent:"), value)
}/* debug [instance_properties/setter]: pskcPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/securityPolicyPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SecurityPolicyPresent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("securityPolicyPresent"))
	return rv
}/* debug [instance_properties/getter]: securityPolicyPresent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents/securityPolicyPresent
func (m_ MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents) SetSecurityPolicyPresent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecurityPolicyPresent:"), value)
}/* debug [instance_properties/setter]: securityPolicyPresent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterOperationalDatasetComponents */



