// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MusicPlayerApplicationController] class.
var (
	MusicPlayerApplicationControllerClass     _MusicPlayerApplicationControllerClass
	MusicPlayerApplicationControllerClassOnce sync.Once
)

func getMusicPlayerApplicationControllerClass() _MusicPlayerApplicationControllerClass {
	MusicPlayerApplicationControllerClassOnce.Do(func() {
		MusicPlayerApplicationControllerClass = _MusicPlayerApplicationControllerClass{objc.GetClass("MPMusicPlayerApplicationController")}
	})
	return MusicPlayerApplicationControllerClass
}

type _MusicPlayerApplicationControllerClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerApplicationController] class.
type IMusicPlayerApplicationController interface {
	IMusicPlayerController
	PerformQueueTransactionCompletionHandler(queueTransaction unsafe.Pointer, completionHandler unsafe.Pointer)
}

// A media player object that you use to revise the queue that’s currently playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerApplicationController
type MusicPlayerApplicationController struct {
	MusicPlayerController
}

// MusicPlayerApplicationControllerFrom constructs a [MusicPlayerApplicationController] from an unsafe.Pointer.
//
// A media player object that you use to revise the queue that’s currently playing.
func MusicPlayerApplicationControllerFrom(ptr unsafe.Pointer) MusicPlayerApplicationController {
	return MusicPlayerApplicationController{
		MusicPlayerController: MusicPlayerControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerApplicationControllerClass) Alloc() MusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerApplicationControllerClass) New() MusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerApplicationController) Init() MusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerApplicationController) Autorelease() MusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerApplicationController creates a new MusicPlayerApplicationController instance.
func NewMusicPlayerApplicationController() MusicPlayerApplicationController {
	return getMusicPlayerApplicationControllerClass().New()
}


// Changes the contents of the media items in the queue.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerApplicationController/perform(queueTransaction:completionHandler:)
func (m_ MusicPlayerApplicationController) PerformQueueTransactionCompletionHandler(queueTransaction unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("performQueueTransaction:completionHandler:"), queueTransaction, completionHandler)
}



