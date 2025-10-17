// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExtensionItem] class.
var extensionItemClass = _ExtensionItemClass{objc.GetClass("NSExtensionItem")}

type _ExtensionItemClass struct {
	class objc.Class
}

// An interface definition for the [ExtensionItem] class.
type IExtensionItem interface {
	objectivec.IObject
}

// An immutable collection of values representing different aspects of an item for an extension to act upon. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem

type ExtensionItem struct {
	objectivec.Object
}

// ExtensionItemFrom constructs a [ExtensionItem] from an unsafe.Pointer.
//
// An immutable collection of values representing different aspects of an item for an extension to act upon.
func ExtensionItemFrom(ptr unsafe.Pointer) ExtensionItem {
	return ExtensionItem{objectivec.Object{objc.ID(ptr)}}
}



