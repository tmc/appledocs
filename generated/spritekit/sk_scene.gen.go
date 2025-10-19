// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKScene] class.
var (
	sKSceneClass     _SKSceneClass
	sKSceneClassOnce sync.Once
)

func getSKSceneClass() _SKSceneClass {
	sKSceneClassOnce.Do(func() {
		sKSceneClass = _SKSceneClass{objc.GetClass("SKScene")}
	})
	return sKSceneClass
}

type _SKSceneClass struct {
	class objc.Class
}

// An interface definition for the [SKScene] class.
type ISKScene interface {
	ISKEffectNode
	ConvertPointFromView(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToView(point unsafe.Pointer) unsafe.Pointer
	DidApplyConstraints()
	DidChangeSize(oldSize unsafe.Pointer)
	DidEvaluateActions()
	DidFinishUpdate()
	DidMoveToView(view unsafe.Pointer)
	DidSimulatePhysics()
	SceneDidLoad()
	Update(currentTime TimeInterval)
	WillMoveFromView(view unsafe.Pointer)
}

// An object that organizes all of the active SpriteKit content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene
type SKScene struct {
	SKEffectNode
}

// SKSceneFrom constructs a [SKScene] from an unsafe.Pointer.
//
// An object that organizes all of the active SpriteKit content.
func SKSceneFrom(ptr unsafe.Pointer) SKScene {
	return SKScene{
		SKEffectNode: SKEffectNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKSceneClass) Alloc() SKScene {
	rv := objc.Send[SKScene](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKSceneClass) New() SKScene {
	rv := objc.Send[SKScene](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKScene) Init() SKScene {
	rv := objc.Send[SKScene](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKScene) Autorelease() SKScene {
	rv := objc.Send[SKScene](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKScene creates a new SKScene instance.
func NewSKScene() SKScene {
	return getSKSceneClass().New()
}


// Initializes a new scene object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/init(size:)
func NewSKSceneWithSize(size unsafe.Pointer) SKScene {
	instance := getSKSceneClass().Alloc()
	rv := objc.Send[SKScene](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}


// Creates and returns a new scene object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/sceneWithSize:
func (sc _SKSceneClass) SceneWithSize(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sceneWithSize:"), size)
	return rv
}
// Converts a point from view coordinates to scene coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/convertPoint(fromView:)
func (s_ SKScene) ConvertPointFromView(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("convertPointFromView:"), point)
	return rv
}
// Converts a point from scene coordinates to view coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/convertPoint(toView:)
func (s_ SKScene) ConvertPointToView(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("convertPointToView:"), point)
	return rv
}
// Tells your app to peform any necessary logic after constraints are applied. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didApplyConstraints()
func (s_ SKScene) DidApplyConstraints() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didApplyConstraints"))
}
// Tells you when the scene’s size has changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didChangeSize(_:)
func (s_ SKScene) DidChangeSize(oldSize unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("didChangeSize:"), oldSize)
}
// Tells your app to peform any necessary logic after scene actions are evaluated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didEvaluateActions()
func (s_ SKScene) DidEvaluateActions() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didEvaluateActions"))
}
// Tells your app to peform any necessary logic after the scene has finished all of the steps required to process animations. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didFinishUpdate()
func (s_ SKScene) DidFinishUpdate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didFinishUpdate"))
}
// Tells you when the scene is presented by a view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didMove(to:)
func (s_ SKScene) DidMoveToView(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("didMoveToView:"), view)
}
// Tells your app to peform any necessary logic after physics simulations are performed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/didSimulatePhysics()
func (s_ SKScene) DidSimulatePhysics() {
	objc.Send[objc.ID](s_.ID, objc.Sel("didSimulatePhysics"))
}
// Tells you when the scene is presented. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/sceneDidLoad()
func (s_ SKScene) SceneDidLoad() {
	objc.Send[objc.ID](s_.ID, objc.Sel("sceneDidLoad"))
}
// Tells your app to perform any app-specific logic to update your scene. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/update(_:)
func (s_ SKScene) Update(currentTime TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("update:"), currentTime)
}
// Tells you when the scene is about to be removed from a view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKScene/willMove(from:)
func (s_ SKScene) WillMoveFromView(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("willMoveFromView:"), view)
}

