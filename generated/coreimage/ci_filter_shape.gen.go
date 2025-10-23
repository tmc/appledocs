// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterShape] class.
var (
	FilterShapeClass     _FilterShapeClass
	FilterShapeClassOnce sync.Once
)

func getFilterShapeClass() _FilterShapeClass {
	FilterShapeClassOnce.Do(func() {
		FilterShapeClass = _FilterShapeClass{objc.GetClass("CIFilterShape")}
	})
	return FilterShapeClass
}

type _FilterShapeClass struct {
	class objc.Class
}

// An interface definition for the [FilterShape] class.
type IFilterShape interface {
	objectivec.IObject
	// properties:
	Extent() coregraphics.CGRect
	// methods:
	InsetByXY(dx int /* primitive/slice/pointer. */, dy int /* primitive/slice/pointer. */) IFilterShape
	IntersectWithRect(r coregraphics.CGRect) IFilterShape
	IntersectWith(s2 ICIFilterShape) IFilterShape
	TransformByInterior(m coregraphics.CGAffineTransform, flag bool /* primitive/slice/pointer. */) IFilterShape
	UnionWith(s2 ICIFilterShape) IFilterShape
	UnionWithRect(r coregraphics.CGRect) IFilterShape
}

// A description of the bounding shape of a filter and the domain of definition for a filter operation.
//
// You use objects in conjunction with Core Image classes, such as , , and , to create custom filters.


// A description of the bounding shape of a filter and the domain of definition for a filter operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape
type FilterShape struct {
	objectivec.Object
}

// FilterShapeFrom constructs a [FilterShape] from an unsafe.Pointer.
//
// A description of the bounding shape of a filter and the domain of definition for a filter operation.
func FilterShapeFrom(ptr unsafe.Pointer) FilterShape {
	return FilterShape{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterShapeClass) Alloc() FilterShape {
	rv := objc.Send[FilterShape](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterShapeClass) New() FilterShape {
	rv := objc.Send[FilterShape](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterShape) Init() FilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterShape) Autorelease() FilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterShape creates a new FilterShape instance.
func NewFilterShape() FilterShape {
	return getFilterShapeClass().New()
}



// Initializes a filter shape object with a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/init(rect:)
func NewFilterShapeWithRect(r coregraphics.CGRect) FilterShape {
	instance := getFilterShapeClass().Alloc()
	rv := objc.Send[FilterShape](instance.ID, objc.Sel("initWithRect:"), r)
	rv.Autorelease()
	return rv
}



// Creates a filter shape object and initializes it with a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/shapeWithRect:
func (fc _FilterShapeClass) ShapeWithRect(r coregraphics.CGRect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("shapeWithRect:"), r)
	return rv
}


// Modifies a filter shape object so that it is inset by the specified x and y values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/insetBy(x:y:)
func (f_ FilterShape) InsetByXY(dx int /* primitive/slice/pointer. */, dy int /* primitive/slice/pointer. */) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("insetByX:Y:"), dx, dy)
	return rv
}


// Creates a filter shape that represents the intersection of the current filter shape and a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-2o2n8
func (f_ FilterShape) IntersectWithRect(r coregraphics.CGRect) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("intersectWithRect:"), r)
	return rv
}


// Creates a filter shape object that represents the intersection of the current filter shape and the specified filter shape object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/intersect(with:)-8iw
func (f_ FilterShape) IntersectWith(s2 ICIFilterShape) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("intersectWith:"), s2)
	return rv
}


// Creates a filter shape that results from applying a transform to the current filter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/transform(by:interior:)
func (f_ FilterShape) TransformByInterior(m coregraphics.CGAffineTransform, flag bool /* primitive/slice/pointer. */) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("transformBy:interior:"), m, flag)
	return rv
}


// Creates a filter shape that results from the union of the current filter shape and another filter shape object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-52mnd
func (f_ FilterShape) UnionWith(s2 ICIFilterShape) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("unionWith:"), s2)
	return rv
}


// Creates a filter shape that results from the union of the current filter shape and a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/union(with:)-75ebo
func (f_ FilterShape) UnionWithRect(r coregraphics.CGRect) IFilterShape {
	rv := objc.Send[FilterShape](f_.ID, objc.Sel("unionWithRect:"), r)
	return rv
}


// The extent of the filter shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterShape/extent
func (f_ FilterShape) Extent() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](f_.ID, objc.Sel("extent"))
	return rv
}


