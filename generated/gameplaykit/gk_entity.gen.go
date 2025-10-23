// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Entity] class.
var (
	EntityClass     _EntityClass
	EntityClassOnce sync.Once
)

func getEntityClass() _EntityClass {
	EntityClassOnce.Do(func() {
		EntityClass = _EntityClass{objc.GetClass("GKEntity")}
	})
	return EntityClass
}

type _EntityClass struct {
	class objc.Class
}

// An interface definition for the [Entity] class.
type IEntity interface {
	objectivec.IObject
	// properties:
	Components() []Component /* primitive/slice/pointer. */
	// methods:
	AddComponent(component IGKComponent)
	ComponentForClass(componentClass objc.Class) IComponent
	RemoveComponentForClass(componentClass objc.Class)
	UpdateWithDeltaTime(seconds foundation.TimeInterval /* not a class type */)
}

// An object relevant to gameplay, with functionality entirely provided by a collection of component objects.
//
// A object represents an entity in games with Entity-Component architecture. In this design, an is a general type for objects relevant to the game. Entities typically define no functionality of their own—instead, you define an entity’s features through composition, by adding that each handle specific aspects of an entity’s behavior in a general way. Because components ( subclasses) are general and reusable, you can add many kinds of entities to a game by combining components in different ways, without needing to design new entity classes. For more information on Entity-Component architecture, read in .


// An object relevant to gameplay, with functionality entirely provided by a collection of component objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity
type Entity struct {
	objectivec.Object
}

// EntityFrom constructs a [Entity] from an unsafe.Pointer.
//
// An object relevant to gameplay, with functionality entirely provided by a collection of component objects.
func EntityFrom(ptr unsafe.Pointer) Entity {
	return Entity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EntityClass) Alloc() Entity {
	rv := objc.Send[Entity](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EntityClass) New() Entity {
	rv := objc.Send[Entity](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Entity) Init() Entity {
	rv := objc.Send[Entity](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Entity) Autorelease() Entity {
	rv := objc.Send[Entity](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEntity creates a new Entity instance.
func NewEntity() Entity {
	return getEntityClass().New()
}




// Creates a new entity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/entity
func (ec _EntityClass) Entity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("entity"))
	return rv
}


// Adds a component to the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/addComponent(_:)
func (e_ Entity) AddComponent(component IGKComponent) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addComponent:"), component)
}


// Returns the entity’s component for the specified component class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/componentForClass:
func (e_ Entity) ComponentForClass(componentClass objc.Class) IComponent {
	rv := objc.Send[Component](e_.ID, objc.Sel("componentForClass:"), componentClass)
	return rv
}


// Removes the component of the specified class from the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/removeComponentForClass:
func (e_ Entity) RemoveComponentForClass(componentClass objc.Class) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeComponentForClass:"), componentClass)
}


// Performs periodic updates for each of the entity’s components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/update(deltaTime:)
func (e_ Entity) UpdateWithDeltaTime(seconds foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}


// The entity’s list of components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/components
func (e_ Entity) Components() []Component /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Component](e_.ID, objc.Sel("components"))
	return rv
}


