// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	// methods:
	AnnotationSelectedForCandidate(annotationString objc.IObject /* cross-framework: AttributedString */, candidateString objc.IObject /* cross-framework: AttributedString */)
	CancelComposition()
	CandidateSelected(candidateString objc.IObject /* cross-framework: AttributedString */)
	CandidateSelectionChanged(candidateString objc.IObject /* cross-framework: AttributedString */)
	Client() objc.ID
	CompositionAttributesAtRange(range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: MutableDictionary */
	Delegate() objc.ID
	DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary objc.IObject /* cross-framework: NSDictionary */)
	HidePalettes()
	InputControllerWillClose()
	MarkForStyleAtRange(style int, range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Dictionary */
	Menu() objc.IObject /* cross-framework: Menu */
	ReplacementRange() objc.IObject /* cross-framework: Range */
	SelectionRange() objc.IObject /* cross-framework: Range */
	Server() IMKServer
	SetDelegate(newDelegate objectivec.IObject)
	UpdateComposition()
}

// The class provides a base class for custom input controller classes. The class, which is allocated in the main function of an input method, creates an input controller object for each input session created by a client application. For every input session there is a corresponding object.
//
// An object controls text input on the input method side. It manages events and text from the applications and converted text from the input method engine. implements fully the and protocols. Typically you do not need to override this class, but you do need to provide a delegate object that implements the methods that your are interested in. The versions of the protocol methods check whether the delegate object implements a method, and calls the delegate version if it exists.


// The class provides a base class for custom input controller classes. The class, which is allocated in the main function of an input method, creates an input controller object for each input session created by a client application. For every input session there is a corresponding object.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/init(server:delegate:client:)
func NewIMKInputControllerWithServerDelegateClient(server IMKServer, delegate objectivec.IObject, inputClient objectivec.IObject) IMKInputController {
	instance := getIMKInputControllerClass().Alloc()
	rv := objc.Send[IMKInputController](instance.ID, objc.Sel("initWithServer:delegate:client:"), server, delegate, inputClient)
	rv.Autorelease()
	return rv
}



// Sends the selected candidate string and annotation string to the input controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/annotationSelected(_:forCandidate:)
func (i_ IMKInputController) AnnotationSelectedForCandidate(annotationString objc.IObject /* cross-framework: AttributedString */, candidateString objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("annotationSelected:forCandidate:"), annotationString, candidateString)
}


// Stops the current composition and replaces marked text with the original text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/cancelComposition()
func (i_ IMKInputController) CancelComposition() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelComposition"))
}


// Informs an input controller that a new candidate is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelected(_:)
func (i_ IMKInputController) CandidateSelected(candidateString objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("candidateSelected:"), candidateString)
}


// Informs an input controller that the current candidate selection in the candidate window has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelectionChanged(_:)
func (i_ IMKInputController) CandidateSelectionChanged(candidateString objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("candidateSelectionChanged:"), candidateString)
}


// Returns the client object associated with the input controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/client()
func (i_ IMKInputController) Client() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("client"))
	return rv
}


// Returns a dictionary of text attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/compositionAttributes(at:)
func (i_ IMKInputController) CompositionAttributesAtRange(range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: MutableDictionary */ {
	rv := objc.Send[foundation.MutableDictionary](i_.ID, objc.Sel("compositionAttributesAtRange:"), range_)
	return rv
}


// Returns the delegate for input controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/delegate()
func (i_ IMKInputController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}


// Passes commands that are not generated as part of the text input process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/doCommand(by:command:)
func (i_ IMKInputController) DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("doCommandBySelector:commandDictionary:"), aSelector, infoDictionary)
}


// Informs an input method that it should close any visible user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/hidePalettes()
func (i_ IMKInputController) HidePalettes() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hidePalettes"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/inputControllerWillClose()
func (i_ IMKInputController) InputControllerWillClose() {
	objc.Send[objc.ID](i_.ID, objc.Sel("inputControllerWillClose"))
}


// Returns a dictionary of text attributes that can mark a range of an attributed string to send to a client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/mark(forStyle:at:)
func (i_ IMKInputController) MarkForStyleAtRange(style int, range_ objc.IObject /* cross-framework: Range */) objc.IObject /* cross-framework: Dictionary */ {
	rv := objc.Send[foundation.Dictionary](i_.ID, objc.Sel("markForStyle:atRange:"), style, range_)
	return rv
}


// Returns a menu of commands that are specific to an input method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/menu()
func (i_ IMKInputController) Menu() objc.IObject /* cross-framework: Menu */ {
	rv := objc.Send[appkit.Menu](i_.ID, objc.Sel("menu"))
	return rv
}


// Returns the range in the client document that the text should replace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/replacementRange()
func (i_ IMKInputController) ReplacementRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](i_.ID, objc.Sel("replacementRange"))
	return rv
}


// Returns where the range of the selection that should be placed inside marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/selectionRange()
func (i_ IMKInputController) SelectionRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[corefoundation.Range](i_.ID, objc.Sel("selectionRange"))
	return rv
}


// Returns the server object that manages the input controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/server()
func (i_ IMKInputController) Server() IMKServer {
	rv := objc.Send[IMKServer](i_.ID, objc.Sel("server"))
	return rv
}


// Sets the delegate for input controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/setDelegate(_:)
func (i_ IMKInputController) SetDelegate(newDelegate objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), newDelegate)
}


// Informs the input controller that the composition has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/updateComposition()
func (i_ IMKInputController) UpdateComposition() {
	objc.Send[objc.ID](i_.ID, objc.Sel("updateComposition"))
}


