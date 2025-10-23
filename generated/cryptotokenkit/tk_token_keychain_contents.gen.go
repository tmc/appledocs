// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKTokenKeychainContents] class.
var (
	TKTokenKeychainContentsClass     _TKTokenKeychainContentsClass
	TKTokenKeychainContentsClassOnce sync.Once
)

func getTKTokenKeychainContentsClass() _TKTokenKeychainContentsClass {
	TKTokenKeychainContentsClassOnce.Do(func() {
		TKTokenKeychainContentsClass = _TKTokenKeychainContentsClass{objc.GetClass("TKTokenKeychainContents")}
	})
	return TKTokenKeychainContentsClass
}

type _TKTokenKeychainContentsClass struct {
	class objc.Class
}

// An interface definition for the [TKTokenKeychainContents] class.
type ITKTokenKeychainContents interface {
	objectivec.IObject
	// properties:
	KeychainContents() ITKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
	Items() unsafe.Pointer
	SetItems(value unsafe.Pointer)
	// methods:
}

// A representation of the state of the keychain for a particular token.


// A representation of the state of the keychain for a particular token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents
type TKTokenKeychainContents struct {
	objectivec.Object
}

// TKTokenKeychainContentsFrom constructs a [TKTokenKeychainContents] from an unsafe.Pointer.
//
// A representation of the state of the keychain for a particular token.
func TKTokenKeychainContentsFrom(ptr unsafe.Pointer) TKTokenKeychainContents {
	return TKTokenKeychainContents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeychainContentsClass) Alloc() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKTokenKeychainContentsClass) New() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeychainContents) Init() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeychainContents) Autorelease() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainContents creates a new TKTokenKeychainContents instance.
func NewTKTokenKeychainContents() TKTokenKeychainContents {
	return getTKTokenKeychainContentsClass().New()
}



// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainContents) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainContents) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}


// Returns all items for token in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenkeychaincontents/items
func (t_ TKTokenKeychainContents) Items() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("items"))
	return rv
}


// Returns all items for token in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenkeychaincontents/items
func (t_ TKTokenKeychainContents) SetItems(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItems:"), value)
}



