
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsContentTypes] class.
var showsContentTypesClass _showsContentTypesClass

func init() {
	showsContentTypesClass = _showsContentTypesClass{objc.GetClass("showsContentTypes")}
}

type _showsContentTypesClass struct {
	objc.Class
}

// An interface definition for the [showsContentTypes] class.
type IshowsContentTypes interface {
	ID() objc.ID
}

type showsContentTypes struct {
	id objc.ID
}

func showsContentTypesFrom(ptr unsafe.Pointer) showsContentTypes {
	return showsContentTypes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsContentTypes) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsContentTypesClass) Alloc() showsContentTypes {
	rv := objc.Send[showsContentTypes](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsContentTypesClass) New() showsContentTypes {
	rv := objc.Send[showsContentTypes](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsContentTypes creates and returns a new initialized instance.
func NewshowsContentTypes() showsContentTypes {
	return showsContentTypesClass.New()
}

// Init initializes the instance.
func (s_ showsContentTypes) Init() showsContentTypes {
	rv := objc.Send[showsContentTypes](s_.ID(), selInit)
	return rv
}
