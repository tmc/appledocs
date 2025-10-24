// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAAutomaticAnchoringStrategy */


/* debug [class_header]: Header for CAAutomaticAnchoringStrategy */
// The class instance for the [AutomaticAnchoringStrategy] class.
var (
	AutomaticAnchoringStrategyClass     _AutomaticAnchoringStrategyClass
	AutomaticAnchoringStrategyClassOnce sync.Once
)

func getAutomaticAnchoringStrategyClass() _AutomaticAnchoringStrategyClass {
	AutomaticAnchoringStrategyClassOnce.Do(func() {
		AutomaticAnchoringStrategyClass = _AutomaticAnchoringStrategyClass{objc.GetClass("CAAutomaticAnchoringStrategy")}
	})
	return AutomaticAnchoringStrategyClass
}

type _AutomaticAnchoringStrategyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AutomaticAnchoringStrategy */
// An interface definition for the [AutomaticAnchoringStrategy] class.
type IAutomaticAnchoringStrategy interface {
	IAnchoringStrategy
	
/* debug [class_interface_properties]: Properties for AutomaticAnchoringStrategy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AutomaticAnchoringStrategy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AutomaticAnchoringStrategy */
// Alloc allocates a new instance without initialization.
func (ac _AutomaticAnchoringStrategyClass) Alloc() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AutomaticAnchoringStrategyClass) New() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AutomaticAnchoringStrategy) Init() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AutomaticAnchoringStrategy) Autorelease() AutomaticAnchoringStrategy {
	rv := objc.Send[AutomaticAnchoringStrategy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAutomaticAnchoringStrategy creates a new AutomaticAnchoringStrategy instance.
func NewAutomaticAnchoringStrategy() AutomaticAnchoringStrategy {
	return getAutomaticAnchoringStrategyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AutomaticAnchoringStrategy */
// A system-defined anchoring strategy.


// A system-defined anchoring strategy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CAAutomaticAnchoringStrategy
type AutomaticAnchoringStrategy struct {
	AnchoringStrategy
}

// AutomaticAnchoringStrategyFrom constructs a [AutomaticAnchoringStrategy] from an unsafe.Pointer.
//
// A system-defined anchoring strategy.
func AutomaticAnchoringStrategyFrom(ptr unsafe.Pointer) AutomaticAnchoringStrategy {
	return AutomaticAnchoringStrategy{
		AnchoringStrategy: AnchoringStrategyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AutomaticAnchoringStrategy */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AutomaticAnchoringStrategy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AutomaticAnchoringStrategy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AutomaticAnchoringStrategy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AutomaticAnchoringStrategy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAAutomaticAnchoringStrategy */


