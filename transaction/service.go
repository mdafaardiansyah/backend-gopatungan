package transaction

import (
	"Gopatungan/campaign"
	"Gopatungan/helper"
	"errors"
	"fmt"
	"time"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != (input.User.ID) {
		return []Transaction{}, errors.New("Unauthorized")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {

	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}

	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	transaction.Code = "ORD" + time.Now().Format("20060102") + "-" + fmt.Sprintf("%04d", helper.RandomInt(1, 9999))

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
