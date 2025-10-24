// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INVocabulary] class.
var (
	INVocabularyClass     _INVocabularyClass
	INVocabularyClassOnce sync.Once
)

func getINVocabularyClass() _INVocabularyClass {
	INVocabularyClassOnce.Do(func() {
		INVocabularyClass = _INVocabularyClass{objc.GetClass("INVocabulary")}
	})
	return INVocabularyClass
}

type _INVocabularyClass struct {
	class objc.Class
}

// An interface definition for the [INVocabulary] class.
type IINVocabulary interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object for registering user-specific vocabulary that Siri requests might include.
//
// The object lets you augment your app’s global vocabulary with terms that are both unique to your app and to the current user of your app. Registering custom terms provides Siri with hints it needs to apply those terms appropriately to the corresponding intent objects. You may register custom terms only for specific types of content, including users of your app, custom workout names, or custom tags applied to a photo. Some tips for specifying custom vocabulary include: Be selective about the terms that you register for users. Include words and phrases only when their use in your app by the current user might differ from everyday usage. Order terms from most important to least important. If you register a large number of entries, Siri may ingest only the ones at the beginning of your list. Don’t register contact names that you retrieved from the user’s Contacts database. Register contacts only if your app manages contact information separately from the system databases. Don’t use this class to register terms that are common to all users of your app. Include vocabulary that’s common to all users of your app in your app’s global vocabulary file. For information about specifying your app’s global vocabulary file, see . Your Intents extension must support at least one intent that uses the registered terms. It’s a programmer error to register terms that aren’t supported by any of your extension’s intents. Use the shared vocabulary object only in your iOS app. Don’t try to register vocabulary from your Intents extension or Intents UI extension.

// An object for registering user-specific vocabulary that Siri requests might include.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INVocabulary
type INVocabulary struct {
	objectivec.Object
}

// INVocabularyFrom constructs a [INVocabulary] from an unsafe.Pointer.
//
// An object for registering user-specific vocabulary that Siri requests might include.
func INVocabularyFrom(ptr unsafe.Pointer) INVocabulary {
	return INVocabulary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INVocabularyClass) Alloc() INVocabulary {
	rv := objc.Send[INVocabulary](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INVocabularyClass) New() INVocabulary {
	rv := objc.Send[INVocabulary](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INVocabulary) Init() INVocabulary {
	rv := objc.Send[INVocabulary](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INVocabulary) Autorelease() INVocabulary {
	rv := objc.Send[INVocabulary](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINVocabulary creates a new INVocabulary instance.
func NewINVocabulary() INVocabulary {
	return getINVocabularyClass().New()
}
