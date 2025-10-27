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
	

	// properties:
	Identifier() objc.ID
	KeySystem() ContentKeySystem
	Options() foundation.IDictionary


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ContentKeySpecifierClass) Alloc() ContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that uniquely identifies a content key.


// An object that uniquely identifies a content key.
//
// [Full Topic]
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






// Creates a content key specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/init(forKeySystem:identifier:options:)
func NewContentKeySpecifierForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.IDictionary) ContentKeySpecifier {
	instance := getContentKeySpecifierClass().Alloc()
	rv := objc.Send[ContentKeySpecifier](instance.ID, objc.Sel("initForKeySystem:identifier:options:"), keySystem, contentKeyIdentifier, options)
	rv.Autorelease()
	return rv
}







// A convenience initializer to create a content key specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/contentKeySpecifierForKeySystem:identifier:options:
func (cc _ContentKeySpecifierClass) ContentKeySpecifierForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contentKeySpecifierForKeySystem:identifier:options:"), keySystem, contentKeyIdentifier, options)
	return rv
}

















// The container and protocol-specific key identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/identifier
func (c_ ContentKeySpecifier) Identifier() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("identifier"))
	return rv
}


// The key system that generates content keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/keySystem
func (c_ ContentKeySpecifier) KeySystem() ContentKeySystem {
	rv := objc.Send[ContentKeySystem](c_.ID, objc.Sel("keySystem"))
	return rv
}


// A dictionary of options with which you initialized the specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/options
func (c_ ContentKeySpecifier) Options() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("options"))
	return rv
}







