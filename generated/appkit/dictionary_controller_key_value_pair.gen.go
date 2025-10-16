
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DictionaryControllerKeyValuePair] class.
var DictionaryControllerKeyValuePairClass _DictionaryControllerKeyValuePairClass

func init() {
	DictionaryControllerKeyValuePairClass = _DictionaryControllerKeyValuePairClass{objc.GetClass("NSDictionaryControllerKeyValuePair")}
}

type _DictionaryControllerKeyValuePairClass struct {
	objc.Class
}

// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	ID() objc.ID
}

type DictionaryControllerKeyValuePair struct {
	id objc.ID
}

func DictionaryControllerKeyValuePairFrom(ptr unsafe.Pointer) DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePair{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DictionaryControllerKeyValuePair) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryControllerKeyValuePairClass) Alloc() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DictionaryControllerKeyValuePairClass) New() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDictionaryControllerKeyValuePair creates and returns a new initialized instance.
func NewDictionaryControllerKeyValuePair() DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePairClass.New()
}

// Init initializes the instance.
func (d_ DictionaryControllerKeyValuePair) Init() DictionaryControllerKeyValuePair {
	rv := objc.Send[DictionaryControllerKeyValuePair](d_.ID(), selInit)
	return rv
}
