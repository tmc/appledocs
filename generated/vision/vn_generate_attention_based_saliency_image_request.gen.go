// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GenerateAttentionBasedSaliencyImageRequest] class.
var (
	GenerateAttentionBasedSaliencyImageRequestClass     _GenerateAttentionBasedSaliencyImageRequestClass
	GenerateAttentionBasedSaliencyImageRequestClassOnce sync.Once
)

func getGenerateAttentionBasedSaliencyImageRequestClass() _GenerateAttentionBasedSaliencyImageRequestClass {
	GenerateAttentionBasedSaliencyImageRequestClassOnce.Do(func() {
		GenerateAttentionBasedSaliencyImageRequestClass = _GenerateAttentionBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateAttentionBasedSaliencyImageRequest")}
	})
	return GenerateAttentionBasedSaliencyImageRequestClass
}

type _GenerateAttentionBasedSaliencyImageRequestClass struct {
	class objc.Class
}

// An interface definition for the [GenerateAttentionBasedSaliencyImageRequest] class.
type IGenerateAttentionBasedSaliencyImageRequest interface {
	IImageBasedRequest
}

// An object that produces a heat map that identifies the parts of an image most likely to draw attention.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateAttentionBasedSaliencyImageRequest
type GenerateAttentionBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

// GenerateAttentionBasedSaliencyImageRequestFrom constructs a [GenerateAttentionBasedSaliencyImageRequest] from an unsafe.Pointer.
//
// An object that produces a heat map that identifies the parts of an image most likely to draw attention.
func GenerateAttentionBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateAttentionBasedSaliencyImageRequest {
	return GenerateAttentionBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GenerateAttentionBasedSaliencyImageRequestClass) Alloc() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GenerateAttentionBasedSaliencyImageRequestClass) New() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateAttentionBasedSaliencyImageRequest) Init() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateAttentionBasedSaliencyImageRequest) Autorelease() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateAttentionBasedSaliencyImageRequest creates a new GenerateAttentionBasedSaliencyImageRequest instance.
func NewGenerateAttentionBasedSaliencyImageRequest() GenerateAttentionBasedSaliencyImageRequest {
	return getGenerateAttentionBasedSaliencyImageRequestClass().New()
}




