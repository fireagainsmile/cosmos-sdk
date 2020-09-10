package simapp

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
	"github.com/spf13/cobra"
)
var(
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

func AddPeerId() *cobra.Command  {
	cmd := &cobra.Command{
		Use: "add-peer [peer list]",
		Short: "config allowed peer id list when starting the node",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			storedIDList = args[0]
		},
	}
	return cmd
}


