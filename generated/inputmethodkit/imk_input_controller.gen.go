// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [IMKInputController] class.
var (
	IMKInputControllerClass     _IMKInputControllerClass
	IMKInputControllerClassOnce sync.Once
)

func getIMKInputControllerClass() _IMKInputControllerClass {
	IMKInputControllerClassOnce.Do(func() {
		IMKInputControllerClass = _IMKInputControllerClass{objc.GetClass("IMKInputController")}
	})
	return IMKInputControllerClass
}

type _IMKInputControllerClass struct {
	class objc.Class
}

// An interface definition for the [IMKInputController] class.
type IIMKInputController interface {
	objectivec.IObject
	AnnotationSelectedForCandidate(annotationString unsafe.Pointer, candidateString unsafe.Pointer)
	CancelComposition()
	CandidateSelected(candidateString unsafe.Pointer)
	CandidateSelectionChanged(candidateString unsafe.Pointer)
	Client() objc.ID
	CompositionAttributesAtRange(range_ Range) unsafe.Pointer
	Delegate() objc.ID
	DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary objc.ID)
	HidePalettes()
	InputControllerWillClose()
	MarkForStyleAtRange(style int, range_ Range) unsafe.Pointer
	Menu() unsafe.Pointer
	ReplacementRange() Range
	SelectionRange() Range
	Server() unsafe.Pointer
	SetDelegate(newDelegate objc.ID)
	UpdateComposition()
}

// The class provides a base class for custom input controller classes. The class, which is allocated in the main function of an input method, creates an input controller object for each input session created by a client application. For every input session there is a corresponding object.
//
// An object controls text input on the input method side. It manages events and text from the applications and converted text from the input method engine. implements fully the and protocols. Typically you do not need to override this class, but you do need to provide a delegate object that implements the methods that your are interested in. The versions of the protocol methods check whether the delegate object implements a method, and calls the delegate version if it exists.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController
type IMKInputController struct {
	objectivec.Object
}

// IMKInputControllerFrom constructs a [IMKInputController] from an unsafe.Pointer.
//
// The class provides a base class for custom input controller classes. The class, which is allocated in the main function of an input method, creates an input controller object for each input session created by a client application. For every input session there is a corresponding object.
func IMKInputControllerFrom(ptr unsafe.Pointer) IMKInputController {
	return IMKInputController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IMKInputControllerClass) Alloc() IMKInputController {
	rv := objc.Send[IMKInputController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IMKInputControllerClass) New() IMKInputController {
	rv := objc.Send[IMKInputController](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IMKInputController) Init() IMKInputController {
	rv := objc.Send[IMKInputController](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IMKInputController) Autorelease() IMKInputController {
	rv := objc.Send[IMKInputController](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKInputController creates a new IMKInputController instance.
func NewIMKInputController() IMKInputController {
	return getIMKInputControllerClass().New()
}




// Initializes the input control by setting the delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/init(server:delegate:client:)
func NewIMKInputControllerWithServerDelegateClient(server unsafe.Pointer, delegate objc.ID, inputClient objc.ID) IMKInputController {
	instance := getIMKInputControllerClass().Alloc()
	rv := objc.Send[IMKInputController](instance.ID, objc.Sel("initWithServer:delegate:client:"), server, delegate, inputClient)
	rv.Autorelease()
	return rv
}


// Sends the selected candidate string and annotation string to the input controller.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/annotationSelected(_:forCandidate:)
func (i_ IMKInputController) AnnotationSelectedForCandidate(annotationString unsafe.Pointer, candidateString unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("annotationSelected:forCandidate:"), annotationString, candidateString)
}

// Stops the current composition and replaces marked text with the original text.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/cancelComposition()
func (i_ IMKInputController) CancelComposition() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelComposition"))
}

// Informs an input controller that a new candidate is selected.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelected(_:)
func (i_ IMKInputController) CandidateSelected(candidateString unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("candidateSelected:"), candidateString)
}

// Informs an input controller that the current candidate selection in the candidate window has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelectionChanged(_:)
func (i_ IMKInputController) CandidateSelectionChanged(candidateString unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("candidateSelectionChanged:"), candidateString)
}

// Returns the client object associated with the input controller.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/client()
func (i_ IMKInputController) Client() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("client"))
	return rv
}

// Returns a dictionary of text attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/compositionAttributes(at:)
func (i_ IMKInputController) CompositionAttributesAtRange(range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("compositionAttributesAtRange:"), range_)
	return rv
}

// Returns the delegate for input controller object.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/delegate()
func (i_ IMKInputController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}

// Passes commands that are not generated as part of the text input process.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/doCommand(by:command:)
func (i_ IMKInputController) DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("doCommandBySelector:commandDictionary:"), aSelector, infoDictionary)
}

// Informs an input method that it should close any visible user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/hidePalettes()
func (i_ IMKInputController) HidePalettes() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hidePalettes"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/inputControllerWillClose()
func (i_ IMKInputController) InputControllerWillClose() {
	objc.Send[objc.ID](i_.ID, objc.Sel("inputControllerWillClose"))
}

// Returns a dictionary of text attributes that can mark a range of an attributed string to send to a client.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/mark(forStyle:at:)
func (i_ IMKInputController) MarkForStyleAtRange(style int, range_ Range) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("markForStyle:atRange:"), style, range_)
	return rv
}

// Returns a menu of commands that are specific to an input method.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/menu()
func (i_ IMKInputController) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("menu"))
	return rv
}

// Returns the range in the client document that the text should replace.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/replacementRange()
func (i_ IMKInputController) ReplacementRange() Range {
	rv := objc.Send[Range](i_.ID, objc.Sel("replacementRange"))
	return rv
}

// Returns where the range of the selection that should be placed inside marked text.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/selectionRange()
func (i_ IMKInputController) SelectionRange() Range {
	rv := objc.Send[Range](i_.ID, objc.Sel("selectionRange"))
	return rv
}

// Returns the server object that manages the input controller.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/server()
func (i_ IMKInputController) Server() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("server"))
	return rv
}

// Sets the delegate for input controller object.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/setDelegate(_:)
func (i_ IMKInputController) SetDelegate(newDelegate objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), newDelegate)
}

// Informs the input controller that the composition has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/updateComposition()
func (i_ IMKInputController) UpdateComposition() {
	objc.Send[objc.ID](i_.ID, objc.Sel("updateComposition"))
}


