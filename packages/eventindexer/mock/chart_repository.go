package mock

import (
	"context"

	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
)

type ChartRepository struct {
}

func NewChartRepository() *ChartRepository {
	return &ChartRepository{}
}

func (r *ChartRepository) Find(
	ctx context.Context,
	task string,
	start string,
	end string) (*eventindexer.ChartResponse, error) {
	return nil, nil
}
