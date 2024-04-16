// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/confirm-payment": {
            "post": {
                "description": "Confirm Payment by writing otp code.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "epg"
                ],
                "summary": "Confirm Payment",
                "parameters": [
                    {
                        "description": "Payment Confirmation request body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/confirpayment.ConfirmPaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A string of success payment",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Some Confirmation Error",
                        "schema": {
                            "$ref": "#/definitions/constants.HttpError"
                        }
                    }
                }
            }
        },
        "/register-order": {
            "post": {
                "description": "Register an order for payment processing.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "epg"
                ],
                "summary": "Order Registration",
                "parameters": [
                    {
                        "description": "Order Registration Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orderregistration.OrderRegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful order registration",
                        "schema": {
                            "$ref": "#/definitions/orderregistration.OrderRegistrationResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or parameters",
                        "schema": {
                            "$ref": "#/definitions/constants.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "confirpayment.ConfirmPaymentRequest": {
            "type": "object",
            "properties": {
                "MDORDER": {
                    "description": "md order which was created on order registration",
                    "type": "string"
                },
                "passwordEdit": {
                    "description": "otp password that comes to phone number",
                    "type": "string"
                },
                "request_id": {
                    "description": "Request id comes from previouse request of card submission",
                    "type": "string"
                }
            }
        },
        "constants.HttpError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "orderregistration.OrderRegistrationRequest": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "password",
                "returnUrl",
                "userName"
            ],
            "properties": {
                "amount": {
                    "description": "Order amount in the minor denomination (for example, cents).",
                    "type": "integer"
                },
                "currency": {
                    "description": "Payment currency code in the ISO 4217 format",
                    "type": "string",
                    "maxLength": 3,
                    "minLength": 3
                },
                "description": {
                    "description": "Free form description of the order.",
                    "type": "string",
                    "maxLength": 512
                },
                "language": {
                    "description": "Optional fields\n\n\t\tLanguage code in the ISO 639-1 format.\n\t\tIf unspecified, SmartVista E-commerce Payment Gateway uses\n\t\tthe default language from the merchant settings.\n\n\t\tError messages are returned in this language.",
                    "type": "string",
                    "maxLength": 2,
                    "minLength": 2
                },
                "orderNumber": {
                    "description": "Order details\n\n\t\tNumber (identifier) of the order in the merchant’s online store system.\n\t\tIt is unique for every store in the system and is generated on the order registration.",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 1
                },
                "password": {
                    "description": "User’s password.",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 1
                },
                "returnUrl": {
                    "description": "URL to which the customer is redirected after a successful payment.",
                    "type": "string",
                    "maxLength": 512,
                    "minLength": 1
                },
                "sessionTimeoutSecs": {
                    "description": "The order lifespan duration can be taken from the following parameters (from the highest priority to the lowest):\n\t\t· sessionTimeoutSecs transferred in a request. It overrides all other order timeout settings.\n\t\t· If the sessionTimeoutSecs parameter is not specified, the value from the merchant’s settings is used.\n\t\tIt is configured by the Alternative session timeout option that must be enabled and\n\t\tthe Session duration additional parameter that must be specified.\n\t\t· If none of the above mentioned settings is specified\n\t\t(neither sessionTimeoutSecs nor merchant’s timeout),\n\t\tthe default value set in Administration \u003e System settings by the default.session.timeout.milliseconds\n\t\tsystem setting is used (the default value is 1200 seconds) .\n\t\tIf the request contains the expirationDate parameter, the sessionTimeoutSecs parameter is ignored.",
                    "type": "integer"
                },
                "userName": {
                    "description": "User credentials\nLogin of the API user on whose behalf requests are processed for a particular merchant.",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 1
                }
            }
        },
        "orderregistration.OrderRegistrationResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "description": "Response code",
                    "type": "integer"
                },
                "errorMessage": {
                    "description": "Error message (in requested language, empty on success)",
                    "type": "string"
                },
                "formUrl": {
                    "description": "URL of the payment page (absent on error)",
                    "type": "string"
                },
                "orderId": {
                    "description": "Unique order ID generated by EPG (absent on error)",
                    "type": "string"
                },
                "recurrenceId": {
                    "description": "Identifier for recurring payments (only for recurring payments)",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
