// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContentKeySpecifier] class.
var (
	ContentKeySpecifierClass     _ContentKeySpecifierClass
	ContentKeySpecifierClassOnce sync.Once
)

func getContentKeySpecifierClass() _ContentKeySpecifierClass {
	ContentKeySpecifierClassOnce.Do(func() {
		ContentKeySpecifierClass = _ContentKeySpecifierClass{objc.GetClass("AVContentKeySpecifier")}
	})
	return ContentKeySpecifierClass
}

type _ContentKeySpecifierClass struct {
	class objc.Class
}

// An interface definition for the [ContentKeySpecifier] class.
type IContentKeySpecifier interface {
	objectivec.IObject
}

// An object that uniquely identifies a content key.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier
type ContentKeySpecifier struct {
	objectivec.Object
}

// ContentKeySpecifierFrom constructs a [ContentKeySpecifier] from an unsafe.Pointer.
//
// An object that uniquely identifies a content key.
func ContentKeySpecifierFrom(ptr unsafe.Pointer) ContentKeySpecifier {
	return ContentKeySpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentKeySpecifierClass) Alloc() ContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentKeySpecifierClass) New() ContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentKeySpecifier) Init() ContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentKeySpecifier) Autorelease() ContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentKeySpecifier creates a new ContentKeySpecifier instance.
func NewContentKeySpecifier() ContentKeySpecifier {
	return getContentKeySpecifierClass().New()
}


// The container and protocol-specific key identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/identifier
func (c_ ContentKeySpecifier) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The container and protocol-specific key identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/identifier
func (c_ ContentKeySpecifier) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}

// The key system that generates content keys.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/keysystem
func (c_ ContentKeySpecifier) KeySystem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("keySystem"))
	return rv
}


// SetKeySystem sets the value of the keySystem property.
// The key system that generates content keys.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/keysystem
func (c_ ContentKeySpecifier) SetKeySystem(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySystem:"), value)
}

// A dictionary of options with which you initialized the specifier.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/options
func (c_ ContentKeySpecifier) Options() string {
	rv := objc.Send[string](c_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// A dictionary of options with which you initialized the specifier.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/options
func (c_ ContentKeySpecifier) SetOptions(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptions:"), objc.String(value))
}



