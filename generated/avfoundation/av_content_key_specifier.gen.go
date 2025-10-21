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




