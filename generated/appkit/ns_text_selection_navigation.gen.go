// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextSelectionNavigation */


/* debug [class_header]: Header for NSTextSelectionNavigation */
// The class instance for the [TextSelectionNavigation] class.
var (
	TextSelectionNavigationClass     _TextSelectionNavigationClass
	TextSelectionNavigationClassOnce sync.Once
)

func getTextSelectionNavigationClass() _TextSelectionNavigationClass {
	TextSelectionNavigationClassOnce.Do(func() {
		TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}
	})
	return TextSelectionNavigationClass
}

type _TextSelectionNavigationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextSelectionNavigation */
// An interface definition for the [TextSelectionNavigation] class.
type ITextSelectionNavigation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextSelectionNavigation */
	// properties:
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
	TextSelectionDataSource() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextSelectionNavigation */
	// methods:
	DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange
	DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) ITextSelection
	FlushLayoutCache()
	ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection ITextSelection, writingDirection TextSelectionNavigationWritingDirection) unsafe.Pointer
	TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) ITextSelection
	TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity TextSelectionGranularity, point corefoundation.CGPoint, location unsafe.Pointer) ITextSelection
	TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point corefoundation.CGPoint, containerLocation unsafe.Pointer, anchors []TextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds corefoundation.CGRect) []TextSelection
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextSelectionNavigation */
// Alloc allocates a new instance without initialization.
func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextSelectionNavigation) Autorelease() TextSelectionNavigation {
	rv := objc.Send[TextSelectionNavigation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextSelectionNavigation creates a new TextSelectionNavigation instance.
func NewTextSelectionNavigation() TextSelectionNavigation {
	return getTextSelectionNavigationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextSelectionNavigation */
// An interface you use to expose methods for obtaining results from actions performed on text selections.


// An interface you use to expose methods for obtaining results from actions performed on text selections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation
type TextSelectionNavigation struct {
	objectivec.Object
}

// TextSelectionNavigationFrom constructs a [TextSelectionNavigation] from an unsafe.Pointer.
//
// An interface you use to expose methods for obtaining results from actions performed on text selections.
func TextSelectionNavigationFrom(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextSelectionNavigation */

// Creates a new object using the text selection data source you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/init(dataSource:)
func NewTextSelectionNavigationWithDataSource(dataSource unsafe.Pointer) TextSelectionNavigation {
	instance := getTextSelectionNavigationClass().Alloc()
	rv := objc.Send[TextSelectionNavigation](instance.ID, objc.Sel("initWithDataSource:"), dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextSelectionNavigationWithDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextSelectionNavigation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextSelectionNavigation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextSelectionNavigation */

// Returns the ranges for deleting the text based on the current selection and movement arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/deletionRanges(for:direction:destination:allowsDecomposition:)
func (t_ TextSelectionNavigation) DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange {
	rv := objc.Send[[]TextRange](t_.ID, objc.Sel("deletionRangesForTextSelection:direction:destination:allowsDecomposition:"), textSelection, direction, destination, allowsDecomposition)
	return rv
}/* debug [instance_methods/method]: DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition */


// Returns a new selection that results from applying the navigation operations you specify to the text selection you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/destinationSelection(for:direction:destination:extending:confined:)
func (t_ TextSelectionNavigation) DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) ITextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("destinationSelectionForTextSelection:direction:destination:extending:confined:"), textSelection, direction, destination, extending, confined)
	return rv
}/* debug [instance_methods/method]: DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined */


// Flushes cached layout information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/flushLayoutCache()
func (t_ TextSelectionNavigation) FlushLayoutCache() {
	objc.Send[objc.ID](t_.ID, objc.Sel("flushLayoutCache"))
}/* debug [instance_methods/method]: FlushLayoutCache */


// Returns the location for inserting the next input depending on the state of the current and secondary selections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/resolvedInsertionLocation(for:writingDirection:)
func (t_ TextSelectionNavigation) ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection ITextSelection, writingDirection TextSelectionNavigationWritingDirection) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("resolvedInsertionLocationForTextSelection:writingDirection:"), textSelection, writingDirection)
	return rv
}/* debug [instance_methods/method]: ResolvedInsertionLocationForTextSelectionWritingDirection */


// Returns a text selection expanded to the nearest boundaries for the selection granularity and enclosing text selection text ranges you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelection(for:enclosing:)
func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) ITextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("textSelectionForSelectionGranularity:enclosingTextSelection:"), selectionGranularity, textSelection)
	return rv
}/* debug [instance_methods/method]: TextSelectionForSelectionGranularityEnclosingTextSelection */


// Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelection(for:enclosing:inContainerAt:)
func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity TextSelectionGranularity, point corefoundation.CGPoint, location unsafe.Pointer) ITextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("textSelectionForSelectionGranularity:enclosingPoint:inContainerAtLocation:"), selectionGranularity, point, location)
	return rv
}/* debug [instance_methods/method]: TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation */


// Returns an array of text selections produced by a tap or click at the point you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelections(interactingAt:inContainerAt:anchors:modifiers:selecting:bounds:)
func (t_ TextSelectionNavigation) TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point corefoundation.CGPoint, containerLocation unsafe.Pointer, anchors []TextSelection, modifiers TextSelectionNavigationModifier, selecting bool, bounds corefoundation.CGRect) []TextSelection {
	rv := objc.Send[[]TextSelection](t_.ID, objc.Sel("textSelectionsInteractingAtPoint:inContainerAtLocation:anchors:modifiers:selecting:bounds:"), point, containerLocation, anchors, modifiers, selecting, bounds)
	return rv
}/* debug [instance_methods/method]: TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextSelectionNavigation */

// Determines if the instance could produce selections with multiple noncontiguous selections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/allowsNonContiguousRanges
func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsNonContiguousRanges"))
	return rv
}/* debug [instance_properties/getter]: allowsNonContiguousRanges */


// Determines if the instance could produce selections with multiple noncontiguous selections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/allowsNonContiguousRanges
func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsNonContiguousRanges:"), value)
}/* debug [instance_properties/setter]: allowsNonContiguousRanges */


// Determines if the framework rotates the coordinate system to match the layout orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/rotatesCoordinateSystemForLayoutOrientation
func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rotatesCoordinateSystemForLayoutOrientation"))
	return rv
}/* debug [instance_properties/getter]: rotatesCoordinateSystemForLayoutOrientation */


// Determines if the framework rotates the coordinate system to match the layout orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/rotatesCoordinateSystemForLayoutOrientation
func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRotatesCoordinateSystemForLayoutOrientation:"), value)
}/* debug [instance_properties/setter]: rotatesCoordinateSystemForLayoutOrientation */


// The data source associated with this selection navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelectionDataSource
func (t_ TextSelectionNavigation) TextSelectionDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textSelectionDataSource"))
	return rv
}/* debug [instance_properties/getter]: textSelectionDataSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextSelectionNavigation */


