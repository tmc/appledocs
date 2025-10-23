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
	// properties:
	URL() IURL
	EncodedHost() string /* primitive/slice/pointer */
	SetEncodedHost(value string /* primitive/slice/pointer */)
	Fragment() string /* primitive/slice/pointer */
	SetFragment(value string /* primitive/slice/pointer */)
	Host() string /* primitive/slice/pointer */
	SetHost(value string /* primitive/slice/pointer */)
	Password() string /* primitive/slice/pointer */
	SetPassword(value string /* primitive/slice/pointer */)
	Path() string /* primitive/slice/pointer */
	SetPath(value string /* primitive/slice/pointer */)
	PercentEncodedFragment() string /* primitive/slice/pointer */
	SetPercentEncodedFragment(value string /* primitive/slice/pointer */)
	PercentEncodedHost() string /* primitive/slice/pointer */
	SetPercentEncodedHost(value string /* primitive/slice/pointer */)
	PercentEncodedPassword() string /* primitive/slice/pointer */
	SetPercentEncodedPassword(value string /* primitive/slice/pointer */)
	PercentEncodedPath() string /* primitive/slice/pointer */
	SetPercentEncodedPath(value string /* primitive/slice/pointer */)
	PercentEncodedQuery() string /* primitive/slice/pointer */
	SetPercentEncodedQuery(value string /* primitive/slice/pointer */)
	PercentEncodedQueryItems() IURLQueryItem
	SetPercentEncodedQueryItems(value IURLQueryItem)
	PercentEncodedUser() string /* primitive/slice/pointer */
	SetPercentEncodedUser(value string /* primitive/slice/pointer */)
	Port() INumber
	SetPort(value INumber)
	Query() string /* primitive/slice/pointer */
	SetQuery(value string /* primitive/slice/pointer */)
	QueryItems() IURLQueryItem
	SetQueryItems(value IURLQueryItem)
	RangeOfFragment() Range /* foo */
	SetRangeOfFragment(value Range /* foo */)
	RangeOfHost() Range /* foo */
	SetRangeOfHost(value Range /* foo */)
	RangeOfPassword() Range /* foo */
	SetRangeOfPassword(value Range /* foo */)
	RangeOfPath() Range /* foo */
	SetRangeOfPath(value Range /* foo */)
	RangeOfPort() Range /* foo */
	SetRangeOfPort(value Range /* foo */)
	RangeOfQuery() Range /* foo */
	SetRangeOfQuery(value Range /* foo */)
	RangeOfScheme() Range /* foo */
	SetRangeOfScheme(value Range /* foo */)
	RangeOfUser() Range /* foo */
	SetRangeOfUser(value Range /* foo */)
	Scheme() string /* primitive/slice/pointer */
	SetScheme(value string /* primitive/slice/pointer */)
	String() string /* primitive/slice/pointer */
	SetString(value string /* primitive/slice/pointer */)
	User() string /* primitive/slice/pointer */
	SetUser(value string /* primitive/slice/pointer */)
	// methods:
}

// An object that parses URLs into and constructs URLs from their constituent parts.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class is a class that is designed to parse URLs based on and to construct URLs from their constituent parts. Its behavior differs subtly from the class, which conforms to older RFCs. However, you can easily obtain an object based on the contents of a URL components object or vice versa. You create a URL components object in one of three ways: from an object that contains a URL, from an object, or from scratch by using the default initializer. From there, you can modify the URL’s individual components and subcomponents by modifying various properties, either in unencoded form or in URL-encoded form. If you set the unencoded property, you can then obtain the encoded equivalent by reading the encoded property value and vice versa.


// An object that parses URLs into and constructs URLs from their constituent parts.
//
// [Full Topic]
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




// A URL object derived from the components object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLComponents/url
func (u_ URLComponents) URL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("URL"))
	return rv
}


// The host subcomponent, percent-encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/encodedhost
func (u_ URLComponents) EncodedHost() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("encodedHost"))
	return rv
}


// The host subcomponent, percent-encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/encodedhost
func (u_ URLComponents) SetEncodedHost(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEncodedHost:"), objc.String(value))
}


// The fragment URL component (the part after a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/fragment
func (u_ URLComponents) Fragment() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("fragment"))
	return rv
}


// The fragment URL component (the part after a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/fragment
func (u_ URLComponents) SetFragment(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFragment:"), objc.String(value))
}


// The host URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/host
func (u_ URLComponents) Host() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("host"))
	return rv
}


// The host URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/host
func (u_ URLComponents) SetHost(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHost:"), objc.String(value))
}


// The password URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/password
func (u_ URLComponents) Password() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}


// The password URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/password
func (u_ URLComponents) SetPassword(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPassword:"), objc.String(value))
}


// The path URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/path
func (u_ URLComponents) Path() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("path"))
	return rv
}


// The path URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/path
func (u_ URLComponents) SetPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPath:"), objc.String(value))
}


// The fragment URL component (the part after a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedfragment
func (u_ URLComponents) PercentEncodedFragment() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedFragment"))
	return rv
}


// The fragment URL component (the part after a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedfragment
func (u_ URLComponents) SetPercentEncodedFragment(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedFragment:"), objc.String(value))
}


// The host URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedhost
func (u_ URLComponents) PercentEncodedHost() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedHost"))
	return rv
}


// The host URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedhost
func (u_ URLComponents) SetPercentEncodedHost(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedHost:"), objc.String(value))
}


// The password URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpassword
func (u_ URLComponents) PercentEncodedPassword() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedPassword"))
	return rv
}


// The password URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpassword
func (u_ URLComponents) SetPercentEncodedPassword(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedPassword:"), objc.String(value))
}


// The path URL component expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpath
func (u_ URLComponents) PercentEncodedPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedPath"))
	return rv
}


// The path URL component expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedpath
func (u_ URLComponents) SetPercentEncodedPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedPath:"), objc.String(value))
}


// The query URL component expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedquery
func (u_ URLComponents) PercentEncodedQuery() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedQuery"))
	return rv
}


// The query URL component expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedquery
func (u_ URLComponents) SetPercentEncodedQuery(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedQuery:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedqueryitems
func (u_ URLComponents) PercentEncodedQueryItems() IURLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("percentEncodedQueryItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodedqueryitems
func (u_ URLComponents) SetPercentEncodedQueryItems(value IURLQueryItem) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedQueryItems:"), value)
}


// The username URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodeduser
func (u_ URLComponents) PercentEncodedUser() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("percentEncodedUser"))
	return rv
}


// The username URL subcomponent expressed as a URL-encoded string, or
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/percentencodeduser
func (u_ URLComponents) SetPercentEncodedUser(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPercentEncodedUser:"), objc.String(value))
}


// The port number URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/port
func (u_ URLComponents) Port() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("port"))
	return rv
}


// The port number URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/port
func (u_ URLComponents) SetPort(value INumber) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPort:"), value)
}


// The query URL component as a string, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/query
func (u_ URLComponents) Query() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("query"))
	return rv
}


// The query URL component as a string, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/query
func (u_ URLComponents) SetQuery(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQuery:"), objc.String(value))
}


// The query URL component as an array of name/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLComponents) QueryItems() IURLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("queryItems"))
	return rv
}


// The query URL component as an array of name/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLComponents) SetQueryItems(value IURLQueryItem) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueryItems:"), value)
}


// Returns the character range of the fragment in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeoffragment
func (u_ URLComponents) RangeOfFragment() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfFragment"))
	return rv
}


// Returns the character range of the fragment in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeoffragment
func (u_ URLComponents) SetRangeOfFragment(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfFragment:"), value)
}


// Returns the character range of the host in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofhost
func (u_ URLComponents) RangeOfHost() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfHost"))
	return rv
}


// Returns the character range of the host in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofhost
func (u_ URLComponents) SetRangeOfHost(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfHost:"), value)
}


// Returns the character range of the password in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpassword
func (u_ URLComponents) RangeOfPassword() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPassword"))
	return rv
}


// Returns the character range of the password in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpassword
func (u_ URLComponents) SetRangeOfPassword(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPassword:"), value)
}


// Returns the character range of the path in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpath
func (u_ URLComponents) RangeOfPath() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPath"))
	return rv
}


// Returns the character range of the path in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofpath
func (u_ URLComponents) SetRangeOfPath(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPath:"), value)
}


// Returns the character range of the port in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofport
func (u_ URLComponents) RangeOfPort() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfPort"))
	return rv
}


// Returns the character range of the port in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofport
func (u_ URLComponents) SetRangeOfPort(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfPort:"), value)
}


// Returns the character range of the query in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofquery
func (u_ URLComponents) RangeOfQuery() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfQuery"))
	return rv
}


// Returns the character range of the query in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofquery
func (u_ URLComponents) SetRangeOfQuery(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfQuery:"), value)
}


// Returns the character range of the scheme in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofscheme
func (u_ URLComponents) RangeOfScheme() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfScheme"))
	return rv
}


// Returns the character range of the scheme in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofscheme
func (u_ URLComponents) SetRangeOfScheme(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfScheme:"), value)
}


// Returns the character range of the user in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofuser
func (u_ URLComponents) RangeOfUser() Range /* foo */ {
	rv := objc.Send[Range](u_.ID, objc.Sel("rangeOfUser"))
	return rv
}


// Returns the character range of the user in the string returned by the string property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/rangeofuser
func (u_ URLComponents) SetRangeOfUser(value Range /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRangeOfUser:"), value)
}


// The scheme URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/scheme
func (u_ URLComponents) Scheme() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("scheme"))
	return rv
}


// The scheme URL component, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/scheme
func (u_ URLComponents) SetScheme(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setScheme:"), objc.String(value))
}


// A URL derived from the components object, in string form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/string
func (u_ URLComponents) String() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("string"))
	return rv
}


// A URL derived from the components object, in string form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/string
func (u_ URLComponents) SetString(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setString:"), objc.String(value))
}


// The username URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/user
func (u_ URLComponents) User() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("user"))
	return rv
}


// The username URL subcomponent, or nil if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/user
func (u_ URLComponents) SetUser(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUser:"), objc.String(value))
}


