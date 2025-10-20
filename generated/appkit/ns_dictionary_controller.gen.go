// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DictionaryController] class.
var (
	DictionaryControllerClass     _DictionaryControllerClass
	DictionaryControllerClassOnce sync.Once
)

func getDictionaryControllerClass() _DictionaryControllerClass {
	DictionaryControllerClassOnce.Do(func() {
		DictionaryControllerClass = _DictionaryControllerClass{objc.GetClass("NSDictionaryController")}
	})
	return DictionaryControllerClass
}

type _DictionaryControllerClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryController] class.
type IDictionaryController interface {
	IArrayController
}

// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs.
//
// transforms the contents of a dictionary into an array of key-value pairs that can be bound to user interface items such as the columns of an . The content of an instance is specified using the inherited method or by binding an instance to the binding. New key/value pairs inserted into the dictionary are created using the method. The initial key name is set to the string returned by . The initial key name is copied to the newly inserted object, while the object returned by is simply retained. As new items are inserted the controller enumerates the initial key name, resulting in key names such as “key”, “key1”, “key2”, and so on. This behavior can be customized by overriding . An instance can be configured to exclude specified keys in a dictionary from being returned by using the property. Similarly, you can specify an array of key names that are always included in the arranged objects, even if they are not present in the content dictionary, using the property. supports providing localized key names for the keys in the dictionary, allowing a user-friendly representation of the key name to be displayed. The localized key names are specified by a dictionary (using ) or by providing a strings table (using ). The method returns an array of objects that implement the informal protocol. User interface controls are bound to the arranged objects array using key paths such as: (displays the key name), (displays the value for the key), or (displays the localized key name). See for more information. overrides to return an array of objects that implement the informal protocol. See and for more information. The constants listed below are used to specify a binding to , , , and . See the for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDictionaryController
type DictionaryController struct {
	ArrayController
}

// DictionaryControllerFrom constructs a [DictionaryController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs.
func DictionaryControllerFrom(ptr unsafe.Pointer) DictionaryController {
	return DictionaryController{
		ArrayController: ArrayControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryControllerClass) Alloc() DictionaryController {
	rv := objc.Send[DictionaryController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DictionaryControllerClass) New() DictionaryController {
	rv := objc.Send[DictionaryController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DictionaryController) Init() DictionaryController {
	rv := objc.Send[DictionaryController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DictionaryController) Autorelease() DictionaryController {
	rv := objc.Send[DictionaryController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionaryController creates a new DictionaryController instance.
func NewDictionaryController() DictionaryController {
	return getDictionaryControllerClass().New()
}




