// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Bundle] class.
var BundleClass objc.Class

func init() {
	BundleClass = objc.GetClass("NSBundle")
}

type Bundle struct {
	objc.ID
}

func BundleFrom(ptr unsafe.Pointer) Bundle {
	return Bundle{
		ID: objc.ID(ptr),
	}
}


// Returns the   object with which the specified class is associated. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Bundle/init(for:)
func (bc Bundle) BundleForClass(aClass objc.Class) unsafe.Pointer {
	sel := objc.RegisterName("bundleForClass:")
	ret := objc.ID(BundleClass).Send(sel, aClass)
	return unsafe.Pointer(ret)
}
// Returns the value associated with the specified key in the receiver’s information property list. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.ID {
	sel := objc.RegisterName("objectForInfoDictionaryKey:")
	ret := b_.ID.Send(sel, key)
	return ret
}

