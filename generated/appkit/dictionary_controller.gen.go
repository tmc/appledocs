// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DictionaryController] class.
var (
	dictionaryControllerClass     _DictionaryControllerClass
	dictionaryControllerClassOnce sync.Once
)

func getDictionaryControllerClass() _DictionaryControllerClass {
	dictionaryControllerClassOnce.Do(func() {
		dictionaryControllerClass = _DictionaryControllerClass{objc.GetClass("NSDictionaryController")}
	})
	return dictionaryControllerClass
}

type _DictionaryControllerClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryController] class.
type IDictionaryController interface {
	IArrayController
}

// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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




