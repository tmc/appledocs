// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrix */


/* debug [class_header]: Header for MPSMatrix */
// The class instance for the [Matrix] class.
var (
	MatrixClass     _MatrixClass
	MatrixClassOnce sync.Once
)

func getMatrixClass() _MatrixClass {
	MatrixClassOnce.Do(func() {
		MatrixClass = _MatrixClass{objc.GetClass("MPSMatrix")}
	})
	return MatrixClass
}

type _MatrixClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Matrix */
// An interface definition for the [Matrix] class.
type IMatrix interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Matrix */
	// properties:
	DataType() DataType get /* not a class type */
	SetDataType(value DataType get /* not a class type */)
	Data() Buffer get /* not a class type */
	SetData(value Buffer get /* not a class type */)
	Columns() objectivec.IObject
	SetColumns(value objectivec.IObject)
	RowBytes() objectivec.IObject
	SetRowBytes(value objectivec.IObject)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
	Rows() objectivec.IObject
	SetRows(value objectivec.IObject)
	Matrices() objectivec.IObject
	SetMatrices(value objectivec.IObject)
	MatrixBytes() objectivec.IObject
	SetMatrixBytes(value objectivec.IObject)
	Offset() objectivec.IObject
	SetOffset(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Matrix */
	// methods:
	ResourceSize()
	Synchronize()
	SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Matrix */
// Alloc allocates a new instance without initialization.
func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixClass) New() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Matrix) Init() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Matrix) Autorelease() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrix creates a new Matrix instance.
func NewMatrix() Matrix {
	return getMatrixClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Matrix */
// A 2D array of data that stores the data’s values.
//
// objects serve as inputs and outputs of objects. Matrix data is assumed to be stored in row-major order.


// A 2D array of data that stores the data’s values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrix
type Matrix struct {
	objectivec.Object
}

// MatrixFrom constructs a [Matrix] from an unsafe.Pointer.
//
// A 2D array of data that stores the data’s values.
func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Matrix */

// Initializes a matrix with a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143201-initwithbuffer
func NewMatrixWithBufferDescriptor(buffer unsafe.Pointer, descriptor IMatrixDescriptor) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithBufferDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3229863-initwithbuffer
func NewMatrixWithBufferOffsetDescriptor(buffer unsafe.Pointer, offset uint, descriptor IMatrixDescriptor) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithBufferOffsetDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942567-initwithdevice
func NewMatrixWithDeviceDescriptor(device unsafe.Pointer, descriptor IMatrixDescriptor) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithDeviceDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Matrix */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Matrix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Matrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942569-resourcesize
func (m_ Matrix) ResourceSize() {
	objc.Send[objc.ID](m_.ID, objc.Sel("resourceSize"))
}/* debug [instance_methods/method]: ResourceSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942571-synchronize
func (m_ Matrix) Synchronize() {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronize"))
}/* debug [instance_methods/method]: Synchronize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942571-synchronizeoncommandbuffer
func (m_ Matrix) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}/* debug [instance_methods/method]: SynchronizeOnCommandBuffer */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Matrix */

// The type of the values in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143197-datatype
func (m_ Matrix) DataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The type of the values in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143197-datatype
func (m_ Matrix) SetDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The buffer that stores the matrix data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143205-data
func (m_ Matrix) Data() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The buffer that stores the matrix data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143205-data
func (m_ Matrix) SetData(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143207-columns
func (m_ Matrix) Columns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("columns"))
	return rv
}/* debug [instance_properties/getter]: columns */


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143207-columns
func (m_ Matrix) SetColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColumns:"), value)
}/* debug [instance_properties/setter]: columns */


// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143208-rowbytes
func (m_ Matrix) RowBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("rowBytes"))
	return rv
}/* debug [instance_properties/getter]: rowBytes */


// The stride, in bytes, between corresponding elements of consecutive rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143208-rowbytes
func (m_ Matrix) SetRowBytes(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRowBytes:"), value)
}/* debug [instance_properties/setter]: rowBytes */


// The device on which the matrix will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143209-device
func (m_ Matrix) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device on which the matrix will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143209-device
func (m_ Matrix) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143210-rows
func (m_ Matrix) Rows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("rows"))
	return rv
}/* debug [instance_properties/getter]: rows */


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143210-rows
func (m_ Matrix) SetRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRows:"), value)
}/* debug [instance_properties/setter]: rows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873334-matrices
func (m_ Matrix) Matrices() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("matrices"))
	return rv
}/* debug [instance_properties/getter]: matrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873334-matrices
func (m_ Matrix) SetMatrices(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatrices:"), value)
}/* debug [instance_properties/setter]: matrices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873344-matrixbytes
func (m_ Matrix) MatrixBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("matrixBytes"))
	return rv
}/* debug [instance_properties/getter]: matrixBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873344-matrixbytes
func (m_ Matrix) SetMatrixBytes(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatrixBytes:"), value)
}/* debug [instance_properties/setter]: matrixBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3375740-offset
func (m_ Matrix) Offset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3375740-offset
func (m_ Matrix) SetOffset(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrix */


