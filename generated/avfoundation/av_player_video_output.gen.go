// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerVideoOutput */


/* debug [class_header]: Header for AVPlayerVideoOutput */
// The class instance for the [PlayerVideoOutput] class.
var (
	PlayerVideoOutputClass     _PlayerVideoOutputClass
	PlayerVideoOutputClassOnce sync.Once
)

func getPlayerVideoOutputClass() _PlayerVideoOutputClass {
	PlayerVideoOutputClassOnce.Do(func() {
		PlayerVideoOutputClass = _PlayerVideoOutputClass{objc.GetClass("AVPlayerVideoOutput")}
	})
	return PlayerVideoOutputClass
}

type _PlayerVideoOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerVideoOutput */
// An interface definition for the [PlayerVideoOutput] class.
type IPlayerVideoOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerVideoOutput */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerVideoOutput */
	// methods:
	CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime objc.IObject /* cross-framework: Time */, presentationTimeStampOut objc.IObject /* cross-framework: Time */, activeConfigurationOut IAVPlayerVideoOutputConfiguration) TaggedBufferGroupRef /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerVideoOutput */
// Alloc allocates a new instance without initialization.
func (pc _PlayerVideoOutputClass) Alloc() PlayerVideoOutput {
	rv := objc.Send[PlayerVideoOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerVideoOutputClass) New() PlayerVideoOutput {
	rv := objc.Send[PlayerVideoOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerVideoOutput) Init() PlayerVideoOutput {
	rv := objc.Send[PlayerVideoOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerVideoOutput) Autorelease() PlayerVideoOutput {
	rv := objc.Send[PlayerVideoOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerVideoOutput creates a new PlayerVideoOutput instance.
func NewPlayerVideoOutput() PlayerVideoOutput {
	return getPlayerVideoOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerVideoOutput */
// An object that receives video data from a player object.
//
// Attach a video output to an object to access the player’s video data as objects.


// An object that receives video data from a player object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput
type PlayerVideoOutput struct {
	objectivec.Object
}

// PlayerVideoOutputFrom constructs a [PlayerVideoOutput] from an unsafe.Pointer.
//
// An object that receives video data from a player object.
func PlayerVideoOutputFrom(ptr unsafe.Pointer) PlayerVideoOutput {
	return PlayerVideoOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerVideoOutput */

// Creates a video output from a specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/init(specification:)
func NewPlayerVideoOutputWithSpecification(specification IAVVideoOutputSpecification) PlayerVideoOutput {
	instance := getPlayerVideoOutputClass().Alloc()
	rv := objc.Send[PlayerVideoOutput](instance.ID, objc.Sel("initWithSpecification:"), specification)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerVideoOutputWithSpecification */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerVideoOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerVideoOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerVideoOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:
func (p_ PlayerVideoOutput) CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime objc.IObject /* cross-framework: Time */, presentationTimeStampOut objc.IObject /* cross-framework: Time */, activeConfigurationOut IAVPlayerVideoOutputConfiguration) TaggedBufferGroupRef /* not a class type */ {
	rv := objc.Send[TaggedBufferGroupRef](p_.ID, objc.Sel("copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:"), hostTime, presentationTimeStampOut, activeConfigurationOut)
	return rv
}/* debug [instance_methods/method]: CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerVideoOutput */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerVideoOutput */


