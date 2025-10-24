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

/* debug [class.gen.go]: Generating class NSTextFinder */


/* debug [class_header]: Header for NSTextFinder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextFinder */
// An interface definition for the [TextFinder] class.
type ITextFinder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextFinder */
	// properties:
	Client() unsafe.Pointer
	SetClient(value unsafe.Pointer)
	FindBarContainer() unsafe.Pointer
	SetFindBarContainer(value unsafe.Pointer)
	FindIndicatorNeedsUpdate() bool
	SetFindIndicatorNeedsUpdate(value bool)
	IncrementalMatchRanges() []foundation.Value
	IncrementalSearchingShouldDimContentView() bool
	SetIncrementalSearchingShouldDimContentView(value bool)
	IncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	FindBarPosition() objectivec.IObject
	SetFindBarPosition(value objectivec.IObject)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextFinder */
	// methods:
	CancelFindIndicator()
	NoteClientStringWillChange()
	PerformAction(op TextFinderAction)
	ValidateAction(op TextFinderAction) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextFinder */
// Alloc allocates a new instance without initialization.
func (tc _TextFinderClass) Alloc() TextFinder {
	rv := objc.Send[TextFinder](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextFinder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextFinder */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/init(coder:)
func NewTextFinderWithCoder(coder foundation.Coder) TextFinder {
	instance := getTextFinderClass().Alloc()
	rv := objc.Send[TextFinder](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextFinderWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextFinder */

// Override this method to draw custom highlighting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/drawIncrementalMatchHighlight(in:)
func (tc _TextFinderClass) DrawIncrementalMatchHighlightInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("drawIncrementalMatchHighlightInRect:"), rect)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DrawIncrementalMatchHighlightInRect) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextFinder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextFinder */

// Cancels the find indicator immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/cancelFindIndicator()
func (t_ TextFinder) CancelFindIndicator() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancelFindIndicator"))
}/* debug [instance_methods/method]: CancelFindIndicator */


// Invoke this method when the searched content will change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/noteClientStringWillChange()
func (t_ TextFinder) NoteClientStringWillChange() {
	objc.Send[objc.ID](t_.ID, objc.Sel("noteClientStringWillChange"))
}/* debug [instance_methods/method]: NoteClientStringWillChange */


// Performs the specified text finding action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/performAction(_:)
func (t_ TextFinder) PerformAction(op TextFinderAction) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performAction:"), op)
}/* debug [instance_methods/method]: PerformAction */


// Allows validation of the find action before performing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/validateAction(_:)
func (t_ TextFinder) ValidateAction(op TextFinderAction) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("validateAction:"), op)
	return rv
}/* debug [instance_methods/method]: ValidateAction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextFinder */

// The object that provides the target search string, find bar location, and feedback methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/client
func (t_ TextFinder) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("client"))
	return rv
}/* debug [instance_properties/getter]: client */


// The object that provides the target search string, find bar location, and feedback methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/client
func (t_ TextFinder) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClient:"), value)
}/* debug [instance_properties/setter]: client */


// Specifies the find bar container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findBarContainer
func (t_ TextFinder) FindBarContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("findBarContainer"))
	return rv
}/* debug [instance_properties/getter]: findBarContainer */


// Specifies the find bar container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findBarContainer
func (t_ TextFinder) SetFindBarContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarContainer:"), value)
}/* debug [instance_properties/setter]: findBarContainer */


// Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findIndicatorNeedsUpdate
func (t_ TextFinder) FindIndicatorNeedsUpdate() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("findIndicatorNeedsUpdate"))
	return rv
}/* debug [instance_properties/getter]: findIndicatorNeedsUpdate */


// Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/findIndicatorNeedsUpdate
func (t_ TextFinder) SetFindIndicatorNeedsUpdate(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindIndicatorNeedsUpdate:"), value)
}/* debug [instance_properties/setter]: findIndicatorNeedsUpdate */


// Array of incremental search matches posted on the main queue, which have been found during a background search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalMatchRanges
func (t_ TextFinder) IncrementalMatchRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("incrementalMatchRanges"))
	return rv
}/* debug [instance_properties/getter]: incrementalMatchRanges */


// Determines the type of incremental search feedback to be presented
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalSearchingShouldDimContentView
func (t_ TextFinder) IncrementalSearchingShouldDimContentView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingShouldDimContentView"))
	return rv
}/* debug [instance_properties/getter]: incrementalSearchingShouldDimContentView */


// Determines the type of incremental search feedback to be presented
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalSearchingShouldDimContentView
func (t_ TextFinder) SetIncrementalSearchingShouldDimContentView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingShouldDimContentView:"), value)
}/* debug [instance_properties/setter]: incrementalSearchingShouldDimContentView */


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/isIncrementalSearchingEnabled
func (t_ TextFinder) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingEnabled"))
	return rv
}/* debug [instance_properties/getter]: incrementalSearchingEnabled */


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/isIncrementalSearchingEnabled
func (t_ TextFinder) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}/* debug [instance_properties/setter]: incrementalSearchingEnabled */


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/findbarposition-swift.property
func (t_ TextFinder) FindBarPosition() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("findBarPosition"))
	return rv
}/* debug [instance_properties/getter]: findBarPosition */


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/findbarposition-swift.property
func (t_ TextFinder) SetFindBarPosition(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarPosition:"), value)
}/* debug [instance_properties/setter]: findBarPosition */


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/isincrementalsearchingenabled
func (t_ TextFinder) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isIncrementalSearchingEnabled */


// Determines if incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/isincrementalsearchingenabled
func (t_ TextFinder) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}/* debug [instance_properties/setter]: isIncrementalSearchingEnabled */


// The view assigned by the text bar as the find bar view for the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/findbarview
func (t_ TextFinder) FindBarView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("findBarView"))
	return rv
}/* debug [instance_properties/getter]: findBarView */


// The view assigned by the text bar as the find bar view for the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/findbarview
func (t_ TextFinder) SetFindBarView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFindBarView:"), value)
}/* debug [instance_properties/setter]: findBarView */


// Returns whether the container should display its find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/isfindbarvisible
func (t_ TextFinder) IsFindBarVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFindBarVisible"))
	return rv
}/* debug [instance_properties/getter]: isFindBarVisible */


// Returns whether the container should display its find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/isfindbarvisible
func (t_ TextFinder) SetIsFindBarVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFindBarVisible:"), value)
}/* debug [instance_properties/setter]: isFindBarVisible */


// Returns whether multiple items can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/allowsmultipleselection
func (t_ TextFinder) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsMultipleSelection */


// Returns whether multiple items can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/allowsmultipleselection
func (t_ TextFinder) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}/* debug [instance_properties/setter]: allowsMultipleSelection */


// Returns the currently selected range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/firstselectedrange
func (t_ TextFinder) FirstSelectedRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("firstSelectedRange"))
	return rv
}/* debug [instance_properties/getter]: firstSelectedRange */


// Returns the currently selected range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/firstselectedrange
func (t_ TextFinder) SetFirstSelectedRange(value corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFirstSelectedRange:"), value)
}/* debug [instance_properties/setter]: firstSelectedRange */


// Returns whether the text is selectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/isselectable
func (t_ TextFinder) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}/* debug [instance_properties/getter]: isSelectable */


// Returns whether the text is selectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/isselectable
func (t_ TextFinder) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}/* debug [instance_properties/setter]: isSelectable */


// Returns an array of selected ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/selectedranges
func (t_ TextFinder) SelectedRanges() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}/* debug [instance_properties/getter]: selectedRanges */


// Returns an array of selected ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/selectedranges
func (t_ TextFinder) SetSelectedRanges(value foundation.Value) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:"), value)
}/* debug [instance_properties/setter]: selectedRanges */


// An array of visible character ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/visiblecharacterranges
func (t_ TextFinder) VisibleCharacterRanges() foundation.Value {
	rv := objc.Send[foundation.Value](t_.ID, objc.Sel("visibleCharacterRanges"))
	return rv
}/* debug [instance_properties/getter]: visibleCharacterRanges */


// An array of visible character ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderclient/visiblecharacterranges
func (t_ TextFinder) SetVisibleCharacterRanges(value foundation.Value) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibleCharacterRanges:"), value)
}/* debug [instance_properties/setter]: visibleCharacterRanges */


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextFinder) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}/* debug [instance_properties/getter]: usesFindBar */


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextFinder) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}/* debug [instance_properties/setter]: usesFindBar */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextFinder */


