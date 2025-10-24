// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayDescriptor */


/* debug [class_header]: Header for MPSNDArrayDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayDescriptor */
// An interface definition for the [NDArrayDescriptor] class.
type INDArrayDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NDArrayDescriptor */
	// properties:
	DataType() DataType get set /* not a class type */
	SetDataType(value DataType get set /* not a class type */)
	NumberOfDimensions() objectivec.IObject
	SetNumberOfDimensions(value objectivec.IObject)
	PreferPackedRows() objectivec.IObject
	SetPreferPackedRows(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayDescriptor */
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
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayDescriptor
type NDArrayDescriptor struct {
	objectivec.Object
}

// NDArrayDescriptorFrom constructs a [NDArrayDescriptor] from an unsafe.Pointer.
func NDArrayDescriptorFrom(ptr unsafe.Pointer) NDArrayDescriptor {
	return NDArrayDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114063-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeDimensionCountDimensionSizes(dataType DataType, numberOfDimensions uint, dimensionSizes uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:dimensionCount:dimensionSizes:"), dataType, numberOfDimensions, dimensionSizes)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithDataTypeDimensionCountDimensionSizes) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114064-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeDimensionSizes(dataType DataType, dimension0 uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:dimensionSizes:"), dataType, dimension0)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithDataTypeDimensionSizes) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143491-descriptorwithdatatype
func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeShape(dataType DataType, shape unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("descriptorWithDataType:shape:"), dataType, shape)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithDataTypeShape) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114065-dimensionorder
func (n_ NDArrayDescriptor) DimensionOrder() {
	objc.Send[objc.ID](n_.ID, objc.Sel("dimensionOrder"))
}/* debug [instance_methods/method]: DimensionOrder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114066-length
func (n_ NDArrayDescriptor) Length() {
	objc.Send[objc.ID](n_.ID, objc.Sel("length"))
}/* debug [instance_methods/method]: Length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114066-lengthofdimension
func (n_ NDArrayDescriptor) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}/* debug [instance_methods/method]: LengthOfDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114069-slicedimension
func (n_ NDArrayDescriptor) SliceDimension() {
	objc.Send[objc.ID](n_.ID, objc.Sel("sliceDimension"))
}/* debug [instance_methods/method]: SliceDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114070-slicerange
func (n_ NDArrayDescriptor) SliceRange() {
	objc.Send[objc.ID](n_.ID, objc.Sel("sliceRange"))
}/* debug [instance_methods/method]: SliceRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114070-slicerangefordimension
func (n_ NDArrayDescriptor) SliceRangeForDimension(dimensionIndex uint) objc.IObject /* cross-framework: MPSDimensionSlice */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("sliceRangeForDimension:"), dimensionIndex)
	return rv
}/* debug [instance_methods/method]: SliceRangeForDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114071-transposedimension
func (n_ NDArrayDescriptor) TransposeDimension() {
	objc.Send[objc.ID](n_.ID, objc.Sel("transposeDimension"))
}/* debug [instance_methods/method]: TransposeDimension */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143492-reshape
func (n_ NDArrayDescriptor) Reshape() {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshape"))
}/* debug [instance_methods/method]: Reshape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143492-reshapewithdimensioncount
func (n_ NDArrayDescriptor) ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshapeWithDimensionCount:dimensionSizes:"), numberOfDimensions, dimensionSizes)
}/* debug [instance_methods/method]: ReshapeWithDimensionCountDimensionSizes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143493-reshapewithshape
func (n_ NDArrayDescriptor) ReshapeWithShape(shape unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reshapeWithShape:"), shape)
}/* debug [instance_methods/method]: ReshapeWithShape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423112-getshape
func (n_ NDArrayDescriptor) GetShape() {
	objc.Send[objc.ID](n_.ID, objc.Sel("getShape"))
}/* debug [instance_methods/method]: GetShape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423113-permute
func (n_ NDArrayDescriptor) Permute() {
	objc.Send[objc.ID](n_.ID, objc.Sel("permute"))
}/* debug [instance_methods/method]: Permute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423113-permutewithdimensionorder
func (n_ NDArrayDescriptor) PermuteWithDimensionOrder(dimensionOrder uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("permuteWithDimensionOrder:"), dimensionOrder)
}/* debug [instance_methods/method]: PermuteWithDimensionOrder */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype
func (n_ NDArrayDescriptor) DataType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype
func (n_ NDArrayDescriptor) SetDataType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions
func (n_ NDArrayDescriptor) NumberOfDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("numberOfDimensions"))
	return rv
}/* debug [instance_properties/getter]: numberOfDimensions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions
func (n_ NDArrayDescriptor) SetNumberOfDimensions(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberOfDimensions:"), value)
}/* debug [instance_properties/setter]: numberOfDimensions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423114-preferpackedrows
func (n_ NDArrayDescriptor) PreferPackedRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("preferPackedRows"))
	return rv
}/* debug [instance_properties/getter]: preferPackedRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/4423114-preferpackedrows
func (n_ NDArrayDescriptor) SetPreferPackedRows(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPreferPackedRows:"), value)
}/* debug [instance_properties/setter]: preferPackedRows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayDescriptor */



