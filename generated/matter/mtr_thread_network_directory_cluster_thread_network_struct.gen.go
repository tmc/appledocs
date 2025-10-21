// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ActiveTimestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("activeTimestamp"))
	return rv
}


// SetActiveTimestamp sets the value of the activeTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetActiveTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Channel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetChannel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ExtendedPanID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("extendedPanID"))
	return rv
}


// SetExtendedPanID sets the value of the extendedPanID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetExtendedPanID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/networkName
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) NetworkName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("networkName"))
	return rv
}


// SetNetworkName sets the value of the networkName property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/networkName
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetNetworkName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), objc.String(value))
}



