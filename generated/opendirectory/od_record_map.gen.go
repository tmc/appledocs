// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ODRecordMap] class.
var (
	ODRecordMapClass     _ODRecordMapClass
	ODRecordMapClassOnce sync.Once
)

func getODRecordMapClass() _ODRecordMapClass {
	ODRecordMapClassOnce.Do(func() {
		ODRecordMapClass = _ODRecordMapClass{objc.GetClass("ODRecordMap")}
	})
	return ODRecordMapClass
}

type _ODRecordMapClass struct {
	class objc.Class
}

// An interface definition for the [ODRecordMap] class.
type IODRecordMap interface {
	objectivec.IObject
	AttributeMapForStandardAttribute(standardAttribute string) unsafe.Pointer
	SetAttributeMapForStandardAttribute(attributeMap unsafe.Pointer, standardAttribute string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap
type ODRecordMap struct {
	objectivec.Object
}

// ODRecordMapFrom constructs a [ODRecordMap] from an unsafe.Pointer.
func ODRecordMapFrom(ptr unsafe.Pointer) ODRecordMap {
	return ODRecordMap{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODRecordMapClass) Alloc() ODRecordMap {
	rv := objc.Send[ODRecordMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODRecordMapClass) New() ODRecordMap {
	rv := objc.Send[ODRecordMap](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODRecordMap) Init() ODRecordMap {
	rv := objc.Send[ODRecordMap](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODRecordMap) Autorelease() ODRecordMap {
	rv := objc.Send[ODRecordMap](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODRecordMap creates a new ODRecordMap instance.
func NewODRecordMap() ODRecordMap {
	return getODRecordMapClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/recordMap
func (oc _ODRecordMapClass) RecordMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("recordMap"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributeMap(forStandardAttribute:)
func (o_ ODRecordMap) AttributeMapForStandardAttribute(standardAttribute string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("attributeMapForStandardAttribute:"), objc.String(standardAttribute))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/setAttribute(_:forStandardAttribute:)
func (o_ ODRecordMap) SetAttributeMapForStandardAttribute(attributeMap unsafe.Pointer, standardAttribute string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAttributeMap:forStandardAttribute:"), attributeMap, objc.String(standardAttribute))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-swift.property
func (o_ ODRecordMap) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("attributes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) Native() string {
	rv := objc.Send[string](o_.ID, objc.Sel("native"))
	return rv
}


// SetNative sets the value of the native property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) SetNative(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNative:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) OdPredicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("odPredicate"))
	return rv
}


// SetOdPredicate sets the value of the odPredicate property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) SetOdPredicate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOdPredicate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/standardAttributeTypes
func (o_ ODRecordMap) StandardAttributeTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("standardAttributeTypes"))
	return rv
}



