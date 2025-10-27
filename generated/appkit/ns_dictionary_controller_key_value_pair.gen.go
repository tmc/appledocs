// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	objectivec.IObject
	

	// properties:
	ExplicitlyIncluded() bool
	Key() foundation.foundation.INSString
	SetKey(value foundation.foundation.INSString)
	LocalizedKey() foundation.foundation.INSString
	SetLocalizedKey(value foundation.foundation.INSString)
	Value() objc.ID
	SetValue(value objc.ID)
	IsExplicitlyIncluded() bool
	SetIsExplicitlyIncluded(value bool)


	

	// methods:


}





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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/isExplicitlyIncluded
func (d_ DictionaryControllerKeyValuePair) ExplicitlyIncluded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("explicitlyIncluded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/key
func (d_ DictionaryControllerKeyValuePair) Key() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("key"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/key
func (d_ DictionaryControllerKeyValuePair) SetKey(value foundation.foundation.INSString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/localizedKey
func (d_ DictionaryControllerKeyValuePair) LocalizedKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("localizedKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/localizedKey
func (d_ DictionaryControllerKeyValuePair) SetLocalizedKey(value foundation.foundation.INSString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/value
func (d_ DictionaryControllerKeyValuePair) Value() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryControllerKeyValuePair/value
func (d_ DictionaryControllerKeyValuePair) SetValue(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/isexplicitlyincluded
func (d_ DictionaryControllerKeyValuePair) IsExplicitlyIncluded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isExplicitlyIncluded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/isexplicitlyincluded
func (d_ DictionaryControllerKeyValuePair) SetIsExplicitlyIncluded(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsExplicitlyIncluded:"), value)
}








