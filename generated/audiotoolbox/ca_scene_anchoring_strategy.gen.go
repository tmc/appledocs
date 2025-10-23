// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [SceneAnchoringStrategy] class.
type ISceneAnchoringStrategy interface {
	IAnchoringStrategy
	SceneIdentifier() string
}

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

// Alloc allocates a new instance without initialization.
func (sc _SceneAnchoringStrategyClass) Alloc() SceneAnchoringStrategy {
	rv := objc.Send[SceneAnchoringStrategy](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASceneAnchoringStrategy/initWithSceneIdentifier:
func NewSceneAnchoringStrategyWithSceneIdentifier(sceneIdentifier string) SceneAnchoringStrategy {
	instance := getSceneAnchoringStrategyClass().Alloc()
	rv := objc.Send[SceneAnchoringStrategy](instance.ID, objc.Sel("initWithSceneIdentifier:"), objc.String(sceneIdentifier))
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASceneAnchoringStrategy/sceneIdentifier
func (s_ SceneAnchoringStrategy) SceneIdentifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("sceneIdentifier"))
	return rv
}


