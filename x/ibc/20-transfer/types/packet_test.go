package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFungibleTokenPacketDataValidateBasic tests ValidateBasic for FungibleTokenPacketData
func TestFungibleTokenPacketDataValidateBasic(t *testing.T) {
	testPacketDataTransfer := []FungibleTokenPacketData{
		NewFungibleTokenPacketData(coins, addr1, addr2, true, 100),             // valid msg
		NewFungibleTokenPacketData(invalidDenomCoins, addr1, addr2, true, 100), // invalid amount
		NewFungibleTokenPacketData(negativeCoins, addr1, addr2, false, 100),    // amount contains negative coin
		NewFungibleTokenPacketData(coins, emptyAddr, addr2, false, 100),        // missing sender address
		NewFungibleTokenPacketData(coins, addr1, emptyAddr, false, 100),
		NewFungibleTokenPacketData(coins, addr1, emptyAddr, false, 0), // missing recipient address
	}

	testCases := []struct {
		packetData FungibleTokenPacketData
		expPass    bool
		errMsg     string
	}{
		{testPacketDataTransfer[0], true, ""},
		{testPacketDataTransfer[1], false, "invalid amount"},
		{testPacketDataTransfer[2], false, "amount contains negative coin"},
		{testPacketDataTransfer[3], false, "missing sender address"},
		{testPacketDataTransfer[4], false, "missing recipient address"},
		{testPacketDataTransfer[5], false, "timeout is 0"},
	}

	for i, tc := range testCases {
		err := tc.packetData.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "PacketDataTransfer %d failed: %v", i, err)
		} else {
			require.Error(t, err, "Invalid PacketDataTransfer %d passed: %s", i, tc.errMsg)
		}
	}
}