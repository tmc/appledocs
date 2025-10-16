
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prompt] class.
var promptClass _promptClass

func init() {
	promptClass = _promptClass{objc.GetClass("prompt")}
}

type _promptClass struct {
	objc.Class
}

// An interface definition for the [prompt] class.
type Iprompt interface {
	ID() objc.ID
}

type prompt struct {
	id objc.ID
}

func promptFrom(ptr unsafe.Pointer) prompt {
	return prompt{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prompt) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _promptClass) Alloc() prompt {
	rv := objc.Send[prompt](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _promptClass) New() prompt {
	rv := objc.Send[prompt](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newprompt creates and returns a new initialized instance.
func Newprompt() prompt {
	return promptClass.New()
}

// Init initializes the instance.
func (p_ prompt) Init() prompt {
	rv := objc.Send[prompt](p_.ID(), selInit)
	return rv
}
