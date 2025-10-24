// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerControllerQueue */


/* debug [class_header]: Header for MPMusicPlayerControllerQueue */
// The class instance for the [MusicPlayerControllerQueue] class.
var (
	MusicPlayerControllerQueueClass     _MusicPlayerControllerQueueClass
	MusicPlayerControllerQueueClassOnce sync.Once
)

func getMusicPlayerControllerQueueClass() _MusicPlayerControllerQueueClass {
	MusicPlayerControllerQueueClassOnce.Do(func() {
		MusicPlayerControllerQueueClass = _MusicPlayerControllerQueueClass{objc.GetClass("MPMusicPlayerControllerQueue")}
	})
	return MusicPlayerControllerQueueClass
}

type _MusicPlayerControllerQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerControllerQueue */
// An interface definition for the [MusicPlayerControllerQueue] class.
type IMusicPlayerControllerQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MusicPlayerControllerQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerControllerQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerControllerQueue */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerQueueClass) Alloc() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicPlayerControllerQueueClass) New() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerControllerQueue) Init() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerControllerQueue) Autorelease() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerControllerQueue creates a new MusicPlayerControllerQueue instance.
func NewMusicPlayerControllerQueue() MusicPlayerControllerQueue {
	return getMusicPlayerControllerQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerControllerQueue */
// An immutable queue containing the media items to play.
//
// An object contains the current queue for an application queue music player. To add or remove media items from a playing queue, use . The results of the method is an object that updates the playing queue. You don’t create your own instance of this class.


// An immutable queue containing the media items to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerQueue
type MusicPlayerControllerQueue struct {
	objectivec.Object
}

// MusicPlayerControllerQueueFrom constructs a [MusicPlayerControllerQueue] from an unsafe.Pointer.
//
// An immutable queue containing the media items to play.
func MusicPlayerControllerQueueFrom(ptr unsafe.Pointer) MusicPlayerControllerQueue {
	return MusicPlayerControllerQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerControllerQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerControllerQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerControllerQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerControllerQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerControllerQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerControllerQueue */


