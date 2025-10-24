// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GraphObject] class.
var (
	GraphObjectClass     _GraphObjectClass
	GraphObjectClassOnce sync.Once
)

func getGraphObjectClass() _GraphObjectClass {
	GraphObjectClassOnce.Do(func() {
		GraphObjectClass = _GraphObjectClass{objc.GetClass("MPSGraphObject")}
	})
	return GraphObjectClass
}

type _GraphObjectClass struct {
	class objc.Class
}

// An interface definition for the [GraphObject] class.
type IGraphObject interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The common base class for all Metal Performance Shaders Graph objects.
//
// Only the child classes should be used.


// The common base class for all Metal Performance Shaders Graph objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphObject
type GraphObject struct {
	objectivec.Object
}

// GraphObjectFrom constructs a [GraphObject] from an unsafe.Pointer.
//
// The common base class for all Metal Performance Shaders Graph objects.
func GraphObjectFrom(ptr unsafe.Pointer) GraphObject {
	return GraphObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphObjectClass) Alloc() GraphObject {
	rv := objc.Send[GraphObject](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphObjectClass) New() GraphObject {
	rv := objc.Send[GraphObject](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphObject) Init() GraphObject {
	rv := objc.Send[GraphObject](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphObject) Autorelease() GraphObject {
	rv := objc.Send[GraphObject](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphObject creates a new GraphObject instance.
func NewGraphObject() GraphObject {
	return getGraphObjectClass().New()
}




