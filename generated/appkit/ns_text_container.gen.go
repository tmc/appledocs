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

// The class instance for the [TextContainer] class.
var (
	TextContainerClass     _TextContainerClass
	TextContainerClassOnce sync.Once
)

func getTextContainerClass() _TextContainerClass {
	TextContainerClassOnce.Do(func() {
		TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}
	})
	return TextContainerClass
}

type _TextContainerClass struct {
	class objc.Class
}

// An interface definition for the [TextContainer] class.
type ITextContainer interface {
	objectivec.IObject
	// properties:
	ContainerSize() objc.IObject /* cross-framework: Size */
	SetContainerSize(value objc.IObject /* cross-framework: Size */)
	ExclusionPaths() []BezierPath
	SetExclusionPaths(value []BezierPath)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	SimpleRectangularTextContainer() bool
	LayoutManager() ILayoutManager
	SetLayoutManager(value ILayoutManager)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	TextLayoutManager() ITextLayoutManager
	TextView() ITextView
	SetTextView(value ITextView)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	IsSimpleRectangularTextContainer() bool
	SetIsSimpleRectangularTextContainer(value bool)
	// methods:
	LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect corefoundation.CGRect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect corefoundation.CGRect) corefoundation.CGRect
	LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect(proposedRect objc.IObject /* cross-framework: Rect */, sweepDirection LineSweepDirection, movementDirection LineMovementDirection, remainingRect RectPointer /* not a class type */) objc.IObject /* cross-framework: Rect */
	ReplaceLayoutManager(newLayoutManager ILayoutManager)
}

// A region where text layout occurs.
//
// An uses to determine where to break lines, lay out portions of text, and so on. An object typically defines rectangular regions, but you can define exclusion paths inside the text container to create regions where text doesn’t flow. You can also subclass to create text containers with nonrectangular regions, such as circular regions, regions with holes in them, or regions that flow alongside graphics. You can access instances of the , , and classes from threads other than the main thread as long as the app guarantees access from only one thread at a time.


// A region where text layout occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer
type TextContainer struct {
	objectivec.Object
}

// TextContainerFrom constructs a [TextContainer] from an unsafe.Pointer.
//
// A region where text layout occurs.
func TextContainerFrom(ptr unsafe.Pointer) TextContainer {
	return TextContainer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextContainerClass) New() TextContainer {
	rv := objc.Send[TextContainer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextContainer) Init() TextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextContainer) Autorelease() TextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextContainer creates a new TextContainer instance.
func NewTextContainer() TextContainer {
	return getTextContainerClass().New()
}



// Creates a text container from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(coder:)
func NewTextContainerWithCoder(coder foundation.Coder) TextContainer {
	instance := getTextContainerClass().Alloc()
	rv := objc.Send[TextContainer](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a text container with a specified bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(containerSize:)
func NewTextContainerWithContainerSize(aContainerSize objc.IObject /* cross-framework: Size */) TextContainer {
	instance := getTextContainerClass().Alloc()
	rv := objc.Send[TextContainer](instance.ID, objc.Sel("initWithContainerSize:"), aContainerSize)
	rv.Autorelease()
	return rv
}


// Initializes a text container with a specified bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(size:)
func NewTextContainerWithSize(size corefoundation.CGSize) TextContainer {
	instance := getTextContainerClass().Alloc()
	rv := objc.Send[TextContainer](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}



// Returns the bounds of a line fragment rectangle inside the text container for the proposed rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentRect(forProposedRect:at:writingDirection:remaining:)
func (t_ TextContainer) LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect corefoundation.CGRect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("lineFragmentRectForProposedRect:atIndex:writingDirection:remainingRect:"), proposedRect, characterIndex, baseWritingDirection, remainingRect)
	return rv
}


// Calculates and returns the longest rectangle available in the proposed rectangle for displaying text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentRect(forProposedRect:sweepDirection:movementDirection:remaining:)
func (t_ TextContainer) LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect(proposedRect objc.IObject /* cross-framework: Rect */, sweepDirection LineSweepDirection, movementDirection LineMovementDirection, remainingRect RectPointer /* not a class type */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](t_.ID, objc.Sel("lineFragmentRectForProposedRect:sweepDirection:movementDirection:remainingRect:"), proposedRect, sweepDirection, movementDirection, remainingRect)
	return rv
}


// Replaces the layout manager for the group of text system objects that contains the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/replaceLayoutManager(_:)
func (t_ TextContainer) ReplaceLayoutManager(newLayoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceLayoutManager:"), newLayoutManager)
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/containerSize
func (t_ TextContainer) ContainerSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("containerSize"))
	return rv
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/containerSize
func (t_ TextContainer) SetContainerSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContainerSize:"), value)
}


// An array of path objects that represents the regions where text doesn’t display in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/exclusionPaths
func (t_ TextContainer) ExclusionPaths() []BezierPath {
	rv := objc.Send[[]BezierPath](t_.ID, objc.Sel("exclusionPaths"))
	return rv
}


// An array of path objects that represents the regions where text doesn’t display in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/exclusionPaths
func (t_ TextContainer) SetExclusionPaths(value []BezierPath) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setExclusionPaths:"), nsArray)
}


// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/heightTracksTextView
func (t_ TextContainer) HeightTracksTextView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("heightTracksTextView"))
	return rv
}


// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/heightTracksTextView
func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeightTracksTextView:"), value)
}


// A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/isSimpleRectangularTextContainer
func (t_ TextContainer) SimpleRectangularTextContainer() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("simpleRectangularTextContainer"))
	return rv
}


// The text container’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager
func (t_ TextContainer) LayoutManager() ILayoutManager {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// The text container’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager
func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}


// The behavior of the last line inside the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineBreakMode
func (t_ TextContainer) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](t_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The behavior of the last line inside the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineBreakMode
func (t_ TextContainer) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The value for the text inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentPadding
func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// The value for the text inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentPadding
func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// The maximum number of lines that the text container can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/maximumNumberOfLines
func (t_ TextContainer) MaximumNumberOfLines() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}


// The maximum number of lines that the text container can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/maximumNumberOfLines
func (t_ TextContainer) SetMaximumNumberOfLines(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/size
func (t_ TextContainer) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("size"))
	return rv
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/size
func (t_ TextContainer) SetSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/textLayoutManager
func (t_ TextContainer) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// The text container’s text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/textView
func (t_ TextContainer) TextView() ITextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("textView"))
	return rv
}


// The text container’s text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/textView
func (t_ TextContainer) SetTextView(value ITextView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextView:"), value)
}


// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/widthTracksTextView
func (t_ TextContainer) WidthTracksTextView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("widthTracksTextView"))
	return rv
}


// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/widthTracksTextView
func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidthTracksTextView:"), value)
}


// A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/issimplerectangulartextcontainer
func (t_ TextContainer) IsSimpleRectangularTextContainer() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSimpleRectangularTextContainer"))
	return rv
}


// A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/issimplerectangulartextcontainer
func (t_ TextContainer) SetIsSimpleRectangularTextContainer(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSimpleRectangularTextContainer:"), value)
}


