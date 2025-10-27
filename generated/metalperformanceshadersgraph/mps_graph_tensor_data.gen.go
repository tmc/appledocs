// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	DataType() DataType /* not a class type */
	Device() IMPSGraphDevice
	Shape() Shape /* not a class type */


	

	// methods:
	Mpsndarray() metalperformanceshaders.NDArray


}





// Alloc allocates a new instance without initialization.
func (gc _GraphTensorDataClass) Alloc() GraphTensorData {
	rv := objc.Send[GraphTensorData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// The representation of a compute data type.
//
// Pass data to a graph using a tensor data, a reference will be taken to your data and used just in time when the graph is run.


// The representation of a compute data type.
//
// [Full Topic]
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






// Initializes the tensor data with an on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(device:data:shape:dataType:)
func NewGraphTensorDataWithDeviceDataShapeDataType(device IMPSGraphDevice, data foundation.foundation.INSData, shape Shape /* not a class type */, dataType DataType /* not a class type */) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithDevice:data:shape:dataType:"), device, data, shape, dataType)
	rv.Autorelease()
	return rv
}


// Initializes a tensor data with an MPS image batch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-511a
func NewGraphTensorDataWithMPSImageBatch(imageBatch ImageBatch /* not a class type */) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSImageBatch:"), imageBatch)
	rv.Autorelease()
	return rv
}


// Initializes a tensor data with an MPS matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-2go2
func NewGraphTensorDataWithMPSMatrix(matrix appkit.Matrix) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSMatrix:"), matrix)
	rv.Autorelease()
	return rv
}


// Initializes a tensor data with an MPS matrix enforcing rank of the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1lnxg
func NewGraphTensorDataWithMPSMatrixRank(matrix appkit.Matrix, rank uint) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSMatrix:rank:"), matrix, rank)
	rv.Autorelease()
	return rv
}


// Initializes an MPSGraphTensorData with an MPS ndarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-4bnfb
func NewGraphTensorDataWithMPSNDArray(ndarray metalperformanceshaders.NDArray) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSNDArray:"), ndarray)
	rv.Autorelease()
	return rv
}


// Initializes a tensor data with an MPS vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-9kgoe
func NewGraphTensorDataWithMPSVector(vector corefoundation.corefoundation.IVector) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSVector:"), vector)
	rv.Autorelease()
	return rv
}


// Initializes a tensor data with an MPS vector enforcing rank of the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:rank:)-1e4ks
func NewGraphTensorDataWithMPSVectorRank(vector corefoundation.corefoundation.IVector, rank uint) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMPSVector:rank:"), vector, rank)
	rv.Autorelease()
	return rv
}


// Initializes an tensor data with a metal buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:)
func NewGraphTensorDataWithMTLBufferShapeDataType(buffer unsafe.Pointer, shape Shape /* not a class type */, dataType DataType /* not a class type */) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMTLBuffer:shape:dataType:"), buffer, shape, dataType)
	rv.Autorelease()
	return rv
}


// Initializes an tensor data with a metal buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:shape:dataType:rowBytes:)
func NewGraphTensorDataWithMTLBufferShapeDataTypeRowBytes(buffer unsafe.Pointer, shape Shape /* not a class type */, dataType DataType /* not a class type */, rowBytes uint) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMTLBuffer:shape:dataType:rowBytes:"), buffer, shape, dataType, rowBytes)
	rv.Autorelease()
	return rv
}


// Initializes an MPSGraphTensorData with an MTLTensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/init(_:)-60j6x
func NewGraphTensorDataWithMTLTensor(tensor unsafe.Pointer) GraphTensorData {
	instance := getGraphTensorDataClass().Alloc()
	rv := objc.Send[GraphTensorData](instance.ID, objc.Sel("initWithMTLTensor:"), tensor)
	rv.Autorelease()
	return rv
}

















// Return an mpsndarray object will copy contents if the contents are not stored in an MPS ndarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/mpsndarray()
func (g_ GraphTensorData) Mpsndarray() metalperformanceshaders.NDArray {
	rv := objc.Send[metalperformanceshaders.NDArray](g_.ID, objc.Sel("mpsndarray"))
	return rv
}







// The data type of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/dataType
func (g_ GraphTensorData) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}


// The device of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/device
func (g_ GraphTensorData) Device() IMPSGraphDevice {
	rv := objc.Send[GraphDevice](g_.ID, objc.Sel("device"))
	return rv
}


// The shape of the tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensorData/shape
func (g_ GraphTensorData) Shape() Shape /* not a class type */ {
	rv := objc.Send[Shape](g_.ID, objc.Sel("shape"))
	return rv
}







