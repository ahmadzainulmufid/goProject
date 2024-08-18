package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	CreateAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreateAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}

type userTransactionFormatter struct {
	ID       int               `json:"id"`
	Amount   int               `json:"amount"`
	Status   string            `json:"status"`
	CreateAt time.Time         `json:"created_at"`
	Campaign CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) userTransactionFormatter {
	formatter := userTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreateAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []userTransactionFormatter {
	if len(transactions) == 0 {
		return []userTransactionFormatter{}
	}

	var transactionsFormatter []userTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}
