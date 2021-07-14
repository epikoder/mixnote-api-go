package merchants

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/mixnote/mixnote-api-go/configs"
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
)
type _stripe struct {}
var once sync.Once
var Stripe *_stripe
func init() {
	once.Do(func() {
		Stripe = new(_stripe)
	})
}

func (*_stripe) Name() string {
	return "Stripe"
}

func (*_stripe) Identifier() string {
	return "stripe"
}

func (*_stripe) Secret() string {
	if configs.App.Env != "production" {
		return "rk_test_51ItfBTC5pbNQYknt95V2HrScnBXfGmKMWR9zAtjGpJ7kqg8JExNsLMR5JKTJgGlrSvxwtv6gnWUhhMWfa2Khr9iM00QAl6QT9t"
	}
	return "rk_test_51ItfBTC5pbNQYknt95V2HrScnBXfGmKMWR9zAtjGpJ7kqg8JExNsLMR5JKTJgGlrSvxwtv6gnWUhhMWfa2Khr9iM00QAl6QT9t"
}

func (*_stripe) Customer(user *models.User) (*stripe.Customer, error){
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	phone := strconv.Itoa(user.Phone)
	params := &stripe.CustomerParams{
		Email: &user.Email,
		Name: &name,
		Phone: &phone,
	}
	return customer.New(params)
}