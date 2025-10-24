// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODRecordMap */


/* debug [class_header]: Header for ODRecordMap */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODRecordMap */
// An interface definition for the [ODRecordMap] class.
type IODRecordMap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODRecordMap */
	// properties:
	Attributes() objc.IObject /* cross-framework: NSDictionary */
	Native() objc.IObject /* cross-framework: NSString */
	SetNative(value objc.IObject /* cross-framework: NSString */)
	OdPredicate() objc.IObject /* cross-framework: NSDictionary */
	SetOdPredicate(value objc.IObject /* cross-framework: NSDictionary */)
	StandardAttributeTypes() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODRecordMap */
	// methods:
	AttributeMapForStandardAttribute(standardAttribute objc.IObject /* cross-framework: NSString */) IODAttributeMap
	SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODRecordMap */
// Alloc allocates a new instance without initialization.
func (oc _ODRecordMapClass) Alloc() ODRecordMap {
	rv := objc.Send[ODRecordMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODRecordMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap
type ODRecordMap struct {
	objectivec.Object
}

// ODRecordMapFrom constructs a [ODRecordMap] from an unsafe.Pointer.
func ODRecordMapFrom(ptr unsafe.Pointer) ODRecordMap {
	return ODRecordMap{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODRecordMap *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODRecordMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/recordMap
func (oc _ODRecordMapClass) RecordMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("recordMap"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecordMap) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODRecordMap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODRecordMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributeMap(forStandardAttribute:)
func (o_ ODRecordMap) AttributeMapForStandardAttribute(standardAttribute objc.IObject /* cross-framework: NSString */) IODAttributeMap {
	rv := objc.Send[ODAttributeMap](o_.ID, objc.Sel("attributeMapForStandardAttribute:"), standardAttribute)
	return rv
}/* debug [instance_methods/method]: AttributeMapForStandardAttribute */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/setAttribute(_:forStandardAttribute:)
func (o_ ODRecordMap) SetAttributeMapForStandardAttribute(attributeMap IODAttributeMap, standardAttribute objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAttributeMap:forStandardAttribute:"), attributeMap, standardAttribute)
}/* debug [instance_methods/method]: SetAttributeMapForStandardAttribute */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODRecordMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-swift.property
func (o_ ODRecordMap) Attributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](o_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) Native() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("native"))
	return rv
}/* debug [instance_properties/getter]: native */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-swift.property
func (o_ ODRecordMap) SetNative(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNative:"), value)
}/* debug [instance_properties/setter]: native */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) OdPredicate() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](o_.ID, objc.Sel("odPredicate"))
	return rv
}/* debug [instance_properties/getter]: odPredicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-swift.property
func (o_ ODRecordMap) SetOdPredicate(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOdPredicate:"), value)
}/* debug [instance_properties/setter]: odPredicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/standardAttributeTypes
func (o_ ODRecordMap) StandardAttributeTypes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("standardAttributeTypes"))
	return rv
}/* debug [instance_properties/getter]: standardAttributeTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODRecordMap */



