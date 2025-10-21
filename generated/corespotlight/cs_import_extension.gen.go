// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CSImportExtension] class.
var (
	CSImportExtensionClass     _CSImportExtensionClass
	CSImportExtensionClassOnce sync.Once
)

func getCSImportExtensionClass() _CSImportExtensionClass {
	CSImportExtensionClassOnce.Do(func() {
		CSImportExtensionClass = _CSImportExtensionClass{objc.GetClass("CSImportExtension")}
	})
	return CSImportExtensionClass
}

type _CSImportExtensionClass struct {
	class objc.Class
}

// An interface definition for the [CSImportExtension] class.
type ICSImportExtension interface {
	objectivec.IObject
	UpdateAttributesForFileAtURLError(attributes unsafe.Pointer, contentURL unsafe.Pointer, error_ unsafe.Pointer) bool
}

// An object that provides searchable attributes for file types that the app supports.
//
// To create a Spotlight File Importer extension, add a target to your app using the Spotlight File Importer template in Xcode. The template project contains a subclass of . To index content on a user’s device, Core Spotlight loads your extension and invokes the method. Core Spotlight passes a and URL of a file to the extension, and you set properties that are relevant for the file. Typically, your extension loads details about the file and uses that information to set properties of the attribute set. For example, if your app contains files that are notes the user creates, it does the following: To specify the file types your app supports, set the value of in your extension’s file to an array of file type identifiers. For more information about file type identifiers, see . The app in the previous example configures the extension’s as follows:
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSImportExtension
type CSImportExtension struct {
	objectivec.Object
}

// CSImportExtensionFrom constructs a [CSImportExtension] from an unsafe.Pointer.
//
// An object that provides searchable attributes for file types that the app supports.
func CSImportExtensionFrom(ptr unsafe.Pointer) CSImportExtension {
	return CSImportExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSImportExtensionClass) Alloc() CSImportExtension {
	rv := objc.Send[CSImportExtension](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSImportExtensionClass) New() CSImportExtension {
	rv := objc.Send[CSImportExtension](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSImportExtension) Init() CSImportExtension {
	rv := objc.Send[CSImportExtension](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSImportExtension) Autorelease() CSImportExtension {
	rv := objc.Send[CSImportExtension](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSImportExtension creates a new CSImportExtension instance.
func NewCSImportExtension() CSImportExtension {
	return getCSImportExtensionClass().New()
}


// Provides searchable attributes for a file at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSImportExtension/update(_:forFileAt:)
func (c_ CSImportExtension) UpdateAttributesForFileAtURLError(attributes unsafe.Pointer, contentURL unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("updateAttributes:forFileAtURL:error:"), attributes, contentURL, error_)
	return rv
}



