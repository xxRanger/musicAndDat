package models

type NftMarketTable struct {
	NftLdefIndex string `orm:"pk;unique"`
	OwnerWalletAddress string
	OwnerUserName string
	MpId string
	NftAdminId string
	Price int
	Qty int
	NumSold int
	Active bool
	ActiveTicker string
}

func (this *NftMarketTable) TableIndex() [][]string {
	return [][]string {
		[]string{"NftAdminId"},
		[]string{"OwnerWalletAddress"},
		[]string{"OwnerUserName"},
	}
}


