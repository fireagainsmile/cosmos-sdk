package simapp

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)
const(
	storedIDList=""
	storedAddressList=""
)

var IDFilter = func(info string) abci.ResponseQuery {
	peerIdList := strings.Split(storedIDList,",")
	res := abci.ResponseQuery{}
	for _, x := range peerIdList{
		if info == x {
			res.Code = abci.CodeTypeOK
			return res
		}
	}
	res.Code = 1
	return res
}

var AddressFilter = func(info string) abci.ResponseQuery{
	_ = storedAddressList
	return  abci.ResponseQuery{
		Code:abci.CodeTypeOK,
	}
}


