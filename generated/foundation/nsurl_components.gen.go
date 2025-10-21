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
	URLRelativeToURL(baseURL URL) URL
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
func NewURLComponentsWithURLResolvingAgainstBaseURL(url URL, resolve bool) URLComponents {
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
func (uc _URLComponentsClass) ComponentsWithURLResolvingAgainstBaseURL(url URL, resolve bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("componentsWithURL:resolvingAgainstBaseURL:"), url, resolve)
	return rv
}

// Returns a URL object derived from the components object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/url(relativeTo:)
func (u_ URLComponents) URLRelativeToURL(baseURL URL) URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URLRelativeToURL:"), baseURL)
	return rv
}

// The password URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpassword
func (u_ URLComponents) PercentEncodedPassword() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedPassword"))
	return rv
}


// SetPercentEncodedPassword sets the value of the percentEncodedPassword property.
// The password URL subcomponent expressed as a URL-encoded string, or

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpassword
func (u_ URLComponents) SetPercentEncodedPassword(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedPassword:"), objc.String(value))
}

// The username URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodeduser
func (u_ URLComponents) PercentEncodedUser() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedUser"))
	return rv
}


// SetPercentEncodedUser sets the value of the percentEncodedUser property.
// The username URL subcomponent expressed as a URL-encoded string, or

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodeduser
func (u_ URLComponents) SetPercentEncodedUser(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedUser:"), objc.String(value))
}

// The host URL subcomponent, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/host
func (u_ URLComponents) Host() string {
	rv := objc.Send[string](u_.ID, objc.Sel("host"))
	return rv
}


// SetHost sets the value of the host property.
// The host URL subcomponent, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/host
func (u_ URLComponents) SetHost(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHost:"), objc.String(value))
}

// Returns the character range of the port in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofport
func (u_ URLComponents) RangeOfPort() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPort"))
	return rv
}


// SetRangeOfPort sets the value of the rangeOfPort property.
// Returns the character range of the port in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofport
func (u_ URLComponents) SetRangeOfPort(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPort:"), value)
}

// Returns the character range of the host in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofhost
func (u_ URLComponents) RangeOfHost() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfHost"))
	return rv
}


// SetRangeOfHost sets the value of the rangeOfHost property.
// Returns the character range of the host in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofhost
func (u_ URLComponents) SetRangeOfHost(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfHost:"), value)
}

// The fragment URL component (the part after a
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedfragment
func (u_ URLComponents) PercentEncodedFragment() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedFragment"))
	return rv
}


// SetPercentEncodedFragment sets the value of the percentEncodedFragment property.
// The fragment URL component (the part after a

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedfragment
func (u_ URLComponents) SetPercentEncodedFragment(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedFragment:"), objc.String(value))
}

// The password URL subcomponent, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/password
func (u_ URLComponents) Password() string {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}


// SetPassword sets the value of the password property.
// The password URL subcomponent, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/password
func (u_ URLComponents) SetPassword(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPassword:"), objc.String(value))
}

// The fragment URL component (the part after a
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/fragment
func (u_ URLComponents) Fragment() string {
	rv := objc.Send[string](u_.ID, objc.Sel("fragment"))
	return rv
}


// SetFragment sets the value of the fragment property.
// The fragment URL component (the part after a

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/fragment
func (u_ URLComponents) SetFragment(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFragment:"), objc.String(value))
}

// Returns the character range of the path in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpath
func (u_ URLComponents) RangeOfPath() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPath"))
	return rv
}


// SetRangeOfPath sets the value of the rangeOfPath property.
// Returns the character range of the path in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpath
func (u_ URLComponents) SetRangeOfPath(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPath:"), value)
}

// The port number URL component, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/port
func (u_ URLComponents) Port() Number {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}


// SetPort sets the value of the port property.
// The port number URL component, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/port
func (u_ URLComponents) SetPort(value Number) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPort:"), value)
}

// The query URL component as a string, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/query
func (u_ URLComponents) Query() string {
	rv := objc.Send[string](u_.ID, objc.Sel("query"))
	return rv
}


// SetQuery sets the value of the query property.
// The query URL component as a string, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/query
func (u_ URLComponents) SetQuery(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQuery:"), objc.String(value))
}

// Returns the character range of the scheme in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofscheme
func (u_ URLComponents) RangeOfScheme() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfScheme"))
	return rv
}


// SetRangeOfScheme sets the value of the rangeOfScheme property.
// Returns the character range of the scheme in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofscheme
func (u_ URLComponents) SetRangeOfScheme(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfScheme:"), value)
}

// The scheme URL component, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/scheme
func (u_ URLComponents) Scheme() string {
	rv := objc.Send[string](u_.ID, objc.Sel("scheme"))
	return rv
}


// SetScheme sets the value of the scheme property.
// The scheme URL component, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/scheme
func (u_ URLComponents) SetScheme(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setScheme:"), objc.String(value))
}

// The query URL component as an array of name/value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLComponents) QueryItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queryItems"))
	return rv
}


// SetQueryItems sets the value of the queryItems property.
// The query URL component as an array of name/value pairs.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLComponents) SetQueryItems(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueryItems:"), value)
}

// Returns the character range of the user in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofuser
func (u_ URLComponents) RangeOfUser() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfUser"))
	return rv
}


// SetRangeOfUser sets the value of the rangeOfUser property.
// Returns the character range of the user in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofuser
func (u_ URLComponents) SetRangeOfUser(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfUser:"), value)
}

// Returns the character range of the query in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofquery
func (u_ URLComponents) RangeOfQuery() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfQuery"))
	return rv
}


// SetRangeOfQuery sets the value of the rangeOfQuery property.
// Returns the character range of the query in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofquery
func (u_ URLComponents) SetRangeOfQuery(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfQuery:"), value)
}

// The query URL component expressed as a URL-encoded string, or
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedquery
func (u_ URLComponents) PercentEncodedQuery() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedQuery"))
	return rv
}


// SetPercentEncodedQuery sets the value of the percentEncodedQuery property.
// The query URL component expressed as a URL-encoded string, or

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedquery
func (u_ URLComponents) SetPercentEncodedQuery(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedQuery:"), objc.String(value))
}

// The path URL component expressed as a URL-encoded string, or
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpath
func (u_ URLComponents) PercentEncodedPath() string {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedPath"))
	return rv
}


// SetPercentEncodedPath sets the value of the percentEncodedPath property.
// The path URL component expressed as a URL-encoded string, or

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpath
func (u_ URLComponents) SetPercentEncodedPath(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedPath:"), objc.String(value))
}

// The path URL component, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/path
func (u_ URLComponents) Path() string {
	rv := objc.Send[string](u_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The path URL component, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/path
func (u_ URLComponents) SetPath(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPath:"), objc.String(value))
}

// The username URL subcomponent, or nil if not present.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/user
func (u_ URLComponents) User() string {
	rv := objc.Send[string](u_.ID, objc.Sel("user"))
	return rv
}


// SetUser sets the value of the user property.
// The username URL subcomponent, or nil if not present.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/user
func (u_ URLComponents) SetUser(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUser:"), objc.String(value))
}

// Returns the character range of the password in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpassword
func (u_ URLComponents) RangeOfPassword() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPassword"))
	return rv
}


// SetRangeOfPassword sets the value of the rangeOfPassword property.
// Returns the character range of the password in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpassword
func (u_ URLComponents) SetRangeOfPassword(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPassword:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedqueryitems
func (u_ URLComponents) PercentEncodedQueryItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("percentEncodedQueryItems"))
	return rv
}


// SetPercentEncodedQueryItems sets the value of the percentEncodedQueryItems property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedqueryitems
func (u_ URLComponents) SetPercentEncodedQueryItems(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedQueryItems:"), value)
}

// Returns the character range of the fragment in the string returned by the string property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeoffragment
func (u_ URLComponents) RangeOfFragment() Range {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfFragment"))
	return rv
}


// SetRangeOfFragment sets the value of the rangeOfFragment property.
// Returns the character range of the fragment in the string returned by the string property.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeoffragment
func (u_ URLComponents) SetRangeOfFragment(value Range) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfFragment:"), value)
}

// The host subcomponent, percent-encoded.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/encodedhost
func (u_ URLComponents) EncodedHost() string {
	rv := objc.Send[string](u_.ID, objc.Sel("encodedHost"))
	return rv
}


// SetEncodedHost sets the value of the encodedHost property.
// The host subcomponent, percent-encoded.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/encodedhost
func (u_ URLComponents) SetEncodedHost(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEncodedHost:"), objc.String(value))
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
func (u_ URLComponents) String_() string {
	rv := objc.Send[string](u_.ID, objc.Sel("string"))
	return rv
}

// A URL object derived from the components object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/url
func (u_ URLComponents) URL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URL"))
	return rv
}


