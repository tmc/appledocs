
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URL] class.
var URLClass _URLClass

func init() {
	URLClass = _URLClass{objc.GetClass("URL")}
}

type _URLClass struct {
	objc.Class
}

// An interface definition for the [URL] class.
type IURL interface {
	ID() objc.ID
}

type URL struct {
	id objc.ID
}

func URLFrom(ptr unsafe.Pointer) URL {
	return URL{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ URL) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _URLClass) Alloc() URL {
	rv := objc.Send[URL](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _URLClass) New() URL {
	rv := objc.Send[URL](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewURL creates and returns a new initialized instance.
func NewURL() URL {
	return URLClass.New()
}

// Init initializes the instance.
func (u_ URL) Init() URL {
	rv := objc.Send[URL](u_.ID(), selInit)
	return rv
}
