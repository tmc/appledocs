// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerQueueDescriptor */


/* debug [class_header]: Header for MPMusicPlayerQueueDescriptor */
// The class instance for the [MusicPlayerQueueDescriptor] class.
var (
	MusicPlayerQueueDescriptorClass     _MusicPlayerQueueDescriptorClass
	MusicPlayerQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerQueueDescriptorClass() _MusicPlayerQueueDescriptorClass {
	MusicPlayerQueueDescriptorClassOnce.Do(func() {
		MusicPlayerQueueDescriptorClass = _MusicPlayerQueueDescriptorClass{objc.GetClass("MPMusicPlayerQueueDescriptor")}
	})
	return MusicPlayerQueueDescriptorClass
}

type _MusicPlayerQueueDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerQueueDescriptor */
// An interface definition for the [MusicPlayerQueueDescriptor] class.
type IMusicPlayerQueueDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MusicPlayerQueueDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerQueueDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerQueueDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerQueueDescriptorClass) Alloc() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicPlayerQueueDescriptorClass) New() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerQueueDescriptor) Init() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerQueueDescriptor) Autorelease() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerQueueDescriptor creates a new MusicPlayerQueueDescriptor instance.
func NewMusicPlayerQueueDescriptor() MusicPlayerQueueDescriptor {
	return getMusicPlayerQueueDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerQueueDescriptor */
// The abstract base class for audio media item and store queue descriptors.


// The abstract base class for audio media item and store queue descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerQueueDescriptor
type MusicPlayerQueueDescriptor struct {
	objectivec.Object
}

// MusicPlayerQueueDescriptorFrom constructs a [MusicPlayerQueueDescriptor] from an unsafe.Pointer.
//
// The abstract base class for audio media item and store queue descriptors.
func MusicPlayerQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerQueueDescriptor {
	return MusicPlayerQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerQueueDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerQueueDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerQueueDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerQueueDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerQueueDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerQueueDescriptor */



