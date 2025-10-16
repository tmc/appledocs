
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [configurationDictionary] class.
var configurationDictionaryClass _configurationDictionaryClass

func init() {
	configurationDictionaryClass = _configurationDictionaryClass{objc.GetClass("configurationDictionary")}
}

type _configurationDictionaryClass struct {
	objc.Class
}

// An interface definition for the [configurationDictionary] class.
type IconfigurationDictionary interface {
	ID() objc.ID
}

type configurationDictionary struct {
	id objc.ID
}

func configurationDictionaryFrom(ptr unsafe.Pointer) configurationDictionary {
	return configurationDictionary{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ configurationDictionary) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _configurationDictionaryClass) Alloc() configurationDictionary {
	rv := objc.Send[configurationDictionary](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _configurationDictionaryClass) New() configurationDictionary {
	rv := objc.Send[configurationDictionary](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewconfigurationDictionary creates and returns a new initialized instance.
func NewconfigurationDictionary() configurationDictionary {
	return configurationDictionaryClass.New()
}

// Init initializes the instance.
func (c_ configurationDictionary) Init() configurationDictionary {
	rv := objc.Send[configurationDictionary](c_.ID(), selInit)
	return rv
}
