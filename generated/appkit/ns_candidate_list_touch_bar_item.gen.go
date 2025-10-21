// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A bar item that, along with its delegate, provides a list of textual suggestions for the current text view.
//
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


// A block that converts a candidate object into an attributed string for display in the candidate list item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) AttributedStringForCandidate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributedStringForCandidate"))
	return rv
}


// SetAttributedStringForCandidate sets the value of the attributedStringForCandidate property.
// A block that converts a candidate object into an attributed string for display in the candidate list item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c_ CandidateListTouchBarItem) SetAttributedStringForCandidate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringForCandidate:"), value)
}

// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) Collapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("collapsed"))
	return rv
}


// SetCollapsed sets the value of the collapsed property.
// A Boolean value that controls the visibility of the candidate list.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollapsed:"), value)
}

// A Boolean value that specifies whether the item can be collapsed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowscollapsing
func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsCollapsing"))
	return rv
}


// SetAllowsCollapsing sets the value of the allowsCollapsing property.
// A Boolean value that specifies whether the item can be collapsed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowscollapsing
func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsCollapsing:"), value)
}

// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowstextinputcontextcandidates
func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsTextInputContextCandidates"))
	return rv
}


// SetAllowsTextInputContextCandidates sets the value of the allowsTextInputContextCandidates property.
// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/allowstextinputcontextcandidates
func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsTextInputContextCandidates:"), value)
}

// The array of candidate objects previously set by
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/candidates
func (c_ CandidateListTouchBarItem) Candidates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("candidates"))
	return rv
}


// SetCandidates sets the value of the candidates property.
// The array of candidate objects previously set by

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/candidates
func (c_ CandidateListTouchBarItem) SetCandidates(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCandidates:"), value)
}

// The client object for the candidate list item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/client
func (c_ CandidateListTouchBarItem) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("client"))
	return rv
}


// SetClient sets the value of the client property.
// The client object for the candidate list item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/client
func (c_ CandidateListTouchBarItem) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClient:"), value)
}

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/customizationlabel
func (c_ CandidateListTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/customizationlabel
func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

// The delegate of the candidate list item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/delegate
func (c_ CandidateListTouchBarItem) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the candidate list item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/delegate
func (c_ CandidateListTouchBarItem) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that represents the visibility of this item’s candidate list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCandidateListVisible"))
	return rv
}


// SetIsCandidateListVisible sets the value of the isCandidateListVisible property.
// A Boolean value that represents the visibility of this item’s candidate list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscandidatelistvisible
func (c_ CandidateListTouchBarItem) SetIsCandidateListVisible(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCandidateListVisible:"), value)
}

// A Boolean value that controls the visibility of the candidate list.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCollapsed"))
	return rv
}


// SetIsCollapsed sets the value of the isCollapsed property.
// A Boolean value that controls the visibility of the candidate list.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritem/iscollapsed
func (c_ CandidateListTouchBarItem) SetIsCollapsed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCollapsed:"), value)
}



