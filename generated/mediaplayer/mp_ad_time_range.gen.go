// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPAdTimeRange */


/* debug [class_header]: Header for MPAdTimeRange */
// The class instance for the [AdTimeRange] class.
var (
	AdTimeRangeClass     _AdTimeRangeClass
	AdTimeRangeClassOnce sync.Once
)

func getAdTimeRangeClass() _AdTimeRangeClass {
	AdTimeRangeClassOnce.Do(func() {
		AdTimeRangeClass = _AdTimeRangeClass{objc.GetClass("MPAdTimeRange")}
	})
	return AdTimeRangeClass
}

type _AdTimeRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AdTimeRange */
// An interface definition for the [AdTimeRange] class.
type IAdTimeRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AdTimeRange */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AdTimeRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AdTimeRange */
// Alloc allocates a new instance without initialization.
func (ac _AdTimeRangeClass) Alloc() AdTimeRange {
	rv := objc.Send[AdTimeRange](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdTimeRangeClass) New() AdTimeRange {
	rv := objc.Send[AdTimeRange](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTimeRange) Init() AdTimeRange {
	rv := objc.Send[AdTimeRange](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTimeRange) Autorelease() AdTimeRange {
	rv := objc.Send[AdTimeRange](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTimeRange creates a new AdTimeRange instance.
func NewAdTimeRange() AdTimeRange {
	return getAdTimeRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AdTimeRange */
// An object that represents a time range where an ad break exists in the current player.
//
// This value must be in bounds of the duration of the current player item.


// An object that represents a time range where an ad break exists in the current player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPAdTimeRange
type AdTimeRange struct {
	objectivec.Object
}

// AdTimeRangeFrom constructs a [AdTimeRange] from an unsafe.Pointer.
//
// An object that represents a time range where an ad break exists in the current player.
func AdTimeRangeFrom(ptr unsafe.Pointer) AdTimeRange {
	return AdTimeRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AdTimeRange */

// Creates a Media Player time range that indicates where an ad break exists in the current player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPAdTimeRange/init(_:)
func NewAdTimeRangeWithTimeRange(timeRange TimeRange /* not a class type */) AdTimeRange {
	instance := getAdTimeRangeClass().Alloc()
	rv := objc.Send[AdTimeRange](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAdTimeRangeWithTimeRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AdTimeRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AdTimeRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AdTimeRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AdTimeRange */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPAdTimeRange */


