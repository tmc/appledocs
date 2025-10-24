// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKPath */


/* debug [class_header]: Header for GKPath */
// The class instance for the [Path] class.
var (
	PathClass     _PathClass
	PathClassOnce sync.Once
)

func getPathClass() _PathClass {
	PathClassOnce.Do(func() {
		PathClass = _PathClass{objc.GetClass("GKPath")}
	})
	return PathClass
}

type _PathClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Path */
// An interface definition for the [Path] class.
type IPath interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Path */
	// properties:
	Cyclical() bool
	SetCyclical(value bool)
	NumPoints() uint
	Radius() float32
	SetRadius(value float32)
	IsCyclical() bool
	SetIsCyclical(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Path */
	// methods:
	Float2AtIndex(index uint) objectivec.IObject
	Float3AtIndex(index uint) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Path */
// Alloc allocates a new instance without initialization.
func (pc _PathClass) Alloc() Path {
	rv := objc.Send[Path](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PathClass) New() Path {
	rv := objc.Send[Path](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Path) Init() Path {
	rv := objc.Send[Path](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Path) Autorelease() Path {
	rv := objc.Send[Path](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPath creates a new Path instance.
func NewPath() Path {
	return getPathClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Path */
// A polygonal path that can be followed by an agent.
//
// To make an agent move to or stay within the area defined by a path, create a goal with the method; to make an agent traverse along a path, create a goal with the method. A path can be expressed as a sequence of either 2D points or 3D points. Use the former to create paths for use by objects, and the latter to create paths for objects to follow. To learn more about using goals and agents, see in .


// A polygonal path that can be followed by an agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath
type Path struct {
	objectivec.Object
}

// PathFrom constructs a [Path] from an unsafe.Pointer.
//
// A polygonal path that can be followed by an agent.
func PathFrom(ptr unsafe.Pointer) Path {
	return Path{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Path */

// Initializes a path with the specified array of 3D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/initWithFloat3Points:count:radius:cyclical:
func NewPathWithFloat3PointsCountRadiusCyclical(points objectivec.IObject, count uintptr /* not a class type */, radius float32, cyclical bool) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithFloat3Points:count:radius:cyclical:"), points, count, radius, cyclical)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPathWithFloat3PointsCountRadiusCyclical */


// Initializes a path using the positions of the specified graph nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/init(graphNodes:radius:)
func NewPathWithGraphNodesRadius(graphNodes []GraphNode, radius float32) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithGraphNodes:radius:"), graphNodes, radius)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPathWithGraphNodesRadius */


// Initializes a path with the specified array of 2D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/initWithPoints:count:radius:cyclical:
func NewPathWithPointsCountRadiusCyclical(points objectivec.IObject, count uintptr /* not a class type */, radius float32, cyclical bool) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithPoints:count:radius:cyclical:"), points, count, radius, cyclical)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPathWithPointsCountRadiusCyclical */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Path */

// Creates a path with the specified array of 3D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithFloat3Points:count:radius:cyclical:
func (pc _PathClass) PathWithFloat3PointsCountRadiusCyclical(points objectivec.IObject, count uintptr /* not a class type */, radius float32, cyclical bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pathWithFloat3Points:count:radius:cyclical:"), points, count, radius, cyclical)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PathWithFloat3PointsCountRadiusCyclical) */


// Creates a path using the positions of the specified graph nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithGraphNodes:radius:
func (pc _PathClass) PathWithGraphNodesRadius(graphNodes []GraphNode, radius float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pathWithGraphNodes:radius:"), graphNodes, radius)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PathWithGraphNodesRadius) */


// Creates a path with the specified array of 2D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithPoints:count:radius:cyclical:
func (pc _PathClass) PathWithPointsCountRadiusCyclical(points objectivec.IObject, count uintptr /* not a class type */, radius float32, cyclical bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pathWithPoints:count:radius:cyclical:"), points, count, radius, cyclical)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PathWithPointsCountRadiusCyclical) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Path */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Path */

// Returns the 2D point at the specified index in the path’s list of vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/float2(at:)
func (p_ Path) Float2AtIndex(index uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("float2AtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: Float2AtIndex */


// Returns the 3D point at the specified index in the path’s list of vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/float3(at:)
func (p_ Path) Float3AtIndex(index uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("float3AtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: Float3AtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Path */

// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/isCyclical
func (p_ Path) Cyclical() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cyclical"))
	return rv
}/* debug [instance_properties/getter]: cyclical */


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/isCyclical
func (p_ Path) SetCyclical(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCyclical:"), value)
}/* debug [instance_properties/setter]: cyclical */


// The number of vertices in the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/numPoints
func (p_ Path) NumPoints() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numPoints"))
	return rv
}/* debug [instance_properties/getter]: numPoints */


// The radius of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/radius
func (p_ Path) Radius() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// The radius of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/radius
func (p_ Path) SetRadius(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRadius:"), value)
}/* debug [instance_properties/setter]: radius */


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkpath/iscyclical
func (p_ Path) IsCyclical() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCyclical"))
	return rv
}/* debug [instance_properties/getter]: isCyclical */


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkpath/iscyclical
func (p_ Path) SetIsCyclical(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCyclical:"), value)
}/* debug [instance_properties/setter]: isCyclical */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKPath */


