// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemMediaDataCollector */


/* debug [class_header]: Header for AVPlayerItemMediaDataCollector */
// The class instance for the [PlayerItemMediaDataCollector] class.
var (
	PlayerItemMediaDataCollectorClass     _PlayerItemMediaDataCollectorClass
	PlayerItemMediaDataCollectorClassOnce sync.Once
)

func getPlayerItemMediaDataCollectorClass() _PlayerItemMediaDataCollectorClass {
	PlayerItemMediaDataCollectorClassOnce.Do(func() {
		PlayerItemMediaDataCollectorClass = _PlayerItemMediaDataCollectorClass{objc.GetClass("AVPlayerItemMediaDataCollector")}
	})
	return PlayerItemMediaDataCollectorClass
}

type _PlayerItemMediaDataCollectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemMediaDataCollector */
// An interface definition for the [PlayerItemMediaDataCollector] class.
type IPlayerItemMediaDataCollector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemMediaDataCollector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemMediaDataCollector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemMediaDataCollector */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemMediaDataCollectorClass) Alloc() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemMediaDataCollectorClass) New() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemMediaDataCollector) Init() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemMediaDataCollector) Autorelease() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemMediaDataCollector creates a new PlayerItemMediaDataCollector instance.
func NewPlayerItemMediaDataCollector() PlayerItemMediaDataCollector {
	return getPlayerItemMediaDataCollectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemMediaDataCollector */
// The abstract base for media data collectors.


// The abstract base for media data collectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMediaDataCollector
type PlayerItemMediaDataCollector struct {
	objectivec.Object
}

// PlayerItemMediaDataCollectorFrom constructs a [PlayerItemMediaDataCollector] from an unsafe.Pointer.
//
// The abstract base for media data collectors.
func PlayerItemMediaDataCollectorFrom(ptr unsafe.Pointer) PlayerItemMediaDataCollector {
	return PlayerItemMediaDataCollector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemMediaDataCollector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemMediaDataCollector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemMediaDataCollector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemMediaDataCollector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemMediaDataCollector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemMediaDataCollector */



