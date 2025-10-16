
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [string] class.
var stringClass _stringClass

func init() {
	stringClass = _stringClass{objc.GetClass("string")}
}

type _stringClass struct {
	objc.Class
}

// An interface definition for the [string] class.
type Istring interface {
	ID() objc.ID
}

type string struct {
	id objc.ID
}

func stringFrom(ptr unsafe.Pointer) string {
	return string{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ string) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _stringClass) Alloc() string {
	rv := objc.Send[string](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _stringClass) New() string {
	rv := objc.Send[string](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newstring creates and returns a new initialized instance.
func Newstring() string {
	return stringClass.New()
}

// Init initializes the instance.
func (s_ string) Init() string {
	rv := objc.Send[string](s_.ID(), selInit)
	return rv
}
