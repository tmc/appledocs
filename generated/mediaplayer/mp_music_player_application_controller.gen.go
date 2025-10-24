// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerApplicationController */


/* debug [class_header]: Header for MPMusicPlayerApplicationController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerApplicationController */
// An interface definition for the [MusicPlayerApplicationController] class.
type IMusicPlayerApplicationController interface {
	IMusicPlayerController
	
/* debug [class_interface_properties]: Properties for MusicPlayerApplicationController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerApplicationController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerApplicationController */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerApplicationControllerClass) Alloc() MusicPlayerApplicationController {
	rv := objc.Send[MusicPlayerApplicationController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerApplicationController */
// A media player object that you use to revise the queue that’s currently playing.


// A media player object that you use to revise the queue that’s currently playing.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerApplicationController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerApplicationController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerApplicationController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerApplicationController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerApplicationController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerApplicationController */


