// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Vector] class.
var (
	VectorClass     _VectorClass
	VectorClassOnce sync.Once
)

func getVectorClass() _VectorClass {
	VectorClassOnce.Do(func() {
		VectorClass = _VectorClass{objc.GetClass("MPSVector")}
	})
	return VectorClass
}

type _VectorClass struct {
	class objc.Class
}





// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
	

	// properties:
	DataType() DataType get /* not a class type */
	SetDataType(value DataType get /* not a class type */)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
	VectorBytes() objectivec.IObject
	SetVectorBytes(value objectivec.IObject)
	Vectors() objectivec.IObject
	SetVectors(value objectivec.IObject)
	Length() objectivec.IObject
	SetLength(value objectivec.IObject)
	Data() Buffer get /* not a class type */
	SetData(value Buffer get /* not a class type */)
	Offset() objectivec.IObject
	SetOffset(value objectivec.IObject)


	

	// methods:
	Synchronize()
	SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer)
	ResourceSize()


}





// Alloc allocates a new instance without initialization.
func (vc _VectorClass) Alloc() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VectorClass) New() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Vector) Init() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Vector) Autorelease() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVector creates a new Vector instance.
func NewVector() Vector {
	return getVectorClass().New()
}





// A 1D array of data that stores the data’s values.


// A 1D array of data that stores the data’s values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVector
type Vector struct {
	objectivec.Object
}

// VectorFrom constructs a [Vector] from an unsafe.Pointer.
//
// A 1D array of data that stores the data’s values.
func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873346-initwithbuffer
func NewVectorWithBufferDescriptor(buffer unsafe.Pointer, descriptor IVectorDescriptor) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithBuffer:descriptor:"), buffer, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/3229864-initwithbuffer
func NewVectorWithBufferOffsetDescriptor(buffer unsafe.Pointer, offset uint, descriptor IVectorDescriptor) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithBuffer:offset:descriptor:"), buffer, offset, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942566-initwithdevice
func NewVectorWithDeviceDescriptor(device unsafe.Pointer, descriptor IVectorDescriptor) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942568-synchronize
func (v_ Vector) Synchronize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("synchronize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942568-synchronizeoncommandbuffer
func (v_ Vector) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942570-resourcesize
func (v_ Vector) ResourceSize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("resourceSize"))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873336-datatype
func (v_ Vector) DataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("dataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873336-datatype
func (v_ Vector) SetDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873338-device
func (v_ Vector) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("device"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873338-device
func (v_ Vector) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDevice:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873340-vectorbytes
func (v_ Vector) VectorBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("vectorBytes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873340-vectorbytes
func (v_ Vector) SetVectorBytes(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVectorBytes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873388-vectors
func (v_ Vector) Vectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("vectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873388-vectors
func (v_ Vector) SetVectors(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVectors:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873392-length
func (v_ Vector) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("length"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873392-length
func (v_ Vector) SetLength(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLength:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873393-data
func (v_ Vector) Data() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873393-data
func (v_ Vector) SetData(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/3375741-offset
func (v_ Vector) Offset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("offset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/3375741-offset
func (v_ Vector) SetOffset(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOffset:"), value)
}







