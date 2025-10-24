// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [PlayerLayer] class.
var (
	PlayerLayerClass     _PlayerLayerClass
	PlayerLayerClassOnce sync.Once
)

func getPlayerLayerClass() _PlayerLayerClass {
	PlayerLayerClassOnce.Do(func() {
		PlayerLayerClass = _PlayerLayerClass{objc.GetClass("AVPlayerLayer")}
	})
	return PlayerLayerClass
}

type _PlayerLayerClass struct {
	class objc.Class
}





// An interface definition for the [PlayerLayer] class.
type IPlayerLayer interface {
	ILayer
	

	// properties:
	ReadyForDisplay() bool
	PixelBufferAttributes() foundation.IDictionary
	SetPixelBufferAttributes(value foundation.IDictionary)
	Player() IAVPlayer
	SetPlayer(value IAVPlayer)
	VideoGravity() LayerVideoGravity /* typedef */
	SetVideoGravity(value LayerVideoGravity /* typedef */)
	VideoRect() corefoundation.CGRect
	IsReadyForDisplay() bool
	SetIsReadyForDisplay(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerLayerClass) Alloc() PlayerLayer {
	rv := objc.Send[PlayerLayer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerLayerClass) New() PlayerLayer {
	rv := objc.Send[PlayerLayer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerLayer) Init() PlayerLayer {
	rv := objc.Send[PlayerLayer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerLayer) Autorelease() PlayerLayer {
	rv := objc.Send[PlayerLayer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerLayer creates a new PlayerLayer instance.
func NewPlayerLayer() PlayerLayer {
	return getPlayerLayerClass().New()
}





// An object that presents the visual contents of a player object.
//
// A common way to use this object in iOS or tvOS is as the backing layer for a , as the following example shows:


// An object that presents the visual contents of a player object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer
type PlayerLayer struct {
	Layer
}

// PlayerLayerFrom constructs a [PlayerLayer] from an unsafe.Pointer.
//
// An object that presents the visual contents of a player object.
func PlayerLayerFrom(ptr unsafe.Pointer) PlayerLayer {
	return PlayerLayer{
		Layer: LayerFrom(ptr),
	}
}






// Creates a layer object to present the visual contents of a player’s current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/init(player:)
func NewPlayerLayerWithPlayer(player IAVPlayer) PlayerLayer {
	rv := objc.Send[PlayerLayer](objc.ID(getPlayerLayerClass().class), objc.Sel("playerLayerWithPlayer:"), player)
	return rv
}







// Creates a layer object to present the visual contents of a player’s current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/init(player:)
func (pc _PlayerLayerClass) PlayerLayerWithPlayer(player IAVPlayer) IPlayerLayer {
	rv := objc.Send[PlayerLayer](objc.ID(pc.class), objc.Sel("playerLayerWithPlayer:"), player)
	return rv
}

















// A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/isReadyForDisplay
func (p_ PlayerLayer) ReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readyForDisplay"))
	return rv
}


// The attributes of the visual output that displays in the player layer during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/pixelBufferAttributes
func (p_ PlayerLayer) PixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("pixelBufferAttributes"))
	return rv
}


// The attributes of the visual output that displays in the player layer during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/pixelBufferAttributes
func (p_ PlayerLayer) SetPixelBufferAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelBufferAttributes:"), value)
}


// The player whose visual content the layer displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/player
func (p_ PlayerLayer) Player() IAVPlayer {
	rv := objc.Send[Player](p_.ID, objc.Sel("player"))
	return rv
}


// The player whose visual content the layer displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/player
func (p_ PlayerLayer) SetPlayer(value IAVPlayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayer:"), value)
}


// A value that specifies how the layer displays the player’s visual content within the layer’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/videoGravity
func (p_ PlayerLayer) VideoGravity() LayerVideoGravity /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("videoGravity"))
	return rv
}


// A value that specifies how the layer displays the player’s visual content within the layer’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/videoGravity
func (p_ PlayerLayer) SetVideoGravity(value LayerVideoGravity /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoGravity:"), value)
}


// The current size and position of the video image that displays within the layer’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/videoRect
func (p_ PlayerLayer) VideoRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("videoRect"))
	return rv
}


// A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/isreadyfordisplay
func (p_ PlayerLayer) IsReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerlayer/isreadyfordisplay
func (p_ PlayerLayer) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}







