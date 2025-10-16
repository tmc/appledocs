
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sharingType] class.
var sharingTypeClass _sharingTypeClass

func init() {
	sharingTypeClass = _sharingTypeClass{objc.GetClass("sharingType")}
}

type _sharingTypeClass struct {
	objc.Class
}

// An interface definition for the [sharingType] class.
type IsharingType interface {
	ID() objc.ID
}

type sharingType struct {
	id objc.ID
}

func sharingTypeFrom(ptr unsafe.Pointer) sharingType {
	return sharingType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sharingType) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sharingTypeClass) Alloc() sharingType {
	rv := objc.Send[sharingType](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sharingTypeClass) New() sharingType {
	rv := objc.Send[sharingType](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsharingType creates and returns a new initialized instance.
func NewsharingType() sharingType {
	return sharingTypeClass.New()
}

// Init initializes the instance.
func (s_ sharingType) Init() sharingType {
	rv := objc.Send[sharingType](s_.ID(), selInit)
	return rv
}
