// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CSCustomAttributeKey] class.
type ICSCustomAttributeKey interface {
	objectivec.IObject
	MultiValued() bool
	Searchable() bool
	SearchableByDefault() bool
	Unique() bool
	KeyName() string
	IsMultiValued() bool
	SetIsMultiValued(value bool)
	IsSearchable() bool
	SetIsSearchable(value bool)
	IsSearchableByDefault() bool
	SetIsSearchableByDefault(value bool)
	IsUnique() bool
	SetIsUnique(value bool)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CSCustomAttributeKeyClass) Alloc() CSCustomAttributeKey {
	rv := objc.Send[CSCustomAttributeKey](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a new custom attribute key with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:)
func NewCSCustomAttributeKeyWithKeyName(keyName string) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[CSCustomAttributeKey](instance.ID, objc.Sel("initWithKeyName:"), objc.String(keyName))
	rv.Autorelease()
	return rv
}


// Returns a new custom attribute key with the specified name and properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/init(keyName:searchable:searchableByDefault:unique:multiValued:)
func NewCSCustomAttributeKeyWithKeyNameSearchableSearchableByDefaultUniqueMultiValued(keyName string, searchable bool, searchableByDefault bool, unique bool, multiValued bool) CSCustomAttributeKey {
	instance := getCSCustomAttributeKeyClass().Alloc()
	rv := objc.Send[CSCustomAttributeKey](instance.ID, objc.Sel("initWithKeyName:searchable:searchableByDefault:unique:multiValued:"), objc.String(keyName), searchable, searchableByDefault, unique, multiValued)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isMultiValued
func (c_ CSCustomAttributeKey) MultiValued() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multiValued"))
	return rv
}


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchable
func (c_ CSCustomAttributeKey) Searchable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("searchable"))
	return rv
}


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isSearchableByDefault
func (c_ CSCustomAttributeKey) SearchableByDefault() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("searchableByDefault"))
	return rv
}


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/isUnique
func (c_ CSCustomAttributeKey) Unique() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unique"))
	return rv
}


// The name of the custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSCustomAttributeKey/keyName
func (c_ CSCustomAttributeKey) KeyName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("keyName"))
	return rv
}


// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/ismultivalued
func (c_ CSCustomAttributeKey) IsMultiValued() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultiValued"))
	return rv
}


// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/ismultivalued
func (c_ CSCustomAttributeKey) SetIsMultiValued(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultiValued:"), value)
}


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchable
func (c_ CSCustomAttributeKey) IsSearchable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSearchable"))
	return rv
}


// A Boolean value that indicates if the custom attribute can be specified as a search term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchable
func (c_ CSCustomAttributeKey) SetIsSearchable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSearchable:"), value)
}


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchablebydefault
func (c_ CSCustomAttributeKey) IsSearchableByDefault() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSearchableByDefault"))
	return rv
}


// A Boolean value that indicates if the custom attribute should be searchable by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/issearchablebydefault
func (c_ CSCustomAttributeKey) SetIsSearchableByDefault(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSearchableByDefault:"), value)
}


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/isunique
func (c_ CSCustomAttributeKey) IsUnique() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUnique"))
	return rv
}


// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/isunique
func (c_ CSCustomAttributeKey) SetIsUnique(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUnique:"), value)
}


