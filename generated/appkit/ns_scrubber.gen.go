// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrubber */


/* debug [class_header]: Header for NSScrubber */
// The class instance for the [Scrubber] class.
var (
	ScrubberClass     _ScrubberClass
	ScrubberClassOnce sync.Once
)

func getScrubberClass() _ScrubberClass {
	ScrubberClassOnce.Do(func() {
		ScrubberClass = _ScrubberClass{objc.GetClass("NSScrubber")}
	})
	return ScrubberClass
}

type _ScrubberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Scrubber */
// An interface definition for the [Scrubber] class.
type IScrubber interface {
	IView
	
/* debug [class_interface_properties]: Properties for Scrubber */
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BackgroundView() IView
	SetBackgroundView(value IView)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FloatsSelectionViews() bool
	SetFloatsSelectionViews(value bool)
	HighlightedIndex() int
	Continuous() bool
	SetContinuous(value bool)
	ItemAlignment() ScrubberAlignment
	SetItemAlignment(value ScrubberAlignment)
	Mode() ScrubberMode
	SetMode(value ScrubberMode)
	NumberOfItems() int
	ScrubberLayout() IScrubberLayout
	SetScrubberLayout(value IScrubberLayout)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionBackgroundStyle() IScrubberSelectionStyle
	SetSelectionBackgroundStyle(value IScrubberSelectionStyle)
	SelectionOverlayStyle() IScrubberSelectionStyle
	SetSelectionOverlayStyle(value IScrubberSelectionStyle)
	ShowsAdditionalContentIndicators() bool
	SetShowsAdditionalContentIndicators(value bool)
	ShowsArrowButtons() bool
	SetShowsArrowButtons(value bool)
	IsContinuous() bool
	SetIsContinuous(value bool)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageView() IImageView
	SetImageView(value IImageView)
	ScrubberContentSize() Size /* not a class type */
	SetScrubberContentSize(value Size /* not a class type */)
	ShouldInvalidateLayoutForHighlightChange() bool
	SetShouldInvalidateLayoutForHighlightChange(value bool)
	ShouldInvalidateLayoutForSelectionChange() bool
	SetShouldInvalidateLayoutForSelectionChange(value bool)
	Alpha() float64
	SetAlpha(value float64)
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	ItemIndex() int
	SetItemIndex(value int)
	TextField() ITextField
	SetTextField(value ITextField)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Scrubber */
	// methods:
	InsertItemsAtIndexes(indexes foundation.IndexSet)
	ItemViewForItemAtIndex(index int) IScrubberItemView
	MakeItemWithIdentifierOwner(itemIdentifier UserInterfaceItemIdentifier /* typedef */, owner objc.IObject) IScrubberItemView
	MoveItemAtIndexToIndex(oldIndex int, newIndex int)
	PerformSequentialBatchUpdates(updateBlock unsafe.Pointer)
	RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier UserInterfaceItemIdentifier /* typedef */)
	RegisterNibForItemIdentifier(nib INib, itemIdentifier UserInterfaceItemIdentifier /* typedef */)
	ReloadData()
	ReloadItemsAtIndexes(indexes foundation.IndexSet)
	RemoveItemsAtIndexes(indexes foundation.IndexSet)
	ScrollItemAtIndexToAlignment(index int, alignment ScrubberAlignment)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Scrubber */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberClass) Alloc() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberClass) New() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scrubber) Init() Scrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scrubber) Autorelease() Scrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubber creates a new Scrubber instance.
func NewScrubber() Scrubber {
	return getScrubberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Scrubber */
// A customizable item picker control for the Touch Bar.
//
// On supported MacBook Pro models, you can use a scrubber (an instance of the class) to provide a horizontally-oriented, item-picker control in the Touch Bar. Use a scrubber to let the user pick an item from a related collection, such as a photo from a library or a date from a date range. Refer to the following sample code projects which demonstrate how to use and related classes, including the class: Each item that appears in a scrubber is a specialized view that supports selection and scrubber-appropriate decorations. The scrubber keeps track of its items by their index positions. There are many classes in the scrubber API, as well as a delegate protocol, a data source protocol, and a callback-based layout API. The design pattern is reminiscent of that used for a collection view (an instance of the class). You might find it helpful to refer to the overview for background. Be aware, though of the differences. For example, while scrubbers and collection views both employ a method, and both employ a reuse queue, a scrubber is subclassed from the class while a collection view is subclassed from the class. A scrubber employs: The itself (an instance of the class), which serves as a container view that shows a subview for each scrubber item, and which employs a reuse-queue pattern for efficiency and performance. A (conforming to the protocol), which provides scrubber items to the scrubber, on demand, from an associated data collection in your app. Specify the data source in the scrubber’s property A (conforming to the protocol), which responds to user interaction — such as with its and methods. Specify the delegate in the scrubber’s property. You can also use the delegate to respond to the highlighting and selection of scrubber items, and to respond to changes in which items are visible in the scrubber. A (an instance of a subclass of the abstract class, typically the concrete subclass). You implement a layout to respond to calls, from the system, to return view specifications for the items to be displayed in the scrubber. The layout, in this way, assists in arranging and decorating the scrubber’s contained items, and in providing appearance changes in response to user interaction. Specify the layout in the scrubber’s property. Before learning how to use a scrubber in the Touch Bar, be sure you read the overview for the class.


// A customizable item picker control for the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber
type Scrubber struct {
	View
}

// ScrubberFrom constructs a [Scrubber] from an unsafe.Pointer.
//
// A customizable item picker control for the Touch Bar.
func ScrubberFrom(ptr unsafe.Pointer) Scrubber {
	return Scrubber{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Scrubber */

// Initializes and returns a newly allocated scrubber object from a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/init(coder:)
func NewScrubberWithCoder(coder foundation.Coder) Scrubber {
	instance := getScrubberClass().Alloc()
	rv := objc.Send[Scrubber](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberWithCoder */


// Initializes and returns a newly allocated scrubber object with the specified frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/init(frame:)
func NewScrubberWithFrame(frameRect Rect /* not a class type */) Scrubber {
	instance := getScrubberClass().Alloc()
	rv := objc.Send[Scrubber](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Scrubber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Scrubber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Scrubber */

// Inserts new items at the specified indexes into the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/insertItems(at:)
func (s_ Scrubber) InsertItemsAtIndexes(indexes foundation.IndexSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertItemsAtIndexes:"), indexes)
}/* debug [instance_methods/method]: InsertItemsAtIndexes */


// Returns the view for the item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/itemViewForItem(at:)
func (s_ Scrubber) ItemViewForItemAtIndex(index int) IScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("itemViewForItemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ItemViewForItemAtIndex */


// Creates or returns a reusable item object with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/makeItem(withIdentifier:owner:)
func (s_ Scrubber) MakeItemWithIdentifierOwner(itemIdentifier UserInterfaceItemIdentifier /* typedef */, owner objc.IObject) IScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("makeItemWithIdentifier:owner:"), itemIdentifier, owner)
	return rv
}/* debug [instance_methods/method]: MakeItemWithIdentifierOwner */


// Moves an item from one index to another in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/moveItem(at:to:)
func (s_ Scrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("moveItemAtIndex:toIndex:"), oldIndex, newIndex)
}/* debug [instance_methods/method]: MoveItemAtIndexToIndex */


// Combines multiple scrubber content updates into a single action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/performSequentialBatchUpdates(_:)
func (s_ Scrubber) PerformSequentialBatchUpdates(updateBlock unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("performSequentialBatchUpdates:"), updateBlock)
}/* debug [instance_methods/method]: PerformSequentialBatchUpdates */


// Registers a class for the scrubber to use when it creates new items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-2rb69
func (s_ Scrubber) RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerClass:forItemIdentifier:"), itemViewClass, itemIdentifier)
}/* debug [instance_methods/method]: RegisterClassForItemIdentifier */


// Registers a nib file for the scrubber to use when it creates new items in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-6jye0
func (s_ Scrubber) RegisterNibForItemIdentifier(nib INib, itemIdentifier UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerNib:forItemIdentifier:"), nib, itemIdentifier)
}/* debug [instance_methods/method]: RegisterNibForItemIdentifier */


// Reloads the content of the entire scrubber, and deselects the currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadData()
func (s_ Scrubber) ReloadData() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Reloads the items at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadItems(at:)
func (s_ Scrubber) ReloadItemsAtIndexes(indexes foundation.IndexSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("reloadItemsAtIndexes:"), indexes)
}/* debug [instance_methods/method]: ReloadItemsAtIndexes */


// Removes the items at the specified indexes from the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/removeItems(at:)
func (s_ Scrubber) RemoveItemsAtIndexes(indexes foundation.IndexSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeItemsAtIndexes:"), indexes)
}/* debug [instance_methods/method]: RemoveItemsAtIndexes */


// Scrolls an item to a specified alignment within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/scrollItem(at:to:)
func (s_ Scrubber) ScrollItemAtIndexToAlignment(index int, alignment ScrubberAlignment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollItemAtIndex:toAlignment:"), index, alignment)
}/* debug [instance_methods/method]: ScrollItemAtIndexToAlignment */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Scrubber */

// The color displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundColor
func (s_ Scrubber) BackgroundColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundColor
func (s_ Scrubber) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// A view that is displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundView
func (s_ Scrubber) BackgroundView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("backgroundView"))
	return rv
}/* debug [instance_properties/getter]: backgroundView */


// A view that is displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundView
func (s_ Scrubber) SetBackgroundView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundView:"), value)
}/* debug [instance_properties/setter]: backgroundView */


// The object that provides the data for the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/dataSource
func (s_ Scrubber) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The object that provides the data for the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/dataSource
func (s_ Scrubber) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// The object that acts as the delegate of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/delegate
func (s_ Scrubber) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object that acts as the delegate of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/delegate
func (s_ Scrubber) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/floatsSelectionViews
func (s_ Scrubber) FloatsSelectionViews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("floatsSelectionViews"))
	return rv
}/* debug [instance_properties/getter]: floatsSelectionViews */


// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/floatsSelectionViews
func (s_ Scrubber) SetFloatsSelectionViews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFloatsSelectionViews:"), value)
}/* debug [instance_properties/setter]: floatsSelectionViews */


// The index of the highlighted item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/highlightedIndex
func (s_ Scrubber) HighlightedIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("highlightedIndex"))
	return rv
}/* debug [instance_properties/getter]: highlightedIndex */


// A Boolean value that, together with the property, determines scrubber interaction style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/isContinuous
func (s_ Scrubber) Continuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("continuous"))
	return rv
}/* debug [instance_properties/getter]: continuous */


// A Boolean value that, together with the property, determines scrubber interaction style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/isContinuous
func (s_ Scrubber) SetContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContinuous:"), value)
}/* debug [instance_properties/setter]: continuous */


// A setting that specifies the snapping behavior of items in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/itemAlignment
func (s_ Scrubber) ItemAlignment() ScrubberAlignment {
	rv := objc.Send[ScrubberAlignment](s_.ID, objc.Sel("itemAlignment"))
	return rv
}/* debug [instance_properties/getter]: itemAlignment */


// A setting that specifies the snapping behavior of items in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/itemAlignment
func (s_ Scrubber) SetItemAlignment(value ScrubberAlignment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemAlignment:"), value)
}/* debug [instance_properties/setter]: itemAlignment */


// A setting that determines whether interaction with the scrubber is fixed or free.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/mode-swift.property
func (s_ Scrubber) Mode() ScrubberMode {
	rv := objc.Send[ScrubberMode](s_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// A setting that determines whether interaction with the scrubber is fixed or free.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/mode-swift.property
func (s_ Scrubber) SetMode(value ScrubberMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// The number of items represented by the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/numberOfItems
func (s_ Scrubber) NumberOfItems() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfItems */


// An object used to describe the layout of items within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/scrubberLayout
func (s_ Scrubber) ScrubberLayout() IScrubberLayout {
	rv := objc.Send[ScrubberLayout](s_.ID, objc.Sel("scrubberLayout"))
	return rv
}/* debug [instance_properties/getter]: scrubberLayout */


// An object used to describe the layout of items within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/scrubberLayout
func (s_ Scrubber) SetScrubberLayout(value IScrubberLayout) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubberLayout:"), value)
}/* debug [instance_properties/setter]: scrubberLayout */


// The index of the selected item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectedIndex
func (s_ Scrubber) SelectedIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// The index of the selected item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectedIndex
func (s_ Scrubber) SetSelectedIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */


// The style applied to the background of selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionBackgroundStyle
func (s_ Scrubber) SelectionBackgroundStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("selectionBackgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: selectionBackgroundStyle */


// The style applied to the background of selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionBackgroundStyle
func (s_ Scrubber) SetSelectionBackgroundStyle(value IScrubberSelectionStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectionBackgroundStyle:"), value)
}/* debug [instance_properties/setter]: selectionBackgroundStyle */


// The style overlaid on selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionOverlayStyle
func (s_ Scrubber) SelectionOverlayStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("selectionOverlayStyle"))
	return rv
}/* debug [instance_properties/getter]: selectionOverlayStyle */


// The style overlaid on selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionOverlayStyle
func (s_ Scrubber) SetSelectionOverlayStyle(value IScrubberSelectionStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectionOverlayStyle:"), value)
}/* debug [instance_properties/setter]: selectionOverlayStyle */


// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/showsAdditionalContentIndicators
func (s_ Scrubber) ShowsAdditionalContentIndicators() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsAdditionalContentIndicators"))
	return rv
}/* debug [instance_properties/getter]: showsAdditionalContentIndicators */


// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/showsAdditionalContentIndicators
func (s_ Scrubber) SetShowsAdditionalContentIndicators(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsAdditionalContentIndicators:"), value)
}/* debug [instance_properties/setter]: showsAdditionalContentIndicators */


// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/showsArrowButtons
func (s_ Scrubber) ShowsArrowButtons() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsArrowButtons"))
	return rv
}/* debug [instance_properties/getter]: showsArrowButtons */


// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/showsArrowButtons
func (s_ Scrubber) SetShowsArrowButtons(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsArrowButtons:"), value)
}/* debug [instance_properties/setter]: showsArrowButtons */


// A Boolean value that, together with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/iscontinuous
func (s_ Scrubber) IsContinuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value that, together with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/iscontinuous
func (s_ Scrubber) SetIsContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imagealignment
func (s_ Scrubber) ImageAlignment() ImageAlignment {
	rv := objc.Send[ImageAlignment](s_.ID, objc.Sel("imageAlignment"))
	return rv
}/* debug [instance_properties/getter]: imageAlignment */


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imagealignment
func (s_ Scrubber) SetImageAlignment(value ImageAlignment) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageAlignment:"), value)
}/* debug [instance_properties/setter]: imageAlignment */


// The image view that the scrubber item uses to display its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imageview
func (s_ Scrubber) ImageView() IImageView {
	rv := objc.Send[ImageView](s_.ID, objc.Sel("imageView"))
	return rv
}/* debug [instance_properties/getter]: imageView */


// The image view that the scrubber item uses to display its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imageview
func (s_ Scrubber) SetImageView(value IImageView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageView:"), value)
}/* debug [instance_properties/setter]: imageView */


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ Scrubber) ScrubberContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("scrubberContentSize"))
	return rv
}/* debug [instance_properties/getter]: scrubberContentSize */


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ Scrubber) SetScrubberContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubberContentSize:"), value)
}/* debug [instance_properties/setter]: scrubberContentSize */


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ Scrubber) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}/* debug [instance_properties/getter]: shouldInvalidateLayoutForHighlightChange */


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ Scrubber) SetShouldInvalidateLayoutForHighlightChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForHighlightChange:"), value)
}/* debug [instance_properties/setter]: shouldInvalidateLayoutForHighlightChange */


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ Scrubber) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}/* debug [instance_properties/getter]: shouldInvalidateLayoutForSelectionChange */


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ Scrubber) SetShouldInvalidateLayoutForSelectionChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForSelectionChange:"), value)
}/* debug [instance_properties/setter]: shouldInvalidateLayoutForSelectionChange */


// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/alpha
func (s_ Scrubber) Alpha() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/alpha
func (s_ Scrubber) SetAlpha(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/frame
func (s_ Scrubber) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/frame
func (s_ Scrubber) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/itemindex
func (s_ Scrubber) ItemIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("itemIndex"))
	return rv
}/* debug [instance_properties/getter]: itemIndex */


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/itemindex
func (s_ Scrubber) SetItemIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemIndex:"), value)
}/* debug [instance_properties/setter]: itemIndex */


// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ Scrubber) TextField() ITextField {
	rv := objc.Send[TextField](s_.ID, objc.Sel("textField"))
	return rv
}/* debug [instance_properties/getter]: textField */


// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ Scrubber) SetTextField(value ITextField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextField:"), value)
}/* debug [instance_properties/setter]: textField */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubber */


