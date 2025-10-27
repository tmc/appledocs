// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore


// Enum types and constants

// CAAutoresizingMask - These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask
type CAAutoresizingMask uint

const (
	// kCALayerNotSizable - The receiver cannot be resized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/kCALayerNotSizable
	kCALayerNotSizable CAAutoresizingMask = 0
	// kCALayerHeightSizable - The receiver’s height is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerHeightSizable
	kCALayerHeightSizable CAAutoresizingMask = 0
	// kCALayerMaxXMargin - The right margin between the receiver and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerMaxXMargin
	kCALayerMaxXMargin CAAutoresizingMask = 0
	// kCALayerMaxYMargin - The top margin between the receiver and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerMaxYMargin
	kCALayerMaxYMargin CAAutoresizingMask = 0
	// kCALayerMinXMargin - The left margin between the receiver and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerMinXMargin
	kCALayerMinXMargin CAAutoresizingMask = 0
	// kCALayerMinYMargin - The bottom margin between the receiver and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerMinYMargin
	kCALayerMinYMargin CAAutoresizingMask = 0
	// kCALayerWidthSizable - The receiver’s width is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/layerWidthSizable
	kCALayerWidthSizable CAAutoresizingMask = 0
)


// CAConstraintAttribute - The constraint attribute type.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute
type CAConstraintAttribute uint

const (
	// kCAConstraintHeight - The height of a layer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/height
	kCAConstraintHeight CAConstraintAttribute = 0
	// kCAConstraintMaxX - The right edge of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/maxX
	kCAConstraintMaxX CAConstraintAttribute = 0
	// kCAConstraintMaxY - The top edge of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/maxY
	kCAConstraintMaxY CAConstraintAttribute = 0
	// kCAConstraintMidX - The horizontal location of the center of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/midX
	kCAConstraintMidX CAConstraintAttribute = 0
	// kCAConstraintMidY - The vertical location of the center of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/midY
	kCAConstraintMidY CAConstraintAttribute = 0
	// kCAConstraintMinX - The left edge of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/minX
	kCAConstraintMinX CAConstraintAttribute = 0
	// kCAConstraintMinY - The bottom edge of a layer’s frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/minY
	kCAConstraintMinY CAConstraintAttribute = 0
	// kCAConstraintWidth - The width of a layer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/width
	kCAConstraintWidth CAConstraintAttribute = 0
)


// CACornerMask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACornerMask
type CACornerMask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACornerMask/layerMaxXMaxYCorner
	kCALayerMaxXMaxYCorner CACornerMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACornerMask/layerMaxXMinYCorner
	kCALayerMaxXMinYCorner CACornerMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACornerMask/layerMinXMaxYCorner
	kCALayerMinXMaxYCorner CACornerMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACornerMask/layerMinXMinYCorner
	kCALayerMinXMinYCorner CACornerMask = 0
)


// CAEdgeAntialiasingMask - This mask is used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask
type CAEdgeAntialiasingMask uint

const (
	// kCALayerBottomEdge - Specifies that the bottom edge of the receiver’s content should be antialiased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask/layerBottomEdge
	kCALayerBottomEdge CAEdgeAntialiasingMask = 0
	// kCALayerLeftEdge - Specifies that the left edge of the receiver’s content should be antialiased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask/layerLeftEdge
	kCALayerLeftEdge CAEdgeAntialiasingMask = 0
	// kCALayerRightEdge - Specifies that the right edge of the receiver’s content should be antialiased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask/layerRightEdge
	kCALayerRightEdge CAEdgeAntialiasingMask = 0
	// kCALayerTopEdge - Specifies that the top edge of the receiver’s content should be antialiased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask/layerTopEdge
	kCALayerTopEdge CAEdgeAntialiasingMask = 0
)


