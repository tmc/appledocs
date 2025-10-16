
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numberOfSections] class.
var numberOfSectionsClass _numberOfSectionsClass

func init() {
	numberOfSectionsClass = _numberOfSectionsClass{objc.GetClass("numberOfSections")}
}

type _numberOfSectionsClass struct {
	objc.Class
}

// An interface definition for the [numberOfSections] class.
type InumberOfSections interface {
	ID() objc.ID
}

type numberOfSections struct {
	id objc.ID
}

func numberOfSectionsFrom(ptr unsafe.Pointer) numberOfSections {
	return numberOfSections{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ numberOfSections) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _numberOfSectionsClass) Alloc() numberOfSections {
	rv := objc.Send[numberOfSections](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _numberOfSectionsClass) New() numberOfSections {
	rv := objc.Send[numberOfSections](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnumberOfSections creates and returns a new initialized instance.
func NewnumberOfSections() numberOfSections {
	return numberOfSectionsClass.New()
}

// Init initializes the instance.
func (n_ numberOfSections) Init() numberOfSections {
	rv := objc.Send[numberOfSections](n_.ID(), selInit)
	return rv
}
