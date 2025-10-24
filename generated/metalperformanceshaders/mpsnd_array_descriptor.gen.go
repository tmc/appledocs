// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NDArrayDescriptor] class.
var (
	NDArrayDescriptorClass     _NDArrayDescriptorClass
	NDArrayDescriptorClassOnce sync.Once
)

func getNDArrayDescriptorClass() _NDArrayDescriptorClass {
	NDArrayDescriptorClassOnce.Do(func() {
		NDArrayDescriptorClass = _NDArrayDescriptorClass{objc.GetClass("MPSNDArrayDescriptor")}
	})
	return NDArrayDescriptorClass
}

type _NDArrayDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [NDArrayDescriptor] class.
type INDArrayDescriptor interface {
	objectivec.IObject
	

	// properties:
	DataType() DataType get set /* not a class type */
	SetDataType(value DataType get set /* not a class type */)
	NumberOfDimensions() objectivec.IObject
	SetNumberOfDimensions(value objectivec.IObject)
	PreferPackedRows() objectivec.IObject
	SetPreferPackedRows(value objectivec.IObject)


	

	// methods:
	DimensionOrder()
	Length()
	LengthOfDimension(dimensionIndex uint) uint
	SliceDimension()
	SliceRange()
	SliceRangeForDimension(dimensionIndex uint) objc.IObject /* cross-framework: MPSDimensionSlice */
	TransposeDimension()
	Reshape()
	ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes uint)
	ReshapeWithShape(shape unsafe.Pointer)
	GetShape()
	Permute()
	PermuteWithDimensionOrder(dimensionOrder uint)


}





// Alloc allocates a new instance without initialization.
func (nc _NDArrayDescriptorClass) Alloc() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayDescriptorClass) New() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayDescriptor) Init() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayDescriptor) Autorelease() NDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayDescriptor creates a new NDArrayDescriptor instance.
func NewNDArrayDescriptor() NDArrayDescriptor {
	return getNDArrayDescriptorClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor
type NDArrayDescriptor struct {
	objectivec.Object
}

// NDArrayDescriptorFrom constructs a [NDArrayDescriptor] from an unsafe.Pointer.
func NDArrayDescriptorFrom(ptr unsafe.Pointer) NDArrayDescriptor {
	return NDArrayDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114063-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeDimensionCountDimensionSizes(dataType DataType, numberOfDimensions uint, dimensionSizes uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:dimensionCount:dimensionSizes:"), dataType, numberOfDimensions, dimensionSizes)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114064-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeDimensionSizes(dataType DataType, dimension0 uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:dimensionSizes:"), dataType, dimension0)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143491-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeShape(dataType DataType, shape unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:shape:"), dataType, shape)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114065-dimensionorder
func (n_ NDArrayDescriptor) DimensionOrder() {
	objc.Send[objc.ID](n_.ID, objc.Sel("dimensionOrder"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114066-length
func (n_ NDArrayDescriptor) Length() {
	objc.Send[objc.ID](n_.ID, objc.Sel("length"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114066-lengthofdimension
func (n_ NDArrayDescriptor) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114069-slicedimension
func (n_ NDArrayDescriptor) SliceDimension() {
	objc.Send[objc.ID](n_.ID, objc.Sel("sliceDimension"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114070-slicerange
func (n_ NDArrayDescriptor) SliceRange() {
	objc.Send[objc.ID](n_.ID, objc.Sel("sliceRange"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114070-slicerangefordimension
func (n_ NDArrayDescriptor) SliceRangeForDimension(dimensionIndex uint) objc.IObject /* cross-framework: MPSDimensionSlice */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("sliceRangeForDimension:"), dimensionIndex)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114071-transposedimension
func (n_ NDArrayDescriptor) TransposeDimension() {
	objc.Send[objc.ID](n_.ID, objc.Sel("transposeDimension"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143492-reshape
func (n_ NDArrayDescriptor) Reshape() {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshape"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143492-reshapewithdimensioncount
func (n_ NDArrayDescriptor) ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshapeWithDimensionCount:dimensionSizes:"), numberOfDimensions, dimensionSizes)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143493-reshapewithshape
func (n_ NDArrayDescriptor) ReshapeWithShape(shape unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshapeWithShape:"), shape)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423112-getshape
func (n_ NDArrayDescriptor) GetShape() {
	objc.Send[objc.ID](n_.ID, objc.Sel("getShape"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423113-permute
func (n_ NDArrayDescriptor) Permute() {
	objc.Send[objc.ID](n_.ID, objc.Sel("permute"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423113-permutewithdimensionorder
func (n_ NDArrayDescriptor) PermuteWithDimensionOrder(dimensionOrder uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("permuteWithDimensionOrder:"), dimensionOrder)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype
func (n_ NDArrayDescriptor) DataType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("dataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype
func (n_ NDArrayDescriptor) SetDataType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions
func (n_ NDArrayDescriptor) NumberOfDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions
func (n_ NDArrayDescriptor) SetNumberOfDimensions(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423114-preferpackedrows
func (n_ NDArrayDescriptor) PreferPackedRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("preferPackedRows"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423114-preferpackedrows
func (n_ NDArrayDescriptor) SetPreferPackedRows(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPreferPackedRows:"), value)
}








