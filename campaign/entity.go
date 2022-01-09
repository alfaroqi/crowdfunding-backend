package campaign

import (
	"bwastartup/user"
	"time"

	"github.com/leekchan/accounting"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

func (c Campaign) GoalAmountInIDR() string {
	a := accounting.Accounting{Symbol: "Rp ", Precision: 2, Thousand: ".", Decimal: ","}
	return a.FormatMoney(float64(c.GoalAmount))
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
