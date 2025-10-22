// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComponentSystem] class.
var (
	ComponentSystemClass     _ComponentSystemClass
	ComponentSystemClassOnce sync.Once
)

func getComponentSystemClass() _ComponentSystemClass {
	ComponentSystemClassOnce.Do(func() {
		ComponentSystemClass = _ComponentSystemClass{objc.GetClass("GKComponentSystem")}
	})
	return ComponentSystemClass
}

type _ComponentSystemClass struct {
	class objc.Class
}

// An interface definition for the [ComponentSystem] class.
type IComponentSystem interface {
	objectivec.IObject
	AddComponent(component unsafe.Pointer)
	AddComponentWithEntity(entity IGKEntity)
	ClassForGenericArgumentAtIndex(index uint) objc.Class
	RemoveComponent(component unsafe.Pointer)
	RemoveComponentWithEntity(entity IGKEntity)
	ObjectAtIndexedSubscript(idx uint) unsafe.Pointer
	UpdateWithDeltaTime(seconds foundation.ITimeInterval)
	ComponentClass() objc.Class
	Components() []Component
}

// Manages periodic update messages for all component objects of a specified class.
//
// A object manages periodic update messages for components in a game that uses Entity-Component architecture. Use a component system to perform per-frame logic for all components of a specific class without traversing your game’s object hierarchy to dispatch update messages. Each object manages components of a specific subclass. You create a component system with the initializer, specifying the component class it will work with. Then, you register the components used by the entities in your game with the or methods. The component system will then forward any component-specific messages it receives to all registered instances of its component class. The most important of the component-specific messages is the method. Call this method from your game’s update/render loop—that is, from a method such as (SpriteKit) or (SceneKit), or from a (iOS) or (macOS) timer in a custom rendering engine. The component system then forwards to the method of all the subclass instances it manages, allowing those objects to perform per-frame update logic. For more information on Entity-Component architecture, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem
type ComponentSystem struct {
	objectivec.Object
}

// ComponentSystemFrom constructs a [ComponentSystem] from an unsafe.Pointer.
//
// Manages periodic update messages for all component objects of a specified class.
func ComponentSystemFrom(ptr unsafe.Pointer) ComponentSystem {
	return ComponentSystem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComponentSystemClass) Alloc() ComponentSystem {
	rv := objc.Send[ComponentSystem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComponentSystemClass) New() ComponentSystem {
	rv := objc.Send[ComponentSystem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComponentSystem) Init() ComponentSystem {
	rv := objc.Send[ComponentSystem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComponentSystem) Autorelease() ComponentSystem {
	rv := objc.Send[ComponentSystem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComponentSystem creates a new ComponentSystem instance.
func NewComponentSystem() ComponentSystem {
	return getComponentSystemClass().New()
}




// Initializes a component system to manage components of the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/init(componentClass:)
func NewComponentSystemWithComponentClass(cls objc.Class) ComponentSystem {
	instance := getComponentSystemClass().Alloc()
	rv := objc.Send[ComponentSystem](instance.ID, objc.Sel("initWithComponentClass:"), cls)
	rv.Autorelease()
	return rv
}


// Adds a component instance to the component system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/addComponent(_:)
func (c_ ComponentSystem) AddComponent(component unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addComponent:"), component)
}

// Adds any instances of the component system’s component class in the specified entity to the component system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/addComponent(foundIn:)
func (c_ ComponentSystem) AddComponentWithEntity(entity IGKEntity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addComponentWithEntity:"), entity)
}

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/classForGenericArgument(at:)
func (c_ ComponentSystem) ClassForGenericArgumentAtIndex(index uint) objc.Class {
	rv := objc.Send[objc.Class](c_.ID, objc.Sel("classForGenericArgumentAtIndex:"), index)
	return rv
}

// Removes the specified component instance from the component system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/removeComponent(_:)
func (c_ ComponentSystem) RemoveComponent(component unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeComponent:"), component)
}

// Removes any instances of the component system’s component class in the specified entity from the component system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/removeComponent(foundIn:)
func (c_ ComponentSystem) RemoveComponentWithEntity(entity IGKEntity) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeComponentWithEntity:"), entity)
}

// Returns the component at the specified index in the system’s list of components.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/subscript(_:)
func (c_ ComponentSystem) ObjectAtIndexedSubscript(idx uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Tells all component instances managed by the system to perform their custom periodic actions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/update(deltaTime:)
func (c_ ComponentSystem) UpdateWithDeltaTime(seconds foundation.ITimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}

// The class of components managed by the component system.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/componentClass
func (c_ ComponentSystem) ComponentClass() objc.Class {
	rv := objc.Send[objc.Class](c_.ID, objc.Sel("componentClass"))
	return rv
}

// The component system’s list of components.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponentSystem/components
func (c_ ComponentSystem) Components() []Component {
	rv := objc.Send[[]Component](c_.ID, objc.Sel("components"))
	return rv
}


