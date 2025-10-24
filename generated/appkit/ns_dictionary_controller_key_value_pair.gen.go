// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDictionaryControllerKeyValuePair */


/* debug [class_header]: Header for NSDictionaryControllerKeyValuePair */
// The class instance for the [DictionaryControllerKeyValuePair] class.
var (
	DictionaryControllerKeyValuePairClass     _DictionaryControllerKeyValuePairClass
	DictionaryControllerKeyValuePairClassOnce sync.Once
)

func getDictionaryControllerKeyValuePairClass() _DictionaryControllerKeyValuePairClass {
	DictionaryControllerKeyValuePairClassOnce.Do(func() {
		DictionaryControllerKeyValuePairClass = _DictionaryControllerKeyValuePairClass{objc.GetClass("NSDictionaryControllerKeyValuePair")}
	})
	return DictionaryControllerKeyValuePairClass
}

type _DictionaryControllerKeyValuePairClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DictionaryControllerKeyValuePair */
// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DictionaryControllerKeyValuePair */
	// properties:
	ExplicitlyIncluded() bool
	Key() objc.IObject /* cross-framework: NSString */
	SetKey(value objc.IObject /* cross-framework: NSString */)
	LocalizedKey() objc.IObject /* cross-framework: NSString */
	SetLocalizedKey(value objc.IObject /* cross-framework: NSString */)
	Value() objc.ID
	SetValue(value objc.ID)
	IsExplicitlyIncluded() bool
	SetIsExplicitlyIncluded(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DictionaryControllerKeyValuePair */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DictionaryControllerKeyValuePair */
// Alloc allocates a new instance without initialization.
func (dc _DictionaryControllerKeyValuePairClass) Alloc() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DictionaryControllerKeyValuePairClass) New() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DictionaryControllerKeyValuePair) Init() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DictionaryControllerKeyValuePair) Autorelease() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionaryControllerKeyValuePair creates a new DictionaryControllerKeyValuePair instance.
func NewDictionaryControllerKeyValuePair() DictionaryControllerKeyValuePair {
	return getDictionaryControllerKeyValuePairClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DictionaryControllerKeyValuePair */
// A set of methods implemented by arranged objects to give access to information about those objects.
//
// is an informal protocol that is implemented by objects returned by the method arrangedObjects. See for more information.


// A set of methods implemented by arranged objects to give access to information about those objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair
type DictionaryControllerKeyValuePair struct {
	objectivec.Object
}

// DictionaryControllerKeyValuePairFrom constructs a [DictionaryControllerKeyValuePair] from an unsafe.Pointer.
//
// A set of methods implemented by arranged objects to give access to information about those objects.
func DictionaryControllerKeyValuePairFrom(ptr unsafe.Pointer) DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePair{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DictionaryControllerKeyValuePair *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DictionaryControllerKeyValuePair */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DictionaryControllerKeyValuePair */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DictionaryControllerKeyValuePair */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DictionaryControllerKeyValuePair */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/isExplicitlyIncluded
func (d_ DictionaryControllerKeyValuePair) ExplicitlyIncluded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("explicitlyIncluded"))
	return rv
}/* debug [instance_properties/getter]: explicitlyIncluded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/key
func (d_ DictionaryControllerKeyValuePair) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/key
func (d_ DictionaryControllerKeyValuePair) SetKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKey:"), value)
}/* debug [instance_properties/setter]: key */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/localizedKey
func (d_ DictionaryControllerKeyValuePair) LocalizedKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("localizedKey"))
	return rv
}/* debug [instance_properties/getter]: localizedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/localizedKey
func (d_ DictionaryControllerKeyValuePair) SetLocalizedKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedKey:"), value)
}/* debug [instance_properties/setter]: localizedKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/value
func (d_ DictionaryControllerKeyValuePair) Value() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/value
func (d_ DictionaryControllerKeyValuePair) SetValue(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/isexplicitlyincluded
func (d_ DictionaryControllerKeyValuePair) IsExplicitlyIncluded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isExplicitlyIncluded"))
	return rv
}/* debug [instance_properties/getter]: isExplicitlyIncluded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/isexplicitlyincluded
func (d_ DictionaryControllerKeyValuePair) SetIsExplicitlyIncluded(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsExplicitlyIncluded:"), value)
}/* debug [instance_properties/setter]: isExplicitlyIncluded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDictionaryControllerKeyValuePair */



