// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GraphType] class.
var (
	GraphTypeClass     _GraphTypeClass
	GraphTypeClassOnce sync.Once
)

func getGraphTypeClass() _GraphTypeClass {
	GraphTypeClassOnce.Do(func() {
		GraphTypeClass = _GraphTypeClass{objc.GetClass("MPSGraphType")}
	})
	return GraphTypeClass
}

type _GraphTypeClass struct {
	class objc.Class
}





// An interface definition for the [GraphType] class.
type IGraphType interface {
	IGraphObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GraphTypeClass) Alloc() GraphType {
	rv := objc.Send[GraphType](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphTypeClass) New() GraphType {
	rv := objc.Send[GraphType](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphType) Init() GraphType {
	rv := objc.Send[GraphType](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphType) Autorelease() GraphType {
	rv := objc.Send[GraphType](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphType creates a new GraphType instance.
func NewGraphType() GraphType {
	return getGraphTypeClass().New()
}





// The base type class for types on tensors.


// The base type class for types on tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphType
type GraphType struct {
	GraphObject
}

// GraphTypeFrom constructs a [GraphType] from an unsafe.Pointer.
//
// The base type class for types on tensors.
func GraphTypeFrom(ptr unsafe.Pointer) GraphType {
	return GraphType{
		GraphObject: GraphObjectFrom(ptr),
	}
}































