// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class IMKCandidates */


/* debug [class_header]: Header for IMKCandidates */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IMKCandidates */
// An interface definition for the [IMKCandidates] class.
type IIMKCandidates interface {
	appkit.IResponder
	
/* debug [class_interface_properties]: Properties for IMKCandidates */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IMKCandidates */
	// methods:
	AttachChildToCandidateType(child IMKCandidates, candidateIdentifier int, theType IMKStyleType)
	Attributes() foundation.Dictionary
	CandidateFrame() Rect /* not a class type */
	CandidateIdentifierAtLineNumber(lineNumber int) int
	CandidateStringIdentifier(candidateString objc.IObject) int
	ClearSelection()
	DetachChild(candidateIdentifier int)
	DismissesAutomatically() bool
	Hide()
	HideChild()
	IsVisible() bool
	LineNumberForCandidateWithIdentifier(candidateIdentifier int) int
	PanelType() IMKCandidatePanelType
	SelectCandidate(candidateIdentifier int)
	SelectCandidateWithIdentifier(candidateIdentifier int) bool
	SelectedCandidate() int
	SelectedCandidateString() foundation.AttributedString
	SelectionKeys() foundation.Array
	SelectionKeysKeylayout() unsafe.Pointer
	SetAttributes(attributes objc.IObject /* cross-framework: NSDictionary */)
	SetCandidateData(candidatesArray objc.IObject /* cross-framework: NSArray */)
	SetCandidateFrameTopLeft(point vision.Point)
	SetDismissesAutomatically(flag bool)
	SetPanelType(panelType IMKCandidatePanelType)
	SetSelectionKeys(keyCodes objc.IObject /* cross-framework: NSArray */)
	SetSelectionKeysKeylayout(layout unsafe.Pointer)
	ShowCandidates()
	Show(locationHint IMKCandidatesLocationHint)
	ShowAnnotation(annotationString foundation.AttributedString)
	ShowChild()
	ShowSublistSubListDelegate(candidates objc.IObject /* cross-framework: NSArray */, delegate objc.IObject)
	UpdateCandidates()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IMKCandidates */
// Alloc allocates a new instance without initialization.
func (ic _IMKCandidatesClass) Alloc() IMKCandidates {
	rv := objc.Send[IMKCandidates](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IMKCandidates */
// The class presents candidates to users and notifies the appropriate object when the user selects a candidate. are alternate characters for a given input sequence. The class supports using a candidates window in your input method; using is optional. Not all input methods require them.
//
// When you create an object, you attach it to the object for your input method. You then need to override the methods and as well as implement a candidates method in your delegate object. The subclass supplies candidates to the object by implementing the candidates method. When you are ready to display a candidates window, call the candidates method to update candidates and to show the candidates window.


// The class presents candidates to users and notifies the appropriate object when the user selects a candidate. are alternate characters for a given input sequence. The class supports using a candidates window in your input method; using is optional. Not all input methods require them.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IMKCandidates */

// Returns the initialized object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:)
func NewIMKCandidatesWithServerPanelType(server IMKServer, panelType IMKCandidatePanelType) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[IMKCandidates](instance.ID, objc.Sel("initWithServer:panelType:"), server, panelType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIMKCandidatesWithServerPanelType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:styleType:)
func NewIMKCandidatesWithServerPanelTypeStyleType(server IMKServer, panelType IMKCandidatePanelType, style IMKStyleType) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[IMKCandidates](instance.ID, objc.Sel("initWithServer:panelType:styleType:"), server, panelType, style)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIMKCandidatesWithServerPanelTypeStyleType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IMKCandidates */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IMKCandidates */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IMKCandidates */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attachChild(_:toCandidate:type:)
func (i_ IMKCandidates) AttachChildToCandidateType(child IMKCandidates, candidateIdentifier int, theType IMKStyleType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("attachChild:toCandidate:type:"), child, candidateIdentifier, theType)
}/* debug [instance_methods/method]: AttachChildToCandidateType */


// Returns a dictionary of the style attributes used for the candidates window..
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attributes()
func (i_ IMKCandidates) Attributes() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](i_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_methods/method]: Attributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateFrame()
func (i_ IMKCandidates) CandidateFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("candidateFrame"))
	return rv
}/* debug [instance_methods/method]: CandidateFrame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateIdentifier(atLineNumber:)
func (i_ IMKCandidates) CandidateIdentifierAtLineNumber(lineNumber int) int {
	rv := objc.Send[int](i_.ID, objc.Sel("candidateIdentifierAtLineNumber:"), lineNumber)
	return rv
}/* debug [instance_methods/method]: CandidateIdentifierAtLineNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateStringIdentifier(_:)
func (i_ IMKCandidates) CandidateStringIdentifier(candidateString objc.IObject) int {
	rv := objc.Send[int](i_.ID, objc.Sel("candidateStringIdentifier:"), candidateString)
	return rv
}/* debug [instance_methods/method]: CandidateStringIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/clearSelection()
func (i_ IMKCandidates) ClearSelection() {
	objc.Send[objc.ID](i_.ID, objc.Sel("clearSelection"))
}/* debug [instance_methods/method]: ClearSelection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/detachChild(_:)
func (i_ IMKCandidates) DetachChild(candidateIdentifier int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("detachChild:"), candidateIdentifier)
}/* debug [instance_methods/method]: DetachChild */


// Returns the state of the flag that determines whether the candidates window dismisses automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/dismissesAutomatically()
func (i_ IMKCandidates) DismissesAutomatically() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("dismissesAutomatically"))
	return rv
}/* debug [instance_methods/method]: DismissesAutomatically */


// Hides a candidates window, if it is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hide()
func (i_ IMKCandidates) Hide() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hide"))
}/* debug [instance_methods/method]: Hide */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hideChild()
func (i_ IMKCandidates) HideChild() {
	objc.Send[objc.ID](i_.ID, objc.Sel("hideChild"))
}/* debug [instance_methods/method]: HideChild */


// Returns whether or not the candidates window is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/isVisible()
func (i_ IMKCandidates) IsVisible() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_methods/method]: IsVisible */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/lineNumberForCandidate(withIdentifier:)
func (i_ IMKCandidates) LineNumberForCandidateWithIdentifier(candidateIdentifier int) int {
	rv := objc.Send[int](i_.ID, objc.Sel("lineNumberForCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}/* debug [instance_methods/method]: LineNumberForCandidateWithIdentifier */


// Returns the style of the candidates window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/panelType()
func (i_ IMKCandidates) PanelType() IMKCandidatePanelType {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("panelType"))
	return rv
}/* debug [instance_methods/method]: PanelType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(_:)
func (i_ IMKCandidates) SelectCandidate(candidateIdentifier int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("selectCandidate:"), candidateIdentifier)
}/* debug [instance_methods/method]: SelectCandidate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(withIdentifier:)
func (i_ IMKCandidates) SelectCandidateWithIdentifier(candidateIdentifier int) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("selectCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}/* debug [instance_methods/method]: SelectCandidateWithIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidate()
func (i_ IMKCandidates) SelectedCandidate() int {
	rv := objc.Send[int](i_.ID, objc.Sel("selectedCandidate"))
	return rv
}/* debug [instance_methods/method]: SelectedCandidate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidateString()
func (i_ IMKCandidates) SelectedCandidateString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](i_.ID, objc.Sel("selectedCandidateString"))
	return rv
}/* debug [instance_methods/method]: SelectedCandidateString */


// Returns an array of objects where each object represents a virtual key code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeys()
func (i_ IMKCandidates) SelectionKeys() foundation.Array {
	rv := objc.Send[foundation.Array](i_.ID, objc.Sel("selectionKeys"))
	return rv
}/* debug [instance_methods/method]: SelectionKeys */


// Returns the key layout that maps virtual key codes to selection keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeysKeylayout()
func (i_ IMKCandidates) SelectionKeysKeylayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectionKeysKeylayout"))
	return rv
}/* debug [instance_methods/method]: SelectionKeysKeylayout */


// Sets the style attributes for the candidates window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setAttributes(_:)
func (i_ IMKCandidates) SetAttributes(attributes objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttributes:"), attributes)
}/* debug [instance_methods/method]: SetAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateData(_:)
func (i_ IMKCandidates) SetCandidateData(candidatesArray objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCandidateData:"), candidatesArray)
}/* debug [instance_methods/method]: SetCandidateData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateFrameTopLeft(_:)
func (i_ IMKCandidates) SetCandidateFrameTopLeft(point vision.Point) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCandidateFrameTopLeft:"), point)
}/* debug [instance_methods/method]: SetCandidateFrameTopLeft */


// Sets the state of the flag that determines whether the candidates window dismisses automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setDismissesAutomatically(_:)
func (i_ IMKCandidates) SetDismissesAutomatically(flag bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDismissesAutomatically:"), flag)
}/* debug [instance_methods/method]: SetDismissesAutomatically */


// Sets the style of the candidates window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setPanelType(_:)
func (i_ IMKCandidates) SetPanelType(panelType IMKCandidatePanelType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPanelType:"), panelType)
}/* debug [instance_methods/method]: SetPanelType */


// Sets the selection keys for the candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeys(_:)
func (i_ IMKCandidates) SetSelectionKeys(keyCodes objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectionKeys:"), keyCodes)
}/* debug [instance_methods/method]: SetSelectionKeys */


// Sets the key layout that is used to map virtual key codes to characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeysKeylayout(_:)
func (i_ IMKCandidates) SetSelectionKeysKeylayout(layout unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectionKeysKeylayout:"), layout)
}/* debug [instance_methods/method]: SetSelectionKeysKeylayout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show()
func (i_ IMKCandidates) ShowCandidates() {
	objc.Send[objc.ID](i_.ID, objc.Sel("showCandidates"))
}/* debug [instance_methods/method]: ShowCandidates */


// Shows the candidates window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show(_:)
func (i_ IMKCandidates) Show(locationHint IMKCandidatesLocationHint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("show:"), locationHint)
}/* debug [instance_methods/method]: Show */


// Displays an annotation string in an annotation window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showAnnotation(_:)
func (i_ IMKCandidates) ShowAnnotation(annotationString foundation.AttributedString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("showAnnotation:"), annotationString)
}/* debug [instance_methods/method]: ShowAnnotation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showChild()
func (i_ IMKCandidates) ShowChild() {
	objc.Send[objc.ID](i_.ID, objc.Sel("showChild"))
}/* debug [instance_methods/method]: ShowChild */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showSublist(_:subListDelegate:)
func (i_ IMKCandidates) ShowSublistSubListDelegate(candidates objc.IObject /* cross-framework: NSArray */, delegate objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("showSublist:subListDelegate:"), candidates, delegate)
}/* debug [instance_methods/method]: ShowSublistSubListDelegate */


// Updates the candidates that are displayed in the candidates window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/update()
func (i_ IMKCandidates) UpdateCandidates() {
	objc.Send[objc.ID](i_.ID, objc.Sel("updateCandidates"))
}/* debug [instance_methods/method]: UpdateCandidates */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IMKCandidates */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IMKCandidates */


