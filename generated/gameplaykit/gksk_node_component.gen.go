// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKNodeComponent] class.
var (
	SKNodeComponentClass     _SKNodeComponentClass
	SKNodeComponentClassOnce sync.Once
)

func getSKNodeComponentClass() _SKNodeComponentClass {
	SKNodeComponentClassOnce.Do(func() {
		SKNodeComponentClass = _SKNodeComponentClass{objc.GetClass("GKSKNodeComponent")}
	})
	return SKNodeComponentClass
}

type _SKNodeComponentClass struct {
	class objc.Class
}

// An interface definition for the [SKNodeComponent] class.
type ISKNodeComponent interface {
	IComponent
	Node() unsafe.Pointer
	SetNode(value unsafe.Pointer)
}

// A component that manages a SpriteKit node.
//
// Adding a object to an entity automatically updates the property of the component’s SpriteKit node (an object) to point to that entity. When you add entities and components to a node in the Xcode SpriteKit scene editor, Xcode automatically creates a object to manage the relationship between that SpriteKit node and the object that node represents. Load the scene file with the class to access these entities and components. For more information on Entity-Component architecture, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent
type SKNodeComponent struct {
	Component
}

// SKNodeComponentFrom constructs a [SKNodeComponent] from an unsafe.Pointer.
//
// A component that manages a SpriteKit node.
func SKNodeComponentFrom(ptr unsafe.Pointer) SKNodeComponent {
	return SKNodeComponent{
		Component: ComponentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _SKNodeComponentClass) Alloc() SKNodeComponent {
	rv := objc.Send[SKNodeComponent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _SKNodeComponentClass) New() SKNodeComponent {
	rv := objc.Send[SKNodeComponent](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ SKNodeComponent) Init() SKNodeComponent {
	rv := objc.Send[SKNodeComponent](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ SKNodeComponent) Autorelease() SKNodeComponent {
	rv := objc.Send[SKNodeComponent](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKNodeComponent creates a new SKNodeComponent instance.
func NewSKNodeComponent() SKNodeComponent {
	return getSKNodeComponentClass().New()
}




// Initializes a component to manage the specified SpriteKit node.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent/init(node:)
func NewSKNodeComponentWithNode(node unsafe.Pointer) SKNodeComponent {
	instance := getSKNodeComponentClass().Alloc()
	rv := objc.Send[SKNodeComponent](instance.ID, objc.Sel("initWithNode:"), node)
	rv.Autorelease()
	return rv
}


// Creates a component to manage the specified SpriteKit node.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent/componentWithNode:
func (nc _SKNodeComponentClass) ComponentWithNode(node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("componentWithNode:"), node)
	return rv
}

// The SpriteKit node managed by the component.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent/node
func (n_ SKNodeComponent) Node() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("node"))
	return rv
}


// SetNode sets the value of the node property.
// The SpriteKit node managed by the component.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent/node
func (n_ SKNodeComponent) SetNode(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNode:"), value)
}


