package service_chain

import (
	"api/conf"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

var Client *ethclient.Client
var sourceAccount = common.HexToAddress("0xAa152bba2885BC2c746a2d7804FfaB821d957aa7")

func init()  {
	var err error
	Client, err = ethclient.Dial(conf.Conf.Geth.Url)
	if err != nil {
		panic(err)
	}
}

func GetSourceBalance() (*big.Int, error) {
	var balance *big.Int
	balance, err := Client.BalanceAt(context.Background(), sourceAccount,nil)
	result := balance.Div(balance,big.NewInt(int64(math.Pow10(18))))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBlockNumber() (uint64, error) {
	result, err := Client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return result, nil
}