package service

import (
	"github.com/CavnHan/wallet-rpc-service/database"
	models "github.com/CavnHan/wallet-rpc-service/services/rest/model"
)

type Service interface {
	GetSupportCoins(*models.ChainRequest) (*models.SupportChainResponse, error)
	GetWalletAddress(*models.ChainRequest) (*models.WalletAddressResponse, error)
}

type HandleSrv struct {
	v        *Validator
	keysView database.KeysView
}

func (h HandleSrv) GetSupportCoins(request *models.ChainRequest) (*models.SupportChainResponse, error) {
	ok := h.v.VerifyWalletAddress(request.Chain, request.Network)
	if ok {
		return &models.SupportChainResponse{
			Support: true,
		}, nil
	} else {
		return &models.SupportChainResponse{
			Support: false,
		}, nil
	}

}

func (h HandleSrv) GetWalletAddress(request *models.ChainRequest) (*models.WalletAddressResponse, error) {
	return &models.WalletAddressResponse{
		PublicKey: "publicKey",
		Address:   "0x00sdfwe",
	}, nil
}

func NewHandleSrv(v *Validator, ksv database.KeysView) Service {
	return &HandleSrv{
		v:        v,
		keysView: ksv,
	}
}
