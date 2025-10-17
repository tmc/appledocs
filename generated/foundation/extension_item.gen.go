// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ExtensionItem] class.
var ExtensionItemClass objc.Class

func init() {
	ExtensionItemClass = objc.GetClass("NSExtensionItem")
}

type ExtensionItem struct {
	objc.ID
}

func ExtensionItemFrom(ptr unsafe.Pointer) ExtensionItem {
	return ExtensionItem{
		ID: objc.ID(ptr),
	}
}



