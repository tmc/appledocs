
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TokenField] class.
var TokenFieldClass _TokenFieldClass

func init() {
	TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}
}

type _TokenFieldClass struct {
	objc.Class
}

// An interface definition for the [TokenField] class.
type ITokenField interface {
	ID() objc.ID
}

type TokenField struct {
	id objc.ID
}

func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TokenField) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TokenFieldClass) New() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTokenField creates and returns a new initialized instance.
func NewTokenField() TokenField {
	return TokenFieldClass.New()
}

// Init initializes the instance.
func (t_ TokenField) Init() TokenField {
	rv := objc.Send[TokenField](t_.ID(), selInit)
	return rv
}
