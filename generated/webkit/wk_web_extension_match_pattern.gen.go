// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionMatchPattern */


/* debug [class_header]: Header for WKWebExtensionMatchPattern */
// The class instance for the [WebExtensionMatchPattern] class.
var (
	WebExtensionMatchPatternClass     _WebExtensionMatchPatternClass
	WebExtensionMatchPatternClassOnce sync.Once
)

func getWebExtensionMatchPatternClass() _WebExtensionMatchPatternClass {
	WebExtensionMatchPatternClassOnce.Do(func() {
		WebExtensionMatchPatternClass = _WebExtensionMatchPatternClass{objc.GetClass("WKWebExtensionMatchPattern")}
	})
	return WebExtensionMatchPatternClass
}

type _WebExtensionMatchPatternClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionMatchPattern */
// An interface definition for the [WebExtensionMatchPattern] class.
type IWebExtensionMatchPattern interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionMatchPattern */
	// properties:
	Host() objc.IObject /* cross-framework: NSString */
	MatchesAllHosts() bool
	MatchesAllURLs() bool
	Path() objc.IObject /* cross-framework: NSString */
	Scheme() objc.IObject /* cross-framework: NSString */
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionMatchPattern */
	// methods:
	MatchesURL(url objc.IObject /* cross-framework: NSURL */) bool
	MatchesPattern(pattern IWKWebExtensionMatchPattern) bool
	MatchesURLOptions(url objc.IObject /* cross-framework: NSURL */, options WebExtensionMatchPatternOptions) bool
	MatchesPatternOptions(pattern IWKWebExtensionMatchPattern, options WebExtensionMatchPatternOptions) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionMatchPattern */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionMatchPatternClass) Alloc() WebExtensionMatchPattern {
	rv := objc.Send[WebExtensionMatchPattern](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionMatchPatternClass) New() WebExtensionMatchPattern {
	rv := objc.Send[WebExtensionMatchPattern](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionMatchPattern) Init() WebExtensionMatchPattern {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionMatchPattern) Autorelease() WebExtensionMatchPattern {
	rv := objc.Send[WebExtensionMatchPattern](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionMatchPattern creates a new WebExtensionMatchPattern instance.
func NewWebExtensionMatchPattern() WebExtensionMatchPattern {
	return getWebExtensionMatchPatternClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionMatchPattern */
// An object that represents a way to specify groups of URLs.
//
// All match patterns are specified as strings. Apart from the special pattern, match patterns consist of three parts: scheme, host, and path.


// An object that represents a way to specify groups of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern
type WebExtensionMatchPattern struct {
	objectivec.Object
}

// WebExtensionMatchPatternFrom constructs a [WebExtensionMatchPattern] from an unsafe.Pointer.
//
// An object that represents a way to specify groups of URLs.
func WebExtensionMatchPatternFrom(ptr unsafe.Pointer) WebExtensionMatchPattern {
	return WebExtensionMatchPattern{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionMatchPattern */

// Returns a pattern object for the specified scheme, host, and path strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(scheme:host:path:)
func NewWebExtensionMatchPatternWithSchemeHostPathError(scheme objc.IObject /* cross-framework: NSString */, host objc.IObject /* cross-framework: NSString */, path objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) WebExtensionMatchPattern {
	instance := getWebExtensionMatchPatternClass().Alloc()
	rv := objc.Send[WebExtensionMatchPattern](instance.ID, objc.Sel("initWithScheme:host:path:error:"), scheme, host, path, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebExtensionMatchPatternWithSchemeHostPathError */


// Returns a pattern object for the specified pattern string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(string:)
func NewWebExtensionMatchPatternWithStringError(string_ objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) WebExtensionMatchPattern {
	instance := getWebExtensionMatchPatternClass().Alloc()
	rv := objc.Send[WebExtensionMatchPattern](instance.ID, objc.Sel("initWithString:error:"), string_, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebExtensionMatchPatternWithStringError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionMatchPattern */

// Returns a pattern object that has for scheme, host, and path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/allHostsAndSchemes()
func (wc _WebExtensionMatchPatternClass) AllHostsAndSchemesMatchPattern() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("allHostsAndSchemesMatchPattern"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllHostsAndSchemesMatchPattern) */


// Returns a pattern object for .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/allURLs()
func (wc _WebExtensionMatchPatternClass) AllURLsMatchPattern() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("allURLsMatchPattern"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllURLsMatchPattern) */


// Registers a custom URL scheme that can be used in match patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/registerCustomURLScheme(_:)
func (wc _WebExtensionMatchPatternClass) RegisterCustomURLScheme(urlScheme objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("registerCustomURLScheme:"), urlScheme)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterCustomURLScheme) */


// Returns a pattern object for the specified scheme, host, and path strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionMatchPattern/matchPatternWithScheme:host:path:
func (wc _WebExtensionMatchPatternClass) MatchPatternWithSchemeHostPath(scheme objc.IObject /* cross-framework: NSString */, host objc.IObject /* cross-framework: NSString */, path objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("matchPatternWithScheme:host:path:"), scheme, host, path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MatchPatternWithSchemeHostPath) */


// Returns a pattern object for the specified pattern string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionMatchPattern/matchPatternWithString:
func (wc _WebExtensionMatchPatternClass) MatchPatternWithString(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("matchPatternWithString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MatchPatternWithString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionMatchPattern */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionMatchPattern */

// Matches the receiver pattern against the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:)-471rf
func (w_ WebExtensionMatchPattern) MatchesURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesURL:"), url)
	return rv
}/* debug [instance_methods/method]: MatchesURL */


// Matches the receiver pattern against the specified pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:)-4d84f
func (w_ WebExtensionMatchPattern) MatchesPattern(pattern IWKWebExtensionMatchPattern) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesPattern:"), pattern)
	return rv
}/* debug [instance_methods/method]: MatchesPattern */


// Matches the receiver pattern against the specified URL with options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:options:)-5wo3g
func (w_ WebExtensionMatchPattern) MatchesURLOptions(url objc.IObject /* cross-framework: NSURL */, options WebExtensionMatchPatternOptions) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesURL:options:"), url, options)
	return rv
}/* debug [instance_methods/method]: MatchesURLOptions */


// Matches the receiver pattern against the specified pattern with options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:options:)-fnde
func (w_ WebExtensionMatchPattern) MatchesPatternOptions(pattern IWKWebExtensionMatchPattern, options WebExtensionMatchPatternOptions) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesPattern:options:"), pattern, options)
	return rv
}/* debug [instance_methods/method]: MatchesPatternOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionMatchPattern */

// The host part of the pattern string, unless is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/host
func (w_ WebExtensionMatchPattern) Host() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("host"))
	return rv
}/* debug [instance_properties/getter]: host */


// A Boolean value that indicates if the pattern is or has as the host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllHosts
func (w_ WebExtensionMatchPattern) MatchesAllHosts() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesAllHosts"))
	return rv
}/* debug [instance_properties/getter]: matchesAllHosts */


// A Boolean value that indicates if the pattern is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs
func (w_ WebExtensionMatchPattern) MatchesAllURLs() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesAllURLs"))
	return rv
}/* debug [instance_properties/getter]: matchesAllURLs */


// The path part of the pattern string, unless is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/path
func (w_ WebExtensionMatchPattern) Path() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The scheme part of the pattern string, unless is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/scheme
func (w_ WebExtensionMatchPattern) Scheme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("scheme"))
	return rv
}/* debug [instance_properties/getter]: scheme */


// The original pattern string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/string
func (w_ WebExtensionMatchPattern) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionMatchPattern */


