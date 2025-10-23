// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Component] class.
var (
	ComponentClass     _ComponentClass
	ComponentClassOnce sync.Once
)

func getComponentClass() _ComponentClass {
	ComponentClassOnce.Do(func() {
		ComponentClass = _ComponentClass{objc.GetClass("GKComponent")}
	})
	return ComponentClass
}

type _ComponentClass struct {
	class objc.Class
}

// An interface definition for the [Component] class.
type IComponent interface {
	objectivec.IObject
	Entity() IGKEntity
	DidAddToEntity()
	UpdateWithDeltaTime(seconds foundation.TimeInterval)
	WillRemoveFromEntity()
}

// The abstract superclass for creating objects that add specific gameplay functionality to an entity.
//
// is the abstract superclass for custom component classes you create when building a game with Entity-Component architecture. In this architecture, an is an object relevant to the game, and a is an object that handles specific aspects of an entity’s behavior in a general way. Because a component’s scope of functionality is limited, you can reuse the same component class for many different kinds of entities. You create components by subclassing to implement reusable behavior. Then, you build game entities by creating objects and using the method to attach instances of your custom component classes. At runtime, a component-based game needs to dispatch periodic logic—from an update/render loop method such as (SpriteKit) or (SceneKit), or a (iOS) or (macOS) timer in a custom rendering engine—to each of its components. GameplayKit provides two mechanisms for dispatching updates: Per-entity. Call each entity’s method, which will then forward to the method of each component. This option can be quickly implemented in games with a small number of entities and components. Per-component. Use a object to handle all instances of a specific component class. When you call a component system’s method, it forwards to the method of all the component objects it manages. Because a component system needs no knowledge of your game’s entity/component hierarchy, this option works well for games with complex object graphs. For more information on Entity-Component architecture, read in .


// The abstract superclass for creating objects that add specific gameplay functionality to an entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponent
type Component struct {
	objectivec.Object
}

// ComponentFrom constructs a [Component] from an unsafe.Pointer.
//
// The abstract superclass for creating objects that add specific gameplay functionality to an entity.
func ComponentFrom(ptr unsafe.Pointer) Component {
	return Component{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComponentClass) Alloc() Component {
	rv := objc.Send[Component](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComponentClass) New() Component {
	rv := objc.Send[Component](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Component) Init() Component {
	rv := objc.Send[Component](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Component) Autorelease() Component {
	rv := objc.Send[Component](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComponent creates a new Component instance.
func NewComponent() Component {
	return getComponentClass().New()
}



// Notifies the component that it has been assigned to an entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponent/didAddToEntity()
func (c_ Component) DidAddToEntity() {
	objc.Send[objc.ID](c_.ID, objc.Sel("didAddToEntity"))
}


// Performs any custom periodic actions defined by the component subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponent/update(deltaTime:)
func (c_ Component) UpdateWithDeltaTime(seconds foundation.TimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}


// Notifies the component that it has been removed from an entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponent/willRemoveFromEntity()
func (c_ Component) WillRemoveFromEntity() {
	objc.Send[objc.ID](c_.ID, objc.Sel("willRemoveFromEntity"))
}


// The entity that owns this component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKComponent/entity
func (c_ Component) Entity() IGKEntity {
	rv := objc.Send[Entity](c_.ID, objc.Sel("entity"))
	return rv
}



