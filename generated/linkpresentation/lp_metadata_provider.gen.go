// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

package linkpresentation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LPMetadataProvider */


/* debug [class_header]: Header for LPMetadataProvider */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LPMetadataProvider */
// An interface definition for the [LPMetadataProvider] class.
type ILPMetadataProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LPMetadataProvider */
	// properties:
	ShouldFetchSubresources() bool
	SetShouldFetchSubresources(value bool)
	Timeout() float64
	SetTimeout(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LPMetadataProvider */
	// methods:
	Cancel()
	StartFetchingMetadataForURLCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	StartFetchingMetadataForRequestCompletionHandler(request foundation.URLRequest, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LPMetadataProvider */
// Alloc allocates a new instance without initialization.
func (lc _LPMetadataProviderClass) Alloc() LPMetadataProvider {
	rv := objc.Send[LPMetadataProvider](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LPMetadataProvider */
// An object that retrieves metadata for a URL.
//
// Use to fetch metadata for a URL, including its title, icon, and image or video links. All properties on the resulting instance are optional.


// An object that retrieves metadata for a URL.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LPMetadataProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LPMetadataProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LPMetadataProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LPMetadataProvider */

// Cancels a metadata request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/cancel()
func (l_ LPMetadataProvider) Cancel() {
	objc.Send[objc.ID](l_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Fetches metadata for the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/startFetchingMetadata(for:completionHandler:)-54z5i
func (l_ LPMetadataProvider) StartFetchingMetadataForURLCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startFetchingMetadataForURL:completionHandler:"), URL, completionHandler)
}/* debug [instance_methods/method]: StartFetchingMetadataForURLCompletionHandler */


// Fetches metadata for the given .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/startFetchingMetadata(for:completionHandler:)-9e6s8
func (l_ LPMetadataProvider) StartFetchingMetadataForRequestCompletionHandler(request foundation.URLRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("startFetchingMetadataForRequest:completionHandler:"), request, completionHandler)
}/* debug [instance_methods/method]: StartFetchingMetadataForRequestCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LPMetadataProvider */

// A Boolean value indicating whether to download subresources specified by the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/shouldFetchSubresources
func (l_ LPMetadataProvider) ShouldFetchSubresources() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldFetchSubresources"))
	return rv
}/* debug [instance_properties/getter]: shouldFetchSubresources */


// A Boolean value indicating whether to download subresources specified by the metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/shouldFetchSubresources
func (l_ LPMetadataProvider) SetShouldFetchSubresources(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShouldFetchSubresources:"), value)
}/* debug [instance_properties/setter]: shouldFetchSubresources */


// The time interval after which the request automatically fails if it hasn’t already completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/timeout
func (l_ LPMetadataProvider) Timeout() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("timeout"))
	return rv
}/* debug [instance_properties/getter]: timeout */


// The time interval after which the request automatically fails if it hasn’t already completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPMetadataProvider/timeout
func (l_ LPMetadataProvider) SetTimeout(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimeout:"), value)
}/* debug [instance_properties/setter]: timeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LPMetadataProvider */






