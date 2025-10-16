
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [content] class.
var contentClass _contentClass

func init() {
	contentClass = _contentClass{objc.GetClass("content")}
}

type _contentClass struct {
	objc.Class
}

// An interface definition for the [content] class.
type Icontent interface {
	ID() objc.ID
}

type content struct {
	id objc.ID
}

func contentFrom(ptr unsafe.Pointer) content {
	return content{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ content) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentClass) Alloc() content {
	rv := objc.Send[content](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentClass) New() content {
	rv := objc.Send[content](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newcontent creates and returns a new initialized instance.
func Newcontent() content {
	return contentClass.New()
}

// Init initializes the instance.
func (c_ content) Init() content {
	rv := objc.Send[content](c_.ID(), selInit)
	return rv
}
