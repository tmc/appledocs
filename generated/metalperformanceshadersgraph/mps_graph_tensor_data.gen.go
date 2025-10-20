// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GraphTensorData] class.
var (
	GraphTensorDataClass     _GraphTensorDataClass
	GraphTensorDataClassOnce sync.Once
)

func getGraphTensorDataClass() _GraphTensorDataClass {
	GraphTensorDataClassOnce.Do(func() {
		GraphTensorDataClass = _GraphTensorDataClass{objc.GetClass("MPSGraphTensorData")}
	})
	return GraphTensorDataClass
}

type _GraphTensorDataClass struct {
	class objc.Class
}

// An interface definition for the [GraphTensorData] class.
type IGraphTensorData interface {
	IGraphObject
}

// The representation of a compute data type.
//
// Pass data to a graph using a tensor data, a reference will be taken to your data and used just in time when the graph is run.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData
type GraphTensorData struct {
	GraphObject
}

// GraphTensorDataFrom constructs a [GraphTensorData] from an unsafe.Pointer.
//
// The representation of a compute data type.
func GraphTensorDataFrom(ptr unsafe.Pointer) GraphTensorData {
	return GraphTensorData{
		GraphObject: GraphObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GraphTensorDataClass) Alloc() GraphTensorData {
	rv := objc.Send[GraphTensorData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphTensorDataClass) New() GraphTensorData {
	rv := objc.Send[GraphTensorData](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphTensorData) Init() GraphTensorData {
	rv := objc.Send[GraphTensorData](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphTensorData) Autorelease() GraphTensorData {
	rv := objc.Send[GraphTensorData](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphTensorData creates a new GraphTensorData instance.
func NewGraphTensorData() GraphTensorData {
	return getGraphTensorDataClass().New()
}


// Initializes a tensor data with an MPS vector enforcing rank of the result.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1e4ks
func NewGraphTensorDataWithMPSVectorRank(vector unsafe.Pointer, rank uint) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSVector:rank:"), vector, rank)
	rv.Autorelease()
	return rv
}


// The data type of the tensor data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/dataType
func (g_ GraphTensorData) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}


