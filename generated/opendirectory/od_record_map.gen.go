// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Attributes() objc.IObject /* cross-framework: NSDictionary */
	Native() objc.IObject /* cross-framework: NSString */
	SetNative(value objc.IObject /* cross-framework: NSString */)
	OdPredicate() objc.IObject /* cross-framework: NSDictionary */
	SetOdPredicate(value objc.IObject /* cross-framework: NSDictionary */)
	StandardAttributeTypes() objc.IObject /* cross-framework: NSArray */
	// methods:
	AttributeMapForStandardAttribute(standardAttribute objc.IObject /* cross-framework: NSString */) IODAttributeMap
	SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute objc.IObject /* cross-framework: NSString */)
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/recordMap
func (oc _ODRecordMapClass) RecordMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("recordMap"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributeMap(forStandardAttribute:)
func (o_ ODRecordMap) AttributeMapForStandardAttribute(standardAttribute objc.IObject /* cross-framework: NSString */) IODAttributeMap {
	rv := objc.Send[ODAttributeMap](o_.ID, objc.Sel("attributeMapForStandardAttribute:"), standardAttribute)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/setAttribute(_:forStandardAttribute:)
func (o_ ODRecordMap) SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAttributeMap:forStandardAttribute:"), attributeMap, standardAttribute)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-swift.property
func (o_ ODRecordMap) Attributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](o_.ID, objc.Sel("attributes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) Native() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("native"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) SetNative(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNative:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) OdPredicate() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](o_.ID, objc.Sel("odPredicate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) SetOdPredicate(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOdPredicate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/standardAttributeTypes
func (o_ ODRecordMap) StandardAttributeTypes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("standardAttributeTypes"))
	return rv
}



