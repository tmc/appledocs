
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [printJobTitle] class.
var printJobTitleClass _printJobTitleClass

func init() {
	printJobTitleClass = _printJobTitleClass{objc.GetClass("printJobTitle")}
}

type _printJobTitleClass struct {
	objc.Class
}

// An interface definition for the [printJobTitle] class.
type IprintJobTitle interface {
	ID() objc.ID
}

type printJobTitle struct {
	id objc.ID
}

func printJobTitleFrom(ptr unsafe.Pointer) printJobTitle {
	return printJobTitle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ printJobTitle) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _printJobTitleClass) Alloc() printJobTitle {
	rv := objc.Send[printJobTitle](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _printJobTitleClass) New() printJobTitle {
	rv := objc.Send[printJobTitle](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprintJobTitle creates and returns a new initialized instance.
func NewprintJobTitle() printJobTitle {
	return printJobTitleClass.New()
}

// Init initializes the instance.
func (p_ printJobTitle) Init() printJobTitle {
	rv := objc.Send[printJobTitle](p_.ID(), selInit)
	return rv
}
