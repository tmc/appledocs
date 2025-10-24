// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CandidateListTouchBarItem] class.
type ICandidateListTouchBarItem interface {
	ITouchBarItem
	// properties:
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	AttributedStringForCandidate() objc.IObject /* cross-framework: AttributedString */
	SetAttributedStringForCandidate(value objc.IObject /* cross-framework: AttributedString */)
	Candidates() unsafe.Pointer
	SetCandidates(value unsafe.Pointer)
	Client() TextInputClient /* not a class type */
	SetClient(value TextInputClient /* not a class type */)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Delegate() CandidateListTouchBarItemDelegate /* not a class type */
	SetDelegate(value CandidateListTouchBarItemDelegate /* not a class type */)
	IsCandidateListVisible() bool
	SetIsCandidateListVisible(value bool)
	IsCollapsed() bool
	SetIsCollapsed(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that specifies whether the item can be collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowscollapsing
func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCollapsing"))
	return rv
}


// A Boolean value that specifies whether the item can be collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowscollapsing
func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCollapsing:"), value)
}


// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowstextinputcontextcandidates
func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsTextInputContextCandidates"))
	return rv
}


// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowstextinputcontextcandidates
func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsTextInputContextCandidates:"), value)
}


// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/attributedstringforcandidate
func (c_ CandidateListTouchBarItem) AttributedStringForCandidate() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringForCandidate"))
	return rv
}


// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/attributedstringforcandidate
func (c_ CandidateListTouchBarItem) SetAttributedStringForCandidate(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringForCandidate:"), value)
}


// The array of candidate objects previously set by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/candidates
func (c_ CandidateListTouchBarItem) Candidates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("candidates"))
	return rv
}


// The array of candidate objects previously set by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/candidates
func (c_ CandidateListTouchBarItem) SetCandidates(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCandidates:"), value)
}


// The client object for the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/client
func (c_ CandidateListTouchBarItem) Client() TextInputClient /* not a class type */ {
	rv := objc.Send[TextInputClient](c_.ID, objc.Sel("client"))
	return rv
}


// The client object for the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/client
func (c_ CandidateListTouchBarItem) SetClient(value TextInputClient /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClient:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/customizationlabel
func (c_ CandidateListTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/customizationlabel
func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The delegate of the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/delegate
func (c_ CandidateListTouchBarItem) Delegate() CandidateListTouchBarItemDelegate /* not a class type */ {
	rv := objc.Send[CandidateListTouchBarItemDelegate](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the candidate list item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/delegate
func (c_ CandidateListTouchBarItem) SetDelegate(value CandidateListTouchBarItemDelegate /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCandidateListVisible"))
	return rv
}


// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) SetIsCandidateListVisible(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCandidateListVisible:"), value)
}


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCollapsed"))
	return rv
}


// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) SetIsCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCollapsed:"), value)
}



