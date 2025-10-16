
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [treatsFilePackagesAsDirectories] class.
var treatsFilePackagesAsDirectoriesClass _treatsFilePackagesAsDirectoriesClass

func init() {
	treatsFilePackagesAsDirectoriesClass = _treatsFilePackagesAsDirectoriesClass{objc.GetClass("treatsFilePackagesAsDirectories")}
}

type _treatsFilePackagesAsDirectoriesClass struct {
	objc.Class
}

// An interface definition for the [treatsFilePackagesAsDirectories] class.
type ItreatsFilePackagesAsDirectories interface {
	ID() objc.ID
}

type treatsFilePackagesAsDirectories struct {
	id objc.ID
}

func treatsFilePackagesAsDirectoriesFrom(ptr unsafe.Pointer) treatsFilePackagesAsDirectories {
	return treatsFilePackagesAsDirectories{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ treatsFilePackagesAsDirectories) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _treatsFilePackagesAsDirectoriesClass) Alloc() treatsFilePackagesAsDirectories {
	rv := objc.Send[treatsFilePackagesAsDirectories](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _treatsFilePackagesAsDirectoriesClass) New() treatsFilePackagesAsDirectories {
	rv := objc.Send[treatsFilePackagesAsDirectories](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtreatsFilePackagesAsDirectories creates and returns a new initialized instance.
func NewtreatsFilePackagesAsDirectories() treatsFilePackagesAsDirectories {
	return treatsFilePackagesAsDirectoriesClass.New()
}

// Init initializes the instance.
func (t_ treatsFilePackagesAsDirectories) Init() treatsFilePackagesAsDirectories {
	rv := objc.Send[treatsFilePackagesAsDirectories](t_.ID(), selInit)
	return rv
}
