// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Scene] class.
type IScene interface {
	objectivec.IObject
	// properties:
	Entities() []Entity /* primitive/slice/pointer. */
	Graphs() foundation.IDictionary /* already interface */
	RootNode() objc.ID
	SetRootNode(value objc.ID)
	// methods:
	AddEntity(entity IGKEntity)
	AddGraphName(graph IGKGraph, name string /* primitive/slice/pointer. */)
	RemoveEntity(entity IGKEntity)
	RemoveGraph(name string /* primitive/slice/pointer. */)
}

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

// Alloc allocates a new instance without initialization.
func (sc _SceneClass) Alloc() Scene {
	rv := objc.Send[Scene](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Loads the specified SpriteKit scene file, creating a object containing the SpriteKit scene and associated GameplayKit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:)
func NewSceneWithFileNamed(filename string /* primitive/slice/pointer. */) Scene {
	rv := objc.Send[Scene](objc.ID(getSceneClass().class), objc.Sel("sceneWithFileNamed:"), objc.String(filename))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:rootNode:)
func NewSceneWithFileNamedRootNode(filename string /* primitive/slice/pointer. */, rootNode objectivec.IObject) Scene {
	rv := objc.Send[Scene](objc.ID(getSceneClass().class), objc.Sel("sceneWithFileNamed:rootNode:"), objc.String(filename), rootNode)
	return rv
}



// Loads the specified SpriteKit scene file, creating a object containing the SpriteKit scene and associated GameplayKit objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:)
func (sc _SceneClass) SceneWithFileNamed(filename string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sceneWithFileNamed:"), objc.String(filename))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/init(fileNamed:rootNode:)
func (sc _SceneClass) SceneWithFileNamedRootNode(filename string /* primitive/slice/pointer. */, rootNode objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sceneWithFileNamed:rootNode:"), objc.String(filename), rootNode)
	return rv
}


// Adds a GameplayKit entity to the list of entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/addEntity(_:)
func (s_ Scene) AddEntity(entity IGKEntity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addEntity:"), entity)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/addGraph(_:name:)
func (s_ Scene) AddGraphName(graph IGKGraph, name string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addGraph:name:"), graph, objc.String(name))
}


// Removes a GameplayKit entity from the list of entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/removeEntity(_:)
func (s_ Scene) RemoveEntity(entity IGKEntity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeEntity:"), entity)
}


// Removes a pathfinding graph from the list of graphs managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/removeGraph(_:)
func (s_ Scene) RemoveGraph(name string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeGraph:"), objc.String(name))
}


// The list of GameplayKit entities managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/entities
func (s_ Scene) Entities() []Entity /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Entity](s_.ID, objc.Sel("entities"))
	return rv
}


// The list of pathfinding graph objects managed by the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/graphs
func (s_ Scene) Graphs() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("graphs"))
	return rv
}


// The SpriteKit scene managed by this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/rootNode
func (s_ Scene) RootNode() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("rootNode"))
	return rv
}


// The SpriteKit scene managed by this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKScene/rootNode
func (s_ Scene) SetRootNode(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRootNode:"), value)
}


