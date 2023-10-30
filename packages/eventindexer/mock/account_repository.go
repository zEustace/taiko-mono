package mock

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type AccountRepository struct {
}

func (r *AccountRepository) Save(
	ctx context.Context,
	address common.Address,
	transactedAt time.Time,
) error {
	return nil
}
