// Code generated from Apple documentation for IdentityLookup. DO NOT EDIT.

package identitylookup

// Enum types and constants
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

// ILMessageFilterAction - Responds to a received message with a filter action.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction
type ILMessageFilterAction uint

const (
	// ILMessageFilterActionJunk - Prevents the system from showing the message normally, filtered as a Junk message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterAction/junk
	ILMessageFilterActionJunk ILMessageFilterAction = 0
)

// ILMessageFilterError - IdentityLookup error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterError-swift.struct/Code
type ILMessageFilterError uint

// ILMessageFilterSubAction - Responds to a received message with a filter subaction.
//
// [Full Topic]: https://developer.apple.com/documentation/IdentityLookup/ILMessageFilterSubAction
type ILMessageFilterSubAction uint

const (
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
)


