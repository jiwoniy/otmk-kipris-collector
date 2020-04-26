package crawlers

import (
	"fmt"
	"github.com/onenodeinc/fullnode-gateway/x/types"
)

func NewCrawler(cfg types.CrawlerConfig) (types.Crawler, error) {
	var c types.Crawler
	if cfg.Network == string(types.CosmosGaia13007) {
		c = NewTendermintCrawler(cfg.Nodes, types.CosmosGaia13007)
	} else if cfg.Network == string(types.CosmosHub3) {
		c = NewTendermintCrawler(cfg.Nodes, types.CosmosHub3)
	} else if cfg.Network == string(types.KavaMainnet2) {
		c = NewTendermintCrawler(cfg.Nodes, types.KavaMainnet2)
	} else if cfg.Network == string(types.KavaTestnet4000) {
		c = NewTendermintCrawler(cfg.Nodes, types.KavaTestnet4000)
	} else {
		return nil, fmt.Errorf("cannot find crawler for network[%s]", cfg.Network)
	}
	return c, nil
}