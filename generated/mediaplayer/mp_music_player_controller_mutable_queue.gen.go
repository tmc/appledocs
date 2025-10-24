// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerControllerMutableQueue */


/* debug [class_header]: Header for MPMusicPlayerControllerMutableQueue */
// The class instance for the [MusicPlayerControllerMutableQueue] class.
var (
	MusicPlayerControllerMutableQueueClass     _MusicPlayerControllerMutableQueueClass
	MusicPlayerControllerMutableQueueClassOnce sync.Once
)

func getMusicPlayerControllerMutableQueueClass() _MusicPlayerControllerMutableQueueClass {
	MusicPlayerControllerMutableQueueClassOnce.Do(func() {
		MusicPlayerControllerMutableQueueClass = _MusicPlayerControllerMutableQueueClass{objc.GetClass("MPMusicPlayerControllerMutableQueue")}
	})
	return MusicPlayerControllerMutableQueueClass
}

type _MusicPlayerControllerMutableQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerControllerMutableQueue */
// An interface definition for the [MusicPlayerControllerMutableQueue] class.
type IMusicPlayerControllerMutableQueue interface {
	IMusicPlayerControllerQueue
	
/* debug [class_interface_properties]: Properties for MusicPlayerControllerMutableQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerControllerMutableQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerControllerMutableQueue */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerMutableQueueClass) Alloc() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicPlayerControllerMutableQueueClass) New() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerControllerMutableQueue) Init() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerControllerMutableQueue) Autorelease() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerControllerMutableQueue creates a new MusicPlayerControllerMutableQueue instance.
func NewMusicPlayerControllerMutableQueue() MusicPlayerControllerMutableQueue {
	return getMusicPlayerControllerMutableQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerControllerMutableQueue */
// A mutable queue containing the media items to play.


// A mutable queue containing the media items to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue
type MusicPlayerControllerMutableQueue struct {
	MusicPlayerControllerQueue
}

// MusicPlayerControllerMutableQueueFrom constructs a [MusicPlayerControllerMutableQueue] from an unsafe.Pointer.
//
// A mutable queue containing the media items to play.
func MusicPlayerControllerMutableQueueFrom(ptr unsafe.Pointer) MusicPlayerControllerMutableQueue {
	return MusicPlayerControllerMutableQueue{
		MusicPlayerControllerQueue: MusicPlayerControllerQueueFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerControllerMutableQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerControllerMutableQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerControllerMutableQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerControllerMutableQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerControllerMutableQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerControllerMutableQueue */


