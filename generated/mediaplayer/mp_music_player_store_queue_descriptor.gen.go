// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerStoreQueueDescriptor */


/* debug [class_header]: Header for MPMusicPlayerStoreQueueDescriptor */
// The class instance for the [MusicPlayerStoreQueueDescriptor] class.
var (
	MusicPlayerStoreQueueDescriptorClass     _MusicPlayerStoreQueueDescriptorClass
	MusicPlayerStoreQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerStoreQueueDescriptorClass() _MusicPlayerStoreQueueDescriptorClass {
	MusicPlayerStoreQueueDescriptorClassOnce.Do(func() {
		MusicPlayerStoreQueueDescriptorClass = _MusicPlayerStoreQueueDescriptorClass{objc.GetClass("MPMusicPlayerStoreQueueDescriptor")}
	})
	return MusicPlayerStoreQueueDescriptorClass
}

type _MusicPlayerStoreQueueDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerStoreQueueDescriptor */
// An interface definition for the [MusicPlayerStoreQueueDescriptor] class.
type IMusicPlayerStoreQueueDescriptor interface {
	IMusicPlayerQueueDescriptor
	
/* debug [class_interface_properties]: Properties for MusicPlayerStoreQueueDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerStoreQueueDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerStoreQueueDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerStoreQueueDescriptorClass) Alloc() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicPlayerStoreQueueDescriptorClass) New() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerStoreQueueDescriptor) Init() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerStoreQueueDescriptor) Autorelease() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerStoreQueueDescriptor creates a new MusicPlayerStoreQueueDescriptor instance.
func NewMusicPlayerStoreQueueDescriptor() MusicPlayerStoreQueueDescriptor {
	return getMusicPlayerStoreQueueDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerStoreQueueDescriptor */
// A set of properties and methods for modifying items, based on their store identifier, in the player’s queue.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item to play.


// A set of properties and methods for modifying items, based on their store identifier, in the player’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor
type MusicPlayerStoreQueueDescriptor struct {
	MusicPlayerQueueDescriptor
}

// MusicPlayerStoreQueueDescriptorFrom constructs a [MusicPlayerStoreQueueDescriptor] from an unsafe.Pointer.
//
// A set of properties and methods for modifying items, based on their store identifier, in the player’s queue.
func MusicPlayerStoreQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerStoreQueueDescriptor {
	return MusicPlayerStoreQueueDescriptor{
		MusicPlayerQueueDescriptor: MusicPlayerQueueDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerStoreQueueDescriptor */

// Creates a new queue descriptor using the designated store identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/init(storeIDs:)
func NewMusicPlayerStoreQueueDescriptorWithStoreIDs(storeIDs []string) MusicPlayerStoreQueueDescriptor {
	instance := getMusicPlayerStoreQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](instance.ID, objc.Sel("initWithStoreIDs:"), storeIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMusicPlayerStoreQueueDescriptorWithStoreIDs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerStoreQueueDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerStoreQueueDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerStoreQueueDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerStoreQueueDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerStoreQueueDescriptor */


