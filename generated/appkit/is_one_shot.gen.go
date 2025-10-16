
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isOneShot] class.
var isOneShotClass _isOneShotClass

func init() {
	isOneShotClass = _isOneShotClass{objc.GetClass("isOneShot")}
}

type _isOneShotClass struct {
	objc.Class
}

// An interface definition for the [isOneShot] class.
type IisOneShot interface {
	ID() objc.ID
}

type isOneShot struct {
	id objc.ID
}

func isOneShotFrom(ptr unsafe.Pointer) isOneShot {
	return isOneShot{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isOneShot) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isOneShotClass) Alloc() isOneShot {
	rv := objc.Send[isOneShot](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isOneShotClass) New() isOneShot {
	rv := objc.Send[isOneShot](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisOneShot creates and returns a new initialized instance.
func NewisOneShot() isOneShot {
	return isOneShotClass.New()
}

// Init initializes the instance.
func (i_ isOneShot) Init() isOneShot {
	rv := objc.Send[isOneShot](i_.ID(), selInit)
	return rv
}
