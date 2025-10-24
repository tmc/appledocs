// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKScene */


/* debug [class_header]: Header for GKScene */
// The class instance for the [Scene] class.
var (
	SceneClass     _SceneClass
	SceneClassOnce sync.Once
)

func getSceneClass() _SceneClass {
	SceneClassOnce.Do(func() {
		SceneClass = _SceneClass{objc.GetClass("GKScene")}
	})
	return SceneClass
}

type _SceneClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Scene */
// An interface definition for the [Scene] class.
type IScene interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Scene */
	// properties:
	Entities() []Entity
	Graphs() foundation.IDictionary
	RootNode() unsafe.Pointer
	SetRootNode(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Scene */
	// methods:
	AddEntity(entity IGKEntity)
	AddGraphName(graph IGKGraph, name objc.IObject /* cross-framework: NSString */)
	RemoveEntity(entity IGKEntity)
	RemoveGraph(name objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Scene */
// Alloc allocates a new instance without initialization.
func (sc _SceneClass) Alloc() Scene {
	rv := objc.Send[Scene](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SceneClass) New() Scene {
	rv := objc.Send[Scene](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scene) Init() Scene {
	rv := objc.Send[Scene](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scene) Autorelease() Scene {
	rv := objc.Send[Scene](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScene creates a new Scene instance.
func NewScene() Scene {
	return getSceneClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Scene */
// A container for associating GameplayKit objects with a SpriteKit scene.
//
// When you create a scene in the Xcode SpriteKit scene editor, Xcode automatically creates a object to manage any GameplayKit objects you add to the scene (entities, components, or pathfinding graphs) and archive them alongside the SpriteKit scene content. To use a SpriteKit scene that contains GameplayKit objects, load the scene file with the method. You can then use the and properties to access the (and associated ) objects and objects in the scene, and the property to access the scene’s SpriteKit content. For more information on Entity-Component architecture and pathfinding graphs, see and in .


// A container for associating GameplayKit objects with a SpriteKit scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene
type Scene struct {
	objectivec.Object
}

// SceneFrom constructs a [Scene] from an unsafe.Pointer.
//
// A container for associating GameplayKit objects with a SpriteKit scene.
func SceneFrom(ptr unsafe.Pointer) Scene {
	return Scene{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Scene */

// Loads the specified SpriteKit scene file, creating a object containing the SpriteKit scene and associated GameplayKit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:)
func NewSceneWithFileNamed(filename objc.IObject /* cross-framework: NSString */) Scene {
	rv := objc.Send[Scene](objc.ID(getSceneClass().class), objc.Sel("sceneWithFileNamed:"), filename)
	return rv
}/* debug [class_init_methods/constructor]: NewSceneWithFileNamed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:rootNode:)
func NewSceneWithFileNamedRootNode(filename objc.IObject /* cross-framework: NSString */, rootNode unsafe.Pointer) Scene {
	rv := objc.Send[Scene](objc.ID(getSceneClass().class), objc.Sel("sceneWithFileNamed:rootNode:"), filename, rootNode)
	return rv
}/* debug [class_init_methods/constructor]: NewSceneWithFileNamedRootNode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Scene */

// Loads the specified SpriteKit scene file, creating a object containing the SpriteKit scene and associated GameplayKit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:)
func (sc _SceneClass) SceneWithFileNamed(filename objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sceneWithFileNamed:"), filename)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SceneWithFileNamed) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:rootNode:)
func (sc _SceneClass) SceneWithFileNamedRootNode(filename objc.IObject /* cross-framework: NSString */, rootNode unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sceneWithFileNamed:rootNode:"), filename, rootNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SceneWithFileNamedRootNode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Scene */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Scene */

// Adds a GameplayKit entity to the list of entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/addEntity(_:)
func (s_ Scene) AddEntity(entity IGKEntity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addEntity:"), entity)
}/* debug [instance_methods/method]: AddEntity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/addGraph(_:name:)
func (s_ Scene) AddGraphName(graph IGKGraph, name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addGraph:name:"), graph, name)
}/* debug [instance_methods/method]: AddGraphName */


// Removes a GameplayKit entity from the list of entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/removeEntity(_:)
func (s_ Scene) RemoveEntity(entity IGKEntity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeEntity:"), entity)
}/* debug [instance_methods/method]: RemoveEntity */


// Removes a pathfinding graph from the list of graphs managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/removeGraph(_:)
func (s_ Scene) RemoveGraph(name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeGraph:"), name)
}/* debug [instance_methods/method]: RemoveGraph */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Scene */

// The list of GameplayKit entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/entities
func (s_ Scene) Entities() []Entity {
	rv := objc.Send[[]Entity](s_.ID, objc.Sel("entities"))
	return rv
}/* debug [instance_properties/getter]: entities */


// The list of pathfinding graph objects managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/graphs
func (s_ Scene) Graphs() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("graphs"))
	return rv
}/* debug [instance_properties/getter]: graphs */


// The SpriteKit scene managed by this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/rootNode
func (s_ Scene) RootNode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("rootNode"))
	return rv
}/* debug [instance_properties/getter]: rootNode */


// The SpriteKit scene managed by this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/rootNode
func (s_ Scene) SetRootNode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRootNode:"), value)
}/* debug [instance_properties/setter]: rootNode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKScene */


