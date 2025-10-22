// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextCheckingController] class.
var (
	TextCheckingControllerClass     _TextCheckingControllerClass
	TextCheckingControllerClassOnce sync.Once
)

func getTextCheckingControllerClass() _TextCheckingControllerClass {
	TextCheckingControllerClassOnce.Do(func() {
		TextCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}
	})
	return TextCheckingControllerClass
}

type _TextCheckingControllerClass struct {
	class objc.Class
}

// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	objectivec.IObject
	Client() unsafe.Pointer
	SetClient(value unsafe.Pointer)
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCheckingController
type TextCheckingController struct {
	objectivec.Object
}

// TextCheckingControllerFrom constructs a [TextCheckingController] from an unsafe.Pointer.
func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextCheckingControllerClass) Alloc() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextCheckingControllerClass) New() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextCheckingController) Init() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextCheckingController) Autorelease() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextCheckingController creates a new TextCheckingController instance.
func NewTextCheckingController() TextCheckingController {
	return getTextCheckingControllerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/client
func (t_ TextCheckingController) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("client"))
	return rv
}


// SetClient sets the value of the client property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/client
func (t_ TextCheckingController) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClient:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/spellcheckerdocumenttag
func (t_ TextCheckingController) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// SetSpellCheckerDocumentTag sets the value of the spellCheckerDocumentTag property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/spellcheckerdocumenttag
func (t_ TextCheckingController) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}



