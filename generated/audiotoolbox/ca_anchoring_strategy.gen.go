// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAAnchoringStrategy */


/* debug [class_header]: Header for CAAnchoringStrategy */
// The class instance for the [AnchoringStrategy] class.
var (
	AnchoringStrategyClass     _AnchoringStrategyClass
	AnchoringStrategyClassOnce sync.Once
)

func getAnchoringStrategyClass() _AnchoringStrategyClass {
	AnchoringStrategyClassOnce.Do(func() {
		AnchoringStrategyClass = _AnchoringStrategyClass{objc.GetClass("CAAnchoringStrategy")}
	})
	return AnchoringStrategyClass
}

type _AnchoringStrategyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AnchoringStrategy */
// An interface definition for the [AnchoringStrategy] class.
type IAnchoringStrategy interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AnchoringStrategy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AnchoringStrategy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AnchoringStrategy */
// Alloc allocates a new instance without initialization.
func (ac _AnchoringStrategyClass) Alloc() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AnchoringStrategyClass) New() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnchoringStrategy) Init() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnchoringStrategy) Autorelease() AnchoringStrategy {
	rv := objc.Send[AnchoringStrategy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnchoringStrategy creates a new AnchoringStrategy instance.
func NewAnchoringStrategy() AnchoringStrategy {
	return getAnchoringStrategyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AnchoringStrategy */
// The center of a head-tracked spatial experience.
//
// The Objective-C version of the Swift type.


// The center of a head-tracked spatial experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAnchoringStrategy
type AnchoringStrategy struct {
	objectivec.Object
}

// AnchoringStrategyFrom constructs a [AnchoringStrategy] from an unsafe.Pointer.
//
// The center of a head-tracked spatial experience.
func AnchoringStrategyFrom(ptr unsafe.Pointer) AnchoringStrategy {
	return AnchoringStrategy{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AnchoringStrategy *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AnchoringStrategy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AnchoringStrategy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AnchoringStrategy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AnchoringStrategy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAAnchoringStrategy */



