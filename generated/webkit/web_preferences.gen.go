// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebPreferences */

/* debug [class_header]: Header for WebPreferences */
// The class instance for the [WebPreferences] class.
var (
	WebPreferencesClass     _WebPreferencesClass
	WebPreferencesClassOnce sync.Once
)

func getWebPreferencesClass() _WebPreferencesClass {
	WebPreferencesClassOnce.Do(func() {
		WebPreferencesClass = _WebPreferencesClass{objc.GetClass("WebPreferences")}
	})
	return WebPreferencesClass
}

type _WebPreferencesClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebPreferences */
// An interface definition for the [WebPreferences] class.
type IWebPreferences interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebPreferences */
	// properties:
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	AllowsAnimatedImageLooping() bool
	SetAllowsAnimatedImageLooping(value bool)
	AllowsAnimatedImages() bool
	SetAllowsAnimatedImages(value bool)
	PlugInsEnabled() bool
	SetPlugInsEnabled(value bool)
	Autosaves() bool
	SetAutosaves(value bool)
	CacheModel() WebCacheModel
	SetCacheModel(value WebCacheModel)
	CursiveFontFamily() objc.IObject /* cross-framework: NSString */
	SetCursiveFontFamily(value objc.IObject /* cross-framework: NSString */)
	DefaultFixedFontSize() int
	SetDefaultFixedFontSize(value int)
	DefaultFontSize() int
	SetDefaultFontSize(value int)
	DefaultTextEncodingName() objc.IObject /* cross-framework: NSString */
	SetDefaultTextEncodingName(value objc.IObject /* cross-framework: NSString */)
	FantasyFontFamily() objc.IObject /* cross-framework: NSString */
	SetFantasyFontFamily(value objc.IObject /* cross-framework: NSString */)
	FixedFontFamily() objc.IObject /* cross-framework: NSString */
	SetFixedFontFamily(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	JavaEnabled() bool
	SetJavaEnabled(value bool)
	JavaScriptEnabled() bool
	SetJavaScriptEnabled(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	LoadsImagesAutomatically() bool
	SetLoadsImagesAutomatically(value bool)
	MinimumFontSize() int
	SetMinimumFontSize(value int)
	MinimumLogicalFontSize() int
	SetMinimumLogicalFontSize(value int)
	PrivateBrowsingEnabled() bool
	SetPrivateBrowsingEnabled(value bool)
	SansSerifFontFamily() objc.IObject /* cross-framework: NSString */
	SetSansSerifFontFamily(value objc.IObject /* cross-framework: NSString */)
	SerifFontFamily() objc.IObject /* cross-framework: NSString */
	SetSerifFontFamily(value objc.IObject /* cross-framework: NSString */)
	ShouldPrintBackgrounds() bool
	SetShouldPrintBackgrounds(value bool)
	StandardFontFamily() objc.IObject /* cross-framework: NSString */
	SetStandardFontFamily(value objc.IObject /* cross-framework: NSString */)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	TabsToLinks() bool
	SetTabsToLinks(value bool)
	UserStyleSheetEnabled() bool
	SetUserStyleSheetEnabled(value bool)
	UserStyleSheetLocation() objc.IObject /* cross-framework: NSURL */
	SetUserStyleSheetLocation(value objc.IObject /* cross-framework: NSURL */)
	UsesPageCache() bool
	SetUsesPageCache(value bool)
	ArePlugInsEnabled() bool
	SetArePlugInsEnabled(value bool)
	IsJavaEnabled() bool
	SetIsJavaEnabled(value bool)
	IsJavaScriptEnabled() bool
	SetIsJavaScriptEnabled(value bool)
	PreferencesIdentifier() objc.IObject /* cross-framework: NSString */
	SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebPreferences */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebPreferences */
// Alloc allocates a new instance without initialization.
func (wc _WebPreferencesClass) Alloc() WebPreferences {
	rv := objc.Send[WebPreferences](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebPreferencesClass) New() WebPreferences {
	rv := objc.Send[WebPreferences](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebPreferences) Init() WebPreferences {
	rv := objc.Send[WebPreferences](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebPreferences) Autorelease() WebPreferences {
	rv := objc.Send[WebPreferences](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebPreferences creates a new WebPreferences instance.
func NewWebPreferences() WebPreferences {
	return getWebPreferencesClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebPreferences */
// WebPreferences encapsulates the preferences you can change per WebView object. These preferences include font, text encoding, and image settings. Normally a WebView object uses the standard preferences returned by the class method. However, you can modify the preferences for individual WebView instances too. Use the WebView method to change a WebView object’s preferences, or to share preferences between WebView objects. Use the method to specify if the preferences object should be automatically saved to the user defaults database.
//
// WebPreferences also manages the font preferences for a web view. You can set custom font families for each of the primary web font styles (standard, serif, sans-serif, cursive, and fantasy) as well as their font sizes. The font size preferences alter the display font sizes in a certain way. If the HTML or CSS in the web view’s content specifies font sizes in a relative fashion (such as in HTML or in CSS), the default font size settings (set by the font size methods prefaced with “default”) have an effect. They do not have an effect for font sizes specified absolutely. The values specified by the minimum font size settings (set by the font size methods prefaced with “minimum”) override all the HTML and CSS font size definitions, and so have an effect on the entirety of the content. The values specified by the minimum logical font size settings (set by the font size methods prefaced with “minimumLogical”) affect all relative font size declarations for HTML and CSS, but also override any CSS font size declarations in the content, whether they are relative or absolute. The font size for a web view is different than its logical font size. The minimum logical font size, for example, is the absolute minimum size at which the font will display onscreen. This is meant to be a functional boundary and not a style boundary. For example, the default value for a web view’s minimum logical font size is 9 points, because typical web content looks good in macOS at font sizes of 9 point and above. The constraint assures that web content will always look good in a web view. If you know that your content will look good only at 12 points or above, you should change the minimum font size to 12 points and leave the minimum font size alone. This will assure that your content will never display at sizes less than 12 points, but the functional font size boundary of the web view will remain at 9 points to prevent any chance of displaying unnecessarily small text.

// WebPreferences encapsulates the preferences you can change per WebView object. These preferences include font, text encoding, and image settings. Normally a WebView object uses the standard preferences returned by the class method. However, you can modify the preferences for individual WebView instances too. Use the WebView method to change a WebView object’s preferences, or to share preferences between WebView objects. Use the method to specify if the preferences object should be automatically saved to the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences
type WebPreferences struct {
	objectivec.Object
}

// WebPreferencesFrom constructs a [WebPreferences] from an unsafe.Pointer.
//
// WebPreferences encapsulates the preferences you can change per WebView object. These preferences include font, text encoding, and image settings. Normally a WebView object uses the standard preferences returned by the class method. However, you can modify the preferences for individual WebView instances too. Use the WebView method to change a WebView object’s preferences, or to share preferences between WebView objects. Use the method to specify if the preferences object should be automatically saved to the user defaults database.
func WebPreferencesFrom(ptr unsafe.Pointer) WebPreferences {
	return WebPreferences{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebPreferences */

// Returns an initialized object, creating one if it does not exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/init(identifier:)
func NewWebPreferencesWithIdentifier(anIdentifier objc.IObject /* cross-framework: NSString */) WebPreferences {
	instance := getWebPreferencesClass().Alloc()
	rv := objc.Send[WebPreferences](instance.ID, objc.Sel("initWithIdentifier:"), anIdentifier)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewWebPreferencesWithIdentifier */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebPreferences */

// Returns the standard set of preferences that may be used by all WebView objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/standard()
func (wc _WebPreferencesClass) StandardPreferences() WebPreferences {
	rv := objc.Send[WebPreferences](objc.ID(wc.class), objc.Sel("standardPreferences"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=StandardPreferences) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebPreferences */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebPreferences */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebPreferences */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAirPlayForMediaPlayback
func (w_ WebPreferences) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAirPlayForMediaPlayback"))
	return rv
} /* debug [instance_properties/getter]: allowsAirPlayForMediaPlayback */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAirPlayForMediaPlayback
func (w_ WebPreferences) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAirPlayForMediaPlayback:"), value)
} /* debug [instance_properties/setter]: allowsAirPlayForMediaPlayback */

// A Boolean that indicates whether or not the receiver allows animated images to loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAnimatedImageLooping
func (w_ WebPreferences) AllowsAnimatedImageLooping() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAnimatedImageLooping"))
	return rv
} /* debug [instance_properties/getter]: allowsAnimatedImageLooping */

// A Boolean that indicates whether or not the receiver allows animated images to loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAnimatedImageLooping
func (w_ WebPreferences) SetAllowsAnimatedImageLooping(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAnimatedImageLooping:"), value)
} /* debug [instance_properties/setter]: allowsAnimatedImageLooping */

// A Boolean that indicates whether or not the receiver allows animated images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAnimatedImages
func (w_ WebPreferences) AllowsAnimatedImages() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAnimatedImages"))
	return rv
} /* debug [instance_properties/getter]: allowsAnimatedImages */

// A Boolean that indicates whether or not the receiver allows animated images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/allowsAnimatedImages
func (w_ WebPreferences) SetAllowsAnimatedImages(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAnimatedImages:"), value)
} /* debug [instance_properties/setter]: allowsAnimatedImages */

// A Boolean that indicates whether or not the web view allows plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/arePlugInsEnabled
func (w_ WebPreferences) PlugInsEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("plugInsEnabled"))
	return rv
} /* debug [instance_properties/getter]: plugInsEnabled */

// A Boolean that indicates whether or not the web view allows plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/arePlugInsEnabled
func (w_ WebPreferences) SetPlugInsEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPlugInsEnabled:"), value)
} /* debug [instance_properties/setter]: plugInsEnabled */

// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/autosaves
func (w_ WebPreferences) Autosaves() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autosaves"))
	return rv
} /* debug [instance_properties/getter]: autosaves */

// A Boolean that indicates whether or not the receiver’s attributes are automatically stored in the user defaults database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/autosaves
func (w_ WebPreferences) SetAutosaves(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutosaves:"), value)
} /* debug [instance_properties/setter]: autosaves */

// The cache model for the web views associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/cacheModel
func (w_ WebPreferences) CacheModel() WebCacheModel {
	rv := objc.Send[WebCacheModel](w_.ID, objc.Sel("cacheModel"))
	return rv
} /* debug [instance_properties/getter]: cacheModel */

// The cache model for the web views associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/cacheModel
func (w_ WebPreferences) SetCacheModel(value WebCacheModel) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCacheModel:"), value)
} /* debug [instance_properties/setter]: cacheModel */

// The cursive font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/cursiveFontFamily
func (w_ WebPreferences) CursiveFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("cursiveFontFamily"))
	return rv
} /* debug [instance_properties/getter]: cursiveFontFamily */

// The cursive font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/cursiveFontFamily
func (w_ WebPreferences) SetCursiveFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCursiveFontFamily:"), value)
} /* debug [instance_properties/setter]: cursiveFontFamily */

// The default fixed font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultFixedFontSize
func (w_ WebPreferences) DefaultFixedFontSize() int {
	rv := objc.Send[int](w_.ID, objc.Sel("defaultFixedFontSize"))
	return rv
} /* debug [instance_properties/getter]: defaultFixedFontSize */

// The default fixed font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultFixedFontSize
func (w_ WebPreferences) SetDefaultFixedFontSize(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultFixedFontSize:"), value)
} /* debug [instance_properties/setter]: defaultFixedFontSize */

// The default font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultFontSize
func (w_ WebPreferences) DefaultFontSize() int {
	rv := objc.Send[int](w_.ID, objc.Sel("defaultFontSize"))
	return rv
} /* debug [instance_properties/getter]: defaultFontSize */

// The default font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultFontSize
func (w_ WebPreferences) SetDefaultFontSize(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultFontSize:"), value)
} /* debug [instance_properties/setter]: defaultFontSize */

// The default text encoding of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultTextEncodingName
func (w_ WebPreferences) DefaultTextEncodingName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("defaultTextEncodingName"))
	return rv
} /* debug [instance_properties/getter]: defaultTextEncodingName */

// The default text encoding of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/defaultTextEncodingName
func (w_ WebPreferences) SetDefaultTextEncodingName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultTextEncodingName:"), value)
} /* debug [instance_properties/setter]: defaultTextEncodingName */

// The fantasy font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/fantasyFontFamily
func (w_ WebPreferences) FantasyFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("fantasyFontFamily"))
	return rv
} /* debug [instance_properties/getter]: fantasyFontFamily */

// The fantasy font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/fantasyFontFamily
func (w_ WebPreferences) SetFantasyFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFantasyFontFamily:"), value)
} /* debug [instance_properties/setter]: fantasyFontFamily */

// The fixed font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/fixedFontFamily
func (w_ WebPreferences) FixedFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("fixedFontFamily"))
	return rv
} /* debug [instance_properties/getter]: fixedFontFamily */

// The fixed font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/fixedFontFamily
func (w_ WebPreferences) SetFixedFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFixedFontFamily:"), value)
} /* debug [instance_properties/setter]: fixedFontFamily */

// The receiver’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/identifier
func (w_ WebPreferences) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// A Boolean that indicates whether or not the web view allows Java.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/isJavaEnabled
func (w_ WebPreferences) JavaEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("javaEnabled"))
	return rv
} /* debug [instance_properties/getter]: javaEnabled */

// A Boolean that indicates whether or not the web view allows Java.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/isJavaEnabled
func (w_ WebPreferences) SetJavaEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setJavaEnabled:"), value)
} /* debug [instance_properties/setter]: javaEnabled */

// A Boolean that indicates whether or not the web view allows JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/isJavaScriptEnabled
func (w_ WebPreferences) JavaScriptEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("javaScriptEnabled"))
	return rv
} /* debug [instance_properties/getter]: javaScriptEnabled */

// A Boolean that indicates whether or not the web view allows JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/isJavaScriptEnabled
func (w_ WebPreferences) SetJavaScriptEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setJavaScriptEnabled:"), value)
} /* debug [instance_properties/setter]: javaScriptEnabled */

// A Boolean that indicates whether or not the web view allows JavaScript to open windows automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/javaScriptCanOpenWindowsAutomatically
func (w_ WebPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("javaScriptCanOpenWindowsAutomatically"))
	return rv
} /* debug [instance_properties/getter]: javaScriptCanOpenWindowsAutomatically */

// A Boolean that indicates whether or not the web view allows JavaScript to open windows automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/javaScriptCanOpenWindowsAutomatically
func (w_ WebPreferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setJavaScriptCanOpenWindowsAutomatically:"), value)
} /* debug [instance_properties/setter]: javaScriptCanOpenWindowsAutomatically */

// A Boolean that indicates whether or not the web view allows images to be loaded automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/loadsImagesAutomatically
func (w_ WebPreferences) LoadsImagesAutomatically() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loadsImagesAutomatically"))
	return rv
} /* debug [instance_properties/getter]: loadsImagesAutomatically */

// A Boolean that indicates whether or not the web view allows images to be loaded automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/loadsImagesAutomatically
func (w_ WebPreferences) SetLoadsImagesAutomatically(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLoadsImagesAutomatically:"), value)
} /* debug [instance_properties/setter]: loadsImagesAutomatically */

// The minimum font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/minimumFontSize
func (w_ WebPreferences) MinimumFontSize() int {
	rv := objc.Send[int](w_.ID, objc.Sel("minimumFontSize"))
	return rv
} /* debug [instance_properties/getter]: minimumFontSize */

// The minimum font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/minimumFontSize
func (w_ WebPreferences) SetMinimumFontSize(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinimumFontSize:"), value)
} /* debug [instance_properties/setter]: minimumFontSize */

// The minimum logical font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/minimumLogicalFontSize
func (w_ WebPreferences) MinimumLogicalFontSize() int {
	rv := objc.Send[int](w_.ID, objc.Sel("minimumLogicalFontSize"))
	return rv
} /* debug [instance_properties/getter]: minimumLogicalFontSize */

// The minimum logical font size of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/minimumLogicalFontSize
func (w_ WebPreferences) SetMinimumLogicalFontSize(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinimumLogicalFontSize:"), value)
} /* debug [instance_properties/setter]: minimumLogicalFontSize */

// A Boolean that indicates whether or not private browsing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/privateBrowsingEnabled
func (w_ WebPreferences) PrivateBrowsingEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("privateBrowsingEnabled"))
	return rv
} /* debug [instance_properties/getter]: privateBrowsingEnabled */

// A Boolean that indicates whether or not private browsing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/privateBrowsingEnabled
func (w_ WebPreferences) SetPrivateBrowsingEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPrivateBrowsingEnabled:"), value)
} /* debug [instance_properties/setter]: privateBrowsingEnabled */

// The sans serif font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/sansSerifFontFamily
func (w_ WebPreferences) SansSerifFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("sansSerifFontFamily"))
	return rv
} /* debug [instance_properties/getter]: sansSerifFontFamily */

// The sans serif font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/sansSerifFontFamily
func (w_ WebPreferences) SetSansSerifFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSansSerifFontFamily:"), value)
} /* debug [instance_properties/setter]: sansSerifFontFamily */

// The serif font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/serifFontFamily
func (w_ WebPreferences) SerifFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("serifFontFamily"))
	return rv
} /* debug [instance_properties/getter]: serifFontFamily */

// The serif font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/serifFontFamily
func (w_ WebPreferences) SetSerifFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSerifFontFamily:"), value)
} /* debug [instance_properties/setter]: serifFontFamily */

// A Boolean that indicates whether or not the web view should include backgrounds when printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/shouldPrintBackgrounds
func (w_ WebPreferences) ShouldPrintBackgrounds() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldPrintBackgrounds"))
	return rv
} /* debug [instance_properties/getter]: shouldPrintBackgrounds */

// A Boolean that indicates whether or not the web view should include backgrounds when printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/shouldPrintBackgrounds
func (w_ WebPreferences) SetShouldPrintBackgrounds(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShouldPrintBackgrounds:"), value)
} /* debug [instance_properties/setter]: shouldPrintBackgrounds */

// The standard font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/standardFontFamily
func (w_ WebPreferences) StandardFontFamily() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("standardFontFamily"))
	return rv
} /* debug [instance_properties/getter]: standardFontFamily */

// The standard font family of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/standardFontFamily
func (w_ WebPreferences) SetStandardFontFamily(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setStandardFontFamily:"), value)
} /* debug [instance_properties/setter]: standardFontFamily */

// A Boolean that indicates whether incremental rendering should be suppressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/suppressesIncrementalRendering
func (w_ WebPreferences) SuppressesIncrementalRendering() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("suppressesIncrementalRendering"))
	return rv
} /* debug [instance_properties/getter]: suppressesIncrementalRendering */

// A Boolean that indicates whether incremental rendering should be suppressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/suppressesIncrementalRendering
func (w_ WebPreferences) SetSuppressesIncrementalRendering(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSuppressesIncrementalRendering:"), value)
} /* debug [instance_properties/setter]: suppressesIncrementalRendering */

// A Boolean that indicates whether or not the tab key will focus links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/tabsToLinks
func (w_ WebPreferences) TabsToLinks() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("tabsToLinks"))
	return rv
} /* debug [instance_properties/getter]: tabsToLinks */

// A Boolean that indicates whether or not the tab key will focus links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/tabsToLinks
func (w_ WebPreferences) SetTabsToLinks(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabsToLinks:"), value)
} /* debug [instance_properties/setter]: tabsToLinks */

// A Boolean that indicates whether or not user style sheets are enabled in the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/userStyleSheetEnabled
func (w_ WebPreferences) UserStyleSheetEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("userStyleSheetEnabled"))
	return rv
} /* debug [instance_properties/getter]: userStyleSheetEnabled */

// A Boolean that indicates whether or not user style sheets are enabled in the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/userStyleSheetEnabled
func (w_ WebPreferences) SetUserStyleSheetEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserStyleSheetEnabled:"), value)
} /* debug [instance_properties/setter]: userStyleSheetEnabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/userStyleSheetLocation
func (w_ WebPreferences) UserStyleSheetLocation() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("userStyleSheetLocation"))
	return rv
} /* debug [instance_properties/getter]: userStyleSheetLocation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/userStyleSheetLocation
func (w_ WebPreferences) SetUserStyleSheetLocation(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUserStyleSheetLocation:"), value)
} /* debug [instance_properties/setter]: userStyleSheetLocation */

// A Boolean that indicates whether the web views associated with the receiver should use the shared page cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/usesPageCache
func (w_ WebPreferences) UsesPageCache() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("usesPageCache"))
	return rv
} /* debug [instance_properties/getter]: usesPageCache */

// A Boolean that indicates whether the web views associated with the receiver should use the shared page cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebPreferences/usesPageCache
func (w_ WebPreferences) SetUsesPageCache(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setUsesPageCache:"), value)
} /* debug [instance_properties/setter]: usesPageCache */

// A Boolean that indicates whether or not the web view allows plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/arepluginsenabled
func (w_ WebPreferences) ArePlugInsEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("arePlugInsEnabled"))
	return rv
} /* debug [instance_properties/getter]: arePlugInsEnabled */

// A Boolean that indicates whether or not the web view allows plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/arepluginsenabled
func (w_ WebPreferences) SetArePlugInsEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setArePlugInsEnabled:"), value)
} /* debug [instance_properties/setter]: arePlugInsEnabled */

// A Boolean that indicates whether or not the web view allows Java.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/isjavaenabled
func (w_ WebPreferences) IsJavaEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isJavaEnabled"))
	return rv
} /* debug [instance_properties/getter]: isJavaEnabled */

// A Boolean that indicates whether or not the web view allows Java.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/isjavaenabled
func (w_ WebPreferences) SetIsJavaEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsJavaEnabled:"), value)
} /* debug [instance_properties/setter]: isJavaEnabled */

// A Boolean that indicates whether or not the web view allows JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/isjavascriptenabled
func (w_ WebPreferences) IsJavaScriptEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isJavaScriptEnabled"))
	return rv
} /* debug [instance_properties/getter]: isJavaScriptEnabled */

// A Boolean that indicates whether or not the web view allows JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences/isjavascriptenabled
func (w_ WebPreferences) SetIsJavaScriptEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsJavaScriptEnabled:"), value)
} /* debug [instance_properties/setter]: isJavaScriptEnabled */

// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferencesidentifier
func (w_ WebPreferences) PreferencesIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("preferencesIdentifier"))
	return rv
} /* debug [instance_properties/getter]: preferencesIdentifier */

// The identifier of the receiver’s preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/preferencesidentifier
func (w_ WebPreferences) SetPreferencesIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferencesIdentifier:"), value)
} /* debug [instance_properties/setter]: preferencesIdentifier */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebPreferences */
