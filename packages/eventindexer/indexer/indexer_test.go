package indexer

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/contracts/swap"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/http"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/mock"
)

func newTestService(syncMode SyncMode, watchMode WatchMode) *Indexer {
	srv, _ := http.NewServer(http.NewServerOpts{
		EventRepo:      &mock.EventRepository{},
		StatRepo:       &mock.StatRepository{},
		NFTBalanceRepo: &mock.NFTBalanceRepository{},
		ChartRepo:      &mock.ChartRepository{},
		Echo:           echo.New(),
		CorsOrigins:    []string{},
		EthClient:      &ethclient.Client{},
	})

	return &Indexer{
		accountRepo:        &mock.AccountRepository{},
		eventRepo:          &mock.EventRepository{},
		processedBlockRepo: &mock.ProcessedBlockRepository{},
		statRepo:           &mock.StatRepository{},
		nftBalanceRepo:     &mock.NFTBalanceRepository{},
		txRepo:             &mock.TransactionRepository{},
		blockRepo:          &mock.BlockRepository{},
		ethClient:          &ethclient.Client{},

		processingBlockHeight: 0,

		blockBatchSize: 100,

		// TODO: mock bridge
		// bridge: &mock.Bridge{},
		swaps: []*swap.Swap{},

		httpPort: 4102,
		srv:      srv,

		indexNfts: true,
		layer:     "l1",

		wg:  &sync.WaitGroup{},
		ctx: context.Background(),

		watchMode: watchMode,
		syncMode:  syncMode,
	}
}
