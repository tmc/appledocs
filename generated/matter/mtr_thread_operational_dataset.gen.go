// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadOperationalDataset */


/* debug [class_header]: Header for MTRThreadOperationalDataset */
// The class instance for the [MTRThreadOperationalDataset] class.
var (
	MTRThreadOperationalDatasetClass     _MTRThreadOperationalDatasetClass
	MTRThreadOperationalDatasetClassOnce sync.Once
)

func getMTRThreadOperationalDatasetClass() _MTRThreadOperationalDatasetClass {
	MTRThreadOperationalDatasetClassOnce.Do(func() {
		MTRThreadOperationalDatasetClass = _MTRThreadOperationalDatasetClass{objc.GetClass("MTRThreadOperationalDataset")}
	})
	return MTRThreadOperationalDatasetClass
}

type _MTRThreadOperationalDatasetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadOperationalDataset */
// An interface definition for the [MTRThreadOperationalDataset] class.
type IMTRThreadOperationalDataset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadOperationalDataset */
	// properties:
	Channel() uint16 /* not a class type */
	SetChannel(value uint16 /* not a class type */)
	ChannelNumber() objc.IObject /* cross-framework: NSNumber */
	ExtendedPANID() objc.IObject /* cross-framework: NSData */
	MasterKey() objc.IObject /* cross-framework: NSData */
	NetworkName() objc.IObject /* cross-framework: NSString */
	PanID() objc.IObject /* cross-framework: NSData */
	PSKc() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadOperationalDataset */
	// methods:
	Data() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadOperationalDataset */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadOperationalDatasetClass) Alloc() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadOperationalDatasetClass) New() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadOperationalDataset) Init() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadOperationalDataset) Autorelease() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadOperationalDataset creates a new MTRThreadOperationalDataset instance.
func NewMTRThreadOperationalDataset() MTRThreadOperationalDataset {
	return getMTRThreadOperationalDatasetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadOperationalDataset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset
type MTRThreadOperationalDataset struct {
	objectivec.Object
}

// MTRThreadOperationalDatasetFrom constructs a [MTRThreadOperationalDataset] from an unsafe.Pointer.
func MTRThreadOperationalDatasetFrom(ptr unsafe.Pointer) MTRThreadOperationalDataset {
	return MTRThreadOperationalDataset{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadOperationalDataset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/init(data:)
func NewMTRThreadOperationalDatasetWithData(data objc.IObject /* cross-framework: NSData */) MTRThreadOperationalDataset {
	instance := getMTRThreadOperationalDatasetClass().Alloc()
	rv := objc.Send[MTRThreadOperationalDataset](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRThreadOperationalDatasetWithData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/init(networkName:extendedPANID:masterKey:psKc:channelNumber:panID:)
func NewMTRThreadOperationalDatasetWithNetworkNameExtendedPANIDMasterKeyPSKcChannelNumberPanID(networkName objc.IObject /* cross-framework: NSString */, extendedPANID objc.IObject /* cross-framework: NSData */, masterKey objc.IObject /* cross-framework: NSData */, PSKc objc.IObject /* cross-framework: NSData */, channelNumber objc.IObject /* cross-framework: NSNumber */, panID objc.IObject /* cross-framework: NSData */) MTRThreadOperationalDataset {
	instance := getMTRThreadOperationalDatasetClass().Alloc()
	rv := objc.Send[MTRThreadOperationalDataset](instance.ID, objc.Sel("initWithNetworkName:extendedPANID:masterKey:PSKc:channelNumber:panID:"), networkName, extendedPANID, masterKey, PSKc, channelNumber, panID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRThreadOperationalDatasetWithNetworkNameExtendedPANIDMasterKeyPSKcChannelNumberPanID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/init(networkName:extendedPANID:masterKey:PSKc:channel:panID:)
func NewMTRThreadOperationalDatasetWithNetworkNameExtendedPANIDMasterKeyPSKcChannelPanID(networkName objc.IObject /* cross-framework: NSString */, extendedPANID objc.IObject /* cross-framework: NSData */, masterKey objc.IObject /* cross-framework: NSData */, PSKc objc.IObject /* cross-framework: NSData */, channel uint16 /* not a class type */, panID objc.IObject /* cross-framework: NSData */) MTRThreadOperationalDataset {
	instance := getMTRThreadOperationalDatasetClass().Alloc()
	rv := objc.Send[MTRThreadOperationalDataset](instance.ID, objc.Sel("initWithNetworkName:extendedPANID:masterKey:PSKc:channel:panID:"), networkName, extendedPANID, masterKey, PSKc, channel, panID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRThreadOperationalDatasetWithNetworkNameExtendedPANIDMasterKeyPSKcChannelPanID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadOperationalDataset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadOperationalDataset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadOperationalDataset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/data()
func (m_ MTRThreadOperationalDataset) Data() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_methods/method]: Data */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadOperationalDataset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/channel
func (m_ MTRThreadOperationalDataset) Channel() uint16 /* not a class type */ {
	rv := objc.Send[uint16](m_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/channel
func (m_ MTRThreadOperationalDataset) SetChannel(value uint16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/channelNumber
func (m_ MTRThreadOperationalDataset) ChannelNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelNumber"))
	return rv
}/* debug [instance_properties/getter]: channelNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/extendedPANID
func (m_ MTRThreadOperationalDataset) ExtendedPANID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedPANID"))
	return rv
}/* debug [instance_properties/getter]: extendedPANID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/masterKey
func (m_ MTRThreadOperationalDataset) MasterKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("masterKey"))
	return rv
}/* debug [instance_properties/getter]: masterKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/networkName
func (m_ MTRThreadOperationalDataset) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}/* debug [instance_properties/getter]: networkName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/panID
func (m_ MTRThreadOperationalDataset) PanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("panID"))
	return rv
}/* debug [instance_properties/getter]: panID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset/psKc
func (m_ MTRThreadOperationalDataset) PSKc() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("PSKc"))
	return rv
}/* debug [instance_properties/getter]: PSKc */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadOperationalDataset */


