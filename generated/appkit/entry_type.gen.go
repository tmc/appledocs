
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [entryType] class.
var entryTypeClass _entryTypeClass

func init() {
	entryTypeClass = _entryTypeClass{objc.GetClass("entryType")}
}

type _entryTypeClass struct {
	objc.Class
}

// An interface definition for the [entryType] class.
type IentryType interface {
	ID() objc.ID
}

type entryType struct {
	id objc.ID
}

func entryTypeFrom(ptr unsafe.Pointer) entryType {
	return entryType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ entryType) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _entryTypeClass) Alloc() entryType {
	rv := objc.Send[entryType](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _entryTypeClass) New() entryType {
	rv := objc.Send[entryType](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewentryType creates and returns a new initialized instance.
func NewentryType() entryType {
	return entryTypeClass.New()
}

// Init initializes the instance.
func (e_ entryType) Init() entryType {
	rv := objc.Send[entryType](e_.ID(), selInit)
	return rv
}
