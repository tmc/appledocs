// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArray] class.
var (
	NDArrayClass     _NDArrayClass
	NDArrayClassOnce sync.Once
)

func getNDArrayClass() _NDArrayClass {
	NDArrayClassOnce.Do(func() {
		NDArrayClass = _NDArrayClass{objc.GetClass("MPSNDArray")}
	})
	return NDArrayClass
}

type _NDArrayClass struct {
	class objc.Class
}

// An interface definition for the [NDArray] class.
type INDArray interface {
	objectivec.IObject
	ArrayViewWithShapeStrides(shape unsafe.Pointer, strides unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray
type NDArray struct {
	objectivec.Object
}

// NDArrayFrom constructs a [NDArray] from an unsafe.Pointer.
func NDArrayFrom(ptr unsafe.Pointer) NDArray {
	return NDArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayClass) Alloc() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayClass) New() NDArray {
	rv := objc.Send[NDArray](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArray) Init() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArray) Autorelease() NDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArray creates a new NDArray instance.
func NewNDArray() NDArray {
	return getNDArrayClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArray/arrayView(withShape:strides:)
func (n_ NDArray) ArrayViewWithShapeStrides(shape unsafe.Pointer, strides unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("arrayViewWithShape:strides:"), shape, strides)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/datatype
func (n_ NDArray) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/datatype
func (n_ NDArray) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/datatypesize
func (n_ NDArray) DataTypeSize() int {
	rv := objc.Send[int](n_.ID, objc.Sel("dataTypeSize"))
	return rv
}


// SetDataTypeSize sets the value of the dataTypeSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/datatypesize
func (n_ NDArray) SetDataTypeSize(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataTypeSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/device
func (n_ NDArray) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/device
func (n_ NDArray) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDevice:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/label
func (n_ NDArray) Label() string {
	rv := objc.Send[string](n_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/label
func (n_ NDArray) SetLabel(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/numberofdimensions
func (n_ NDArray) NumberOfDimensions() int {
	rv := objc.Send[int](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}


// SetNumberOfDimensions sets the value of the numberOfDimensions property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/numberofdimensions
func (n_ NDArray) SetNumberOfDimensions(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/parent
func (n_ NDArray) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("parent"))
	return rv
}


// SetParent sets the value of the parent property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/parent
func (n_ NDArray) SetParent(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setParent:"), value)
}



