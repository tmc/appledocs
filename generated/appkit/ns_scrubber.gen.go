// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

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

// An interface definition for the [Scrubber] class.
type IScrubber interface {
	IView
	// properties:
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	BackgroundView() IView
	SetBackgroundView(value IView)
	DataSource() ScrubberDataSource /* not a class type */
	SetDataSource(value ScrubberDataSource /* not a class type */)
	Delegate() ScrubberDelegate /* not a class type */
	SetDelegate(value ScrubberDelegate /* not a class type */)
	FloatsSelectionViews() bool
	SetFloatsSelectionViews(value bool)
	HighlightedIndex() int
	SetHighlightedIndex(value int)
	IsContinuous() bool
	SetIsContinuous(value bool)
	ItemAlignment() unsafe.Pointer
	SetItemAlignment(value unsafe.Pointer)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	NumberOfItems() int
	SetNumberOfItems(value int)
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
	ImageAlignment() ImageAlignment /* not a class type */
	SetImageAlignment(value ImageAlignment /* not a class type */)
	ImageView() objc.IObject /* cross-framework: ImageView */
	SetImageView(value objc.IObject /* cross-framework: ImageView */)
	ScrubberContentSize() objc.IObject /* cross-framework: Size */
	SetScrubberContentSize(value objc.IObject /* cross-framework: Size */)
	ShouldInvalidateLayoutForHighlightChange() bool
	SetShouldInvalidateLayoutForHighlightChange(value bool)
	ShouldInvalidateLayoutForSelectionChange() bool
	SetShouldInvalidateLayoutForSelectionChange(value bool)
	Alpha() float64
	SetAlpha(value float64)
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	ItemIndex() int
	SetItemIndex(value int)
	TextField() ITextField
	SetTextField(value ITextField)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _ScrubberClass) Alloc() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The color displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/backgroundcolor
func (s_ Scrubber) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](s_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/backgroundcolor
func (s_ Scrubber) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}


// A view that is displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/backgroundview
func (s_ Scrubber) BackgroundView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("backgroundView"))
	return rv
}


// A view that is displayed behind the scrubber content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/backgroundview
func (s_ Scrubber) SetBackgroundView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundView:"), value)
}


// The object that provides the data for the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/datasource
func (s_ Scrubber) DataSource() ScrubberDataSource /* not a class type */ {
	rv := objc.Send[ScrubberDataSource](s_.ID, objc.Sel("dataSource"))
	return rv
}


// The object that provides the data for the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/datasource
func (s_ Scrubber) SetDataSource(value ScrubberDataSource /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDataSource:"), value)
}


// The object that acts as the delegate of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/delegate
func (s_ Scrubber) Delegate() ScrubberDelegate /* not a class type */ {
	rv := objc.Send[ScrubberDelegate](s_.ID, objc.Sel("delegate"))
	return rv
}


// The object that acts as the delegate of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/delegate
func (s_ Scrubber) SetDelegate(value ScrubberDelegate /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/floatsselectionviews
func (s_ Scrubber) FloatsSelectionViews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("floatsSelectionViews"))
	return rv
}


// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/floatsselectionviews
func (s_ Scrubber) SetFloatsSelectionViews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFloatsSelectionViews:"), value)
}


// The index of the highlighted item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/highlightedindex
func (s_ Scrubber) HighlightedIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("highlightedIndex"))
	return rv
}


// The index of the highlighted item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/highlightedindex
func (s_ Scrubber) SetHighlightedIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHighlightedIndex:"), value)
}


// A Boolean value that, together with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/iscontinuous
func (s_ Scrubber) IsContinuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value that, together with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/iscontinuous
func (s_ Scrubber) SetIsContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsContinuous:"), value)
}


// A setting that specifies the snapping behavior of items in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/itemalignment
func (s_ Scrubber) ItemAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("itemAlignment"))
	return rv
}


// A setting that specifies the snapping behavior of items in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/itemalignment
func (s_ Scrubber) SetItemAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemAlignment:"), value)
}


// A setting that determines whether interaction with the scrubber is fixed or free.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/mode-swift.property
func (s_ Scrubber) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("mode"))
	return rv
}


// A setting that determines whether interaction with the scrubber is fixed or free.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/mode-swift.property
func (s_ Scrubber) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMode:"), value)
}


// The number of items represented by the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/numberofitems
func (s_ Scrubber) NumberOfItems() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of items represented by the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/numberofitems
func (s_ Scrubber) SetNumberOfItems(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfItems:"), value)
}


// An object used to describe the layout of items within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/scrubberlayout
func (s_ Scrubber) ScrubberLayout() IScrubberLayout {
	rv := objc.Send[ScrubberLayout](s_.ID, objc.Sel("scrubberLayout"))
	return rv
}


// An object used to describe the layout of items within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/scrubberlayout
func (s_ Scrubber) SetScrubberLayout(value IScrubberLayout) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubberLayout:"), value)
}


// The index of the selected item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectedindex
func (s_ Scrubber) SelectedIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedIndex"))
	return rv
}


// The index of the selected item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectedindex
func (s_ Scrubber) SetSelectedIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedIndex:"), value)
}


// The style applied to the background of selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectionbackgroundstyle
func (s_ Scrubber) SelectionBackgroundStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("selectionBackgroundStyle"))
	return rv
}


// The style applied to the background of selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectionbackgroundstyle
func (s_ Scrubber) SetSelectionBackgroundStyle(value IScrubberSelectionStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectionBackgroundStyle:"), value)
}


// The style overlaid on selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectionoverlaystyle
func (s_ Scrubber) SelectionOverlayStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("selectionOverlayStyle"))
	return rv
}


// The style overlaid on selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/selectionoverlaystyle
func (s_ Scrubber) SetSelectionOverlayStyle(value IScrubberSelectionStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectionOverlayStyle:"), value)
}


// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/showsadditionalcontentindicators
func (s_ Scrubber) ShowsAdditionalContentIndicators() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsAdditionalContentIndicators"))
	return rv
}


// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/showsadditionalcontentindicators
func (s_ Scrubber) SetShowsAdditionalContentIndicators(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsAdditionalContentIndicators:"), value)
}


// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/showsarrowbuttons
func (s_ Scrubber) ShowsArrowButtons() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsArrowButtons"))
	return rv
}


// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubber/showsarrowbuttons
func (s_ Scrubber) SetShowsArrowButtons(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsArrowButtons:"), value)
}


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imagealignment
func (s_ Scrubber) ImageAlignment() ImageAlignment /* not a class type */ {
	rv := objc.Send[ImageAlignment](s_.ID, objc.Sel("imageAlignment"))
	return rv
}


// The alignment of the image within the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imagealignment
func (s_ Scrubber) SetImageAlignment(value ImageAlignment /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageAlignment:"), value)
}


// The image view that the scrubber item uses to display its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imageview
func (s_ Scrubber) ImageView() objc.IObject /* cross-framework: ImageView */ {
	rv := objc.Send[ImageView](s_.ID, objc.Sel("imageView"))
	return rv
}


// The image view that the scrubber item uses to display its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imageview
func (s_ Scrubber) SetImageView(value objc.IObject /* cross-framework: ImageView */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageView:"), value)
}


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ Scrubber) ScrubberContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](s_.ID, objc.Sel("scrubberContentSize"))
	return rv
}


// The size required to contain all elements within the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s_ Scrubber) SetScrubberContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrubberContentSize:"), value)
}


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ Scrubber) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}


// Determines whether the scrubber should refresh its layout when an item is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s_ Scrubber) SetShouldInvalidateLayoutForHighlightChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForHighlightChange:"), value)
}


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ Scrubber) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}


// Determines whether the scrubber should refresh its layout when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s_ Scrubber) SetShouldInvalidateLayoutForSelectionChange(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShouldInvalidateLayoutForSelectionChange:"), value)
}


// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/alpha
func (s_ Scrubber) Alpha() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("alpha"))
	return rv
}


// The item’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/alpha
func (s_ Scrubber) SetAlpha(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlpha:"), value)
}


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/frame
func (s_ Scrubber) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](s_.ID, objc.Sel("frame"))
	return rv
}


// The frame of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/frame
func (s_ Scrubber) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrame:"), value)
}


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/itemindex
func (s_ Scrubber) ItemIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("itemIndex"))
	return rv
}


// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/itemindex
func (s_ Scrubber) SetItemIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemIndex:"), value)
}


// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ Scrubber) TextField() ITextField {
	rv := objc.Send[TextField](s_.ID, objc.Sel("textField"))
	return rv
}


// The text field that the scrubber item uses to display its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s_ Scrubber) SetTextField(value ITextField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextField:"), value)
}



