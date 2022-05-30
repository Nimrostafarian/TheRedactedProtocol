package v2_0_0_test

import (
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	budgettypes "github.com/tendermint/budget/x/budget/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/crescent-network/crescent/app"
	"github.com/crescent-network/crescent/app/upgrades/mainnet/v2.0.0"
	"github.com/crescent-network/crescent/cmd/crescentd/cmd"
	utils "github.com/crescent-network/crescent/types"
)

type UpgradeTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.App
}

func (suite *UpgradeTestSuite) SetupTest() {
	cmd.GetConfig()
	suite.app = app.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{Height: 1})

}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

const testUpgradeHeight = 10

var (
	InflationFeeCollector   = "cre17xpfvakm2amg962yls6f84z3kell8c5l53s97s"
	EcosystemIncentive      = "cre1kgshua58cjr2p7hnrvgun68yrqf7ktdzyz2yxv54fqj6uwl4gc4q95txqa"
	EcosystemIncentiveLP    = "cre1wht0xhmuqph4rhzulhejgatthnpeatzjgnnkvqvphq97xr26np0qdvun2s"
	EcosystemIncentiveMM    = "cre1ddn66jv0sjpmck0ptegmhmqtn35qsg2vxyk2hn9sqf4qxtzqz3sq3qhhde"
	EcosystemIncentiveBoost = "cre17zftu6rg7mkmemqxv4whjkvecl0e2ja7j6um9t8qaczp79y72d7q2su2xm"
	DevTeamAddress          = "cre1ge2jm9nkvu2l8cvhc2un4m33d4yy4p0wfag09j"
)

func (suite *UpgradeTestSuite) TestUpgradeV2() {
	testCases := []struct {
		title   string
		before  func()
		upgrade func()
		after   func()
		expPass bool
	}{
		{
			"v2 upgrade mint, budget params",
			func() {
				mintparams := suite.app.MintKeeper.GetParams(suite.ctx)
				fmt.Println(mintparams)
				budgetparams := suite.app.BudgetKeeper.GetParams(suite.ctx)
				genesisTime := utils.ParseTime("2022-04-13T00:00:00Z")
				budgetparams.Budgets = []budgettypes.Budget{
					{
						Name:               "budget-ecosystem-incentive",
						Rate:               sdk.MustNewDecFromStr("0.662500000000000000"),
						SourceAddress:      InflationFeeCollector,
						DestinationAddress: EcosystemIncentive,
						StartTime:          genesisTime,
						EndTime:            genesisTime.AddDate(10, 0, 0),
					},
					{
						Name:               "budget-dev-team",
						Rate:               sdk.MustNewDecFromStr("0.250000000000000000"),
						SourceAddress:      InflationFeeCollector,
						DestinationAddress: DevTeamAddress,
						StartTime:          genesisTime,
						EndTime:            genesisTime.AddDate(10, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-lp-1",
						Rate:               sdk.MustNewDecFromStr("0.500000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveLP,
						StartTime:          genesisTime,
						EndTime:            genesisTime.AddDate(1, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-mm-1",
						Rate:               sdk.MustNewDecFromStr("0.300000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveMM,
						StartTime:          genesisTime,
						EndTime:            genesisTime.AddDate(1, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-boost-1",
						Rate:               sdk.MustNewDecFromStr("0.200000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveBoost,
						StartTime:          genesisTime,
						EndTime:            genesisTime.AddDate(1, 0, 0),
					},

					{
						Name:               "budget-ecosystem-incentive-lp-2",
						Rate:               sdk.MustNewDecFromStr("0.200000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveLP,
						StartTime:          genesisTime.AddDate(1, 0, 0),
						EndTime:            genesisTime.AddDate(2, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-mm-2",
						Rate:               sdk.MustNewDecFromStr("0.300000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveMM,
						StartTime:          genesisTime.AddDate(1, 0, 0),
						EndTime:            genesisTime.AddDate(2, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-boost-2",
						Rate:               sdk.MustNewDecFromStr("0.500000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveBoost,
						StartTime:          genesisTime.AddDate(1, 0, 0),
						EndTime:            genesisTime.AddDate(2, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-lp-3-10",
						Rate:               sdk.MustNewDecFromStr("0.100000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveLP,
						StartTime:          genesisTime.AddDate(2, 0, 0),
						EndTime:            genesisTime.AddDate(10, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-mm-3-10",
						Rate:               sdk.MustNewDecFromStr("0.300000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveMM,
						StartTime:          genesisTime.AddDate(2, 0, 0),
						EndTime:            genesisTime.AddDate(10, 0, 0),
					},
					{
						Name:               "budget-ecosystem-incentive-boost-3-10",
						Rate:               sdk.MustNewDecFromStr("0.600000000000000000"),
						SourceAddress:      EcosystemIncentive,
						DestinationAddress: EcosystemIncentiveBoost,
						StartTime:          genesisTime.AddDate(2, 0, 0),
						EndTime:            genesisTime.AddDate(10, 0, 0),
					},
				}
				suite.app.BudgetKeeper.SetParams(suite.ctx, budgetparams)
				suite.Require().Equal(InflationFeeCollector, budgetparams.Budgets[0].SourceAddress)
				suite.Require().Equal(InflationFeeCollector, budgetparams.Budgets[1].SourceAddress)
				suite.Require().Len(budgetparams.Budgets, 11)
			},
			func() {
				// add test upgrade plan
				suite.ctx = suite.ctx.WithBlockHeight(testUpgradeHeight - 1)
				plan := upgradetypes.Plan{Name: v2_0_0.UpgradeName, Height: testUpgradeHeight}
				err := suite.app.UpgradeKeeper.ScheduleUpgrade(suite.ctx, plan)
				suite.Require().NoError(err)
				_, exists := suite.app.UpgradeKeeper.GetUpgradePlan(suite.ctx)
				suite.Require().True(exists)

				suite.ctx = suite.ctx.WithBlockHeight(testUpgradeHeight)
				suite.Require().NotPanics(func() {
					beginBlockRequest := abci.RequestBeginBlock{}
					suite.app.BeginBlocker(suite.ctx, beginBlockRequest)
				})
			},
			func() {
				mintparams := suite.app.MintKeeper.GetParams(suite.ctx)
				budgetparams := suite.app.BudgetKeeper.GetParams(suite.ctx)
				budgetparamsJson, _ := json.Marshal(budgetparams)

				// TODO: remove debug logging
				fmt.Println(mintparams)
				fmt.Println(string(budgetparamsJson))

				totalRateOfMintPool := sdk.ZeroDec()
				for _, budget := range budgetparams.Budgets {
					if budget.SourceAddress == mintparams.MintPoolAddress {
						totalRateOfMintPool = totalRateOfMintPool.Add(budget.Rate)
					}
				}
				suite.Require().EqualValues(sdk.OneDec(), totalRateOfMintPool)
				suite.Require().Equal(mintparams.MintPoolAddress, budgetparams.Budgets[0].SourceAddress)
				suite.Require().Equal(mintparams.MintPoolAddress, budgetparams.Budgets[1].SourceAddress)
				suite.Require().Equal(mintparams.MintPoolAddress, budgetparams.Budgets[2].SourceAddress)
				suite.Require().Len(budgetparams.Budgets, 12)
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.title, func() {
			suite.SetupTest()

			tc.before()
			tc.upgrade()
			tc.after()
		})
	}
}