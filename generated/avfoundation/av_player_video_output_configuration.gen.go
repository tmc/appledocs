// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerVideoOutputConfiguration] class.
var (
	PlayerVideoOutputConfigurationClass     _PlayerVideoOutputConfigurationClass
	PlayerVideoOutputConfigurationClassOnce sync.Once
)

func getPlayerVideoOutputConfigurationClass() _PlayerVideoOutputConfigurationClass {
	PlayerVideoOutputConfigurationClassOnce.Do(func() {
		PlayerVideoOutputConfigurationClass = _PlayerVideoOutputConfigurationClass{objc.GetClass("AVPlayerVideoOutputConfiguration")}
	})
	return PlayerVideoOutputConfigurationClass
}

type _PlayerVideoOutputConfigurationClass struct {
	class objc.Class
}





// An interface definition for the [PlayerVideoOutputConfiguration] class.
type IPlayerVideoOutputConfiguration interface {
	objectivec.IObject
	

	// properties:
	DataChannelDescriptions() foundation.foundation.INSArray
	ActivationTime() objectivec.IObject
	PreferredTransform() corefoundation.CGAffineTransform
	SourcePlayerItem() IAVPlayerItem
	DataChannelDescription() objectivec.IObject
	SetDataChannelDescription(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerVideoOutputConfigurationClass) Alloc() PlayerVideoOutputConfiguration {
	rv := objc.Send[PlayerVideoOutputConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerVideoOutputConfigurationClass) New() PlayerVideoOutputConfiguration {
	rv := objc.Send[PlayerVideoOutputConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerVideoOutputConfiguration) Init() PlayerVideoOutputConfiguration {
	rv := objc.Send[PlayerVideoOutputConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerVideoOutputConfiguration) Autorelease() PlayerVideoOutputConfiguration {
	rv := objc.Send[PlayerVideoOutputConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerVideoOutputConfiguration creates a new PlayerVideoOutputConfiguration instance.
func NewPlayerVideoOutputConfiguration() PlayerVideoOutputConfiguration {
	return getPlayerVideoOutputConfigurationClass().New()
}





// An object that provides configuration information for the related player item.


// An object that provides configuration information for the related player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration
type PlayerVideoOutputConfiguration struct {
	objectivec.Object
}

// PlayerVideoOutputConfigurationFrom constructs a [PlayerVideoOutputConfiguration] from an unsafe.Pointer.
//
// An object that provides configuration information for the related player item.
func PlayerVideoOutputConfigurationFrom(ptr unsafe.Pointer) PlayerVideoOutputConfiguration {
	return PlayerVideoOutputConfiguration{objectivec.Object{objc.ID(ptr)}}
}

























// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutputConfiguration/dataChannelDescriptions
func (p_ PlayerVideoOutputConfiguration) DataChannelDescriptions() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](p_.ID, objc.Sel("dataChannelDescriptions"))
	return rv
}


// The host time this configuration became active on its associated player object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/activationTime
func (p_ PlayerVideoOutputConfiguration) ActivationTime() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("activationTime"))
	return rv
}


// The preferred transform of the visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/preferredTransform
func (p_ PlayerVideoOutputConfiguration) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](p_.ID, objc.Sel("preferredTransform"))
	return rv
}


// The player item that’s the source of this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/sourcePlayerItem
func (p_ PlayerVideoOutputConfiguration) SourcePlayerItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("sourcePlayerItem"))
	return rv
}


// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayervideooutput/configuration/datachanneldescription
func (p_ PlayerVideoOutputConfiguration) DataChannelDescription() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("dataChannelDescription"))
	return rv
}


// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayervideooutput/configuration/datachanneldescription
func (p_ PlayerVideoOutputConfiguration) SetDataChannelDescription(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataChannelDescription:"), value)
}








