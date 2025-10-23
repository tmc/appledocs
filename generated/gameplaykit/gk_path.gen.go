// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Path] class.
type IPath interface {
	objectivec.IObject
	Cyclical() bool
	SetCyclical(value bool)
	NumPoints() uint
	Radius() float32
	SetRadius(value float32)
	IsCyclical() bool
	SetIsCyclical(value bool)
	Float2AtIndex(index uint) unsafe.Pointer
	Float3AtIndex(index uint) unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (pc _PathClass) Alloc() Path {
	rv := objc.Send[Path](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a path with the specified array of 3D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/initWithFloat3Points:count:radius:cyclical:
func NewPathWithFloat3PointsCountRadiusCyclical(points unsafe.Pointer, count uintptr, radius float32, cyclical bool) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithFloat3Points:count:radius:cyclical:"), points, count, radius, cyclical)
	rv.Autorelease()
	return rv
}


// Initializes a path using the positions of the specified graph nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/init(graphNodes:radius:)
func NewPathWithGraphNodesRadius(graphNodes []GraphNode, radius float32) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithGraphNodes:radius:"), graphNodes, radius)
	rv.Autorelease()
	return rv
}


// Initializes a path with the specified array of 2D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/initWithPoints:count:radius:cyclical:
func NewPathWithPointsCountRadiusCyclical(points unsafe.Pointer, count uintptr, radius float32, cyclical bool) Path {
	instance := getPathClass().Alloc()
	rv := objc.Send[Path](instance.ID, objc.Sel("initWithPoints:count:radius:cyclical:"), points, count, radius, cyclical)
	rv.Autorelease()
	return rv
}



// Creates a path with the specified array of 3D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithFloat3Points:count:radius:cyclical:
func (pc _PathClass) PathWithFloat3PointsCountRadiusCyclical(points unsafe.Pointer, count uintptr, radius float32, cyclical bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pathWithFloat3Points:count:radius:cyclical:"), points, count, radius, cyclical)
	return rv
}


// Creates a path using the positions of the specified graph nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithGraphNodes:radius:
func (pc _PathClass) PathWithGraphNodesRadius(graphNodes []GraphNode, radius float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pathWithGraphNodes:radius:"), graphNodes, radius)
	return rv
}


// Creates a path with the specified array of 2D points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/pathWithPoints:count:radius:cyclical:
func (pc _PathClass) PathWithPointsCountRadiusCyclical(points unsafe.Pointer, count uintptr, radius float32, cyclical bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pathWithPoints:count:radius:cyclical:"), points, count, radius, cyclical)
	return rv
}


// Returns the 2D point at the specified index in the path’s list of vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/float2(at:)
func (p_ Path) Float2AtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("float2AtIndex:"), index)
	return rv
}


// Returns the 3D point at the specified index in the path’s list of vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/float3(at:)
func (p_ Path) Float3AtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("float3AtIndex:"), index)
	return rv
}


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/isCyclical
func (p_ Path) Cyclical() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cyclical"))
	return rv
}


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/isCyclical
func (p_ Path) SetCyclical(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCyclical:"), value)
}


// The number of vertices in the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/numPoints
func (p_ Path) NumPoints() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numPoints"))
	return rv
}


// The radius of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/radius
func (p_ Path) Radius() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("radius"))
	return rv
}


// The radius of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPath/radius
func (p_ Path) SetRadius(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRadius:"), value)
}


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkpath/iscyclical
func (p_ Path) IsCyclical() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCyclical"))
	return rv
}


// A Boolean value that determines whether the path loops around on itself (that is, the path’s end point connects to its start point).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gameplaykit/gkpath/iscyclical
func (p_ Path) SetIsCyclical(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCyclical:"), value)
}


