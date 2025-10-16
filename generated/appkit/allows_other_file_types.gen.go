
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsOtherFileTypes] class.
var allowsOtherFileTypesClass _allowsOtherFileTypesClass

func init() {
	allowsOtherFileTypesClass = _allowsOtherFileTypesClass{objc.GetClass("allowsOtherFileTypes")}
}

type _allowsOtherFileTypesClass struct {
	objc.Class
}

// An interface definition for the [allowsOtherFileTypes] class.
type IallowsOtherFileTypes interface {
	ID() objc.ID
}

type allowsOtherFileTypes struct {
	id objc.ID
}

func allowsOtherFileTypesFrom(ptr unsafe.Pointer) allowsOtherFileTypes {
	return allowsOtherFileTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsOtherFileTypes) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsOtherFileTypesClass) Alloc() allowsOtherFileTypes {
	rv := objc.Send[allowsOtherFileTypes](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsOtherFileTypesClass) New() allowsOtherFileTypes {
	rv := objc.Send[allowsOtherFileTypes](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsOtherFileTypes creates and returns a new initialized instance.
func NewallowsOtherFileTypes() allowsOtherFileTypes {
	return allowsOtherFileTypesClass.New()
}

// Init initializes the instance.
func (a_ allowsOtherFileTypes) Init() allowsOtherFileTypes {
	rv := objc.Send[allowsOtherFileTypes](a_.ID(), selInit)
	return rv
}
