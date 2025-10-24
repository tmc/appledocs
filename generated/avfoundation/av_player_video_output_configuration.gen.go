// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerVideoOutputConfiguration */


/* debug [class_header]: Header for AVPlayerVideoOutputConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerVideoOutputConfiguration */
// An interface definition for the [PlayerVideoOutputConfiguration] class.
type IPlayerVideoOutputConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerVideoOutputConfiguration */
	// properties:
	DataChannelDescriptions() objc.IObject /* cross-framework: NSArray */
	ActivationTime() objc.IObject /* cross-framework: Time */
	PreferredTransform() corefoundation.CGAffineTransform
	SourcePlayerItem() IAVPlayerItem
	DataChannelDescription() Tag /* not a class type */
	SetDataChannelDescription(value Tag /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerVideoOutputConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerVideoOutputConfiguration */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerVideoOutputConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerVideoOutputConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerVideoOutputConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerVideoOutputConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerVideoOutputConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerVideoOutputConfiguration */

// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutputConfiguration/dataChannelDescriptions
func (p_ PlayerVideoOutputConfiguration) DataChannelDescriptions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](p_.ID, objc.Sel("dataChannelDescriptions"))
	return rv
}/* debug [instance_properties/getter]: dataChannelDescriptions */


// The host time this configuration became active on its associated player object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/activationTime
func (p_ PlayerVideoOutputConfiguration) ActivationTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("activationTime"))
	return rv
}/* debug [instance_properties/getter]: activationTime */


// The preferred transform of the visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/preferredTransform
func (p_ PlayerVideoOutputConfiguration) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](p_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The player item that’s the source of this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/sourcePlayerItem
func (p_ PlayerVideoOutputConfiguration) SourcePlayerItem() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("sourcePlayerItem"))
	return rv
}/* debug [instance_properties/getter]: sourcePlayerItem */


// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayervideooutput/configuration/datachanneldescription
func (p_ PlayerVideoOutputConfiguration) DataChannelDescription() Tag /* not a class type */ {
	rv := objc.Send[Tag](p_.ID, objc.Sel("dataChannelDescription"))
	return rv
}/* debug [instance_properties/getter]: dataChannelDescription */


// An array of data channels selected for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayervideooutput/configuration/datachanneldescription
func (p_ PlayerVideoOutputConfiguration) SetDataChannelDescription(value Tag /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDataChannelDescription:"), value)
}/* debug [instance_properties/setter]: dataChannelDescription */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerVideoOutputConfiguration */



