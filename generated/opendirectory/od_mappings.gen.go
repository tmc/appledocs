// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODMappings */


/* debug [class_header]: Header for ODMappings */
// The class instance for the [ODMappings] class.
var (
	ODMappingsClass     _ODMappingsClass
	ODMappingsClassOnce sync.Once
)

func getODMappingsClass() _ODMappingsClass {
	ODMappingsClassOnce.Do(func() {
		ODMappingsClass = _ODMappingsClass{objc.GetClass("ODMappings")}
	})
	return ODMappingsClass
}

type _ODMappingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODMappings */
// An interface definition for the [ODMappings] class.
type IODMappings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODMappings */
	// properties:
	Comment() objc.IObject /* cross-framework: NSString */
	SetComment(value objc.IObject /* cross-framework: NSString */)
	Function() objc.IObject /* cross-framework: NSString */
	SetFunction(value objc.IObject /* cross-framework: NSString */)
	FunctionAttributes() objc.IObject /* cross-framework: NSArray */
	SetFunctionAttributes(value objc.IObject /* cross-framework: NSArray */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	RecordTypes() objc.IObject /* cross-framework: NSArray */
	TemplateName() objc.IObject /* cross-framework: NSString */
	SetTemplateName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODMappings */
	// methods:
	RecordMapForStandardRecordType(stdType objc.IObject /* cross-framework: NSString */) IODRecordMap
	SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODMappings */
// Alloc allocates a new instance without initialization.
func (oc _ODMappingsClass) Alloc() ODMappings {
	rv := objc.Send[ODMappings](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ODMappingsClass) New() ODMappings {
	rv := objc.Send[ODMappings](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODMappings) Init() ODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODMappings) Autorelease() ODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODMappings creates a new ODMappings instance.
func NewODMappings() ODMappings {
	return getODMappingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODMappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings
type ODMappings struct {
	objectivec.Object
}

// ODMappingsFrom constructs a [ODMappings] from an unsafe.Pointer.
func ODMappingsFrom(ptr unsafe.Pointer) ODMappings {
	return ODMappings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODMappings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODMappings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/mappings
func (oc _ODMappingsClass) Mappings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("mappings"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Mappings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODMappings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODMappings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordMap(forStandardRecordType:)
func (o_ ODMappings) RecordMapForStandardRecordType(stdType objc.IObject /* cross-framework: NSString */) IODRecordMap {
	rv := objc.Send[ODRecordMap](o_.ID, objc.Sel("recordMapForStandardRecordType:"), stdType)
	return rv
}/* debug [instance_methods/method]: RecordMapForStandardRecordType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/setRecordMap(_:forStandardRecordType:)
func (o_ ODMappings) SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRecordMap:forStandardRecordType:"), map_, stdType)
}/* debug [instance_methods/method]: SetRecordMapForStandardRecordType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODMappings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) Comment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("comment"))
	return rv
}/* debug [instance_properties/getter]: comment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) SetComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), value)
}/* debug [instance_properties/setter]: comment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) Function() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("function"))
	return rv
}/* debug [instance_properties/getter]: function */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) SetFunction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunction:"), value)
}/* debug [instance_properties/setter]: function */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) FunctionAttributes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("functionAttributes"))
	return rv
}/* debug [instance_properties/getter]: functionAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) SetFunctionAttributes(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunctionAttributes:"), value)
}/* debug [instance_properties/setter]: functionAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-swift.property
func (o_ ODMappings) RecordTypes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("recordTypes"))
	return rv
}/* debug [instance_properties/getter]: recordTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) TemplateName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("templateName"))
	return rv
}/* debug [instance_properties/getter]: templateName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) SetTemplateName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), value)
}/* debug [instance_properties/setter]: templateName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODMappings */



