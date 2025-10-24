// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemOutput */


/* debug [class_header]: Header for AVPlayerItemOutput */
// The class instance for the [PlayerItemOutput] class.
var (
	PlayerItemOutputClass     _PlayerItemOutputClass
	PlayerItemOutputClassOnce sync.Once
)

func getPlayerItemOutputClass() _PlayerItemOutputClass {
	PlayerItemOutputClassOnce.Do(func() {
		PlayerItemOutputClass = _PlayerItemOutputClass{objc.GetClass("AVPlayerItemOutput")}
	})
	return PlayerItemOutputClass
}

type _PlayerItemOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemOutput */
// An interface definition for the [PlayerItemOutput] class.
type IPlayerItemOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemOutput */
	// properties:
	SuppressesPlayerRendering() bool
	SetSuppressesPlayerRendering(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemOutput */
	// methods:
	ItemTimeForCVTimeStamp(timestamp objc.IObject /* cross-framework: TimeStamp */) objc.IObject /* cross-framework: Time */
	ItemTimeForHostTime(hostTimeInSeconds float64) objc.IObject /* cross-framework: Time */
	ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) objc.IObject /* cross-framework: Time */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemOutput */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemOutputClass) Alloc() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemOutputClass) New() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemOutput) Init() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemOutput) Autorelease() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemOutput creates a new PlayerItemOutput instance.
func NewPlayerItemOutput() PlayerItemOutput {
	return getPlayerItemOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemOutput */
// An abstract class that defines the common interface to output media data from a player item.
//
// This class provides basic methods for converting time values to the timebase of the item. It also provides an option to suppress rendering of the output associated with the specific instance of this class.


// An abstract class that defines the common interface to output media data from a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput
type PlayerItemOutput struct {
	objectivec.Object
}

// PlayerItemOutputFrom constructs a [PlayerItemOutput] from an unsafe.Pointer.
//
// An abstract class that defines the common interface to output media data from a player item.
func PlayerItemOutputFrom(ptr unsafe.Pointer) PlayerItemOutput {
	return PlayerItemOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemOutput */

// Converts a Core Video timestamp to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(for:)
func (p_ PlayerItemOutput) ItemTimeForCVTimeStamp(timestamp objc.IObject /* cross-framework: TimeStamp */) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForCVTimeStamp:"), timestamp)
	return rv
}/* debug [instance_methods/method]: ItemTimeForCVTimeStamp */


// Converts a host time, specified in seconds, to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forHostTime:)
func (p_ PlayerItemOutput) ItemTimeForHostTime(hostTimeInSeconds float64) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForHostTime:"), hostTimeInSeconds)
	return rv
}/* debug [instance_methods/method]: ItemTimeForHostTime */


// Converts a Mach host time to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forMachAbsoluteTime:)
func (p_ PlayerItemOutput) ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForMachAbsoluteTime:"), machAbsoluteTime)
	return rv
}/* debug [instance_methods/method]: ItemTimeForMachAbsoluteTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemOutput */

// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/suppressesPlayerRendering
func (p_ PlayerItemOutput) SuppressesPlayerRendering() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("suppressesPlayerRendering"))
	return rv
}/* debug [instance_properties/getter]: suppressesPlayerRendering */


// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/suppressesPlayerRendering
func (p_ PlayerItemOutput) SetSuppressesPlayerRendering(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSuppressesPlayerRendering:"), value)
}/* debug [instance_properties/setter]: suppressesPlayerRendering */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemOutput */



