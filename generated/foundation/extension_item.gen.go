// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExtensionItem] class.
var ExtensionItemClass = _ExtensionItemClass{objc.GetClass("NSExtensionItem")}

type _ExtensionItemClass struct {
	class objc.Class
}

type ExtensionItem struct {
	objc.ID
}

func ExtensionItemFrom(ptr unsafe.Pointer) ExtensionItem {
	return ExtensionItem{
		ID: objc.ID(ptr),
	}
}




