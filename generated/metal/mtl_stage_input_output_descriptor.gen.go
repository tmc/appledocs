// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLStageInputOutputDescriptor */


/* debug [class_header]: Header for MTLStageInputOutputDescriptor */
// The class instance for the [StageInputOutputDescriptor] class.
var (
	StageInputOutputDescriptorClass     _StageInputOutputDescriptorClass
	StageInputOutputDescriptorClassOnce sync.Once
)

func getStageInputOutputDescriptorClass() _StageInputOutputDescriptorClass {
	StageInputOutputDescriptorClassOnce.Do(func() {
		StageInputOutputDescriptorClass = _StageInputOutputDescriptorClass{objc.GetClass("MTLStageInputOutputDescriptor")}
	})
	return StageInputOutputDescriptorClass
}

type _StageInputOutputDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StageInputOutputDescriptor */
// An interface definition for the [StageInputOutputDescriptor] class.
type IStageInputOutputDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StageInputOutputDescriptor */
	// properties:
	Attributes() IMTLAttributeDescriptorArray
	IndexBufferIndex() uint
	SetIndexBufferIndex(value uint)
	IndexType() IndexType
	SetIndexType(value IndexType)
	Layouts() IMTLBufferLayoutDescriptorArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StageInputOutputDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StageInputOutputDescriptor */
// Alloc allocates a new instance without initialization.
func (sc _StageInputOutputDescriptorClass) Alloc() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StageInputOutputDescriptorClass) New() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StageInputOutputDescriptor) Init() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StageInputOutputDescriptor) Autorelease() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStageInputOutputDescriptor creates a new StageInputOutputDescriptor instance.
func NewStageInputOutputDescriptor() StageInputOutputDescriptor {
	return getStageInputOutputDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StageInputOutputDescriptor */
// A description of the input and output data of a function.


// A description of the input and output data of a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor
type StageInputOutputDescriptor struct {
	objectivec.Object
}

// StageInputOutputDescriptorFrom constructs a [StageInputOutputDescriptor] from an unsafe.Pointer.
//
// A description of the input and output data of a function.
func StageInputOutputDescriptorFrom(ptr unsafe.Pointer) StageInputOutputDescriptor {
	return StageInputOutputDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StageInputOutputDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StageInputOutputDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/stageInputOutputDescriptor
func (sc _StageInputOutputDescriptorClass) StageInputOutputDescriptor() IStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("stageInputOutputDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StageInputOutputDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StageInputOutputDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StageInputOutputDescriptor */

// Resets the default state for the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/reset()
func (s_ StageInputOutputDescriptor) Reset() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StageInputOutputDescriptor */

// An array that describes where and how to fetch data for the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/attributes
func (s_ StageInputOutputDescriptor) Attributes() IMTLAttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](s_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// The location of the index buffer for a compute function using indexed thread addressing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexBufferIndex
func (s_ StageInputOutputDescriptor) IndexBufferIndex() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("indexBufferIndex"))
	return rv
}/* debug [instance_properties/getter]: indexBufferIndex */


// The location of the index buffer for a compute function using indexed thread addressing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexBufferIndex
func (s_ StageInputOutputDescriptor) SetIndexBufferIndex(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexBufferIndex:"), value)
}/* debug [instance_properties/setter]: indexBufferIndex */


// The data type of the indices stored in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexType
func (s_ StageInputOutputDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](s_.ID, objc.Sel("indexType"))
	return rv
}/* debug [instance_properties/getter]: indexType */


// The data type of the indices stored in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexType
func (s_ StageInputOutputDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexType:"), value)
}/* debug [instance_properties/setter]: indexType */


// An array that describes how the function fetches data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/layouts
func (s_ StageInputOutputDescriptor) Layouts() IMTLBufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](s_.ID, objc.Sel("layouts"))
	return rv
}/* debug [instance_properties/getter]: layouts */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLStageInputOutputDescriptor */



