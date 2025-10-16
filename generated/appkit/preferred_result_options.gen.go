
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [preferredResultOptions] class.
var preferredResultOptionsClass _preferredResultOptionsClass

func init() {
	preferredResultOptionsClass = _preferredResultOptionsClass{objc.GetClass("preferredResultOptions")}
}

type _preferredResultOptionsClass struct {
	objc.Class
}

// An interface definition for the [preferredResultOptions] class.
type IpreferredResultOptions interface {
	ID() objc.ID
}

type preferredResultOptions struct {
	id objc.ID
}

func preferredResultOptionsFrom(ptr unsafe.Pointer) preferredResultOptions {
	return preferredResultOptions{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ preferredResultOptions) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _preferredResultOptionsClass) Alloc() preferredResultOptions {
	rv := objc.Send[preferredResultOptions](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _preferredResultOptionsClass) New() preferredResultOptions {
	rv := objc.Send[preferredResultOptions](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreferredResultOptions creates and returns a new initialized instance.
func NewpreferredResultOptions() preferredResultOptions {
	return preferredResultOptionsClass.New()
}

// Init initializes the instance.
func (p_ preferredResultOptions) Init() preferredResultOptions {
	rv := objc.Send[preferredResultOptions](p_.ID(), selInit)
	return rv
}
