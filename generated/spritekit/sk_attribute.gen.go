// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKAttribute] class.
var (
	sKAttributeClass     _SKAttributeClass
	sKAttributeClassOnce sync.Once
)

func getSKAttributeClass() _SKAttributeClass {
	sKAttributeClassOnce.Do(func() {
		sKAttributeClass = _SKAttributeClass{objc.GetClass("SKAttribute")}
	})
	return sKAttributeClass
}

type _SKAttributeClass struct {
	class objc.Class
}

// An interface definition for the [SKAttribute] class.
type ISKAttribute interface {
	objectivec.IObject
}

// A specification for dynamic per-node data used with a custom shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttribute
type SKAttribute struct {
	objectivec.Object
}

// SKAttributeFrom constructs a [SKAttribute] from an unsafe.Pointer.
//
// A specification for dynamic per-node data used with a custom shader.
func SKAttributeFrom(ptr unsafe.Pointer) SKAttribute {
	return SKAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKAttributeClass) Alloc() SKAttribute {
	rv := objc.Send[SKAttribute](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKAttributeClass) New() SKAttribute {
	rv := objc.Send[SKAttribute](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKAttribute) Init() SKAttribute {
	rv := objc.Send[SKAttribute](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKAttribute) Autorelease() SKAttribute {
	rv := objc.Send[SKAttribute](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAttribute creates a new SKAttribute instance.
func NewSKAttribute() SKAttribute {
	return getSKAttributeClass().New()
}


// Creates and initializes a new attribute object of a specified type with a name that can be referenced within the shader. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttribute/init(name:type:)
func NewSKAttributeWithNameType(name string, type_ unsafe.Pointer) SKAttribute {
	instance := getSKAttributeClass().Alloc()
	rv := objc.Send[SKAttribute](instance.ID, objc.Sel("initWithName:type:"), objc.String(name), type_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAttribute/attributeWithName:type:
func (sc _SKAttributeClass) AttributeWithNameType(name string, type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("attributeWithName:type:"), objc.String(name), type_)
	return rv
}

