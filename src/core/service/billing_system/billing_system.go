package billingsystem

import (
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"github.com/mixnote/mixnote-api-go/src/core/service/billing_system/merchants"
	// "github.com/mixnote/mixnote-api-go/src/framework/database"
)

type billingSystem struct {}

// var db, _ = database.DBConnection("")
var EnabledMerchants map[string]merchants.IMerchant


var merchant merchants.IMerchant
var user *models.User
func init() {
	
}

func New(user_ *models.User) (billingSystem_ *billingSystem) {
	user = user_
	return
}

func (billingSystem_ *billingSystem) UseMerchant(merchant_ merchants.IMerchant) (*billingSystem) {
	merchant = merchant_
	return billingSystem_
}

func Merchant() (merchants.IMerchant) {
	return merchant
}

func (billingSystem_ *billingSystem)  Customer() {
	merchant.Customer(user)
}