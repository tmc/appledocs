// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCoordinatedPlaybackSuspension */


/* debug [class_header]: Header for AVCoordinatedPlaybackSuspension */
// The class instance for the [CoordinatedPlaybackSuspension] class.
var (
	CoordinatedPlaybackSuspensionClass     _CoordinatedPlaybackSuspensionClass
	CoordinatedPlaybackSuspensionClassOnce sync.Once
)

func getCoordinatedPlaybackSuspensionClass() _CoordinatedPlaybackSuspensionClass {
	CoordinatedPlaybackSuspensionClassOnce.Do(func() {
		CoordinatedPlaybackSuspensionClass = _CoordinatedPlaybackSuspensionClass{objc.GetClass("AVCoordinatedPlaybackSuspension")}
	})
	return CoordinatedPlaybackSuspensionClass
}

type _CoordinatedPlaybackSuspensionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CoordinatedPlaybackSuspension */
// An interface definition for the [CoordinatedPlaybackSuspension] class.
type ICoordinatedPlaybackSuspension interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CoordinatedPlaybackSuspension */
	// properties:
	BeginDate() objc.IObject /* cross-framework: NSDate */
	Reason() CoordinatedPlaybackSuspensionReason /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CoordinatedPlaybackSuspension */
	// methods:
	End()
	EndProposingNewTime(time objc.IObject /* cross-framework: Time */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CoordinatedPlaybackSuspension */
// Alloc allocates a new instance without initialization.
func (cc _CoordinatedPlaybackSuspensionClass) Alloc() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CoordinatedPlaybackSuspensionClass) New() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoordinatedPlaybackSuspension) Init() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoordinatedPlaybackSuspension) Autorelease() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoordinatedPlaybackSuspension creates a new CoordinatedPlaybackSuspension instance.
func NewCoordinatedPlaybackSuspension() CoordinatedPlaybackSuspension {
	return getCoordinatedPlaybackSuspensionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CoordinatedPlaybackSuspension */
// An object that represents a temporary suspension of coordinated playback.
//
// See the playback coordinator’s method for details about suspending playback.


// An object that represents a temporary suspension of coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension
type CoordinatedPlaybackSuspension struct {
	objectivec.Object
}

// CoordinatedPlaybackSuspensionFrom constructs a [CoordinatedPlaybackSuspension] from an unsafe.Pointer.
//
// An object that represents a temporary suspension of coordinated playback.
func CoordinatedPlaybackSuspensionFrom(ptr unsafe.Pointer) CoordinatedPlaybackSuspension {
	return CoordinatedPlaybackSuspension{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CoordinatedPlaybackSuspension *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CoordinatedPlaybackSuspension */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CoordinatedPlaybackSuspension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CoordinatedPlaybackSuspension */

// Ends a suspension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end()
func (c_ CoordinatedPlaybackSuspension) End() {
	objc.Send[objc.ID](c_.ID, objc.Sel("end"))
}/* debug [instance_methods/method]: End */


// Ends a suspension and proposes a new playback time to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end(proposingNewTime:)
func (c_ CoordinatedPlaybackSuspension) EndProposingNewTime(time objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endProposingNewTime:"), time)
}/* debug [instance_methods/method]: EndProposingNewTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CoordinatedPlaybackSuspension */

// The time the suspension begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/beginDate
func (c_ CoordinatedPlaybackSuspension) BeginDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("beginDate"))
	return rv
}/* debug [instance_properties/getter]: beginDate */


// The reason for the suspension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/reason-swift.property
func (c_ CoordinatedPlaybackSuspension) Reason() CoordinatedPlaybackSuspensionReason /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCoordinatedPlaybackSuspension */



