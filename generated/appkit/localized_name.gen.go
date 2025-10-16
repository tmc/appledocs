
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [localizedName] class.
var localizedNameClass _localizedNameClass

func init() {
	localizedNameClass = _localizedNameClass{objc.GetClass("localizedName")}
}

type _localizedNameClass struct {
	objc.Class
}

// An interface definition for the [localizedName] class.
type IlocalizedName interface {
	ID() objc.ID
}

type localizedName struct {
	id objc.ID
}

func localizedNameFrom(ptr unsafe.Pointer) localizedName {
	return localizedName{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ localizedName) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _localizedNameClass) Alloc() localizedName {
	rv := objc.Send[localizedName](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _localizedNameClass) New() localizedName {
	rv := objc.Send[localizedName](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlocalizedName creates and returns a new initialized instance.
func NewlocalizedName() localizedName {
	return localizedNameClass.New()
}

// Init initializes the instance.
func (l_ localizedName) Init() localizedName {
	rv := objc.Send[localizedName](l_.ID(), selInit)
	return rv
}
