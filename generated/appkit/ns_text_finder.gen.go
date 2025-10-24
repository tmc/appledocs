// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextFinder] class.
var (
	TextFinderClass     _TextFinderClass
	TextFinderClassOnce sync.Once
)

func getTextFinderClass() _TextFinderClass {
	TextFinderClassOnce.Do(func() {
		TextFinderClass = _TextFinderClass{objc.GetClass("NSTextFinder")}
	})
	return TextFinderClass
}

type _TextFinderClass struct {
	class objc.Class
}

// An interface definition for the [TextFinder] class.
type ITextFinder interface {
	objectivec.IObject
	// properties:
	Client() objc.ID
	SetClient(value objc.ID)
	FindBarContainer() objc.ID
	SetFindBarContainer(value objc.ID)
	FindIndicatorNeedsUpdate() bool
	SetFindIndicatorNeedsUpdate(value bool)
	IncrementalMatchRanges() []foundation.Value
	IncrementalSearchingShouldDimContentView() bool
	SetIncrementalSearchingShouldDimContentView(value bool)
	IncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	FindBarPosition() unsafe.Pointer
	SetFindBarPosition(value unsafe.Pointer)
	IsIncrementalSearchingEnabled() bool
	SetIsIncrementalSearchingEnabled(value bool)
	FindBarView() IView
	SetFindBarView(value IView)
	IsFindBarVisible() bool
	SetIsFindBarVisible(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	FirstSelectedRange() corefoundation.Range
	SetFirstSelectedRange(value corefoundation.Range)
	IsSelectable() bool
	SetIsSelectable(value bool)
	SelectedRanges() foundation.Value
	SetSelectedRanges(value foundation.Value)
	VisibleCharacterRanges() foundation.Value
	SetVisibleCharacterRanges(value foundation.Value)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	// methods:
	CancelFindIndicator()
	NoteClientStringWillChange()
	PerformAction(op TextFinderAction)
	ValidateAction(op TextFinderAction) bool
}

// An optional search-and-replace find interface inside a view, usually a scroll view.
//
// The class serves as a controller for the standard Cocoa find bar. The class interacts heavily with a object which supports the protocol. The client object provides access to the content being searched and provides visual feedback for a search operation. All menu items related to finding (Find…, Find Next, Find Previous, Use Selection for Find, etc.) should have the same action, , which gets sent down the responder chain in the standard method.


// An optional search-and-replace find interface inside a view, usually a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder
type TextFinder struct {
	objectivec.Object
}

// TextFinderFrom constructs a [TextFinder] from an unsafe.Pointer.
//
// An optional search-and-replace find interface inside a view, usually a scroll view.
func TextFinderFrom(ptr unsafe.Pointer) TextFinder {
	return TextFinder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextFinderClass) Alloc() TextFinder {
	rv := objc.Send[TextFinder](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextFinderClass) New() TextFinder {
	rv := objc.Send[TextFinder](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextFinder) Init() TextFinder {
	rv := objc.Send[TextFinder](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextFinder) Autorelease() TextFinder {
	rv := objc.Send[TextFinder](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextFinder creates a new TextFinder instance.
func NewTextFinder() TextFinder {
	return getTextFinderClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/init(coder:)
func NewTextFinderWithCoder(coder foundation.Coder) TextFinder {
	instance := getTextFinderClass().Alloc()
	rv := objc.Send[TextFinder](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Override this method to draw custom highlighting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/drawIncrementalMatchHighlight(in:)
func (tc _TextFinderClass) DrawIncrementalMatchHighlightInRect(rect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("drawIncrementalMatchHighlightInRect:"), rect)
}


// Cancels the find indicator immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/cancelFindIndicator()
func (t_ TextFinder) CancelFindIndicator() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancelFindIndicator"))
}


// Invoke this method when the searched content will change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/noteClientStringWillChange()
func (t_ TextFinder) NoteClientStringWillChange() {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteClientStringWillChange"))
}


// Performs the specified text finding action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/performAction(_:)
func (t_ TextFinder) PerformAction(op TextFinderAction) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performAction:"), op)
}


// Allows validation of the find action before performing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/validateAction(_:)
func (t_ TextFinder) ValidateAction(op TextFinderAction) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("validateAction:"), op)
	return rv
}


// The object that provides the target search string, find bar location, and feedback methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/client
func (t_ TextFinder) Client() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("client"))
	return rv
}


// The object that provides the target search string, find bar location, and feedback methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/client
func (t_ TextFinder) SetClient(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClient:"), value)
}


// Specifies the find bar container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findBarContainer
func (t_ TextFinder) FindBarContainer() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("findBarContainer"))
	return rv
}


// Specifies the find bar container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findBarContainer
func (t_ TextFinder) SetFindBarContainer(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarContainer:"), value)
}


// Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findIndicatorNeedsUpdate
func (t_ TextFinder) FindIndicatorNeedsUpdate() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("findIndicatorNeedsUpdate"))
	return rv
}


// Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findIndicatorNeedsUpdate
func (t_ TextFinder) SetFindIndicatorNeedsUpdate(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindIndicatorNeedsUpdate:"), value)
}


// Array of incremental search matches posted on the main queue, which have been found during a background search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalMatchRanges
func (t_ TextFinder) IncrementalMatchRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("incrementalMatchRanges"))
	return rv
}


// Determines the type of incremental search feedback to be presented
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalSearchingShouldDimContentView
func (t_ TextFinder) IncrementalSearchingShouldDimContentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingShouldDimContentView"))
	return rv
}


// Determines the type of incremental search feedback to be presented
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalSearchingShouldDimContentView
func (t_ TextFinder) SetIncrementalSearchingShouldDimContentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingShouldDimContentView:"), value)
}


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/isIncrementalSearchingEnabled
func (t_ TextFinder) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingEnabled"))
	return rv
}


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/isIncrementalSearchingEnabled
func (t_ TextFinder) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/findbarposition-swift.property
func (t_ TextFinder) FindBarPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("findBarPosition"))
	return rv
}


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/findbarposition-swift.property
func (t_ TextFinder) SetFindBarPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarPosition:"), value)
}


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/isincrementalsearchingenabled
func (t_ TextFinder) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/isincrementalsearchingenabled
func (t_ TextFinder) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}


// The view assigned by the text bar as the find bar view for the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/findbarview
func (t_ TextFinder) FindBarView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("findBarView"))
	return rv
}


// The view assigned by the text bar as the find bar view for the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/findbarview
func (t_ TextFinder) SetFindBarView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarView:"), value)
}


// Returns whether the container should display its find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/isfindbarvisible
func (t_ TextFinder) IsFindBarVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFindBarVisible"))
	return rv
}


// Returns whether the container should display its find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/isfindbarvisible
func (t_ TextFinder) SetIsFindBarVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFindBarVisible:"), value)
}


// Returns whether multiple items can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/allowsmultipleselection
func (t_ TextFinder) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// Returns whether multiple items can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/allowsmultipleselection
func (t_ TextFinder) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// Returns the currently selected range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/firstselectedrange
func (t_ TextFinder) FirstSelectedRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("firstSelectedRange"))
	return rv
}


// Returns the currently selected range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/firstselectedrange
func (t_ TextFinder) SetFirstSelectedRange(value corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFirstSelectedRange:"), value)
}


// Returns whether the text is selectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/isselectable
func (t_ TextFinder) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// Returns whether the text is selectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/isselectable
func (t_ TextFinder) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// Returns an array of selected ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/selectedranges
func (t_ TextFinder) SelectedRanges() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}


// Returns an array of selected ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/selectedranges
func (t_ TextFinder) SetSelectedRanges(value foundation.Value) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:"), value)
}


// An array of visible character ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/visiblecharacterranges
func (t_ TextFinder) VisibleCharacterRanges() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("visibleCharacterRanges"))
	return rv
}


// An array of visible character ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/visiblecharacterranges
func (t_ TextFinder) SetVisibleCharacterRanges(value foundation.Value) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibleCharacterRanges:"), value)
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextFinder) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextFinder) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}


