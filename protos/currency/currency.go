package currency

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

// Currency struct
type Currency struct {
	log hclog.Logger
}

// NewCurrency constructor
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

// GetRate method
func (c *Currency) GetRate(ctx context.Context, rr *RateRequest) (*RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &RateResponse{Rate: 0.5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {

}
