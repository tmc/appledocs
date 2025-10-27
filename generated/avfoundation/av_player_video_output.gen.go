// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [PlayerVideoOutput] class.
type IPlayerVideoOutput interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime objectivec.IObject, presentationTimeStampOut objectivec.IObject, activeConfigurationOut IAVPlayerVideoOutputConfiguration) TaggedBufferGroupRef /* not a class type */


}





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






// Creates a video output from a specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/init(specification:)
func NewPlayerVideoOutputWithSpecification(specification IAVVideoOutputSpecification) PlayerVideoOutput {
	instance := getPlayerVideoOutputClass().Alloc()
	rv := objc.Send[PlayerVideoOutput](instance.ID, objc.Sel("initWithSpecification:"), specification)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:
func (p_ PlayerVideoOutput) CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime objectivec.IObject, presentationTimeStampOut objectivec.IObject, activeConfigurationOut IAVPlayerVideoOutputConfiguration) TaggedBufferGroupRef /* not a class type */ {
	rv := objc.Send[TaggedBufferGroupRef](p_.ID, objc.Sel("copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:"), hostTime, presentationTimeStampOut, activeConfigurationOut)
	return rv
}












