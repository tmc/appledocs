// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest_test

import (
	"github.com/tmc/appledocs/generated/storekittest"
)

// Suppress unused import errors
var _ = storekittest.NewAdTestPostback

// ExampleNewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL demonstrates how to create a AdTestPostback instance using NewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL.
// Creates a test postback for an in-app ad.
func ExampleNewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL() {
	_ = storekittest.NewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL(
		storekittest.AdTestPostbackVersion{}, // version AdTestPostbackVersion
		"adNetworkIdentifier", // adNetworkIdentifier string
		0, // adCampaignIdentifier int
		0, // appStoreItemIdentifier int
		0, // sourceAppStoreItemIdentifier int
		0, // conversionValue int
		0, // fidelityType int
		false, // isRedownload bool
		false, // didWin bool
		"https://example.com", // postbackURL string
	)
	// Output:
}
// ExampleNewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL demonstrates how to create a AdTestPostback instance using NewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL.
// Creates a test postback for a web ad or an in-app ad.
func ExampleNewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL() {
	_ = storekittest.NewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL(
		storekittest.AdTestPostbackVersion{}, // version AdTestPostbackVersion
		"adNetworkIdentifier", // adNetworkIdentifier string
		"sourceIdentifier", // sourceIdentifier string
		0, // appStoreItemIdentifier int
		0, // sourceAppStoreItemIdentifier int
		"sourceDomain", // sourceDomain string
		0, // fidelityType int
		false, // isRedownload bool
		false, // didWin bool
		"https://example.com", // postbackURL string
	)
	// Output:
}
