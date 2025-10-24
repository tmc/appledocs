// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CWChannel */


/* debug [class_header]: Header for CWChannel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWChannel */
// An interface definition for the [CWChannel] class.
type ICWChannel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWChannel */
	// properties:
	ChannelBand() CWChannelBand
	ChannelNumber() int
	ChannelWidth() CWChannelWidth
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWChannel */
	// methods:
	IsEqualToChannel(channel ICWChannel) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWChannel */
// Alloc allocates a new instance without initialization.
func (cc _CWChannelClass) Alloc() CWChannel {
	rv := objc.Send[CWChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWChannel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWChannel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWChannel */

// Determine CWChannel object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/isEqual(to:)
func (c_ CWChannel) IsEqualToChannel(channel ICWChannel) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToChannel:"), channel)
	return rv
}/* debug [instance_methods/method]: IsEqualToChannel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWChannel */

// The channel band.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelBand
func (c_ CWChannel) ChannelBand() CWChannelBand {
	rv := objc.Send[CWChannelBand](c_.ID, objc.Sel("channelBand"))
	return rv
}/* debug [instance_properties/getter]: channelBand */


// The channel number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelNumber
func (c_ CWChannel) ChannelNumber() int {
	rv := objc.Send[int](c_.ID, objc.Sel("channelNumber"))
	return rv
}/* debug [instance_properties/getter]: channelNumber */


// The channel width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelWidth
func (c_ CWChannel) ChannelWidth() CWChannelWidth {
	rv := objc.Send[CWChannelWidth](c_.ID, objc.Sel("channelWidth"))
	return rv
}/* debug [instance_properties/getter]: channelWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWChannel */



