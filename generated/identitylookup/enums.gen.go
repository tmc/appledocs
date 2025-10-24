// Code generated from Apple documentation for IdentityLookup. DO NOT EDIT.

package identitylookup

/* debug [enums.gen.go]: Generating 4 enums for IdentityLookup */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ILClassificationAction (4 cases) */
// ILClassificationAction - The actions the system can take in response to the reported communication.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILClassificationAction
type ILClassificationAction uint

const (
	// ILClassificationActionNone - No action is required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILClassificationAction/none
	ILClassificationActionNone ILClassificationAction = 0
	// ILClassificationActionReportJunk - The system should report the communication as junk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILClassificationAction/reportJunk
	ILClassificationActionReportJunk ILClassificationAction = 0
	// ILClassificationActionReportJunkAndBlockSender - The system should report the communication as junk and add the number to the system’s block list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILClassificationAction/reportJunkAndBlockSender
	ILClassificationActionReportJunkAndBlockSender ILClassificationAction = 0
	// ILClassificationActionReportNotJunk - The system should report that the communication is not junk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILClassificationAction/reportNotJunk
	ILClassificationActionReportNotJunk ILClassificationAction = 0
)

/* debug [enums.gen.go]: Processing enum ILMessageFilterAction (6 cases) */
// ILMessageFilterAction - Responds to a received message with a filter action.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction
type ILMessageFilterAction uint

const (
	// ILMessageFilterActionAllow - Allows the system to show the message unfiltered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/allow
	ILMessageFilterActionAllow ILMessageFilterAction = 0
	// ILMessageFilterActionFilter - Prevents the system from showing the message unfiltered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/filter
	ILMessageFilterActionFilter ILMessageFilterAction = 0
	// ILMessageFilterActionJunk - Prevents the system from showing the message normally, filtered as a Junk message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/junk
	ILMessageFilterActionJunk ILMessageFilterAction = 0
	// ILMessageFilterActionNone - Allows the system to show the message unfiltered due to insufficient information to determine an action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/none
	ILMessageFilterActionNone ILMessageFilterAction = 0
	// ILMessageFilterActionPromotion - Prevents the system from showing the message normally, filtered as a Promotional message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/promotion
	ILMessageFilterActionPromotion ILMessageFilterAction = 0
	// ILMessageFilterActionTransaction - Prevents the system from showing the message normally, filtered as a Transactional message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/transaction
	ILMessageFilterActionTransaction ILMessageFilterAction = 0
)

/* debug [enums.gen.go]: Processing enum ILMessageFilterError (5 cases) */
// ILMessageFilterError - IdentityLookup error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code
type ILMessageFilterError uint

const (
	// ILMessageFilterErrorInvalidNetworkURL - The network request URL given by the   key in the app extension’s   file is either missing or invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code/invalidNetworkURL
	ILMessageFilterErrorInvalidNetworkURL ILMessageFilterError = 0
	// ILMessageFilterErrorNetworkRequestFailed - The network request failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code/networkRequestFailed
	ILMessageFilterErrorNetworkRequestFailed ILMessageFilterError = 0
	// ILMessageFilterErrorNetworkURLUnauthorized - The app extension’s containing app isn’t authorized to allow the app extension to defer network requests to the host specified in its   file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code/networkURLUnauthorized
	ILMessageFilterErrorNetworkURLUnauthorized ILMessageFilterError = 0
	// ILMessageFilterErrorRedundantNetworkDeferral - The app extension tried to defer a request to its network service more than once, which isn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code/redundantNetworkDeferral
	ILMessageFilterErrorRedundantNetworkDeferral ILMessageFilterError = 0
	// ILMessageFilterErrorSystem - An unspecified system error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code/system
	ILMessageFilterErrorSystem ILMessageFilterError = 0
)

/* debug [enums.gen.go]: Processing enum ILMessageFilterSubAction (13 cases) */
// ILMessageFilterSubAction - Responds to a received message with a filter subaction.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction
type ILMessageFilterSubAction uint

const (
	// ILMessageFilterSubActionNone - Allows the system to show the message unfiltered due to insufficient information to determine an action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/none
	ILMessageFilterSubActionNone ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionPromotionalCoupons - Prevents the system from showing the message normally, filtered as an Coupons message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/promotionalCoupons
	ILMessageFilterSubActionPromotionalCoupons ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionPromotionalOffers - Prevents the system from showing the message normally, filtered as an Offers message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/promotionalOffers
	ILMessageFilterSubActionPromotionalOffers ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionPromotionalOthers - Prevents the system from showing the message normally, filtered as an Others message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/promotionalOthers
	ILMessageFilterSubActionPromotionalOthers ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalCarrier - Prevents the system from showing the message normally, filtered as a Carrier message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalCarrier
	ILMessageFilterSubActionTransactionalCarrier ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalFinance - Prevents the system from showing the message normally, filtered as a Finance message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalFinance
	ILMessageFilterSubActionTransactionalFinance ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalHealth - Prevents the system from showing the message normally, filtered as a Health message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalHealth
	ILMessageFilterSubActionTransactionalHealth ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalOrders - Prevents the system from showing the message normally, filtered as an Orders (eCommerce) message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalOrders
	ILMessageFilterSubActionTransactionalOrders ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalOthers - Prevents the system from showing the message normally, filtered as an Others message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalOthers
	ILMessageFilterSubActionTransactionalOthers ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalPublicServices - Prevents the system from showing the message normally, filtered as a Government message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalPublicServices
	ILMessageFilterSubActionTransactionalPublicServices ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalReminders - Prevents the system from showing the message normally, filtered as a Reminder message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalReminders
	ILMessageFilterSubActionTransactionalReminders ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalRewards - Prevents the system from showing the message normally, filtered as a Rewards message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalRewards
	ILMessageFilterSubActionTransactionalRewards ILMessageFilterSubAction = 0
	// ILMessageFilterSubActionTransactionalWeather - Prevents the system from showing the message normally, filtered as a Weather message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction/transactionalWeather
	ILMessageFilterSubActionTransactionalWeather ILMessageFilterSubAction = 0
)


