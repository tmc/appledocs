// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CASceneAnchoringStrategy */


/* debug [class_header]: Header for CASceneAnchoringStrategy */
// The class instance for the [SceneAnchoringStrategy] class.
var (
	SceneAnchoringStrategyClass     _SceneAnchoringStrategyClass
	SceneAnchoringStrategyClassOnce sync.Once
)

func getSceneAnchoringStrategyClass() _SceneAnchoringStrategyClass {
	SceneAnchoringStrategyClassOnce.Do(func() {
		SceneAnchoringStrategyClass = _SceneAnchoringStrategyClass{objc.GetClass("CASceneAnchoringStrategy")}
	})
	return SceneAnchoringStrategyClass
}

type _SceneAnchoringStrategyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SceneAnchoringStrategy */
// An interface definition for the [SceneAnchoringStrategy] class.
type ISceneAnchoringStrategy interface {
	IAnchoringStrategy
	
/* debug [class_interface_properties]: Properties for SceneAnchoringStrategy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SceneAnchoringStrategy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SceneAnchoringStrategy */
// Alloc allocates a new instance without initialization.
func (sc _SceneAnchoringStrategyClass) Alloc() SceneAnchoringStrategy {
	rv := objc.Send[SceneAnchoringStrategy](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SceneAnchoringStrategyClass) New() SceneAnchoringStrategy {
	rv := objc.Send[SceneAnchoringStrategy](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SceneAnchoringStrategy) Init() SceneAnchoringStrategy {
	rv := objc.Send[SceneAnchoringStrategy](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SceneAnchoringStrategy) Autorelease() SceneAnchoringStrategy {
	rv := objc.Send[SceneAnchoringStrategy](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSceneAnchoringStrategy creates a new SceneAnchoringStrategy instance.
func NewSceneAnchoringStrategy() SceneAnchoringStrategy {
	return getSceneAnchoringStrategyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SceneAnchoringStrategy */
// Anchor to the visual center of a particular UIScene.


// Anchor to the visual center of a particular UIScene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASceneAnchoringStrategy
type SceneAnchoringStrategy struct {
	AnchoringStrategy
}

// SceneAnchoringStrategyFrom constructs a [SceneAnchoringStrategy] from an unsafe.Pointer.
//
// Anchor to the visual center of a particular UIScene.
func SceneAnchoringStrategyFrom(ptr unsafe.Pointer) SceneAnchoringStrategy {
	return SceneAnchoringStrategy{
		AnchoringStrategy: AnchoringStrategyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SceneAnchoringStrategy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASceneAnchoringStrategy/initWithSceneIdentifier:
func NewSceneAnchoringStrategyWithSceneIdentifier(sceneIdentifier objc.IObject /* cross-framework: NSString */) SceneAnchoringStrategy {
	instance := getSceneAnchoringStrategyClass().Alloc()
	rv := objc.Send[SceneAnchoringStrategy](instance.ID, objc.Sel("initWithSceneIdentifier:"), sceneIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSceneAnchoringStrategyWithSceneIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SceneAnchoringStrategy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SceneAnchoringStrategy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SceneAnchoringStrategy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SceneAnchoringStrategy */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CASceneAnchoringStrategy */


