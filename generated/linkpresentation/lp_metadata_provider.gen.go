// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

package linkpresentation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LPMetadataProvider] class.
var (
	LPMetadataProviderClass     _LPMetadataProviderClass
	LPMetadataProviderClassOnce sync.Once
)

func getLPMetadataProviderClass() _LPMetadataProviderClass {
	LPMetadataProviderClassOnce.Do(func() {
		LPMetadataProviderClass = _LPMetadataProviderClass{objc.GetClass("LPMetadataProvider")}
	})
	return LPMetadataProviderClass
}

type _LPMetadataProviderClass struct {
	class objc.Class
}

// An interface definition for the [LPMetadataProvider] class.
type ILPMetadataProvider interface {
	objectivec.IObject
	Cancel()
	StartFetchingMetadataForURLCompletionHandler(URL foundation.IURL, completionHandler unsafe.Pointer)
	StartFetchingMetadataForRequestCompletionHandler(request foundation.IURLRequest, completionHandler unsafe.Pointer)
	ShouldFetchSubresources() bool
	SetShouldFetchSubresources(value bool)
	Timeout() foundation.TimeInterval
	SetTimeout(value foundation.ITimeInterval)
}

// An object that retrieves metadata for a URL.
//
// Use to fetch metadata for a URL, including its title, icon, and image or video links. All properties on the resulting instance are optional.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider
type LPMetadataProvider struct {
	objectivec.Object
}

// LPMetadataProviderFrom constructs a [LPMetadataProvider] from an unsafe.Pointer.
//
// An object that retrieves metadata for a URL.
func LPMetadataProviderFrom(ptr unsafe.Pointer) LPMetadataProvider {
	return LPMetadataProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LPMetadataProviderClass) Alloc() LPMetadataProvider {
	rv := objc.Send[LPMetadataProvider](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LPMetadataProviderClass) New() LPMetadataProvider {
	rv := objc.Send[LPMetadataProvider](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LPMetadataProvider) Init() LPMetadataProvider {
	rv := objc.Send[LPMetadataProvider](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LPMetadataProvider) Autorelease() LPMetadataProvider {
	rv := objc.Send[LPMetadataProvider](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLPMetadataProvider creates a new LPMetadataProvider instance.
func NewLPMetadataProvider() LPMetadataProvider {
	return getLPMetadataProviderClass().New()
}


// Cancels a metadata request.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/cancel()
func (l_ LPMetadataProvider) Cancel() {
	objc.Send[objc.ID](l_.ID, objc.Sel("cancel"))
}

// Fetches metadata for the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/startFetchingMetadata(for:completionHandler:)-54z5i
func (l_ LPMetadataProvider) StartFetchingMetadataForURLCompletionHandler(URL foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startFetchingMetadataForURL:completionHandler:"), URL, completionHandler)
}

// Fetches metadata for the given .
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/startFetchingMetadata(for:completionHandler:)-9e6s8
func (l_ LPMetadataProvider) StartFetchingMetadataForRequestCompletionHandler(request foundation.IURLRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startFetchingMetadataForRequest:completionHandler:"), request, completionHandler)
}

// A Boolean value indicating whether to download subresources specified by the metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/shouldFetchSubresources
func (l_ LPMetadataProvider) ShouldFetchSubresources() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldFetchSubresources"))
	return rv
}


// SetShouldFetchSubresources sets the value of the shouldFetchSubresources property.
// A Boolean value indicating whether to download subresources specified by the metadata.

//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/shouldFetchSubresources
func (l_ LPMetadataProvider) SetShouldFetchSubresources(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShouldFetchSubresources:"), value)
}

// The time interval after which the request automatically fails if it hasn’t already completed.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/timeout
func (l_ LPMetadataProvider) Timeout() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](l_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
// The time interval after which the request automatically fails if it hasn’t already completed.

//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/timeout
func (l_ LPMetadataProvider) SetTimeout(value foundation.ITimeInterval) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimeout:"), value)
}




