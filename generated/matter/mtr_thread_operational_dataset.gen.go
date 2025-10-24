// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRThreadOperationalDataset] class.
type IMTRThreadOperationalDataset interface {
	objectivec.IObject
	// properties:
	Channel() unsafe.Pointer
	SetChannel(value unsafe.Pointer)
	ChannelNumber() objc.IObject /* cross-framework: NSNumber */
	SetChannelNumber(value objc.IObject /* cross-framework: NSNumber */)
	ExtendedPANID() objc.IObject /* cross-framework: Data */
	SetExtendedPANID(value objc.IObject /* cross-framework: Data */)
	MasterKey() objc.IObject /* cross-framework: Data */
	SetMasterKey(value objc.IObject /* cross-framework: Data */)
	NetworkName() objc.IObject /* cross-framework: NSString */
	SetNetworkName(value objc.IObject /* cross-framework: NSString */)
	PanID() objc.IObject /* cross-framework: Data */
	SetPanID(value objc.IObject /* cross-framework: Data */)
	PsKc() objc.IObject /* cross-framework: Data */
	SetPsKc(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadOperationalDataset
type MTRThreadOperationalDataset struct {
	objectivec.Object
}

// MTRThreadOperationalDatasetFrom constructs a [MTRThreadOperationalDataset] from an unsafe.Pointer.
func MTRThreadOperationalDatasetFrom(ptr unsafe.Pointer) MTRThreadOperationalDataset {
	return MTRThreadOperationalDataset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadOperationalDatasetClass) Alloc() MTRThreadOperationalDataset {
	rv := objc.Send[MTRThreadOperationalDataset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/channel
func (m_ MTRThreadOperationalDataset) Channel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("channel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/channel
func (m_ MTRThreadOperationalDataset) SetChannel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/channelnumber
func (m_ MTRThreadOperationalDataset) ChannelNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channelNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/channelnumber
func (m_ MTRThreadOperationalDataset) SetChannelNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannelNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/extendedpanid
func (m_ MTRThreadOperationalDataset) ExtendedPANID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedPANID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/extendedpanid
func (m_ MTRThreadOperationalDataset) SetExtendedPANID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPANID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/masterkey
func (m_ MTRThreadOperationalDataset) MasterKey() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("masterKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/masterkey
func (m_ MTRThreadOperationalDataset) SetMasterKey(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMasterKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/networkname
func (m_ MTRThreadOperationalDataset) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/networkname
func (m_ MTRThreadOperationalDataset) SetNetworkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/panid
func (m_ MTRThreadOperationalDataset) PanID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("panID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/panid
func (m_ MTRThreadOperationalDataset) SetPanID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPanID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/pskc
func (m_ MTRThreadOperationalDataset) PsKc() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("psKc"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadoperationaldataset/pskc
func (m_ MTRThreadOperationalDataset) SetPsKc(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPsKc:"), value)
}



