package orderregistration

import "github.com/milkaxq/bpcpayment/constants"

type OrderRegistrationRequest struct {
	// Order details

	// bank merchant user name
	Username string `json:"userName" binding:"required"`
	// bank merchant password
	Password string `json:"password" binding:"required"`
	/*
		Number (identifier) of the order in the merchant’s online store system.
		It is unique for every store in the system and is generated on the order registration.
	*/
	OrderNumber string `json:"orderNumber" minLength:"1" maxLength:"32"`
	//Order amount in the minor denomination (for example, cents).
	Amount int64 `json:"amount" minLength:"1" maxLength:"19" binding:"required"`
	// Payment currency code in the ISO 4217 format
	Currency string `json:"currency" minLength:"3" maxLength:"3" binding:"required"`
	// URL to which the customer is redirected after a successful payment.
	ReturnUrl string `json:"returnUrl" minLength:"1" maxLength:"512"`
	// Free form description of the order.
	Description string `json:"description,omitempty" maxLength:"512"`

	// Optional fields
	/*
		Language code in the ISO 639-1 format.
		If unspecified, SmartVista E-commerce Payment Gateway uses
		the default language from the merchant settings.

		Error messages are returned in this language.
	*/
	Language string `json:"language,omitempty" minLength:"2" maxLength:"2"`

	/*
		The order lifespan duration can be taken from the following parameters (from the highest priority to the lowest):
		· sessionTimeoutSecs transferred in a request. It overrides all other order timeout settings.
		· If the sessionTimeoutSecs parameter is not specified, the value from the merchant’s settings is used.
		It is configured by the Alternative session timeout option that must be enabled and
		the Session duration additional parameter that must be specified.
		· If none of the above mentioned settings is specified
		(neither sessionTimeoutSecs nor merchant’s timeout),
		the default value set in Administration > System settings by the default.session.timeout.milliseconds
		system setting is used (the default value is 1200 seconds) .
		If the request contains the expirationDate parameter, the sessionTimeoutSecs parameter is ignored.
	*/
	SessionTimeoutSecs int `json:"sessionTimeoutSecs,omitempty" maxLength:"9"`

	// /*
	// 	Customer identifier in the merchant system. It is used in the bindings functionality.
	// 	This parameter may be present, if the merchant has the permissions to create and manage bindings
	// 	(the Merchant is allowed to use bindings merchant option and other,
	// 		see “Main information tab” in the Merchants section of EPG user guide.)
	// */
	// ClientId string `json:"clientId,omitempty" minLength:"1" maxLength:"512"`
	// // If notifying customers is enabled, the customer’s email address to send notifications to.
	// Email string `json:"email,omitempty"`
	// // If notifying customers is enabled, the customer’s phone number to send SMS notifications to.
	// Phone string `json:"phone,omitempty" maxLength:"255"`

	// // Additional information
	// /*
	// 	Fields used to store additional information. The type is as follows:
	// 	{"param":"value","param2":"value2"}
	// */
	// // JsonParams []map[string]string `json:"jsonParams,omitempty" maxLength:"1024"`

	// /*
	// 	Date and time when the order is terminated. The following format is used: yyyy-MM-ddTHH:mm:ss
	// 	If this parameter is not specified, the sessionTimeoutSecs parameter is used to determine the
	// 	date and time when the order is terminated.
	// */
	// ExpirationDate string `json:"expirationDate,omitempty"`

	// // Binding and payment method
	// /*
	// 	Identifier of the binding that was created earlier.
	// 	It may be used only if the merchant has the permission to work with bindings.
	// 	If this parameter is sent in the registerOrder request, the following actions apply:
	// 	1. This order can be paid only by binding.
	// 	2. The payer is redirected to a payment page where the CVC must be entered.
	// */
	// BindingId string `json:"bindingId,omitempty" maxLength:"255"`
	// /*
	// 	Parameter that enables switching on 3-D Secure check for a merchant.
	// 	The available values are:
	// 	· FORCE_SSL — in this case, the SSL payment allowed option must be enabled for the merchant.
	// 	· FORCE_TDS — in this case, the following options must be enabled for the merchant
	// 	for 3-D Secure checks in different networks:
	// 	oMasterCard Identity Check payments oVISA Secure payments
	// 	oMirAccept payments
	// 	oCUP 3SD payments
	// 	oDiners Club ProtectBuy payments oJCB J/Secure payments
	// 	oAmEx SafeKey payments
	// 	See 3DS configuration tab in the Merchants section in SmartVista E-commerce Payment Gateway User Guide.
	// */
	// Features string `json:"features,omitempty"`
	// /*
	// 	Payment method used for the payment.
	// 	The possible values are:
	// 	· CARD — Payment made by entering the card details
	// 	· CARD_BINDING — Payment made using a binding
	// 	· CARD_MOTO — Payment made using the call
	// 	center
	// 	· UPOP_MOTO (CUP UPOP MOTO) — Payment made China UnionPay Express Pay
	// 	· UPOP — Payment made China UnionPay Secure Pay
	// 	· FILE_BINDING — Payment made using a binding uploaded in a file
	// 	· P2P — Payment made when transferring funds from one card to another
	// 	· APPLE_PAY— Payment made using the Apple Pay service
	// 	· MASTERPASS — Payment made using a Masterpass e-wallet
	// 	· OTHER — Payment used for orders processed outside EPG
	// 	· GOOGLE_PAY — Payment made using the Google Pay service
	// 	· SAMSUNG_PAY — Payment made using the Samsung Pay service
	// */
	// PaymentWay string `json:"paymentWay,omitempty" maxLength:"32"`

	// // Recurring payments (mandatory for recurring payments)
	// /*
	// 	Way in which subsequent recurring payments are processed:
	// 	· MANUAL — for manual processing of recurring payments scheduled on the merchant side.
	// 	· AUTO — for automatic handling of payments scheduled on the EPG side.
	// 	This parameter is mandatory to register a recurring payment.
	// */
	// RecurrenceType string `json:"recurrenceType,omitempty"`
	// /*
	// 	Parameter is used only for automatic recurring
	// 	payments. It is not required for manual
	// 	payments. See Specifying reccurenceFrequency
	// 	for information about how to configure this parameter.
	// */
	// RecurrenceFrequency string `json:"recurrenceFrequency,omitempty"`
	// /*
	// 	Date when the recurring payment starts. The date format is YYYYMMDD.
	// 	YYYYMMDD 00:00:00 time is used for the start date
	// */
	// RecurrenceStartDate string `json:"recurrenceStartDate,omitempty" minLength:"8" maxLength:"8"`
	// /*
	// 	Date when the recurring payment ends. The date format is YYYYMMDD.
	// 	The default value (if empty) is infinite.
	// 	YYYYMMDD 23:59:59 time is used for the end date.
	// */
	// RecurrenceEndDate string `json:"recurrenceEndDate,omitempty" minLength:"8" maxLength:"8"`

	// // 3D Secure 2 parameters (conditional)
	// /*
	// 	Parameters of the 3-D Secure 2 protocol
	// 	authentication. The threeDS2Params parameter is
	// 	a JSON-based structure. See 3-D Secure 2
	// 	parameter list
	// 	for details.
	// 	This parameter is used for 3-D Secure 2 authentication (for 3-D Secure 1 it is unnecessary).
	// 	Depending on the type of channel interface that is used to initiate the transaction (deviceChannel), the parameter is mandatory or optional:
	// 	· Browser-based authentication — optional
	// 	· Application-based authentication — mandator
	// 	Data for SDK collected from the order threeDS2Params parameter is stored in the THREEDS2_TRANS_SDK_INFO database table.
	// */
	// ThreeDS2Params interface{} `json:"threeDS2Params,omitempty"`
}

type OrderRegistrationResponse struct {
	// Unique order ID generated by EPG (absent on error)
	OrderId string `json:"orderId,omitempty"`

	// URL of the payment page (absent on error)
	FormUrl string `json:"formUrl,omitempty"`

	// Response code
	ErrorCode int `json:"errorCode"`

	// Error message (in requested language, empty on success)
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Identifier for recurring payments (only for recurring payments)
	RecurrenceId string `json:"recurrenceId,omitempty"`
}

// getErrorString constructs a user-friendly error message based on the ErrorCode and ErrorMessage.
func (c OrderRegistrationResponse) getErrorString() string {
	errorMessage := constants.ErrorCodeToMessage[c.ErrorCode]

	switch c.ErrorCode {
	case constants.ErrorCodeNoError:
		// No error, nothing to report
	case constants.ErrorCodeProcessed:
		errorMessage += "Order already processed or incorrect childId."
	case constants.ErrorCodeNotPaid:
		errorMessage += "Order registered but not paid."
	case constants.ErrorCodeUnknownCurrency:
		errorMessage += "Unknown currency."
	case constants.ErrorCodeMissingParameter:
		errorMessage += "Missing required parameter(s):"
		if c.ErrorMessage != "" {
			errorMessage += "\n  - " + c.ErrorMessage
		}
	case constants.ErrorCodeIncorrectValue:
		errorMessage += "Incorrect value for a request parameter or language code."
	case constants.ErrorCodeAccessDenied:
		errorMessage += "Access denied or other issues:"
		switch c.ErrorMessage {
		case "Merchant must change the password.":
			errorMessage += "\n  - Merchant password needs to be changed."
		case "Invalid jsonParams[]":
			errorMessage += "\n  - Invalid data format in jsonParams."
		default:
			errorMessage += "\n  - Access denied. Please check your credentials or contact support."
		}
	case constants.ErrorCodeSystemError:
		errorMessage += "System error. Please try again later."
	case constants.ErrorCodeInvalidPaymentMethod:
		errorMessage += "Invalid payment method provided."
	default:
		errorMessage += "Unknown error code. Please contact support."
	}

	// Append additional details from ErrorMessage if present
	if c.ErrorMessage != "" && errorMessage != "" {
		errorMessage += "\n  Additional details: " + c.ErrorMessage
	}

	return errorMessage
}
