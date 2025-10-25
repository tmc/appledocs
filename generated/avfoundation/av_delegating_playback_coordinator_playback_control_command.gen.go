// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDelegatingPlaybackCoordinatorPlaybackControlCommand */


/* debug [class_header]: Header for AVDelegatingPlaybackCoordinatorPlaybackControlCommand */
// The class instance for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
var (
	DelegatingPlaybackCoordinatorPlaybackControlCommandClass     _DelegatingPlaybackCoordinatorPlaybackControlCommandClass
	DelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorPlaybackControlCommandClass() _DelegatingPlaybackCoordinatorPlaybackControlCommandClass {
	DelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorPlaybackControlCommandClass = _DelegatingPlaybackCoordinatorPlaybackControlCommandClass{objc.GetClass("AVDelegatingPlaybackCoordinatorPlaybackControlCommand")}
	})
	return DelegatingPlaybackCoordinatorPlaybackControlCommandClass
}

type _DelegatingPlaybackCoordinatorPlaybackControlCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DelegatingPlaybackCoordinatorPlaybackControlCommand */
// An interface definition for the [DelegatingPlaybackCoordinatorPlaybackControlCommand] class.
type IDelegatingPlaybackCoordinatorPlaybackControlCommand interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DelegatingPlaybackCoordinatorPlaybackControlCommand */
	// properties:
	ExpectedCurrentItemIdentifier() objc.IObject /* cross-framework: NSString */
	Originator() IAVCoordinatedPlaybackParticipant
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DelegatingPlaybackCoordinatorPlaybackControlCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DelegatingPlaybackCoordinatorPlaybackControlCommand */
// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) Alloc() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorPlaybackControlCommandClass) New() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Init() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Autorelease() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[DelegatingPlaybackCoordinatorPlaybackControlCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinatorPlaybackControlCommand creates a new DelegatingPlaybackCoordinatorPlaybackControlCommand instance.
func NewDelegatingPlaybackCoordinatorPlaybackControlCommand() DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return getDelegatingPlaybackCoordinatorPlaybackControlCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DelegatingPlaybackCoordinatorPlaybackControlCommand */
// An abstract superclass for playback commands.
//
// Playback commands inherit state that identifies their originator and applicable item.


// An abstract superclass for playback commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand
type DelegatingPlaybackCoordinatorPlaybackControlCommand struct {
	objectivec.Object
}

// DelegatingPlaybackCoordinatorPlaybackControlCommandFrom constructs a [DelegatingPlaybackCoordinatorPlaybackControlCommand] from an unsafe.Pointer.
//
// An abstract superclass for playback commands.
func DelegatingPlaybackCoordinatorPlaybackControlCommandFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinatorPlaybackControlCommand {
	return DelegatingPlaybackCoordinatorPlaybackControlCommand{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DelegatingPlaybackCoordinatorPlaybackControlCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DelegatingPlaybackCoordinatorPlaybackControlCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DelegatingPlaybackCoordinatorPlaybackControlCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DelegatingPlaybackCoordinatorPlaybackControlCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DelegatingPlaybackCoordinatorPlaybackControlCommand */

// An item identifier the coordinator issues the command for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand/expectedCurrentItemIdentifier
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) ExpectedCurrentItemIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("expectedCurrentItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: expectedCurrentItemIdentifier */


// The participant that causes the coordinator to issue the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand/originator
func (d_ DelegatingPlaybackCoordinatorPlaybackControlCommand) Originator() IAVCoordinatedPlaybackParticipant {
	rv := objc.Send[CoordinatedPlaybackParticipant](d_.ID, objc.Sel("originator"))
	return rv
}/* debug [instance_properties/getter]: originator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDelegatingPlaybackCoordinatorPlaybackControlCommand */



