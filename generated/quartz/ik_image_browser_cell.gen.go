// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKImageBrowserCell */


/* debug [class_header]: Header for IKImageBrowserCell */
// The class instance for the [IKImageBrowserCell] class.
var (
	IKImageBrowserCellClass     _IKImageBrowserCellClass
	IKImageBrowserCellClassOnce sync.Once
)

func getIKImageBrowserCellClass() _IKImageBrowserCellClass {
	IKImageBrowserCellClassOnce.Do(func() {
		IKImageBrowserCellClass = _IKImageBrowserCellClass{objc.GetClass("IKImageBrowserCell")}
	})
	return IKImageBrowserCellClass
}

type _IKImageBrowserCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKImageBrowserCell */
// An interface definition for the [IKImageBrowserCell] class.
type IIKImageBrowserCell interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IKImageBrowserCell */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKImageBrowserCell */
	// methods:
	CellState() objectivec.IObject
	Frame() Rect /* not a class type */
	ImageAlignment() ImageAlignment /* not a class type */
	ImageBrowserView() IKImageBrowserView
	ImageContainerFrame() Rect /* not a class type */
	ImageFrame() Rect /* not a class type */
	IndexOfRepresentedItem() uint
	IsSelected() bool
	LayerForType(type_ objc.IObject /* cross-framework: NSString */) avfoundation.Layer
	Opacity() float64
	RepresentedItem() objc.ID
	SelectionFrame() Rect /* not a class type */
	SubtitleFrame() Rect /* not a class type */
	TitleFrame() Rect /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKImageBrowserCell */
// Alloc allocates a new instance without initialization.
func (ic _IKImageBrowserCellClass) Alloc() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKImageBrowserCellClass) New() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageBrowserCell) Init() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageBrowserCell) Autorelease() IKImageBrowserCell {
	rv := objc.Send[IKImageBrowserCell](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageBrowserCell creates a new IKImageBrowserCell instance.
func NewIKImageBrowserCell() IKImageBrowserCell {
	return getIKImageBrowserCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKImageBrowserCell */
// A class used to display a cell.
//
// class that is used to display a cell conforming to the in an .


// A class used to display a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell
type IKImageBrowserCell struct {
	objectivec.Object
}

// IKImageBrowserCellFrom constructs a [IKImageBrowserCell] from an unsafe.Pointer.
//
// A class used to display a cell.
func IKImageBrowserCellFrom(ptr unsafe.Pointer) IKImageBrowserCell {
	return IKImageBrowserCell{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKImageBrowserCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKImageBrowserCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKImageBrowserCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKImageBrowserCell */

// Returns the current cell state of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/cellState()
func (i_ IKImageBrowserCell) CellState() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("cellState"))
	return rv
}/* debug [instance_methods/method]: CellState */


// Returns the receiver’s frame rectangle, which defines its position in its .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/frame()
func (i_ IKImageBrowserCell) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_methods/method]: Frame */


// Returns the position of the cell’s image in the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/imageAlignment()
func (i_ IKImageBrowserCell) ImageAlignment() ImageAlignment /* not a class type */ {
	rv := objc.Send[ImageAlignment](i_.ID, objc.Sel("imageAlignment"))
	return rv
}/* debug [instance_methods/method]: ImageAlignment */


// Returns the view the receiver uses to display the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/imageBrowserView()
func (i_ IKImageBrowserCell) ImageBrowserView() IKImageBrowserView {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("imageBrowserView"))
	return rv
}/* debug [instance_methods/method]: ImageBrowserView */


// Returns the receiver’s image container frame rectangle, which defines the position of the container of the thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/imageContainerFrame()
func (i_ IKImageBrowserCell) ImageContainerFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("imageContainerFrame"))
	return rv
}/* debug [instance_methods/method]: ImageContainerFrame */


// Returns the receiver’s image frame rectangle, which defines the position of the thumbnail in its .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/imageFrame()
func (i_ IKImageBrowserCell) ImageFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("imageFrame"))
	return rv
}/* debug [instance_methods/method]: ImageFrame */


// Returns the index of the receiver’s represented object in the datasource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/indexOfRepresentedItem()
func (i_ IKImageBrowserCell) IndexOfRepresentedItem() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexOfRepresentedItem"))
	return rv
}/* debug [instance_methods/method]: IndexOfRepresentedItem */


// Returns whether the cell is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/isSelected()
func (i_ IKImageBrowserCell) IsSelected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isSelected"))
	return rv
}/* debug [instance_methods/method]: IsSelected */


// Returns a layer for the specified position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/layer(forType:)
func (i_ IKImageBrowserCell) LayerForType(type_ objc.IObject /* cross-framework: NSString */) avfoundation.Layer {
	rv := objc.Send[avfoundation.Layer](i_.ID, objc.Sel("layerForType:"), type_)
	return rv
}/* debug [instance_methods/method]: LayerForType */


// Returns the opacity of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/opacity()
func (i_ IKImageBrowserCell) Opacity() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("opacity"))
	return rv
}/* debug [instance_methods/method]: Opacity */


// Returns the receiver’s represented object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/representedItem()
func (i_ IKImageBrowserCell) RepresentedItem() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("representedItem"))
	return rv
}/* debug [instance_methods/method]: RepresentedItem */


// Returns the receiver’s selection frame rectangle, which defines the position of the selection rectangle in its .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/selectionFrame()
func (i_ IKImageBrowserCell) SelectionFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("selectionFrame"))
	return rv
}/* debug [instance_methods/method]: SelectionFrame */


// Returns the receiver’s subtitle frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/subtitleFrame()
func (i_ IKImageBrowserCell) SubtitleFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("subtitleFrame"))
	return rv
}/* debug [instance_methods/method]: SubtitleFrame */


// Returns the receiver’s title frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageBrowserCell/titleFrame()
func (i_ IKImageBrowserCell) TitleFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](i_.ID, objc.Sel("titleFrame"))
	return rv
}/* debug [instance_methods/method]: TitleFrame */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKImageBrowserCell */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKImageBrowserCell */



