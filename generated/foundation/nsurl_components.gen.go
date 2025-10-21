// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLComponents] class.
var (
	URLComponentsClass     _URLComponentsClass
	URLComponentsClassOnce sync.Once
)

func getURLComponentsClass() _URLComponentsClass {
	URLComponentsClassOnce.Do(func() {
		URLComponentsClass = _URLComponentsClass{objc.GetClass("NSURLComponents")}
	})
	return URLComponentsClass
}

type _URLComponentsClass struct {
	class objc.Class
}

// An interface definition for the [URLComponents] class.
type IURLComponents interface {
	objectivec.IObject
	URLRelativeToURL(baseURL unsafe.Pointer) unsafe.Pointer
}

// An object that parses URLs into and constructs URLs from their constituent parts.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class is a class that is designed to parse URLs based on and to construct URLs from their constituent parts. Its behavior differs subtly from the class, which conforms to older RFCs. However, you can easily obtain an object based on the contents of a URL components object or vice versa. You create a URL components object in one of three ways: from an object that contains a URL, from an object, or from scratch by using the default initializer. From there, you can modify the URL’s individual components and subcomponents by modifying various properties, either in unencoded form or in URL-encoded form. If you set the unencoded property, you can then obtain the encoded equivalent by reading the encoded property value and vice versa.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents
type URLComponents struct {
	objectivec.Object
}

// URLComponentsFrom constructs a [URLComponents] from an unsafe.Pointer.
//
// An object that parses URLs into and constructs URLs from their constituent parts.
func URLComponentsFrom(ptr unsafe.Pointer) URLComponents {
	return URLComponents{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLComponentsClass) Alloc() URLComponents {
	rv := objc.Send[URLComponents](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLComponentsClass) New() URLComponents {
	rv := objc.Send[URLComponents](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLComponents) Init() URLComponents {
	rv := objc.Send[URLComponents](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLComponents) Autorelease() URLComponents {
	rv := objc.Send[URLComponents](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLComponents creates a new URLComponents instance.
func NewURLComponents() URLComponents {
	return getURLComponentsClass().New()
}




// Creates a URL components object by parsing a URL in string form.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:)
func NewURLComponentsWithString(URLString string) URLComponents {
	instance := getURLComponentsClass().Alloc()
	rv := objc.Send[URLComponents](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	rv.Autorelease()
	return rv
}



// Creates a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:encodingInvalidCharacters:)
func NewURLComponentsWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) URLComponents {
	instance := getURLComponentsClass().Alloc()
	rv := objc.Send[URLComponents](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	rv.Autorelease()
	return rv
}



// Creates a URL components object by parsing the URL from an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(url:resolvingAgainstBaseURL:)
func NewURLComponentsWithURLResolvingAgainstBaseURL(url unsafe.Pointer, resolve bool) URLComponents {
	instance := getURLComponentsClass().Alloc()
	rv := objc.Send[URLComponents](instance.ID, objc.Sel("initWithURL:resolvingAgainstBaseURL:"), url, resolve)
	rv.Autorelease()
	return rv
}


// Returns a URL components object by parsing a URL in string form.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithString:
func (uc _URLComponentsClass) ComponentsWithString(URLString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("componentsWithString:"), objc.String(URLString))
	return rv
}

// Returns a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithString:encodingInvalidCharacters:
func (uc _URLComponentsClass) ComponentsWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("componentsWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}

// Returns a URL components object by parsing the URL from an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithURL:resolvingAgainstBaseURL:
func (uc _URLComponentsClass) ComponentsWithURLResolvingAgainstBaseURL(url unsafe.Pointer, resolve bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("componentsWithURL:resolvingAgainstBaseURL:"), url, resolve)
	return rv
}

// Returns a URL object derived from the components object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/url(relativeTo:)
func (u_ URLComponents) URLRelativeToURL(baseURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URLRelativeToURL:"), baseURL)
	return rv
}

// The host URL subcomponent expressed as a URL-encoded string, or if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedHost
func (u_ URLComponents) PercentEncodedHost() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedHost"))
	return rv
}


// SetPercentEncodedHost sets the value of the percentEncodedHost property.
// The host URL subcomponent expressed as a URL-encoded string, or if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedHost
func (u_ URLComponents) SetPercentEncodedHost(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedHost:"), objc.String(value))
}

// A URL derived from the components object, in string form.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/string
func (u_ URLComponents) String() string {
	rv := objc.Send[string](u_.ID, objc.Sel("string"))
	return rv
}

// A URL object derived from the components object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/url
func (u_ URLComponents) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("URL"))
	return rv
}


