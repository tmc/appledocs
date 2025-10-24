// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSVectorDescriptor */


/* debug [class_header]: Header for MPSVectorDescriptor */
// The class instance for the [VectorDescriptor] class.
var (
	VectorDescriptorClass     _VectorDescriptorClass
	VectorDescriptorClassOnce sync.Once
)

func getVectorDescriptorClass() _VectorDescriptorClass {
	VectorDescriptorClassOnce.Do(func() {
		VectorDescriptorClass = _VectorDescriptorClass{objc.GetClass("MPSVectorDescriptor")}
	})
	return VectorDescriptorClass
}

type _VectorDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VectorDescriptor */
// An interface definition for the [VectorDescriptor] class.
type IVectorDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VectorDescriptor */
	// properties:
	Vectors() objectivec.IObject
	SetVectors(value objectivec.IObject)
	VectorBytes() objectivec.IObject
	SetVectorBytes(value objectivec.IObject)
	Length() objectivec.IObject
	SetLength(value objectivec.IObject)
	DataType() DataType get set /* not a class type */
	SetDataType(value DataType get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VectorDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VectorDescriptor */
// Alloc allocates a new instance without initialization.
func (vc _VectorDescriptorClass) Alloc() VectorDescriptor {
	rv := objc.Send[VectorDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VectorDescriptorClass) New() VectorDescriptor {
	rv := objc.Send[VectorDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VectorDescriptor) Init() VectorDescriptor {
	rv := objc.Send[VectorDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VectorDescriptor) Autorelease() VectorDescriptor {
	rv := objc.Send[VectorDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVectorDescriptor creates a new VectorDescriptor instance.
func NewVectorDescriptor() VectorDescriptor {
	return getVectorDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VectorDescriptor */
// A description of the length and data type of a vector.


// A description of the length and data type of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSVectorDescriptor
type VectorDescriptor struct {
	objectivec.Object
}

// VectorDescriptorFrom constructs a [VectorDescriptor] from an unsafe.Pointer.
//
// A description of the length and data type of a vector.
func VectorDescriptorFrom(ptr unsafe.Pointer) VectorDescriptor {
	return VectorDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VectorDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VectorDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873328-vectordescriptorwithlength
func (vc _VectorDescriptorClass) VectorDescriptorWithLengthDataType(length uint, dataType DataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorDescriptorWithLength:dataType:"), length, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorDescriptorWithLengthDataType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873337-vectorbytes
func (vc _VectorDescriptorClass) VectorBytes() {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("vectorBytes"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorBytes) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873337-vectorbytesforlength
func (vc _VectorDescriptorClass) VectorBytesForLengthDataType(length uint, dataType DataType) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](objc.ID(vc.class), objc.Sel("vectorBytesForLength:dataType:"), length, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorBytesForLengthDataType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873348-vectordescriptorwithlength
func (vc _VectorDescriptorClass) VectorDescriptorWithLengthVectorsVectorBytesDataType(length uint, vectors uint, vectorBytes uint, dataType DataType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorDescriptorWithLength:vectors:vectorBytes:dataType:"), length, vectors, vectorBytes, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorDescriptorWithLengthVectorsVectorBytesDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VectorDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VectorDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VectorDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873333-vectors
func (v_ VectorDescriptor) Vectors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("vectors"))
	return rv
}/* debug [instance_properties/getter]: vectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873333-vectors
func (v_ VectorDescriptor) SetVectors(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVectors:"), value)
}/* debug [instance_properties/setter]: vectors */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873335-vectorbytes
func (v_ VectorDescriptor) VectorBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("vectorBytes"))
	return rv
}/* debug [instance_properties/getter]: vectorBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873335-vectorbytes
func (v_ VectorDescriptor) SetVectorBytes(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVectorBytes:"), value)
}/* debug [instance_properties/setter]: vectorBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873345-length
func (v_ VectorDescriptor) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873345-length
func (v_ VectorDescriptor) SetLength(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLength:"), value)
}/* debug [instance_properties/setter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873362-datatype
func (v_ VectorDescriptor) DataType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873362-datatype
func (v_ VectorDescriptor) SetDataType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSVectorDescriptor */





