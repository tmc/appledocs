// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ContentKey] class.
var (
	ContentKeyClass     _ContentKeyClass
	ContentKeyClassOnce sync.Once
)

func getContentKeyClass() _ContentKeyClass {
	ContentKeyClassOnce.Do(func() {
		ContentKeyClass = _ContentKeyClass{objc.GetClass("AVContentKey")}
	})
	return ContentKeyClass
}

type _ContentKeyClass struct {
	class objc.Class
}





// An interface definition for the [ContentKey] class.
type IContentKey interface {
	objectivec.IObject
	

	// properties:
	ContentKeySpecifier() IAVContentKeySpecifier
	ExternalContentProtectionStatus() ExternalContentProtectionStatus


	

	// methods:
	Revoke()


}





// Alloc allocates a new instance without initialization.
func (cc _ContentKeyClass) Alloc() ContentKey {
	rv := objc.Send[ContentKey](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentKeyClass) New() ContentKey {
	rv := objc.Send[ContentKey](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentKey) Init() ContentKey {
	rv := objc.Send[ContentKey](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentKey) Autorelease() ContentKey {
	rv := objc.Send[ContentKey](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentKey creates a new ContentKey instance.
func NewContentKey() ContentKey {
	return getContentKeyClass().New()
}





// An object that represents the content key decryptor.


// An object that represents the content key decryptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKey
type ContentKey struct {
	objectivec.Object
}

// ContentKeyFrom constructs a [ContentKey] from an unsafe.Pointer.
//
// An object that represents the content key decryptor.
func ContentKeyFrom(ptr unsafe.Pointer) ContentKey {
	return ContentKey{objectivec.Object{objc.ID(ptr)}}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKey/revoke()
func (c_ ContentKey) Revoke() {
	objc.Send[objc.ID](c_.ID, objc.Sel("revoke"))
}







// The content key’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKey/contentKeySpecifier
func (c_ ContentKey) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[ContentKeySpecifier](c_.ID, objc.Sel("contentKeySpecifier"))
	return rv
}


// The external protection status for the content key based on all attached displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKey/externalContentProtectionStatus
func (c_ ContentKey) ExternalContentProtectionStatus() ExternalContentProtectionStatus {
	rv := objc.Send[ExternalContentProtectionStatus](c_.ID, objc.Sel("externalContentProtectionStatus"))
	return rv
}








