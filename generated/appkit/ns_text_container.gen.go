// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	LayoutManager() NSLayoutManager
	SetLayoutManager(value ILayoutManager)
	ContainerSize() coregraphics.CGSize
	SetContainerSize(value coregraphics.CGSize)
	ExclusionPaths() NSBezierPath
	SetExclusionPaths(value IBezierPath)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	IsSimpleRectangularTextContainer() bool
	SetIsSimpleRectangularTextContainer(value bool)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	Size() coregraphics.CGSize
	SetSize(value coregraphics.CGSize)
	TextLayoutManager() NSTextLayoutManager
	SetTextLayoutManager(value ITextLayoutManager)
	TextView() NSTextView
	SetTextView(value ITextView)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
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



// The text container’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager

func (t_ TextContainer) LayoutManager() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// The text container’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager

func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/containersize

func (t_ TextContainer) ContainerSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("containerSize"))
	return rv
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/containersize

func (t_ TextContainer) SetContainerSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContainerSize:"), value)
}


// An array of path objects that represents the regions where text doesn’t display in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/exclusionpaths

func (t_ TextContainer) ExclusionPaths() NSBezierPath {
	rv := objc.Send[NSBezierPath](t_.ID, objc.Sel("exclusionPaths"))
	return rv
}


// An array of path objects that represents the regions where text doesn’t display in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/exclusionpaths

func (t_ TextContainer) SetExclusionPaths(value IBezierPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExclusionPaths:"), value)
}


// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/heighttrackstextview

func (t_ TextContainer) HeightTracksTextView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("heightTracksTextView"))
	return rv
}


// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/heighttrackstextview

func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeightTracksTextView:"), value)
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


// The behavior of the last line inside the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/linebreakmode

func (t_ TextContainer) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](t_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The behavior of the last line inside the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/linebreakmode

func (t_ TextContainer) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The value for the text inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/linefragmentpadding

func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// The value for the text inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/linefragmentpadding

func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// The maximum number of lines that the text container can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/maximumnumberoflines

func (t_ TextContainer) MaximumNumberOfLines() int {
	rv := objc.Send[int](t_.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}


// The maximum number of lines that the text container can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/maximumnumberoflines

func (t_ TextContainer) SetMaximumNumberOfLines(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/size

func (t_ TextContainer) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("size"))
	return rv
}


// The size of the text container’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/size

func (t_ TextContainer) SetSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/textlayoutmanager

func (t_ TextContainer) TextLayoutManager() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/textlayoutmanager

func (t_ TextContainer) SetTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}


// The text container’s text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/textview

func (t_ TextContainer) TextView() NSTextView {
	rv := objc.Send[NSTextView](t_.ID, objc.Sel("textView"))
	return rv
}


// The text container’s text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/textview

func (t_ TextContainer) SetTextView(value ITextView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextView:"), value)
}


// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/widthtrackstextview

func (t_ TextContainer) WidthTracksTextView() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("widthTracksTextView"))
	return rv
}


// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontainer/widthtrackstextview

func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidthTracksTextView:"), value)
}



