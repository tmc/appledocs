// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKFieldNode] class.
var (
	sKFieldNodeClass     _SKFieldNodeClass
	sKFieldNodeClassOnce sync.Once
)

func getSKFieldNodeClass() _SKFieldNodeClass {
	sKFieldNodeClassOnce.Do(func() {
		sKFieldNodeClass = _SKFieldNodeClass{objc.GetClass("SKFieldNode")}
	})
	return sKFieldNodeClass
}

type _SKFieldNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKFieldNode] class.
type ISKFieldNode interface {
	ISKNode
}

// A node that applies physics effects to nearby nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode
type SKFieldNode struct {
	SKNode
}

// SKFieldNodeFrom constructs a [SKFieldNode] from an unsafe.Pointer.
//
// A node that applies physics effects to nearby nodes.
func SKFieldNodeFrom(ptr unsafe.Pointer) SKFieldNode {
	return SKFieldNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKFieldNodeClass) Alloc() SKFieldNode {
	rv := objc.Send[SKFieldNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKFieldNodeClass) New() SKFieldNode {
	rv := objc.Send[SKFieldNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKFieldNode) Init() SKFieldNode {
	rv := objc.Send[SKFieldNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKFieldNode) Autorelease() SKFieldNode {
	rv := objc.Send[SKFieldNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKFieldNode creates a new SKFieldNode instance.
func NewSKFieldNode() SKFieldNode {
	return getSKFieldNodeClass().New()
}


// Creates a field node that calculates and applies a custom force to the physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/customField(evaluationBlock:)
func (sc _SKFieldNodeClass) CustomFieldWithEvaluationBlock(block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("customFieldWithEvaluationBlock:"), block)
	return rv
}
// Creates a field node that applies a force that resists the motion of physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/dragField()
func (sc _SKFieldNodeClass) DragField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("dragField"))
	return rv
}
// Creates a field node that applies an electrical force proportional to the electrical charge of physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/electricField()
func (sc _SKFieldNodeClass) ElectricField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("electricField"))
	return rv
}
// Creates a field node that accelerates physics bodies in a specific direction. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/linearGravityField(withVector:)
func (sc _SKFieldNodeClass) LinearGravityFieldWithVector(direction unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("linearGravityFieldWithVector:"), direction)
	return rv
}
// Creates a field node that applies a magnetic force based on the velocity and electrical charge of the physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/magneticField()
func (sc _SKFieldNodeClass) MagneticField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("magneticField"))
	return rv
}
// Creates a field node that applies a randomized acceleration to physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/noiseField(withSmoothness:animationSpeed:)
func (sc _SKFieldNodeClass) NoiseFieldWithSmoothnessAnimationSpeed(smoothness float64, speed float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("noiseFieldWithSmoothness:animationSpeed:"), smoothness, speed)
	return rv
}
// Creates a field node that accelerates physics bodies toward the field node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/radialGravityField()
func (sc _SKFieldNodeClass) RadialGravityField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("radialGravityField"))
	return rv
}
// Creates a field node that applies a spring-like force that pulls physics bodies toward the field node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/springField()
func (sc _SKFieldNodeClass) SpringField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("springField"))
	return rv
}
// Creates a field node that applies a randomized acceleration to physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/turbulenceField(withSmoothness:animationSpeed:)
func (sc _SKFieldNodeClass) TurbulenceFieldWithSmoothnessAnimationSpeed(smoothness float64, speed float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("turbulenceFieldWithSmoothness:animationSpeed:"), smoothness, speed)
	return rv
}
// Creates a field node that sets the velocity of physics bodies that enter the node’s area based on the pixel values of a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/velocityField(with:)
func (sc _SKFieldNodeClass) VelocityFieldWithTexture(velocityTexture unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("velocityFieldWithTexture:"), velocityTexture)
	return rv
}
// Creates a field node that gives physics bodies a constant velocity. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/velocityField(withVector:)
func (sc _SKFieldNodeClass) VelocityFieldWithVector(direction unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("velocityFieldWithVector:"), direction)
	return rv
}
// Creates a field node that applies a perpendicular force to physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/vortexField()
func (sc _SKFieldNodeClass) VortexField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("vortexField"))
	return rv
}


