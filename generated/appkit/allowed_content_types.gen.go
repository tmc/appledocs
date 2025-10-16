
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowedContentTypes] class.
var allowedContentTypesClass _allowedContentTypesClass

func init() {
	allowedContentTypesClass = _allowedContentTypesClass{objc.GetClass("allowedContentTypes")}
}

type _allowedContentTypesClass struct {
	objc.Class
}

// An interface definition for the [allowedContentTypes] class.
type IallowedContentTypes interface {
	ID() objc.ID
}

type allowedContentTypes struct {
	id objc.ID
}

func allowedContentTypesFrom(ptr unsafe.Pointer) allowedContentTypes {
	return allowedContentTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowedContentTypes) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowedContentTypesClass) Alloc() allowedContentTypes {
	rv := objc.Send[allowedContentTypes](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowedContentTypesClass) New() allowedContentTypes {
	rv := objc.Send[allowedContentTypes](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowedContentTypes creates and returns a new initialized instance.
func NewallowedContentTypes() allowedContentTypes {
	return allowedContentTypesClass.New()
}

// Init initializes the instance.
func (a_ allowedContentTypes) Init() allowedContentTypes {
	rv := objc.Send[allowedContentTypes](a_.ID(), selInit)
	return rv
}
