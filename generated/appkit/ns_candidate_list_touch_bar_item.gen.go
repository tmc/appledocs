// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSCandidateListTouchBarItem */


/* debug [class_header]: Header for NSCandidateListTouchBarItem */
// The class instance for the [CandidateListTouchBarItem] class.
var (
	CandidateListTouchBarItemClass     _CandidateListTouchBarItemClass
	CandidateListTouchBarItemClassOnce sync.Once
)

func getCandidateListTouchBarItemClass() _CandidateListTouchBarItemClass {
	CandidateListTouchBarItemClassOnce.Do(func() {
		CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}
	})
	return CandidateListTouchBarItemClass
}

type _CandidateListTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CandidateListTouchBarItem */
// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for CandidateListTouchBarItem */
	// properties:
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	AttributedStringForCandidate() func(objc.ID, unsafe.Pointer) unsafe.Pointer
	SetAttributedStringForCandidate(value func(objc.ID, unsafe.Pointer) unsafe.Pointer)
	Candidates() []objc.ID
	Client() unsafe.Pointer
	SetClient(value unsafe.Pointer)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	CandidateListVisible() bool
	Collapsed() bool
	SetCollapsed(value bool)
	IsCandidateListVisible() bool
	SetIsCandidateListVisible(value bool)
	IsCollapsed() bool
	SetIsCollapsed(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CandidateListTouchBarItem */
	// methods:
	SetCandidatesForSelectedRangeInString(candidates []objc.ID, selectedRange corefoundation.Range, originalString objc.IObject /* cross-framework: NSString */)
	UpdateWithInsertionPointVisibility(isVisible bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CandidateListTouchBarItem */
// Alloc allocates a new instance without initialization.
func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CandidateListTouchBarItem) Autorelease() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCandidateListTouchBarItem creates a new CandidateListTouchBarItem instance.
func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return getCandidateListTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CandidateListTouchBarItem */
// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.


// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem
type CandidateListTouchBarItem struct {
	TouchBarItem
}

// CandidateListTouchBarItemFrom constructs a [CandidateListTouchBarItem] from an unsafe.Pointer.
//
// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.
func CandidateListTouchBarItemFrom(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CandidateListTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CandidateListTouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CandidateListTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CandidateListTouchBarItem */

// Sets an array of candidate objects to be displayed in the candidate list bar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/setCandidates(_:forSelectedRange:in:)
func (c_ CandidateListTouchBarItem) SetCandidatesForSelectedRangeInString(candidates []objc.ID, selectedRange corefoundation.Range, originalString objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCandidates:forSelectedRange:inString:"), candidates, selectedRange, originalString)
}/* debug [instance_methods/method]: SetCandidatesForSelectedRangeInString */


// Updates the candidate list visibility configuration based on the client’s insertion point state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/update(withInsertionPointVisibility:)
func (c_ CandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateWithInsertionPointVisibility:"), isVisible)
}/* debug [instance_methods/method]: UpdateWithInsertionPointVisibility */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CandidateListTouchBarItem */

// A Boolean value that specifies whether the item can be collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsCollapsing
func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCollapsing"))
	return rv
}/* debug [instance_properties/getter]: allowsCollapsing */


// A Boolean value that specifies whether the item can be collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsCollapsing
func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCollapsing:"), value)
}/* debug [instance_properties/setter]: allowsCollapsing */


// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsTextInputContextCandidates
func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsTextInputContextCandidates"))
	return rv
}/* debug [instance_properties/getter]: allowsTextInputContextCandidates */


// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsTextInputContextCandidates
func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsTextInputContextCandidates:"), value)
}/* debug [instance_properties/setter]: allowsTextInputContextCandidates */


// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) AttributedStringForCandidate() func(objc.ID, unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[func(objc.ID, unsafe.Pointer) unsafe.Pointer](c_.ID, objc.Sel("attributedStringForCandidate"))
	return rv
}/* debug [instance_properties/getter]: attributedStringForCandidate */


// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) SetAttributedStringForCandidate(value func(objc.ID, unsafe.Pointer) unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringForCandidate:"), value)
}/* debug [instance_properties/setter]: attributedStringForCandidate */


// The array of candidate objects previously set by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/candidates
func (c_ CandidateListTouchBarItem) Candidates() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("candidates"))
	return rv
}/* debug [instance_properties/getter]: candidates */


// The client object for the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/client
func (c_ CandidateListTouchBarItem) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("client"))
	return rv
}/* debug [instance_properties/getter]: client */


// The client object for the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/client
func (c_ CandidateListTouchBarItem) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClient:"), value)
}/* debug [instance_properties/setter]: client */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/customizationLabel
func (c_ CandidateListTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/customizationLabel
func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// The delegate of the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/delegate
func (c_ CandidateListTouchBarItem) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate of the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/delegate
func (c_ CandidateListTouchBarItem) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCandidateListVisible
func (c_ CandidateListTouchBarItem) CandidateListVisible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("candidateListVisible"))
	return rv
}/* debug [instance_properties/getter]: candidateListVisible */


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) Collapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("collapsed"))
	return rv
}/* debug [instance_properties/getter]: collapsed */


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollapsed:"), value)
}/* debug [instance_properties/setter]: collapsed */


// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCandidateListVisible"))
	return rv
}/* debug [instance_properties/getter]: isCandidateListVisible */


// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) SetIsCandidateListVisible(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCandidateListVisible:"), value)
}/* debug [instance_properties/setter]: isCandidateListVisible */


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCollapsed"))
	return rv
}/* debug [instance_properties/getter]: isCollapsed */


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) SetIsCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCollapsed:"), value)
}/* debug [instance_properties/setter]: isCollapsed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCandidateListTouchBarItem */



