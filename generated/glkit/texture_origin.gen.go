// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [textureOrigin] class.
var (
	TextureOriginClass     _textureOriginClass
	TextureOriginClassOnce sync.Once
)

func gettextureOriginClass() _textureOriginClass {
	TextureOriginClassOnce.Do(func() {
		TextureOriginClass = _textureOriginClass{objc.GetClass("textureOrigin")}
	})
	return TextureOriginClass
}

type _textureOriginClass struct {
	class objc.Class
}

// An interface definition for the [textureOrigin] class.
type ItextureOrigin interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/textureOrigin-c.ivar
type textureOrigin struct {
	objectivec.Object
}

// textureOriginFrom constructs a [textureOrigin] from an unsafe.Pointer.
func textureOriginFrom(ptr unsafe.Pointer) textureOrigin {
	return textureOrigin{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _textureOriginClass) Alloc() textureOrigin {
	rv := objc.Send[textureOrigin](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _textureOriginClass) New() textureOrigin {
	rv := objc.Send[textureOrigin](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ textureOrigin) Init() textureOrigin {
	rv := objc.Send[textureOrigin](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ textureOrigin) Autorelease() textureOrigin {
	rv := objc.Send[textureOrigin](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtextureOrigin creates a new textureOrigin instance.
func NewtextureOrigin() textureOrigin {
	return gettextureOriginClass().New()
}




