// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CPaddingLayer] class.
var (
	CPaddingLayerClass     _CPaddingLayerClass
	CPaddingLayerClassOnce sync.Once
)

func getCPaddingLayerClass() _CPaddingLayerClass {
	CPaddingLayerClassOnce.Do(func() {
		CPaddingLayerClass = _CPaddingLayerClass{objc.GetClass("MLCPaddingLayer")}
	})
	return CPaddingLayerClass
}

type _CPaddingLayerClass struct {
	class objc.Class
}

// An interface definition for the [CPaddingLayer] class.
type ICPaddingLayer interface {
	ICLayer
	ConstantValue() float32
	SetConstantValue(value float32)
	PaddingBottom() int
	SetPaddingBottom(value int)
	PaddingLeft() int
	SetPaddingLeft(value int)
	PaddingRight() int
	SetPaddingRight(value int)
	PaddingTop() int
	SetPaddingTop(value int)
	PaddingType() CPaddingType
	SetPaddingType(value CPaddingType)
}

// A layer that pads a tensor with the padding sizes you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer
type CPaddingLayer struct {
	CLayer
}

// CPaddingLayerFrom constructs a [CPaddingLayer] from an unsafe.Pointer.
//
// A layer that pads a tensor with the padding sizes you specify.
func CPaddingLayerFrom(ptr unsafe.Pointer) CPaddingLayer {
	return CPaddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CPaddingLayerClass) Alloc() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CPaddingLayerClass) New() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPaddingLayer) Init() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPaddingLayer) Autorelease() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPaddingLayer creates a new CPaddingLayer instance.
func NewCPaddingLayer() CPaddingLayer {
	return getCPaddingLayerClass().New()
}


// The constant value you use if padding type is constant.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/constantvalue
func (c_ CPaddingLayer) ConstantValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("constantValue"))
	return rv
}


// SetConstantValue sets the value of the constantValue property.
// The constant value you use if padding type is constant.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/constantvalue
func (c_ CPaddingLayer) SetConstantValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantValue:"), value)
}

// The bottom padding size.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingbottom
func (c_ CPaddingLayer) PaddingBottom() int {
	rv := objc.Send[int](c_.ID, objc.Sel("paddingBottom"))
	return rv
}


// SetPaddingBottom sets the value of the paddingBottom property.
// The bottom padding size.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingbottom
func (c_ CPaddingLayer) SetPaddingBottom(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPaddingBottom:"), value)
}

// The left padding size.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingleft
func (c_ CPaddingLayer) PaddingLeft() int {
	rv := objc.Send[int](c_.ID, objc.Sel("paddingLeft"))
	return rv
}


// SetPaddingLeft sets the value of the paddingLeft property.
// The left padding size.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingleft
func (c_ CPaddingLayer) SetPaddingLeft(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPaddingLeft:"), value)
}

// The right padding size.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingright
func (c_ CPaddingLayer) PaddingRight() int {
	rv := objc.Send[int](c_.ID, objc.Sel("paddingRight"))
	return rv
}


// SetPaddingRight sets the value of the paddingRight property.
// The right padding size.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingright
func (c_ CPaddingLayer) SetPaddingRight(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPaddingRight:"), value)
}

// The top padding size.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingtop
func (c_ CPaddingLayer) PaddingTop() int {
	rv := objc.Send[int](c_.ID, objc.Sel("paddingTop"))
	return rv
}


// SetPaddingTop sets the value of the paddingTop property.
// The top padding size.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingtop
func (c_ CPaddingLayer) SetPaddingTop(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPaddingTop:"), value)
}

// The padding type.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingtype
func (c_ CPaddingLayer) PaddingType() CPaddingType {
	rv := objc.Send[CPaddingType](c_.ID, objc.Sel("paddingType"))
	return rv
}


// SetPaddingType sets the value of the paddingType property.
// The padding type.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcpaddinglayer/paddingtype
func (c_ CPaddingLayer) SetPaddingType(value CPaddingType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPaddingType:"), value)
}



