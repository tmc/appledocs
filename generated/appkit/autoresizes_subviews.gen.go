
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autoresizesSubviews] class.
var autoresizesSubviewsClass _autoresizesSubviewsClass

func init() {
	autoresizesSubviewsClass = _autoresizesSubviewsClass{objc.GetClass("autoresizesSubviews")}
}

type _autoresizesSubviewsClass struct {
	objc.Class
}

// An interface definition for the [autoresizesSubviews] class.
type IautoresizesSubviews interface {
	ID() objc.ID
}

type autoresizesSubviews struct {
	id objc.ID
}

func autoresizesSubviewsFrom(ptr unsafe.Pointer) autoresizesSubviews {
	return autoresizesSubviews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autoresizesSubviews) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autoresizesSubviewsClass) Alloc() autoresizesSubviews {
	rv := objc.Send[autoresizesSubviews](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autoresizesSubviewsClass) New() autoresizesSubviews {
	rv := objc.Send[autoresizesSubviews](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautoresizesSubviews creates and returns a new initialized instance.
func NewautoresizesSubviews() autoresizesSubviews {
	return autoresizesSubviewsClass.New()
}

// Init initializes the instance.
func (a_ autoresizesSubviews) Init() autoresizesSubviews {
	rv := objc.Send[autoresizesSubviews](a_.ID(), selInit)
	return rv
}
