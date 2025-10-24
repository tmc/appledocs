// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPRatingCommand */


/* debug [class_header]: Header for MPRatingCommand */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RatingCommand */
// An interface definition for the [RatingCommand] class.
type IRatingCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for RatingCommand */
	// properties:
	MaximumRating() float32
	SetMaximumRating(value float32)
	MinimumRating() float32
	SetMinimumRating(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RatingCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RatingCommand */
// Alloc allocates a new instance without initialization.
func (rc _RatingCommandClass) Alloc() RatingCommand {
	rv := objc.Send[RatingCommand](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RatingCommand */
// An object that provides a detailed rating for the playing item.


// An object that provides a detailed rating for the playing item.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RatingCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RatingCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RatingCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RatingCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RatingCommand */

// The maximum rating for a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/maximumRating
func (r_ RatingCommand) MaximumRating() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("maximumRating"))
	return rv
}/* debug [instance_properties/getter]: maximumRating */


// The maximum rating for a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/maximumRating
func (r_ RatingCommand) SetMaximumRating(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaximumRating:"), value)
}/* debug [instance_properties/setter]: maximumRating */


// The minimum rating for a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/minimumRating
func (r_ RatingCommand) MinimumRating() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("minimumRating"))
	return rv
}/* debug [instance_properties/getter]: minimumRating */


// The minimum rating for a command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommand/minimumRating
func (r_ RatingCommand) SetMinimumRating(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMinimumRating:"), value)
}/* debug [instance_properties/setter]: minimumRating */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPRatingCommand */



