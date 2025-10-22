// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CWChannel] class.
var (
	CWChannelClass     _CWChannelClass
	CWChannelClassOnce sync.Once
)

func getCWChannelClass() _CWChannelClass {
	CWChannelClassOnce.Do(func() {
		CWChannelClass = _CWChannelClass{objc.GetClass("CWChannel")}
	})
	return CWChannelClass
}

type _CWChannelClass struct {
	class objc.Class
}

// An interface definition for the [CWChannel] class.
type ICWChannel interface {
	objectivec.IObject
	IsEqualToChannel(channel ICWChannel) bool
	ChannelBand() CWChannelBand
	ChannelNumber() int
	ChannelWidth() CWChannelWidth
}

// Encapsulates an IEEE 802.11 channel.


// Encapsulates an IEEE 802.11 channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel

type CWChannel struct {
	objectivec.Object
}

// CWChannelFrom constructs a [CWChannel] from an unsafe.Pointer.
//
// Encapsulates an IEEE 802.11 channel.
func CWChannelFrom(ptr unsafe.Pointer) CWChannel {
	return CWChannel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CWChannelClass) Alloc() CWChannel {
	rv := objc.Send[CWChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWChannelClass) New() CWChannel {
	rv := objc.Send[CWChannel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWChannel) Init() CWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWChannel) Autorelease() CWChannel {
	rv := objc.Send[CWChannel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWChannel creates a new CWChannel instance.
func NewCWChannel() CWChannel {
	return getCWChannelClass().New()
}




// Determine CWChannel object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/isEqual(to:)

func (c_ CWChannel) IsEqualToChannel(channel ICWChannel) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToChannel:"), channel)
	return rv
}


// The channel band.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelBand

func (c_ CWChannel) ChannelBand() CWChannelBand {
	rv := objc.Send[CWChannelBand](c_.ID, objc.Sel("channelBand"))
	return rv
}


// The channel number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelNumber

func (c_ CWChannel) ChannelNumber() int {
	rv := objc.Send[int](c_.ID, objc.Sel("channelNumber"))
	return rv
}


// The channel width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelWidth

func (c_ CWChannel) ChannelWidth() CWChannelWidth {
	rv := objc.Send[CWChannelWidth](c_.ID, objc.Sel("channelWidth"))
	return rv
}



