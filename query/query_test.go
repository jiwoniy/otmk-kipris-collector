package query

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jiwoniy/otmk-kipris-collector/types"
)

// type AccountResult struct {
// 	Address   string        `json:"address"`
// 	Network   NetworkId     `json:"network"`
// 	BackendId string        `json:"backend_id"`
// 	PublicKey string        `json:"public_key"`
// 	Algorithm SignAlgorithm `json:"algorithm"`
// 	Tag       Tags          `json:"tag"`
// 	Disabled  bool          `json:"disabled"`
// }

// func NewAccountResultFromAccountModel(model *AccountModel) *AccountResult {
// 	return &AccountResult{
// 		Address:   model.Address,
// 		Network:   model.Network,
// 		BackendId: model.KeyLocation.BackendId,
// 		PublicKey: hex.EncodeToString(model.KeyLocation.PublicKey),
// 		Algorithm: model.KeyLocation.Algorithm,
// 		Tag:       TagsFromAccountTags(model.Tags),
// 		Disabled:  model.Disabled,
// 	}
// }

func TestQuery(t *testing.T) {
	config := types.QueryConfig{
		DbType:       "mysql",
		DbConnString: "kipris_server:OnthemarkKipris0507!@@(61.97.187.142:3306)/kipris?charset=utf8&parseTime=True&loc=Local",
	}
	query, err := NewApp(config)

	if err != nil {
		t.Errorf(fmt.Sprintf("fail to create query instance %s", err))
	}

	data := query.GetApplicationNumber("4020200000001")

	test, err := json.Marshal(data)
	fmt.Println(test)
	fmt.Println(string(test))
	fmt.Println(err)

	// res := make([]*sdk.AccountResult, len(accounts))
	// 	for i, v := range accounts {
	// 		res[i] = sdk.NewAccountResultFromAccountModel(&v)
	// 	}

}
