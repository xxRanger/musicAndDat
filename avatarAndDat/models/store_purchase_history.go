package models

import "time"

type StorePurchaseHistroy struct {
	PurchaseId string `orm:"pk;unique"`
	BuyerNickname string
	BuyerWalletId string
	SellerNickname string
	SellerWalletId string
	TransactionAddress string
	ActiveTicker string
	TotalPaid int
	NftLdefIndex string
	Timestamp time.Time `orm:"auto_now_add;type(datetime)"`
	Status int
}

func (this *StorePurchaseHistroy) TableIndex() [][]string {
	return [][]string {
		[]string {"BuyerNickname"},
		[]string {"BuyerWalletId"},
		[]string {"NftLdefIndex"},
		[]string {"TransactionAddress"},
		[]string {"SellerNickname"},
		[]string {"SellerWalletId"},
	}
}

