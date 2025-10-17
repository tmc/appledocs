
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Movie] class.
var MovieClass _MovieClass

func init() {
	MovieClass = _MovieClass{objc.GetClass("NSMovie")}
}

type _MovieClass struct {
	objc.Class
}

// An interface definition for the [Movie] class.
type IMovie interface {
	ID() objc.ID
}

type Movie struct {
	id objc.ID
}

func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ Movie) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MovieClass) Alloc() Movie {
	rv := objc.Send[Movie](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MovieClass) New() Movie {
	rv := objc.Send[Movie](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMovie creates and returns a new initialized instance.
func NewMovie() Movie {
	return MovieClass.New()
}

// Init initializes the instance.
func (m_ Movie) Init() Movie {
	rv := objc.Send[Movie](m_.ID(), selInit)
	return rv
}
