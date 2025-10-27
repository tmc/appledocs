// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AttributeDescriptorArray] class.
var (
	AttributeDescriptorArrayClass     _AttributeDescriptorArrayClass
	AttributeDescriptorArrayClassOnce sync.Once
)

func getAttributeDescriptorArrayClass() _AttributeDescriptorArrayClass {
	AttributeDescriptorArrayClassOnce.Do(func() {
		AttributeDescriptorArrayClass = _AttributeDescriptorArrayClass{objc.GetClass("MTLAttributeDescriptorArray")}
	})
	return AttributeDescriptorArrayClass
}

type _AttributeDescriptorArrayClass struct {
	class objc.Class
}





// An interface definition for the [AttributeDescriptorArray] class.
type IAttributeDescriptorArray interface {
	objectivec.IObject
	

	// properties:
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)


	

	// methods:
	SetObjectAtIndexedSubscript(attributeDesc IMTLAttributeDescriptor, index uint)
	ObjectAtIndexedSubscript(index uint) IAttributeDescriptor


}





// Alloc allocates a new instance without initialization.
func (ac _AttributeDescriptorArrayClass) Alloc() AttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AttributeDescriptorArrayClass) New() AttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributeDescriptorArray) Init() AttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributeDescriptorArray) Autorelease() AttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributeDescriptorArray creates a new AttributeDescriptorArray instance.
func NewAttributeDescriptorArray() AttributeDescriptorArray {
	return getAttributeDescriptorArrayClass().New()
}





// An array of attribute descriptor objects.
//
// An defines the data format and index binding for the attribute argument table, using instances.


// An array of attribute descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray
type AttributeDescriptorArray struct {
	objectivec.Object
}

// AttributeDescriptorArrayFrom constructs a [AttributeDescriptorArray] from an unsafe.Pointer.
//
// An array of attribute descriptor objects.
func AttributeDescriptorArrayFrom(ptr unsafe.Pointer) AttributeDescriptorArray {
	return AttributeDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}




















// Sets state for the specified attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray/setObject:atIndexedSubscript:
func (a_ AttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IMTLAttributeDescriptor, index uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setObject:atIndexedSubscript:"), attributeDesc, index)
}


// Returns the state of the specified attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray/subscript(_:)
func (a_ AttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) IAttributeDescriptor {
	rv := objc.Send[AttributeDescriptor](a_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}







// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptorArray) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](a_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptorArray) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStageInputDescriptor:"), value)
}








