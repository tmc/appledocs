
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canCreateDirectories] class.
var canCreateDirectoriesClass _canCreateDirectoriesClass

func init() {
	canCreateDirectoriesClass = _canCreateDirectoriesClass{objc.GetClass("canCreateDirectories")}
}

type _canCreateDirectoriesClass struct {
	objc.Class
}

// An interface definition for the [canCreateDirectories] class.
type IcanCreateDirectories interface {
	ID() objc.ID
}

type canCreateDirectories struct {
	id objc.ID
}

func canCreateDirectoriesFrom(ptr unsafe.Pointer) canCreateDirectories {
	return canCreateDirectories{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canCreateDirectories) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canCreateDirectoriesClass) Alloc() canCreateDirectories {
	rv := objc.Send[canCreateDirectories](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canCreateDirectoriesClass) New() canCreateDirectories {
	rv := objc.Send[canCreateDirectories](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanCreateDirectories creates and returns a new initialized instance.
func NewcanCreateDirectories() canCreateDirectories {
	return canCreateDirectoriesClass.New()
}

// Init initializes the instance.
func (c_ canCreateDirectories) Init() canCreateDirectories {
	rv := objc.Send[canCreateDirectories](c_.ID(), selInit)
	return rv
}
