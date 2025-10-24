// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] class.
var (
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClass     _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterThreadNetworkStructClass() _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass {
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterThreadNetworkStructClass = _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass{objc.GetClass("MTRThreadNetworkDirectoryClusterThreadNetworkStruct")}
	})
	return MTRThreadNetworkDirectoryClusterThreadNetworkStructClass
}

type _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] class.
type IMTRThreadNetworkDirectoryClusterThreadNetworkStruct interface {
	objectivec.IObject
	// properties:
	ActiveTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetActiveTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	Channel() objc.IObject /* cross-framework: NSNumber */
	SetChannel(value objc.IObject /* cross-framework: NSNumber */)
	ExtendedPanID() objc.IObject /* cross-framework: NSData */
	SetExtendedPanID(value objc.IObject /* cross-framework: NSData */)
	NetworkName() objc.IObject /* cross-framework: NSString */
	SetNetworkName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct
type MTRThreadNetworkDirectoryClusterThreadNetworkStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterThreadNetworkStructFrom constructs a [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterThreadNetworkStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	return MTRThreadNetworkDirectoryClusterThreadNetworkStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass) Alloc() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass) New() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Init() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Autorelease() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterThreadNetworkStruct creates a new MTRThreadNetworkDirectoryClusterThreadNetworkStruct instance.
func NewMTRThreadNetworkDirectoryClusterThreadNetworkStruct() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	return getMTRThreadNetworkDirectoryClusterThreadNetworkStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ActiveTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeTimestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetActiveTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestamp:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ExtendedPanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedPanID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetExtendedPanID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/networkName
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/networkName
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetNetworkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), value)
}



