// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSCustomAttributeKey */


/* debug [class_header]: Header for CSCustomAttributeKey */
// The class instance for the [CSCustomAttributeKey] class.
var (
	CSCustomAttributeKeyClass     _CSCustomAttributeKeyClass
	CSCustomAttributeKeyClassOnce sync.Once
)

func getCSCustomAttributeKeyClass() _CSCustomAttributeKeyClass {
	CSCustomAttributeKeyClassOnce.Do(func() {
		CSCustomAttributeKeyClass = _CSCustomAttributeKeyClass{objc.GetClass("CSCustomAttributeKey")}
	})
	return CSCustomAttributeKeyClass
}

type _CSCustomAttributeKeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSCustomAttributeKey */
// An interface definition for the [CSCustomAttributeKey] class.
type ICSCustomAttributeKey interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSCustomAttributeKey */
	// properties:
	MultiValued() bool
	Searchable() bool
	SearchableByDefault() bool
	Unique() bool
	KeyName() objc.IObject /* cross-framework: NSString */
	IsMultiValued() bool
	SetIsMultiValued(value bool)
	IsSearchable() bool
	SetIsSearchable(value bool)
	IsSearchableByDefault() bool
	SetIsSearchableByDefault(value bool)
	IsUnique() bool
	SetIsUnique(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSCustomAttributeKey */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSCustomAttributeKey */
// Alloc allocates a new instance without initialization.
func (cc _CSCustomAttributeKeyClass) Alloc() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSCustomAttributeKeyClass) New() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSCustomAttributeKey) Init() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSCustomAttributeKey) Autorelease() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSCustomAttributeKey creates a new CSCustomAttributeKey instance.
func NewCSCustomAttributeKey() CSCustomAttributeKey {
	return getCSCustomAttributeKeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSCustomAttributeKey */
// A key associated with a custom attribute for a searchable item.
//
// The class defines a key that you can associate with a custom attribute for a searchable item. Item attributes provide metadata about the item that can be indexed and displayed to users in search results. Although the Core Spotlight framework provides several predefined attributes, such as title and description, you can create a object to specify a custom attribute that makes sense in your domain.


// A key associated with a custom attribute for a searchable item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey
type CSCustomAttributeKey struct {
	objectivec.Object
}

// CSCustomAttributeKeyFrom constructs a [CSCustomAttributeKey] from an unsafe.Pointer.
//
// A key associated with a custom attribute for a searchable item.
func CSCustomAttributeKeyFrom(ptr unsafe.Pointer) CSCustomAttributeKey {
	return CSCustomAttributeKey{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSCustomAttributeKey */

// Returns a new custom attribute key with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:)
func NewCSCustomAttributeKeyWithKeyName(keyName objc.IObject /* cross-framework: NSString */) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[CSCustomAttributeKey](instance.ID, objc.Sel("initWithKeyName:"), keyName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSCustomAttributeKeyWithKeyName */


// Returns a new custom attribute key with the specified name and properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:searchable:searchableByDefault:unique:multiValued:)
func NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(keyName objc.IObject /* cross-framework: NSString */, searchable bool, searchableByDefault bool, unique bool, multiValued bool) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[CSCustomAttributeKey](instance.ID, objc.Sel("initWithKeyName:searchable:searchableByDefault:unique:multiValued:"), keyName, searchable, searchableByDefault, unique, multiValued)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSCustomAttributeKey */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSCustomAttributeKey */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSCustomAttributeKey */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSCustomAttributeKey */

// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isMultiValued
func (c_ CSCustomAttributeKey) MultiValued() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multiValued"))
	return rv
}/* debug [instance_properties/getter]: multiValued */


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchable
func (c_ CSCustomAttributeKey) Searchable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("searchable"))
	return rv
}/* debug [instance_properties/getter]: searchable */


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchableByDefault
func (c_ CSCustomAttributeKey) SearchableByDefault() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("searchableByDefault"))
	return rv
}/* debug [instance_properties/getter]: searchableByDefault */


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isUnique
func (c_ CSCustomAttributeKey) Unique() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unique"))
	return rv
}/* debug [instance_properties/getter]: unique */


// The name of the custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/keyName
func (c_ CSCustomAttributeKey) KeyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keyName"))
	return rv
}/* debug [instance_properties/getter]: keyName */


// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/ismultivalued
func (c_ CSCustomAttributeKey) IsMultiValued() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultiValued"))
	return rv
}/* debug [instance_properties/getter]: isMultiValued */


// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/ismultivalued
func (c_ CSCustomAttributeKey) SetIsMultiValued(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultiValued:"), value)
}/* debug [instance_properties/setter]: isMultiValued */


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchable
func (c_ CSCustomAttributeKey) IsSearchable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSearchable"))
	return rv
}/* debug [instance_properties/getter]: isSearchable */


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchable
func (c_ CSCustomAttributeKey) SetIsSearchable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSearchable:"), value)
}/* debug [instance_properties/setter]: isSearchable */


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchablebydefault
func (c_ CSCustomAttributeKey) IsSearchableByDefault() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSearchableByDefault"))
	return rv
}/* debug [instance_properties/getter]: isSearchableByDefault */


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchablebydefault
func (c_ CSCustomAttributeKey) SetIsSearchableByDefault(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSearchableByDefault:"), value)
}/* debug [instance_properties/setter]: isSearchableByDefault */


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/isunique
func (c_ CSCustomAttributeKey) IsUnique() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUnique"))
	return rv
}/* debug [instance_properties/getter]: isUnique */


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/isunique
func (c_ CSCustomAttributeKey) SetIsUnique(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUnique:"), value)
}/* debug [instance_properties/setter]: isUnique */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSCustomAttributeKey */


