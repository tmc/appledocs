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
	dictionaryControllerKeyValuePairClass     _DictionaryControllerKeyValuePairClass
	dictionaryControllerKeyValuePairClassOnce sync.Once
)

func getDictionaryControllerKeyValuePairClass() _DictionaryControllerKeyValuePairClass {
	dictionaryControllerKeyValuePairClassOnce.Do(func() {
		dictionaryControllerKeyValuePairClass = _DictionaryControllerKeyValuePairClass{objc.GetClass("NSDictionaryControllerKeyValuePair")}
	})
	return dictionaryControllerKeyValuePairClass
}

type _DictionaryControllerKeyValuePairClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	objectivec.IObject
}

// A set of methods implemented by arranged objects to give access to information about those objects. [Full Topic]
//
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

// Alloc allocates a new instance without initialization.
func (dc _DictionaryControllerKeyValuePairClass) Alloc() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




