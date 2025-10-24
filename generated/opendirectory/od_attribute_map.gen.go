// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODAttributeMap */


/* debug [class_header]: Header for ODAttributeMap */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODAttributeMap */
// An interface definition for the [ODAttributeMap] class.
type IODAttributeMap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODAttributeMap */
	// properties:
	CustomAttributes() objc.IObject /* cross-framework: NSArray */
	SetCustomAttributes(value objc.IObject /* cross-framework: NSArray */)
	CustomQueryFunction() objc.IObject /* cross-framework: NSString */
	SetCustomQueryFunction(value objc.IObject /* cross-framework: NSString */)
	CustomTranslationFunction() objc.IObject /* cross-framework: NSString */
	SetCustomTranslationFunction(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODAttributeMap */
	// methods:
	SetStaticValue(staticValue objc.IObject /* cross-framework: NSString */)
	SetVariableSubstitution(variableSubstitution objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODAttributeMap */
// Alloc allocates a new instance without initialization.
func (oc _ODAttributeMapClass) Alloc() ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODAttributeMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap
type ODAttributeMap struct {
	objectivec.Object
}

// ODAttributeMapFrom constructs a [ODAttributeMap] from an unsafe.Pointer.
func ODAttributeMapFrom(ptr unsafe.Pointer) ODAttributeMap {
	return ODAttributeMap{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODAttributeMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(staticValue:)
func NewODAttributeMapWithStaticValue(staticValue objc.IObject /* cross-framework: NSString */) ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithStaticValue:"), staticValue)
	return rv
}/* debug [class_init_methods/constructor]: NewODAttributeMapWithStaticValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(value:)
func NewODAttributeMapWithValue(value objc.IObject /* cross-framework: NSString */) ODAttributeMap {
	rv := objc.Send[ODAttributeMap](objc.ID(getODAttributeMapClass().class), objc.Sel("attributeMapWithValue:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewODAttributeMapWithValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODAttributeMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(staticValue:)
func (oc _ODAttributeMapClass) AttributeMapWithStaticValue(staticValue objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("attributeMapWithStaticValue:"), staticValue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributeMapWithStaticValue) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/init(value:)
func (oc _ODAttributeMapClass) AttributeMapWithValue(value objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("attributeMapWithValue:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AttributeMapWithValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODAttributeMap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODAttributeMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setStaticValue(_:)
func (o_ ODAttributeMap) SetStaticValue(staticValue objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStaticValue:"), staticValue)
}/* debug [instance_methods/method]: SetStaticValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/setVariableSubstitution(_:)
func (o_ ODAttributeMap) SetVariableSubstitution(variableSubstitution objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVariableSubstitution:"), variableSubstitution)
}/* debug [instance_methods/method]: SetVariableSubstitution */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODAttributeMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-swift.property
func (o_ ODAttributeMap) CustomAttributes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("customAttributes"))
	return rv
}/* debug [instance_properties/getter]: customAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customAttributes-swift.property
func (o_ ODAttributeMap) SetCustomAttributes(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomAttributes:"), value)
}/* debug [instance_properties/setter]: customAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-swift.property
func (o_ ODAttributeMap) CustomQueryFunction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customQueryFunction"))
	return rv
}/* debug [instance_properties/getter]: customQueryFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customQueryFunction-swift.property
func (o_ ODAttributeMap) SetCustomQueryFunction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomQueryFunction:"), value)
}/* debug [instance_properties/setter]: customQueryFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-swift.property
func (o_ ODAttributeMap) CustomTranslationFunction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customTranslationFunction"))
	return rv
}/* debug [instance_properties/getter]: customTranslationFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/customTranslationFunction-swift.property
func (o_ ODAttributeMap) SetCustomTranslationFunction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomTranslationFunction:"), value)
}/* debug [instance_properties/setter]: customTranslationFunction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-swift.property
func (o_ ODAttributeMap) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-swift.property
func (o_ ODAttributeMap) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODAttributeMap */


