// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEObject] class.
var (
	PHASEObjectClass     _PHASEObjectClass
	PHASEObjectClassOnce sync.Once
)

func getPHASEObjectClass() _PHASEObjectClass {
	PHASEObjectClassOnce.Do(func() {
		PHASEObjectClass = _PHASEObjectClass{objc.GetClass("PHASEObject")}
	})
	return PHASEObjectClass
}

type _PHASEObjectClass struct {
	class objc.Class
}

// An interface definition for the [PHASEObject] class.
type IPHASEObject interface {
	objectivec.IObject
	// properties:
	Children() []IPHASEObject
	Parent() IPHASEObject
	Transform() unsafe.Pointer
	SetTransform(value unsafe.Pointer)
	WorldTransform() unsafe.Pointer
	SetWorldTransform(value unsafe.Pointer)
	// methods:
	AddChildError(child IPHASEObject, error_ unsafe.Pointer) bool
	RemoveChild(child IPHASEObject)
	RemoveChildren()
}

// An object in the scene.
//
// This class models a member of your app’s scene by defining a 3D position and orientation. The following subclasses derive from this class: The array holds instances of this class to position and orient them relatively.


// An object in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject
type PHASEObject struct {
	objectivec.Object
}

// PHASEObjectFrom constructs a [PHASEObject] from an unsafe.Pointer.
//
// An object in the scene.
func PHASEObjectFrom(ptr unsafe.Pointer) PHASEObject {
	return PHASEObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEObjectClass) Alloc() PHASEObject {
	rv := objc.Send[PHASEObject](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEObjectClass) New() PHASEObject {
	rv := objc.Send[PHASEObject](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEObject) Init() PHASEObject {
	rv := objc.Send[PHASEObject](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEObject) Autorelease() PHASEObject {
	rv := objc.Send[PHASEObject](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEObject creates a new PHASEObject instance.
func NewPHASEObject() PHASEObject {
	return getPHASEObjectClass().New()
}



// Creates an object in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/init(engine:)
func NewPHASEObjectWithEngine(engine IPHASEEngine) PHASEObject {
	instance := getPHASEObjectClass().Alloc()
	rv := objc.Send[PHASEObject](instance.ID, objc.Sel("initWithEngine:"), engine)
	rv.Autorelease()
	return rv
}



// A vector that points forward in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/forward
func (pc _PHASEObjectClass) Forward() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("forward"))
	return rv
}

// A vector that points right in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/right
func (pc _PHASEObjectClass) Right() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("right"))
	return rv
}

// A vector that points up in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/up
func (pc _PHASEObjectClass) Up() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("up"))
	return rv
}

// Adds the given object as a child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/addChild(_:)
func (p_ PHASEObject) AddChildError(child IPHASEObject, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("addChild:error:"), child, error_)
	return rv
}


// Removes the given object as a child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/removeChild(_:)
func (p_ PHASEObject) RemoveChild(child IPHASEObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeChild:"), child)
}


// Removes all child objects from the given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/removeChildren()
func (p_ PHASEObject) RemoveChildren() {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeChildren"))
}


// Objects that position and orient in the scene relative to the given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/children
func (p_ PHASEObject) Children() []IPHASEObject {
	rv := objc.Send[[]PHASEObject](p_.ID, objc.Sel("children"))
	return rv
}


// A vector that points forward in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/forward
func (p_ PHASEObject) Forward() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("forward"))
	return rv
}


// The object that this instance positions and orients relative to in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/parent
func (p_ PHASEObject) Parent() IPHASEObject {
	rv := objc.Send[PHASEObject](p_.ID, objc.Sel("parent"))
	return rv
}


// A vector that points right in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/right
func (p_ PHASEObject) Right() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("right"))
	return rv
}


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/transform
func (p_ PHASEObject) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transform"))
	return rv
}


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/transform
func (p_ PHASEObject) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransform:"), value)
}


// A vector that points up in the local coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/up
func (p_ PHASEObject) Up() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("up"))
	return rv
}


// A matrix, in scene coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/worldTransform
func (p_ PHASEObject) WorldTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("worldTransform"))
	return rv
}


// A matrix, in scene coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEObject/worldTransform
func (p_ PHASEObject) SetWorldTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWorldTransform:"), value)
}


