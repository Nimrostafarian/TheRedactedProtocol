package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/crescent-network/crescent/x/farming/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetPlanKey() {
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPlanKey(0))
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetPlanKey(9))
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPlanKey(10))
}

func (s *keysTestSuite) TestGetStakingKey() {
	testCases := []struct {
		stakingCoinDenom string
		farmerAcc        sdk.AccAddress
		expected         []byte
	}{
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			[]byte{0x21, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3,
				0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d},
		},
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			[]byte{0x21, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8,
				0x6b, 0xa1, 0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3},
		},
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer3"))),
			[]byte{0x21, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a, 0xcd,
				0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69, 0x54, 0xa6, 0x1e, 0xe2},
		},
	}

	for _, tc := range testCases {
		key := types.GetStakingKey(tc.stakingCoinDenom, tc.farmerAcc)
		s.Require().Equal(tc.expected, key)

		stakingCoinDenom, farmerAcc := types.ParseStakingKey(key)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
		s.Require().Equal(tc.farmerAcc, farmerAcc)
	}
}

func (s *keysTestSuite) TestGetStakingIndexKey() {

	testCases := []struct {
		farmerAcc        sdk.AccAddress
		stakingCoinDenom string
		expected         []byte
	}{
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			sdk.DefaultBondDenom,
			[]byte{0x22, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf,
				0x1b, 0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			sdk.DefaultBondDenom,
			[]byte{0x22, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8, 0x6b, 0xa1, 0xed, 0xfb,
				0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			sdk.DefaultBondDenom,
			[]byte{0x22, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8, 0x6b, 0xa1, 0xed, 0xfb,
				0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
	}

	for _, tc := range testCases {
		key := types.GetStakingIndexKey(tc.farmerAcc, tc.stakingCoinDenom)
		s.Require().Equal(tc.expected, key)

		farmerAcc, stakingCoinDenom := types.ParseStakingIndexKey(key)
		s.Require().Equal(tc.farmerAcc, farmerAcc)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
	}
}

func (s *keysTestSuite) TestGetStakingsByFarmerPrefix() {
	farmer0 := sdk.AccAddress(crypto.AddressHash([]byte("farmer1")))
	farmer1 := sdk.AccAddress(crypto.AddressHash([]byte("farmer2")))
	farmer2 := sdk.AccAddress(crypto.AddressHash([]byte("farmer3")))
	farmer3 := sdk.AccAddress(crypto.AddressHash([]byte("farmer4")))
	s.Require().Equal([]byte{0x22, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3, 0xaa, 0xe5,
		0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d}, types.GetStakingsByFarmerPrefix(farmer0))
	s.Require().Equal([]byte{0x22, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8, 0x6b, 0xa1,
		0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3}, types.GetStakingsByFarmerPrefix(farmer1))
	s.Require().Equal([]byte{0x22, 0x14, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a, 0xcd, 0xf5, 0x7b,
		0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69, 0x54, 0xa6, 0x1e, 0xe2}, types.GetStakingsByFarmerPrefix(farmer2))
	s.Require().Equal([]byte{0x22, 0x14, 0x98, 0x94, 0x3f, 0x57, 0x25, 0xab, 0x66, 0xef, 0x46,
		0x63, 0x4a, 0xfe, 0xeb, 0x8, 0xc0, 0x4a, 0x53, 0x25, 0x2c, 0x9f}, types.GetStakingsByFarmerPrefix(farmer3))
}

func (s *keysTestSuite) TestGetQueuedStakingKey() {
	testCases := []struct {
		stakingCoinDenom string
		farmerAcc        sdk.AccAddress
		expected         []byte
	}{
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			[]byte{0x23, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3,
				0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d},
		},
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			[]byte{0x23, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8,
				0x6b, 0xa1, 0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3},
		},
		{
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer3"))),
			[]byte{0x23, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a, 0xcd,
				0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69, 0x54, 0xa6, 0x1e, 0xe2},
		},
	}

	for _, tc := range testCases {
		key := types.GetQueuedStakingKey(tc.stakingCoinDenom, tc.farmerAcc)
		s.Require().Equal(tc.expected, key)

		stakingCoinDenom, farmerAcc := types.ParseQueuedStakingKey(key)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
		s.Require().Equal(tc.farmerAcc, farmerAcc)
	}
}

func (s *keysTestSuite) TestGetQueuedStakingIndexKey() {
	testCases := []struct {
		farmerAcc        sdk.AccAddress
		stakingCoinDenom string
		expected         []byte
	}{
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			sdk.DefaultBondDenom,
			[]byte{0x24, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf, 0x1b,
				0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			sdk.DefaultBondDenom,
			[]byte{0x24, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d, 0xe8, 0x6b, 0xa1, 0xed, 0xfb, 0x6f,
				0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer3"))),
			sdk.DefaultBondDenom,
			[]byte{0x24, 0x14, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a, 0xcd, 0xf5, 0x7b, 0xb, 0xe7, 0x69,
				0x75, 0x50, 0x9e, 0x69, 0x54, 0xa6, 0x1e, 0xe2, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
	}

	for _, tc := range testCases {
		key := types.GetQueuedStakingIndexKey(tc.farmerAcc, tc.stakingCoinDenom)
		s.Require().Equal(tc.expected, key)

		farmerAcc, stakingCoinDenom := types.ParseQueuedStakingIndexKey(key)
		s.Require().Equal(tc.farmerAcc, farmerAcc)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
	}
}

func (s *keysTestSuite) TestGetQueuedStakingByFarmerPrefix() {
	farmer0 := sdk.AccAddress(crypto.AddressHash([]byte("farmer1")))
	farmer1 := sdk.AccAddress(crypto.AddressHash([]byte("farmer2")))
	farmer2 := sdk.AccAddress(crypto.AddressHash([]byte("farmer3")))
	farmer3 := sdk.AccAddress(crypto.AddressHash([]byte("farmer4")))
	s.Require().Equal([]byte{0x24, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75,
		0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc,
		0xbd, 0x5c, 0xfe, 0x3d}, types.GetQueuedStakingByFarmerPrefix(farmer0))
	s.Require().Equal([]byte{0x24, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a,
		0x5d, 0xe8, 0x6b, 0xa1, 0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb,
		0xfb, 0x6f, 0x28, 0x66, 0xf3}, types.GetQueuedStakingByFarmerPrefix(farmer1))
	s.Require().Equal([]byte{0x24, 0x14, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6,
		0x9a, 0xcd, 0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69,
		0x54, 0xa6, 0x1e, 0xe2}, types.GetQueuedStakingByFarmerPrefix(farmer2))
	s.Require().Equal([]byte{0x24, 0x14, 0x98, 0x94, 0x3f, 0x57, 0x25,
		0xab, 0x66, 0xef, 0x46, 0x63, 0x4a, 0xfe, 0xeb, 0x8, 0xc0, 0x4a,
		0x53, 0x25, 0x2c, 0x9f}, types.GetQueuedStakingByFarmerPrefix(farmer3))
}

func (s *keysTestSuite) TestGetTotalStakingsKey() {
	for _, tc := range []struct {
		stakingCoinDenom string
		expected         []byte
	}{
		{
			"denom1",
			[]byte{0x25, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x31},
		},
		{
			sdk.DefaultBondDenom,
			[]byte{0x25, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
	} {
		key := types.GetTotalStakingsKey(tc.stakingCoinDenom)
		s.Require().Equal(tc.expected, key)

		stakingCoinDenom := types.ParseTotalStakingsKey(key)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
	}
}

func (s *keysTestSuite) TestGetHistoricalRewardsKey() {
	testCases := []struct {
		stakingCoinDenom string
		epoch            uint64
		expected         []byte
	}{
		{
			sdk.DefaultBondDenom,
			1,
			[]byte{0x31, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			sdk.DefaultBondDenom,
			2,
			[]byte{0x31, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2},
		},
	}

	for _, tc := range testCases {
		key := types.GetHistoricalRewardsKey(tc.stakingCoinDenom, tc.epoch)
		s.Require().Equal(tc.expected, key)

		stakingCoinDenom, epoch := types.ParseHistoricalRewardsKey(key)
		s.Require().Equal(tc.stakingCoinDenom, stakingCoinDenom)
		s.Require().Equal(tc.epoch, epoch)
	}
}

func (s *keysTestSuite) TestGetCurrentEpochKey() {
	// key0
	stakingCoinDenom0 := ""
	key0 := types.GetCurrentEpochKey(stakingCoinDenom0)
	s.Require().Equal([]byte{0x32}, key0)

	stakingCoinDenom := types.ParseCurrentEpochKey(key0)
	s.Require().Equal(stakingCoinDenom, stakingCoinDenom0)

	// key1
	stakingCoinDenom1 := sdk.DefaultBondDenom
	key1 := types.GetCurrentEpochKey(stakingCoinDenom1)
	s.Require().Equal([]byte{0x32, 0x73, 0x74, 0x61, 0x6b, 0x65}, key1)

	stakingCoinDenom = types.ParseCurrentEpochKey(key1)
	s.Require().Equal(stakingCoinDenom, stakingCoinDenom1)
}

func (s *keysTestSuite) TestLengthPrefix() {
	denom0 := sdk.DefaultBondDenom
	denom1 := "uatom"

	testCases := []struct {
		stakingCoinDenom string
		length           int
		expected         []byte
	}{
		{
			denom0,
			len(denom0),
			[]byte{0x5, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			denom1,
			len(denom1),
			[]byte{0x5, 0x75, 0x61, 0x74, 0x6f, 0x6d},
		},
	}

	for _, tc := range testCases {
		bz := types.LengthPrefixString(tc.stakingCoinDenom)
		s.Require().Equal(tc.length, int(bz[0]))
		s.Require().Equal(tc.expected, bz)
	}
}
