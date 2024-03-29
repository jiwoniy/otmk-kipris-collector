package query

import (
	"github.com/jiwoniy/otmk-kipris-collector/kipris/model"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/storage"
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

type kiprisQuery struct {
	storage types.Storage
}

func NewApp(config types.QueryConfig) (types.Query, error) {
	storage, err := storage.NewStorage(types.StorageConfig{
		DbType:       config.DbType,
		DbConnString: config.DbConnString,
	})

	if err != nil {
		return nil, err
	}

	return &kiprisQuery{
		storage: storage,
	}, nil
}

func (k *kiprisQuery) GetApplicationNumber(applicationNumber string) *model.TradeMarkInfo {
	reqTradeMarkInfo := model.TradeMarkInfo{
		ApplicationNumber: applicationNumber,
	}

	reqTrademarkDesignationGoodstInfo := model.TrademarkDesignationGoodstInfo{
		ApplicationNumber: applicationNumber,
	}

	var resTradeMarkInfo model.TradeMarkInfo
	var resTrademarkDesignationGoodstInfos []model.TrademarkDesignationGoodstInfo

	k.storage.GetTradeMarkInfo(reqTradeMarkInfo, &resTradeMarkInfo)
	k.storage.GetTrademarkDesignationGoodstInfo(reqTrademarkDesignationGoodstInfo, &resTrademarkDesignationGoodstInfos)
	resTradeMarkInfo.TrademarkDesignationGoodstInfos = resTrademarkDesignationGoodstInfos

	return &resTradeMarkInfo
}
