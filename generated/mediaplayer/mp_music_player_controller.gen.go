// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MusicPlayerController] class.
var (
	MusicPlayerControllerClass     _MusicPlayerControllerClass
	MusicPlayerControllerClassOnce sync.Once
)

func getMusicPlayerControllerClass() _MusicPlayerControllerClass {
	MusicPlayerControllerClassOnce.Do(func() {
		MusicPlayerControllerClass = _MusicPlayerControllerClass{objc.GetClass("MPMusicPlayerController")}
	})
	return MusicPlayerControllerClass
}

type _MusicPlayerControllerClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerController] class.
type IMusicPlayerController interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that plays audio media items from the device’s Music app library.
//
// Create an instance of a music player to play media items in your app. There are two types of music players: An plays music locally within your app. It isn’t aware of the Music app’s Now Playing item, nor does it affect the Music app’s state. There are two application music players: and . The application queue player provides greater control over the contents of the queue and is the preferred player. The employs the built-in Music app on your behalf. On instantiation, it takes on the current Music app state, such as the identification of the Now Playing item. If a user switches away from your app while music is playing, that music continues to play. The Music app then has your music player’s most recently-set repeat mode, shuffle mode, playback state, and Now Playing item. Creating a new instance of and not specifying the player type returns a system music player.


// An object that plays audio media items from the device’s Music app library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController
type MusicPlayerController struct {
	objectivec.Object
}

// MusicPlayerControllerFrom constructs a [MusicPlayerController] from an unsafe.Pointer.
//
// An object that plays audio media items from the device’s Music app library.
func MusicPlayerControllerFrom(ptr unsafe.Pointer) MusicPlayerController {
	return MusicPlayerController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerClass) Alloc() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerControllerClass) New() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerController) Init() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerController) Autorelease() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerController creates a new MusicPlayerController instance.
func NewMusicPlayerController() MusicPlayerController {
	return getMusicPlayerControllerClass().New()
}



// Returns the application music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/applicationMusicPlayer
func (mc _MusicPlayerControllerClass) ApplicationMusicPlayer() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("applicationMusicPlayer"))
	return rv
}

// Returns the application queue music player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/applicationQueuePlayer
func (mc _MusicPlayerControllerClass) ApplicationQueuePlayer() IMPMusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](objc.ID(mc.class), objc.Sel("applicationQueuePlayer"))
	return rv
}

// Returns the iPod music player, which controls the iPod app’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/iPodMusicPlayer
func (mc _MusicPlayerControllerClass) IPodMusicPlayer() MusicPlayerController {
	rv := objc.Send[MusicPlayerController](objc.ID(mc.class), objc.Sel("iPodMusicPlayer"))
	return rv
}

// Returns the system music player, which controls the Music app’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerController/systemMusicPlayer
func (mc _MusicPlayerControllerClass) SystemMusicPlayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("systemMusicPlayer"))
	return rv
}

