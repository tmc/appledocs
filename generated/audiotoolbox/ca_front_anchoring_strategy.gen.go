// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAFrontAnchoringStrategy */


/* debug [class_header]: Header for CAFrontAnchoringStrategy */
// The class instance for the [FrontAnchoringStrategy] class.
var (
	FrontAnchoringStrategyClass     _FrontAnchoringStrategyClass
	FrontAnchoringStrategyClassOnce sync.Once
)

func getFrontAnchoringStrategyClass() _FrontAnchoringStrategyClass {
	FrontAnchoringStrategyClassOnce.Do(func() {
		FrontAnchoringStrategyClass = _FrontAnchoringStrategyClass{objc.GetClass("CAFrontAnchoringStrategy")}
	})
	return FrontAnchoringStrategyClass
}

type _FrontAnchoringStrategyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FrontAnchoringStrategy */
// An interface definition for the [FrontAnchoringStrategy] class.
type IFrontAnchoringStrategy interface {
	IAnchoringStrategy
	
/* debug [class_interface_properties]: Properties for FrontAnchoringStrategy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FrontAnchoringStrategy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FrontAnchoringStrategy */
// Alloc allocates a new instance without initialization.
func (fc _FrontAnchoringStrategyClass) Alloc() FrontAnchoringStrategy {
	rv := objc.Send[FrontAnchoringStrategy](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FrontAnchoringStrategyClass) New() FrontAnchoringStrategy {
	rv := objc.Send[FrontAnchoringStrategy](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FrontAnchoringStrategy) Init() FrontAnchoringStrategy {
	rv := objc.Send[FrontAnchoringStrategy](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FrontAnchoringStrategy) Autorelease() FrontAnchoringStrategy {
	rv := objc.Send[FrontAnchoringStrategy](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFrontAnchoringStrategy creates a new FrontAnchoringStrategy instance.
func NewFrontAnchoringStrategy() FrontAnchoringStrategy {
	return getFrontAnchoringStrategyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FrontAnchoringStrategy */
// Anchor to the front of the user’s space.


// Anchor to the front of the user’s space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAFrontAnchoringStrategy
type FrontAnchoringStrategy struct {
	AnchoringStrategy
}

// FrontAnchoringStrategyFrom constructs a [FrontAnchoringStrategy] from an unsafe.Pointer.
//
// Anchor to the front of the user’s space.
func FrontAnchoringStrategyFrom(ptr unsafe.Pointer) FrontAnchoringStrategy {
	return FrontAnchoringStrategy{
		AnchoringStrategy: AnchoringStrategyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FrontAnchoringStrategy */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FrontAnchoringStrategy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FrontAnchoringStrategy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FrontAnchoringStrategy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FrontAnchoringStrategy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAFrontAnchoringStrategy */


