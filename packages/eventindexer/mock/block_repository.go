package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type BlockRepository struct {
}

func (r *BlockRepository) Save(
	ctx context.Context,
	block *types.Block,
	chainID *big.Int,
) error {
	return nil
}
