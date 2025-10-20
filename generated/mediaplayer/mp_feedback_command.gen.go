// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FeedbackCommand] class.
var (
	FeedbackCommandClass     _FeedbackCommandClass
	FeedbackCommandClassOnce sync.Once
)

func getFeedbackCommandClass() _FeedbackCommandClass {
	FeedbackCommandClassOnce.Do(func() {
		FeedbackCommandClass = _FeedbackCommandClass{objc.GetClass("MPFeedbackCommand")}
	})
	return FeedbackCommandClass
}

type _FeedbackCommandClass struct {
	class objc.Class
}

// An interface definition for the [FeedbackCommand] class.
type IFeedbackCommand interface {
	IRemoteCommand
}

// An object that reflects the feedback state for the playing item.
//
// The shared object vends feedback objects for liking, disliking, and bookmarking media items. Use these objects to register handlers for the types of feedback your app supports and to perform the appropriate tasks when that feedback changes. When the currently playing item changes, you can also use this object to set the feedback state for the new item. When the state of a feedback item changes, the system delivers an appropriate event to registered handlers of this object. Your handler code must determine which media item receives the feedback and then apply the update the feedback state for that item. You might also perform other tasks related to receiving feedback. For example, if the user likes the currently playing song, you might update the appropriate UI in your app or use the information to recommend similar songs.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommand
type FeedbackCommand struct {
	RemoteCommand
}

// FeedbackCommandFrom constructs a [FeedbackCommand] from an unsafe.Pointer.
//
// An object that reflects the feedback state for the playing item.
func FeedbackCommandFrom(ptr unsafe.Pointer) FeedbackCommand {
	return FeedbackCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FeedbackCommandClass) Alloc() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FeedbackCommandClass) New() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeedbackCommand) Init() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeedbackCommand) Autorelease() FeedbackCommand {
	rv := objc.Send[FeedbackCommand](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeedbackCommand creates a new FeedbackCommand instance.
func NewFeedbackCommand() FeedbackCommand {
	return getFeedbackCommandClass().New()
}




