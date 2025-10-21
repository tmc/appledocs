// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [IMKCandidates] class.
var (
	IMKCandidatesClass     _IMKCandidatesClass
	IMKCandidatesClassOnce sync.Once
)

func getIMKCandidatesClass() _IMKCandidatesClass {
	IMKCandidatesClassOnce.Do(func() {
		IMKCandidatesClass = _IMKCandidatesClass{objc.GetClass("IMKCandidates")}
	})
	return IMKCandidatesClass
}

type _IMKCandidatesClass struct {
	class objc.Class
}

// An interface definition for the [IMKCandidates] class.
type IIMKCandidates interface {
	appkit.IResponder
	AttachChildToCandidateType(child unsafe.Pointer, candidateIdentifier int, theType unsafe.Pointer)
	Attributes() unsafe.Pointer
	CandidateFrame() Rect
	CandidateIdentifierAtLineNumber(lineNumber int) int
	CandidateStringIdentifier(candidateString objc.ID) int
	ClearSelection()
	DetachChild(candidateIdentifier int)
	DismissesAutomatically() bool
	Hide()
	HideChild()
	IsVisible() bool
	LineNumberForCandidateWithIdentifier(candidateIdentifier int) int
	PanelType() unsafe.Pointer
	SelectCandidate(candidateIdentifier int)
	SelectCandidateWithIdentifier(candidateIdentifier int) bool
	SelectedCandidate() int
	SelectedCandidateString() unsafe.Pointer
	SelectionKeys() unsafe.Pointer
	SelectionKeysKeylayout() unsafe.Pointer
	SetAttributes(attributes objc.ID)
	SetCandidateData(candidatesArray objc.ID)
	SetCandidateFrameTopLeft(point Point)
	SetDismissesAutomatically(flag bool)
	SetPanelType(panelType unsafe.Pointer)
	SetSelectionKeys(keyCodes objc.ID)
	SetSelectionKeysKeylayout(layout unsafe.Pointer)
	ShowCandidates()
	Show(locationHint unsafe.Pointer)
	ShowAnnotation(annotationString unsafe.Pointer)
	ShowChild()
	ShowSublistSubListDelegate(candidates objc.ID, delegate objc.ID)
	UpdateCandidates()
}

// The class presents candidates to users and notifies the appropriate object when the user selects a candidate. are alternate characters for a given input sequence. The class supports using a candidates window in your input method; using is optional. Not all input methods require them.
//
// When you create an object, you attach it to the object for your input method. You then need to override the methods and as well as implement a candidates method in your delegate object. The subclass supplies candidates to the object by implementing the candidates method. When you are ready to display a candidates window, call the candidates method to update candidates and to show the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates
type IMKCandidates struct {
	appkit.Responder
}

// IMKCandidatesFrom constructs a [IMKCandidates] from an unsafe.Pointer.
//
// The class presents candidates to users and notifies the appropriate object when the user selects a candidate. are alternate characters for a given input sequence. The class supports using a candidates window in your input method; using is optional. Not all input methods require them.
func IMKCandidatesFrom(ptr unsafe.Pointer) IMKCandidates {
	return IMKCandidates{
		Responder: appkit.ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IMKCandidatesClass) Alloc() IMKCandidates {
	rv := objc.Send[IMKCandidates](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IMKCandidatesClass) New() IMKCandidates {
	rv := objc.Send[IMKCandidates](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IMKCandidates) Init() IMKCandidates {
	rv := objc.Send[IMKCandidates](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IMKCandidates) Autorelease() IMKCandidates {
	rv := objc.Send[IMKCandidates](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKCandidates creates a new IMKCandidates instance.
func NewIMKCandidates() IMKCandidates {
	return getIMKCandidatesClass().New()
}




// Returns the initialized object.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:)
func NewIMKCandidatesWithServerPanelType(server unsafe.Pointer, panelType unsafe.Pointer) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[IMKCandidates](instance.ID, objc.Sel("initWithServer:panelType:"), server, panelType)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:styleType:)
func NewIMKCandidatesWithServerPanelTypeStyleType(server unsafe.Pointer, panelType unsafe.Pointer, style unsafe.Pointer) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[IMKCandidates](instance.ID, objc.Sel("initWithServer:panelType:styleType:"), server, panelType, style)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attachChild(_:toCandidate:type:)
func (i_ IMKCandidates) AttachChildToCandidateType(child unsafe.Pointer, candidateIdentifier int, theType unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("attachChild:toCandidate:type:"), child, candidateIdentifier, theType)
}

// Returns a dictionary of the style attributes used for the candidates window..
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attributes()
func (i_ IMKCandidates) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("attributes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateFrame()
func (i_ IMKCandidates) CandidateFrame() Rect {
	rv := objc.Send[Rect](i_.ID, objc.Sel("candidateFrame"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateIdentifier(atLineNumber:)
func (i_ IMKCandidates) CandidateIdentifierAtLineNumber(lineNumber int) int {
	rv := objc.Send[int](i_.ID, objc.Sel("candidateIdentifierAtLineNumber:"), lineNumber)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateStringIdentifier(_:)
func (i_ IMKCandidates) CandidateStringIdentifier(candidateString objc.ID) int {
	rv := objc.Send[int](i_.ID, objc.Sel("candidateStringIdentifier:"), candidateString)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/clearSelection()
func (i_ IMKCandidates) ClearSelection() {
	objc.Send[objc.ID](i_.ID, objc.Sel("clearSelection"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/detachChild(_:)
func (i_ IMKCandidates) DetachChild(candidateIdentifier int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("detachChild:"), candidateIdentifier)
}

// Returns the state of the flag that determines whether the candidates window dismisses automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/dismissesAutomatically()
func (i_ IMKCandidates) DismissesAutomatically() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("dismissesAutomatically"))
	return rv
}

// Hides a candidates window, if it is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hide()
func (i_ IMKCandidates) Hide() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hide"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hideChild()
func (i_ IMKCandidates) HideChild() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hideChild"))
}

// Returns whether or not the candidates window is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/isVisible()
func (i_ IMKCandidates) IsVisible() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isVisible"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/lineNumberForCandidate(withIdentifier:)
func (i_ IMKCandidates) LineNumberForCandidateWithIdentifier(candidateIdentifier int) int {
	rv := objc.Send[int](i_.ID, objc.Sel("lineNumberForCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}

// Returns the style of the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/panelType()
func (i_ IMKCandidates) PanelType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("panelType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(_:)
func (i_ IMKCandidates) SelectCandidate(candidateIdentifier int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("selectCandidate:"), candidateIdentifier)
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(withIdentifier:)
func (i_ IMKCandidates) SelectCandidateWithIdentifier(candidateIdentifier int) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("selectCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidate()
func (i_ IMKCandidates) SelectedCandidate() int {
	rv := objc.Send[int](i_.ID, objc.Sel("selectedCandidate"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidateString()
func (i_ IMKCandidates) SelectedCandidateString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectedCandidateString"))
	return rv
}

// Returns an array of objects where each object represents a virtual key code.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeys()
func (i_ IMKCandidates) SelectionKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectionKeys"))
	return rv
}

// Returns the key layout that maps virtual key codes to selection keys.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeysKeylayout()
func (i_ IMKCandidates) SelectionKeysKeylayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectionKeysKeylayout"))
	return rv
}

// Sets the style attributes for the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setAttributes(_:)
func (i_ IMKCandidates) SetAttributes(attributes objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttributes:"), attributes)
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateData(_:)
func (i_ IMKCandidates) SetCandidateData(candidatesArray objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCandidateData:"), candidatesArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateFrameTopLeft(_:)
func (i_ IMKCandidates) SetCandidateFrameTopLeft(point Point) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCandidateFrameTopLeft:"), point)
}

// Sets the state of the flag that determines whether the candidates window dismisses automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setDismissesAutomatically(_:)
func (i_ IMKCandidates) SetDismissesAutomatically(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDismissesAutomatically:"), flag)
}

// Sets the style of the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setPanelType(_:)
func (i_ IMKCandidates) SetPanelType(panelType unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPanelType:"), panelType)
}

// Sets the selection keys for the candidates.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeys(_:)
func (i_ IMKCandidates) SetSelectionKeys(keyCodes objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectionKeys:"), keyCodes)
}

// Sets the key layout that is used to map virtual key codes to characters.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeysKeylayout(_:)
func (i_ IMKCandidates) SetSelectionKeysKeylayout(layout unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectionKeysKeylayout:"), layout)
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show()
func (i_ IMKCandidates) ShowCandidates() {
	objc.Send[objc.ID](i_.ID, objc.Sel("showCandidates"))
}

// Shows the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show(_:)
func (i_ IMKCandidates) Show(locationHint unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("show:"), locationHint)
}

// Displays an annotation string in an annotation window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showAnnotation(_:)
func (i_ IMKCandidates) ShowAnnotation(annotationString unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("showAnnotation:"), annotationString)
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showChild()
func (i_ IMKCandidates) ShowChild() {
	objc.Send[objc.ID](i_.ID, objc.Sel("showChild"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showSublist(_:subListDelegate:)
func (i_ IMKCandidates) ShowSublistSubListDelegate(candidates objc.ID, delegate objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("showSublist:subListDelegate:"), candidates, delegate)
}

// Updates the candidates that are displayed in the candidates window.
//
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/update()
func (i_ IMKCandidates) UpdateCandidates() {
	objc.Send[objc.ID](i_.ID, objc.Sel("updateCandidates"))
}


