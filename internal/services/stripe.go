package services

import (
	"errors"

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
