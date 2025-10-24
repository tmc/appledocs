// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixDescriptor */


/* debug [class_header]: Header for MPSMatrixDescriptor */
// The class instance for the [MatrixDescriptor] class.
var (
	MatrixDescriptorClass     _MatrixDescriptorClass
	MatrixDescriptorClassOnce sync.Once
)

func getMatrixDescriptorClass() _MatrixDescriptorClass {
	MatrixDescriptorClassOnce.Do(func() {
		MatrixDescriptorClass = _MatrixDescriptorClass{objc.GetClass("MPSMatrixDescriptor")}
	})
	return MatrixDescriptorClass
}

type _MatrixDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixDescriptor */
// An interface definition for the [MatrixDescriptor] class.
type IMatrixDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MatrixDescriptor */
	// properties:
	Columns() objectivec.IObject
	SetColumns(value objectivec.IObject)
	RowBytes() objectivec.IObject
	SetRowBytes(value objectivec.IObject)
	DataType() DataType get set /* not a class type */
	SetDataType(value DataType get set /* not a class type */)
	Rows() objectivec.IObject
	SetRows(value objectivec.IObject)
	Matrices() objectivec.IObject
	SetMatrices(value objectivec.IObject)
	MatrixBytes() objectivec.IObject
	SetMatrixBytes(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MatrixDescriptorClass) Alloc() MatrixDescriptor {
	rv := objc.Send[MatrixDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixDescriptorClass) New() MatrixDescriptor {
	rv := objc.Send[MatrixDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixDescriptor) Init() MatrixDescriptor {
	rv := objc.Send[MatrixDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixDescriptor) Autorelease() MatrixDescriptor {
	rv := objc.Send[MatrixDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixDescriptor creates a new MatrixDescriptor instance.
func NewMatrixDescriptor() MatrixDescriptor {
	return getMatrixDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixDescriptor */
// A description of attributes used to create an MPS matrix.
//
// Matrix data is assumed to be stored in row-major order.


// A description of attributes used to create an MPS matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixDescriptor
type MatrixDescriptor struct {
	objectivec.Object
}

// MatrixDescriptorFrom constructs a [MatrixDescriptor] from an unsafe.Pointer.
//
// A description of attributes used to create an MPS matrix.
func MatrixDescriptorFrom(ptr unsafe.Pointer) MatrixDescriptor {
	return MatrixDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixDescriptor */

// Determines the recommended matrix row stride, in bytes, for a given number of columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143204-rowbytes
func (mc _MatrixDescriptorClass) RowBytes() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("rowBytes"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RowBytes) */


// Determines the recommended matrix row stride, in bytes, for a given number of columns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143204-rowbytesfromcolumns
func (mc _MatrixDescriptorClass) RowBytesFromColumnsDataType(columns uint, dataType DataType) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](objc.ID(mc.class), objc.Sel("rowBytesFromColumns:dataType:"), columns, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RowBytesFromColumnsDataType) */


// Creates a matrix descriptor with the specified dimensions and data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143206-matrixdescriptorwithdimensions
func (mc _MatrixDescriptorClass) MatrixDescriptorWithDimensionsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType DataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("matrixDescriptorWithDimensions:columns:rowBytes:dataType:"), rows, columns, rowBytes, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MatrixDescriptorWithDimensionsColumnsRowBytesDataType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873331-matrixdescriptorwithrows
func (mc _MatrixDescriptorClass) MatrixDescriptorWithRowsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType DataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("matrixDescriptorWithRows:columns:rowBytes:dataType:"), rows, columns, rowBytes, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MatrixDescriptorWithRowsColumnsRowBytesDataType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873350-matrixdescriptorwithrows
func (mc _MatrixDescriptorClass) MatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType(rows uint, columns uint, matrices uint, rowBytes uint, matrixBytes uint, dataType DataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("matrixDescriptorWithRows:columns:matrices:rowBytes:matrixBytes:dataType:"), rows, columns, matrices, rowBytes, matrixBytes, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873394-rowbytesforcolumns
func (mc _MatrixDescriptorClass) RowBytesForColumnsDataType(columns uint, dataType DataType) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](objc.ID(mc.class), objc.Sel("rowBytesForColumns:dataType:"), columns, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RowBytesForColumnsDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixDescriptor */

// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143196-columns
func (m_ MatrixDescriptor) Columns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("columns"))
	return rv
}/* debug [instance_properties/getter]: columns */


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143196-columns
func (m_ MatrixDescriptor) SetColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColumns:"), value)
}/* debug [instance_properties/setter]: columns */


// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143199-rowbytes
func (m_ MatrixDescriptor) RowBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("rowBytes"))
	return rv
}/* debug [instance_properties/getter]: rowBytes */


// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143199-rowbytes
func (m_ MatrixDescriptor) SetRowBytes(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRowBytes:"), value)
}/* debug [instance_properties/setter]: rowBytes */


// The type of the values in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143202-datatype
func (m_ MatrixDescriptor) DataType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The type of the values in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143202-datatype
func (m_ MatrixDescriptor) SetDataType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143203-rows
func (m_ MatrixDescriptor) Rows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("rows"))
	return rv
}/* debug [instance_properties/getter]: rows */


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143203-rows
func (m_ MatrixDescriptor) SetRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRows:"), value)
}/* debug [instance_properties/setter]: rows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873351-matrices
func (m_ MatrixDescriptor) Matrices() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("matrices"))
	return rv
}/* debug [instance_properties/getter]: matrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873351-matrices
func (m_ MatrixDescriptor) SetMatrices(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatrices:"), value)
}/* debug [instance_properties/setter]: matrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873387-matrixbytes
func (m_ MatrixDescriptor) MatrixBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("matrixBytes"))
	return rv
}/* debug [instance_properties/getter]: matrixBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873387-matrixbytes
func (m_ MatrixDescriptor) SetMatrixBytes(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatrixBytes:"), value)
}/* debug [instance_properties/setter]: matrixBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixDescriptor */



