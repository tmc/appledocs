// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

/* debug [enums.gen.go]: Generating 5 enums for MediaAccessibility */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MACaptionAppearanceBehavior (2 cases) */
// MACaptionAppearanceBehavior - A value that indicates the preferred behavior for a preference setting.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceBehavior
type MACaptionAppearanceBehavior uint

const (
	// kMACaptionAppearanceBehaviorUseContentIfAvailable - The preference setting should be used unless the content media being played has its own custom value for this setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceBehavior/useContentIfAvailable
	kMACaptionAppearanceBehaviorUseContentIfAvailable MACaptionAppearanceBehavior = 0
	// kMACaptionAppearanceBehaviorUseValue - The preference setting should always be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceBehavior/useValue
	kMACaptionAppearanceBehaviorUseValue MACaptionAppearanceBehavior = 0
)

/* debug [enums.gen.go]: Processing enum MACaptionAppearanceDisplayType (3 cases) */
// MACaptionAppearanceDisplayType - A value that specifies the type of captions to display.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDisplayType
type MACaptionAppearanceDisplayType uint

const (
	// kMACaptionAppearanceDisplayTypeAlwaysOn - The most robust available captioning track should always be displayed, whether subtitles, CC, or SDH. This option is selected by a switch labeled “Closed Captions + SDH” (on the Subtitles & Captioning page of iOS) and “Prefer Closed Captions and SDH” checkbox (on the Captions pane of the Accessibility options in macOS).
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDisplayType/alwaysOn
	kMACaptionAppearanceDisplayTypeAlwaysOn MACaptionAppearanceDisplayType = 0
	// kMACaptionAppearanceDisplayTypeAutomatic - If the language of the audio track differs from the system locale, then captions matching the system locale should be displayed (if available). If the language of the audio and the language of the system locale match, no captions are shown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDisplayType/automatic
	kMACaptionAppearanceDisplayTypeAutomatic MACaptionAppearanceDisplayType = 0
	// kMACaptionAppearanceDisplayTypeForcedOnly - Do not display captions unless they are forced for translation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDisplayType/forcedOnly
	kMACaptionAppearanceDisplayTypeForcedOnly MACaptionAppearanceDisplayType = 0
)

/* debug [enums.gen.go]: Processing enum MACaptionAppearanceDomain (2 cases) */
// MACaptionAppearanceDomain - A value that specifies which domain to retrieve a preference setting from.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDomain
type MACaptionAppearanceDomain uint

const (
	// kMACaptionAppearanceDomainDefault - The system default value for the setting should be returned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDomain/default
	kMACaptionAppearanceDomainDefault MACaptionAppearanceDomain = 0
	// kMACaptionAppearanceDomainUser - The user’s preferred value for the setting should be returned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDomain/user
	kMACaptionAppearanceDomainUser MACaptionAppearanceDomain = 0
)

/* debug [enums.gen.go]: Processing enum MACaptionAppearanceFontStyle (8 cases) */
// MACaptionAppearanceFontStyle - A value that specifies a font style.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle
type MACaptionAppearanceFontStyle uint

const (
	// kMACaptionAppearanceFontStyleCasual - The font style preferred for the casual font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/casual
	kMACaptionAppearanceFontStyleCasual MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleCursive - The font style preferred for the cursive font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/cursive
	kMACaptionAppearanceFontStyleCursive MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleDefault - The default font style for all caption text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/default
	kMACaptionAppearanceFontStyleDefault MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleMonospacedWithoutSerif - The font style preferred for the monospaced sans serif font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/monospacedWithoutSerif
	kMACaptionAppearanceFontStyleMonospacedWithoutSerif MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleMonospacedWithSerif - The font style preferred for the monospaced serif font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/monospacedWithSerif
	kMACaptionAppearanceFontStyleMonospacedWithSerif MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleProportionalWithoutSerif - The font style preferred for the proportional sans serif font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/proportionalWithoutSerif
	kMACaptionAppearanceFontStyleProportionalWithoutSerif MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleProportionalWithSerif - The font style preferred for the proportional serif font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/proportionalWithSerif
	kMACaptionAppearanceFontStyleProportionalWithSerif MACaptionAppearanceFontStyle = 0
	// kMACaptionAppearanceFontStyleSmallCapital - The font style preferred for the small capital font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceFontStyle/smallCapital
	kMACaptionAppearanceFontStyleSmallCapital MACaptionAppearanceFontStyle = 0
)

/* debug [enums.gen.go]: Processing enum MACaptionAppearanceTextEdgeStyle (6 cases) */
// MACaptionAppearanceTextEdgeStyle - A value that specifies a style for the outside of the text.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle
type MACaptionAppearanceTextEdgeStyle uint

const (
	// kMACaptionAppearanceTextEdgeStyleDepressed - An edge makes the text appear pushed in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/depressed
	kMACaptionAppearanceTextEdgeStyleDepressed MACaptionAppearanceTextEdgeStyle = 0
	// kMACaptionAppearanceTextEdgeStyleDropShadow - An edge makes the text appear to float above the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/dropShadow
	kMACaptionAppearanceTextEdgeStyleDropShadow MACaptionAppearanceTextEdgeStyle = 0
	// kMACaptionAppearanceTextEdgeStyleNone - The text should not have a styled edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/none
	kMACaptionAppearanceTextEdgeStyleNone MACaptionAppearanceTextEdgeStyle = 0
	// kMACaptionAppearanceTextEdgeStyleRaised - An edge makes the text appear to rise above the background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/raised
	kMACaptionAppearanceTextEdgeStyleRaised MACaptionAppearanceTextEdgeStyle = 0
	// kMACaptionAppearanceTextEdgeStyleUndefined - An edge style has not been specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/undefined
	kMACaptionAppearanceTextEdgeStyleUndefined MACaptionAppearanceTextEdgeStyle = 0
	// kMACaptionAppearanceTextEdgeStyleUniform - A thin outline lies along the edge of the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceTextEdgeStyle/uniform
	kMACaptionAppearanceTextEdgeStyleUniform MACaptionAppearanceTextEdgeStyle = 0
)


