// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKNode] class.
var (
	sKNodeClass     _SKNodeClass
	sKNodeClassOnce sync.Once
)

func getSKNodeClass() _SKNodeClass {
	sKNodeClassOnce.Do(func() {
		sKNodeClass = _SKNodeClass{objc.GetClass("SKNode")}
	})
	return sKNodeClass
}

type _SKNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKNode] class.
type ISKNode interface {
	IUIResponder
	RunAction(action unsafe.Pointer)
	RunActionCompletion(action unsafe.Pointer, block unsafe.Pointer)
	SetValueForAttributeNamed(value unsafe.Pointer, key string)
}

// The base class of all SpriteKit nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode
type SKNode struct {
	UIResponder
}

// SKNodeFrom constructs a [SKNode] from an unsafe.Pointer.
//
// The base class of all SpriteKit nodes.
func SKNodeFrom(ptr unsafe.Pointer) SKNode {
	return SKNode{
		UIResponder: UIResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKNodeClass) Alloc() SKNode {
	rv := objc.Send[SKNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKNodeClass) New() SKNode {
	rv := objc.Send[SKNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKNode) Init() SKNode {
	rv := objc.Send[SKNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKNode) Autorelease() SKNode {
	rv := objc.Send[SKNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKNode creates a new SKNode instance.
func NewSKNode() SKNode {
	return getSKNodeClass().New()
}


// Converts each node into an obstacle by transforming its bounds into the scene’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromNodeBounds:)
func (sc _SKNodeClass) ObstaclesFromNodeBounds(nodes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("obstaclesFromNodeBounds:"), nodes)
	return rv
}
// Converts each node into an obstacle by transforming the node’s physics body shape into the scene’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromNodePhysicsBodies:)
func (sc _SKNodeClass) ObstaclesFromNodePhysicsBodies(nodes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("obstaclesFromNodePhysicsBodies:"), nodes)
	return rv
}
// Turns each node into an obstacle by changing the node’s texture into a physics shape and converting it into the scene’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromSpriteTextures:accuracy:)
func (sc _SKNodeClass) ObstaclesFromSpriteTexturesAccuracy(sprites unsafe.Pointer, accuracy float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("obstaclesFromSpriteTextures:accuracy:"), sprites, accuracy)
	return rv
}
// Adds an action to the list of actions executed by the node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/run(_:)
func (s_ SKNode) RunAction(action unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("runAction:"), action)
}
// Adds an action to the list of actions executed by the node and schedules the argument block to be run upon completion of the action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/run(_:completion:)
func (s_ SKNode) RunActionCompletion(action unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("runAction:completion:"), action, block)
}
// Sets an attribute value for an attached shader [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKNode/setValue(_:forAttribute:)
func (s_ SKNode) SetValueForAttributeNamed(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}


