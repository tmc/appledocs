
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [currentContentType] class.
var currentContentTypeClass _currentContentTypeClass

func init() {
	currentContentTypeClass = _currentContentTypeClass{objc.GetClass("currentContentType")}
}

type _currentContentTypeClass struct {
	objc.Class
}

// An interface definition for the [currentContentType] class.
type IcurrentContentType interface {
	ID() objc.ID
}

type currentContentType struct {
	id objc.ID
}

func currentContentTypeFrom(ptr unsafe.Pointer) currentContentType {
	return currentContentType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ currentContentType) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _currentContentTypeClass) Alloc() currentContentType {
	rv := objc.Send[currentContentType](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _currentContentTypeClass) New() currentContentType {
	rv := objc.Send[currentContentType](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcurrentContentType creates and returns a new initialized instance.
func NewcurrentContentType() currentContentType {
	return currentContentTypeClass.New()
}

// Init initializes the instance.
func (c_ currentContentType) Init() currentContentType {
	rv := objc.Send[currentContentType](c_.ID(), selInit)
	return rv
}
