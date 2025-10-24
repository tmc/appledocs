// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKEntity */


/* debug [class_header]: Header for GKEntity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Entity */
// An interface definition for the [Entity] class.
type IEntity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Entity */
	// properties:
	Components() []Component
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Entity */
	// methods:
	AddComponent(component IGKComponent)
	ComponentForClass(componentClass objc.Class) IComponent
	RemoveComponentForClass(componentClass objc.Class)
	UpdateWithDeltaTime(seconds float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Entity */
// Alloc allocates a new instance without initialization.
func (ec _EntityClass) Alloc() Entity {
	rv := objc.Send[Entity](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Entity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Entity */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Entity */

// Creates a new entity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/entity
func (ec _EntityClass) Entity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("entity"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Entity) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Entity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Entity */

// Adds a component to the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/addComponent(_:)
func (e_ Entity) AddComponent(component IGKComponent) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addComponent:"), component)
}/* debug [instance_methods/method]: AddComponent */


// Returns the entity’s component for the specified component class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/componentForClass:
func (e_ Entity) ComponentForClass(componentClass objc.Class) IComponent {
	rv := objc.Send[Component](e_.ID, objc.Sel("componentForClass:"), componentClass)
	return rv
}/* debug [instance_methods/method]: ComponentForClass */


// Removes the component of the specified class from the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/removeComponentForClass:
func (e_ Entity) RemoveComponentForClass(componentClass objc.Class) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeComponentForClass:"), componentClass)
}/* debug [instance_methods/method]: RemoveComponentForClass */


// Performs periodic updates for each of the entity’s components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/update(deltaTime:)
func (e_ Entity) UpdateWithDeltaTime(seconds float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}/* debug [instance_methods/method]: UpdateWithDeltaTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Entity */

// The entity’s list of components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKEntity/components
func (e_ Entity) Components() []Component {
	rv := objc.Send[[]Component](e_.ID, objc.Sel("components"))
	return rv
}/* debug [instance_properties/getter]: components */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKEntity */


