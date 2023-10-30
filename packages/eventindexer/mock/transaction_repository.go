package mock

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionRepository struct {
}

func (r *TransactionRepository) Save(
	ctx context.Context,
	tx *types.Transaction,
	sender common.Address,
	blockID *big.Int,
	transactedAt time.Time,
	contractAddress common.Address,
) error {
	return nil
}
