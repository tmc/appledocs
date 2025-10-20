// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var BundleClass _BundleClass

func init() {
	BundleClass = _BundleClass{objc.GetClass("NSBundle")}
}

type _BundleClass struct {
	class objc.Class
}

type Bundle struct {
	objc.ID
}

func BundleFrom(ptr unsafe.Pointer) Bundle {
	return Bundle{
		ID: objc.ID(ptr),
	}
}


// Returns the object with which the specified class is associated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}
// Returns the value associated with the specified key in the receiver’s information property list. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("objectForInfoDictionaryKey:"), key)
	return rv
}


