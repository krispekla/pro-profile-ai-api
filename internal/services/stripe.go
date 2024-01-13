package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
)

type CreateCheckoutSessionInput struct {
	ProductIds *[]string
	CustomerId string
}

func CreateCheckoutSession(inp *CreateCheckoutSessionInput) (*stripe.CheckoutSession, error) {
	if *inp.ProductIds == nil || len(*inp.ProductIds) == 0 || inp.CustomerId == "" {
		return nil, errors.New("ProductIds and CustomerId are required")
	}
	// TODO: Move to config
	domain := "http://localhost:5173"
	lineItems := []*stripe.CheckoutSessionLineItemParams{}
	for _, priceId := range *inp.ProductIds {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(priceId),
			Quantity: stripe.Int64(1),
		})
	}
	params := &stripe.CheckoutSessionParams{
		UIMode:    stripe.String("embedded"),
		ReturnURL: stripe.String(domain + "/package/buy/return?session_id={CHECKOUT_SESSION_ID}"),
		LineItems: lineItems,
		Customer:  stripe.String(inp.CustomerId),
		CustomerUpdate: &stripe.CheckoutSessionCustomerUpdateParams{
			Address:  stripe.String("auto"),
			Shipping: stripe.String("auto"),
		},
		Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}
	s, err := session.New(params)
	return s, err
}

type CreateCustomerInput struct {
	UserId   string
	Email    string
	FullName string
}

func CreateCustomer(inp *CreateCustomerInput) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Name:  &inp.FullName,
		Email: &inp.Email,
		Metadata: map[string]string{
			"userId": inp.UserId,
		},
	}
	result, err := customer.New(params)
	return result, err
}

func GetStripeSession(id string) (*stripe.CheckoutSession, error) {
	s, err := session.Get(id, nil)
	return s, err
}

func ProcceesStripeWebhook(payload []byte, reqSignature string) (int, error) {
	// // If you are testing your webhook locally with the Stripe CLI you
	// // can find the endpoint's secret by running `stripe listen`
	// // Otherwise, find your endpoint's secret in your webhook settings
	// // in the Developer Dashboard
	// endpointSecret := "whsec_07428a63dc4da4bb4bf76fa7243125a7369a79158d97e081d1dbfc497e537478"

	// // Pass the request body and Stripe-Signature header to ConstructEvent, along
	// // with the webhook signing key.
	// event, err := webhook.ConstructEvent(payload, reqSignature, endpointSecret)

	// if err != nil {
	// 	// fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
	// 	// w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
	// 	return http.StatusBadRequest, errors.New("error verifying webhook signature")
	// }
	event := stripe.Event{}
	if err := json.Unmarshal(payload, &event); err != nil {
		// fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		return http.StatusBadRequest, errors.New("failed to parse webhook body json")
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			// fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			// w.WriteHeader(http.StatusBadRequest)
			return http.StatusBadRequest, errors.New("error parsing webhook JSON")
		}
		fmt.Println("PaymentIntent was successful!")
	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			// fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			// w.WriteHeader(http.StatusBadRequest)
			return http.StatusBadRequest, errors.New("error parsing webhook JSON")
		}
		fmt.Println("PaymentMethod was attached to a Customer!")
	// ... handle other event types
	default:
		fmt.Fprintf(os.Stderr, "unhandled event type: %s\n", event.Type)
	}
	return http.StatusOK, nil
}
