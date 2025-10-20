// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RatingCommand] class.
var (
	RatingCommandClass     _RatingCommandClass
	RatingCommandClassOnce sync.Once
)

func getRatingCommandClass() _RatingCommandClass {
	RatingCommandClassOnce.Do(func() {
		RatingCommandClass = _RatingCommandClass{objc.GetClass("MPRatingCommand")}
	})
	return RatingCommandClass
}

type _RatingCommandClass struct {
	class objc.Class
}

// An interface definition for the [RatingCommand] class.
type IRatingCommand interface {
	IRemoteCommand
}

// An object that provides a detailed rating for the playing item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand
type RatingCommand struct {
	RemoteCommand
}

// RatingCommandFrom constructs a [RatingCommand] from an unsafe.Pointer.
//
// An object that provides a detailed rating for the playing item.
func RatingCommandFrom(ptr unsafe.Pointer) RatingCommand {
	return RatingCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RatingCommandClass) Alloc() RatingCommand {
	rv := objc.Send[RatingCommand](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RatingCommandClass) New() RatingCommand {
	rv := objc.Send[RatingCommand](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RatingCommand) Init() RatingCommand {
	rv := objc.Send[RatingCommand](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RatingCommand) Autorelease() RatingCommand {
	rv := objc.Send[RatingCommand](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRatingCommand creates a new RatingCommand instance.
func NewRatingCommand() RatingCommand {
	return getRatingCommandClass().New()
}


// The maximum rating for a command.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/maximumRating
func (r_ RatingCommand) MaximumRating() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("maximumRating"))
	return rv
}


// SetMaximumRating sets the value of the maximumRating property.
// The maximum rating for a command.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/maximumRating
func (r_ RatingCommand) SetMaximumRating(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaximumRating:"), value)
}
// The minimum rating for a command.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/minimumRating
func (r_ RatingCommand) MinimumRating() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("minimumRating"))
	return rv
}


// SetMinimumRating sets the value of the minimumRating property.
// The minimum rating for a command.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/minimumRating
func (r_ RatingCommand) SetMinimumRating(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinimumRating:"), value)
}


