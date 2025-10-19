// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKAttributeValue] class.
var sKAttributeValueClass = _SKAttributeValueClass{objc.GetClass("SKAttributeValue")}

type _SKAttributeValueClass struct {
	class objc.Class
}

// An interface definition for the [SKAttributeValue] class.
type ISKAttributeValue interface {
	objectivec.IObject
}

// A container for dynamic shader data associated with a node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue

type SKAttributeValue struct {
	objectivec.Object
}

// SKAttributeValueFrom constructs a [SKAttributeValue] from an unsafe.Pointer.
//
// A container for dynamic shader data associated with a node.
func SKAttributeValueFrom(ptr unsafe.Pointer) SKAttributeValue {
	return SKAttributeValue{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKAttributeValueClass) Alloc() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKAttributeValueClass) New() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKAttributeValue) Init() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKAttributeValue) Autorelease() SKAttributeValue {
	rv := objc.Send[SKAttributeValue](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAttributeValue creates a new SKAttributeValue instance.
func NewSKAttributeValue() SKAttributeValue {
	return sKAttributeValueClass.New()
}


// Creates and initializes a new attribute value object that holds a floating point number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(float:)
func NewValueWithFloat(value float32) SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sKAttributeValueClass.class), objc.Sel("valueWithFloat:"), value)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of two floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat2:)
func NewValueWithVectorFloat2(value unsafe.Pointer) SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sKAttributeValueClass.class), objc.Sel("valueWithVectorFloat2:"), value)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of three floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat3:)
func NewValueWithVectorFloat3(value unsafe.Pointer) SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sKAttributeValueClass.class), objc.Sel("valueWithVectorFloat3:"), value)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of four floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat4:)
func NewValueWithVectorFloat4(value unsafe.Pointer) SKAttributeValue {
	rv := objc.Send[SKAttributeValue](objc.ID(sKAttributeValueClass.class), objc.Sel("valueWithVectorFloat4:"), value)
	rv.Autorelease()
	return rv
}


// Creates and initializes a new attribute value object that holds a floating point number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(float:)
func (sc _SKAttributeValueClass) ValueWithFloat(value float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("valueWithFloat:"), value)
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of two floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat2:)
func (sc _SKAttributeValueClass) ValueWithVectorFloat2(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("valueWithVectorFloat2:"), value)
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of three floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat3:)
func (sc _SKAttributeValueClass) ValueWithVectorFloat3(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("valueWithVectorFloat3:"), value)
	return rv
}
// Creates and initializes a new attribute value object that holds a vector of four floating point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttributeValue/init(vectorFloat4:)
func (sc _SKAttributeValueClass) ValueWithVectorFloat4(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("valueWithVectorFloat4:"), value)
	return rv
}

