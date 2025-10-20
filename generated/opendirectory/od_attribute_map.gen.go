// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ODAttributeMap] class.
var (
	ODAttributeMapClass     _ODAttributeMapClass
	ODAttributeMapClassOnce sync.Once
)

func getODAttributeMapClass() _ODAttributeMapClass {
	ODAttributeMapClassOnce.Do(func() {
		ODAttributeMapClass = _ODAttributeMapClass{objc.GetClass("ODAttributeMap")}
	})
	return ODAttributeMapClass
}

type _ODAttributeMapClass struct {
	class objc.Class
}

// An interface definition for the [ODAttributeMap] class.
type IODAttributeMap interface {
	objectivec.IObject
	SetStaticValue(staticValue string)
	SetVariableSubstitution(variableSubstitution string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap
type ODAttributeMap struct {
	objectivec.Object
}

// ODAttributeMapFrom constructs a [ODAttributeMap] from an unsafe.Pointer.
func ODAttributeMapFrom(ptr unsafe.Pointer) ODAttributeMap {
	return ODAttributeMap{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODAttributeMapClass) Alloc() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODAttributeMapClass) New() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODAttributeMap) Init() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODAttributeMap) Autorelease() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODAttributeMap creates a new ODAttributeMap instance.
func NewODAttributeMap() ODAttributeMap {
	return getODAttributeMapClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(staticValue:)
func NewODAttributeMapWithStaticValue(staticValue string) ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithStaticValue:"), objc.String(staticValue))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(value:)
func NewODAttributeMapWithValue(value string) ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithValue:"), objc.String(value))
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(staticValue:)
func (oc _ODAttributeMapClass) AttributeMapWithStaticValue(staticValue string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("attributeMapWithStaticValue:"), objc.String(staticValue))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(value:)
func (oc _ODAttributeMapClass) AttributeMapWithValue(value string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("attributeMapWithValue:"), objc.String(value))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setStaticValue(_:)
func (o_ ODAttributeMap) SetStaticValue(staticValue string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStaticValue:"), objc.String(staticValue))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setVariableSubstitution(_:)
func (o_ ODAttributeMap) SetVariableSubstitution(variableSubstitution string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVariableSubstitution:"), objc.String(variableSubstitution))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-swift.property
func (o_ ODAttributeMap) CustomAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("customAttributes"))
	return rv
}


// SetCustomAttributes sets the value of the customAttributes property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-swift.property
func (o_ ODAttributeMap) SetCustomAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomAttributes:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-swift.property
func (o_ ODAttributeMap) CustomQueryFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("customQueryFunction"))
	return rv
}


// SetCustomQueryFunction sets the value of the customQueryFunction property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-swift.property
func (o_ ODAttributeMap) SetCustomQueryFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomQueryFunction:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-swift.property
func (o_ ODAttributeMap) CustomTranslationFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("customTranslationFunction"))
	return rv
}


// SetCustomTranslationFunction sets the value of the customTranslationFunction property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-swift.property
func (o_ ODAttributeMap) SetCustomTranslationFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomTranslationFunction:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-swift.property
func (o_ ODAttributeMap) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-swift.property
func (o_ ODAttributeMap) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:"), value)
}

