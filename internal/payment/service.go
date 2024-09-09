package payment

import (
	"Gopatungan/internal/user"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midtransServerKey := os.Getenv("API_MIDTRANS_SERVER_KEY")
	midtransClientKey := os.Getenv("API_MIDTRANS_CLIENT_KEY")

	midtrans.ServerKey = midtransServerKey
	midtrans.ClientKey = midtransClientKey

	midtrans.Environment = midtrans.Sandbox

	snapReq := &snap.Request{
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snap.CreateTransaction(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
