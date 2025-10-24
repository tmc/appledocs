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


