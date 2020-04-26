package crawlers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/onenodeinc/fullnode-gateway/x/types"
	"github.com/onenodeinc/fullnode-gateway/x/utils"
)

type tendermintCrawler struct {
	endpt *utils.EndptSelector
	net   types.Network
}

func NewTendermintCrawler(endpts []string, net types.Network) types.Crawler {
	return &tendermintCrawler{
		endpt: utils.NewEndptSelector(endpts),
		net:   net,
	}
}

func (c *tendermintCrawler) GetNetwork() types.Network {
	return c.net
}

func (c *tendermintCrawler) GetTx(hash string) (*types.TxInfo, error) {
	url := fmt.Sprintf("/txs/%s", hash)
	data := make(map[string]interface{})

	caller, err := utils.BuildRESTCaller(c.endpt.Get()).Build()
	if err != nil {
		return nil, err
	}

	err = caller.Get(url, nil, nil, &data)
	if err != nil {
		return nil, err
	}

	heightStr, err := utils.NewPluck(data).Key("height").PickOneString()
	if err != nil {
		heightStr = "0"
	}
	height, err := strconv.ParseInt(heightStr, 10, 64)

	txType := make([]string, 0)
	msgTypeList := utils.NewPluck(data).Key("tx").Key("value").Key("msg").AllIdx().Key("type").PickAll()

	for _, msgType := range msgTypeList {
		msgTypeStr, ok := msgType.(string)
		if ok {
			txType = append(txType, msgTypeStr)
		}
	}

	relatedAddrs := make([]string, 0)
	relatedAddrSet := make(map[string]int)

	// To extract all names of account fields, run following command after cloning the blockchain proj. github
	// cat `find . -name msg\* | grep -v test | xargs` | grep Address | grep json
	interested := []string{
		"from_address", // Cosmos
		"to_address",
		"address",
		"proposer",
		"depositor",
		"voter",
		"sender",
		"submitter",
		"withdraw_address",
		"delegator_address",
		"validator_address",
		"validator_src_address",
		"validator_dst_address",
	}
	msgAddrList := utils.NewPluck(data).Key("tx").Key("value").Key("msg").AllIdx().Key("value").Keys(interested).PickAll()
	for _, msgAddr := range msgAddrList {
		msgAddrStr, ok := msgAddr.(string)
		if ok {
			relatedAddrSet[msgAddrStr] = 1
		}
	}

	if c.net == types.KavaTestnet4000 || c.net == types.KavaMainnet2 {
		interested := []string{
			"from", // KAVA
			"depositor",
			"owner",
			"bidder",
		}
		msgAddrList := utils.NewPluck(data).Key("tx").Key("value").Key("msg").AllIdx().Key("value").Keys(interested).PickAll()
		for _, msgAddr := range msgAddrList {
			msgAddrStr, ok := msgAddr.(string)
			if ok {
				relatedAddrSet[msgAddrStr] = 1
			}
		}
	}

	for addrStr, _ := range relatedAddrSet {
		relatedAddrs = append(relatedAddrs, addrStr)
	}

	marshaledData, err := json.Marshal(data)

	return &types.TxInfo{
		Height:           uint64(height),
		Hash:             hash,
		TxType:           txType,
		Network:          c.net,
		RelatedAddresses: relatedAddrs,
		Content:          string(marshaledData),
	}, nil
}

func (c *tendermintCrawler) GetLatestHeight() (uint64, error) {
	url := fmt.Sprintf("/blocks/latest")
	data := make(map[string]interface{})

	caller, err := utils.BuildRESTCaller(c.endpt.Get()).Build()
	if err != nil {
		return 0, err
	}

	err = caller.Get(url, nil, nil, &data)
	if err != nil {
		return 0, err
	}

	pickedStr, err := utils.NewPluck(data).Key("block_meta").Key("header").Key("height").PickOneString()
	if err != nil {
		return 0, err
	}

	picked, err := strconv.ParseUint(pickedStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return picked, nil
}

func (c *tendermintCrawler) GetTxsFromHeight(height uint64) ([]string, error) {
	if height < 1 {
		return nil, fmt.Errorf("height should be larger than 0")
	}

	url := fmt.Sprintf("/blocks/%d", height)
	data := make(map[string]interface{})

	caller, err := utils.BuildRESTCaller(c.endpt.Get()).Build()
	if err != nil {
		return nil, err
	}

	err = caller.Get(url, nil, nil, &data)
	if err != nil {
		return nil, err
	}

	pickedAll, err := utils.NewPluck(data).Key("block").Key("data").Key("txs").PickOne()
	if err != nil {
		return nil, err
	}
	if pickedAll == nil {
		return make([]string, 0), nil
	}

	txlst, ok := pickedAll.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected []string, but got %v", txlst)
	}

	txHashList := make([]string, 0)
	for _, txObj := range txlst {
		tx, ok := txObj.(string)
		if ok {
			buff := make([]byte, len(tx))
			decReader := base64.NewDecoder(base64.RawStdEncoding, strings.NewReader(tx))
			readsz, err := decReader.Read(buff)
			if err != nil {
				continue
			}

			slice := buff[0:readsz]
			h := sha256.Sum256(slice)
			txhash := strings.ToUpper(hex.EncodeToString(h[:]))

			txHashList = append(txHashList, txhash)
		}
	}

	return txHashList, nil
}
