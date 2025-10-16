
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxFullScreenContentSize] class.
var maxFullScreenContentSizeClass _maxFullScreenContentSizeClass

func init() {
	maxFullScreenContentSizeClass = _maxFullScreenContentSizeClass{objc.GetClass("maxFullScreenContentSize")}
}

type _maxFullScreenContentSizeClass struct {
	objc.Class
}

// An interface definition for the [maxFullScreenContentSize] class.
type ImaxFullScreenContentSize interface {
	ID() objc.ID
}

type maxFullScreenContentSize struct {
	id objc.ID
}

func maxFullScreenContentSizeFrom(ptr unsafe.Pointer) maxFullScreenContentSize {
	return maxFullScreenContentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxFullScreenContentSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxFullScreenContentSizeClass) Alloc() maxFullScreenContentSize {
	rv := objc.Send[maxFullScreenContentSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxFullScreenContentSizeClass) New() maxFullScreenContentSize {
	rv := objc.Send[maxFullScreenContentSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxFullScreenContentSize creates and returns a new initialized instance.
func NewmaxFullScreenContentSize() maxFullScreenContentSize {
	return maxFullScreenContentSizeClass.New()
}

// Init initializes the instance.
func (m_ maxFullScreenContentSize) Init() maxFullScreenContentSize {
	rv := objc.Send[maxFullScreenContentSize](m_.ID(), selInit)
	return rv
}
