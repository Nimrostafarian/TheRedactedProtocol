package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmosquad-labs/squad/x/claim/types"
	claimtypes "github.com/cosmosquad-labs/squad/x/claim/types"
	liqtypes "github.com/cosmosquad-labs/squad/x/liquidity/types"
	liqstakingtypes "github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

func PrepareGenesisCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-genesis [network-type] [chain-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Prepare a genesis file with initial setup",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Prepare a genesis file with initial setup.

The initial setup includes initial params for squad modules

Example:
$ %s prepare-genesis mainnet squad-1
$ %s prepare-genesis testnet squad-1

The genesis output file is at $HOME/.squadapp/config/genesis.json
`,
				version.AppName,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)

			serverCfg := serverCtx.Config

			// Read genesis file
			genFile := serverCfg.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// Parse genesis params depending on the network type
			networkType := args[0]
			genParams := parseGenesisParams(networkType)
			if genParams == nil {
				return fmt.Errorf("you must choose between mainnet (m) or testnet (t): %s", args[0])
			}

			// Prepare genesis
			chainID := args[1]
			appState, genDoc, err = PrepareGenesis(clientCtx, appState, genDoc, genParams, chainID)
			if err != nil {
				return fmt.Errorf("failed to prepare genesis %w", err)
			}

			if err := mbm.ValidateGenesis(clientCtx.Codec, clientCtx.TxConfig, appState); err != nil {
				return fmt.Errorf("failed to validate genesis file: %w", err)
			}

			// Marshal and save the app state
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON

			// Export the genesis state to a file
			if err := genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return fmt.Errorf("failed to export genesis file %w", err)
			}

			return err
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func PrepareGenesis(
	clientCtx client.Context,
	appState map[string]json.RawMessage,
	genDoc *tmtypes.GenesisDoc,
	genParams *GenesisParams,
	chainID string,
) (map[string]json.RawMessage, *tmtypes.GenesisDoc, error) {
	cdc := clientCtx.Codec

	genDoc.ChainID = chainID
	genDoc.GenesisTime = genParams.GenesisTime
	genDoc.ConsensusParams = genParams.ConsensusParams

	// Bank module app state
	bankGenState := banktypes.DefaultGenesisState()
	bankGenState.Balances = genParams.BankGenesisStates.Balances
	bankGenState.Supply = genParams.BankGenesisStates.Supply
	bankGenStateBz := cdc.MustMarshalJSON(bankGenState)
	appState[banktypes.ModuleName] = bankGenStateBz

	// Claim module app state
	claimGenState := claimtypes.DefaultGenesis()
	claimGenState.Airdrops = genParams.ClaimGenesisState.Airdrops
	claimGenState.ClaimRecords = genParams.ClaimGenesisState.ClaimRecords
	claimGenStateBz := cdc.MustMarshalJSON(claimGenState)
	appState[claimtypes.ModuleName] = claimGenStateBz

	return appState, genDoc, nil
}

type GenesisParams struct {
	AirdropSupply sdk.Coin

	GenesisTime     time.Time
	ChainId         string
	ConsensusParams *tmproto.ConsensusParams

	StakingParams       stakingtypes.Params
	GovParams           govtypes.Params
	LiquidityParams     liqtypes.Params
	LiquidStakingParams liqstakingtypes.Params

	BankGenesisStates banktypes.GenesisState
	ClaimGenesisState claimtypes.GenesisState
}

func TestnetGenesisParams() *GenesisParams {
	genParams := &GenesisParams{}
	genParams.GenesisTime = time.Now()
	genParams.AirdropSupply = sdk.NewCoin("airdrop", sdk.NewInt(15_000_000_000_000)) // 15 milion

	// Set source address balance and add total supply
	genParams.BankGenesisStates.Balances = []banktypes.Balance{
		{
			Address: "cosmos15rz2rwnlgr7nf6eauz52usezffwrxc0mz4pywr",
			Coins:   sdk.NewCoins(genParams.AirdropSupply),
		},
	}
	genParams.BankGenesisStates.Supply = sdk.NewCoins(genParams.AirdropSupply)

	// Set airdrop
	genParams.ClaimGenesisState.Airdrops = []types.Airdrop{
		{
			Id:            1,
			SourceAddress: "cosmos15rz2rwnlgr7nf6eauz52usezffwrxc0mz4pywr",
			Conditions: []claimtypes.ConditionType{
				claimtypes.ConditionTypeDeposit,
				claimtypes.ConditionTypeSwap,
				claimtypes.ConditionTypeFarming,
			},
			StartTime: genParams.GenesisTime,
			EndTime:   genParams.GenesisTime.AddDate(0, 1, 0),
		},
	}

	// Set claim records
	genParams.ClaimGenesisState.ClaimRecords = []claimtypes.ClaimRecord{
		{
			AirdropId:             1,
			Recipient:             "cosmos1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v", // validator
			InitialClaimableCoins: sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(3_000_000_000_000))),
			ClaimableCoins:        sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(3_000_000_000_000))),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             1,
			Recipient:             "cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu", // user1
			InitialClaimableCoins: sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(9_000_000_000_000))),
			ClaimableCoins:        sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(9_000_000_000_000))),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             1,
			Recipient:             "cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny", // user2
			InitialClaimableCoins: sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(3_000_000_000_000))),
			ClaimableCoins:        sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, sdk.NewInt(3_000_000_000_000))),
			ClaimedConditions:     []types.ConditionType{},
		},
	}

	return genParams
}

func MainnetGenesisParams() *GenesisParams {
	genParams := &GenesisParams{}
	genParams.GenesisTime = time.Now()
	genParams.AirdropSupply = sdk.NewCoin("usquad", sdk.NewInt(15_000_000_000_000)) // 15 milion

	// TODO: setup genesis values for mainnet

	// Set source address balance
	genParams.BankGenesisStates.Balances = []banktypes.Balance{
		{
			Address: "cosmos15rz2rwnlgr7nf6eauz52usezffwrxc0mz4pywr",
			Coins:   sdk.NewCoins(genParams.AirdropSupply),
		},
	}
	genParams.BankGenesisStates.Supply = sdk.NewCoins(genParams.AirdropSupply)

	// Set airdrop
	genParams.ClaimGenesisState.Airdrops = []types.Airdrop{
		{
			Id:            1,
			SourceAddress: "cosmos15rz2rwnlgr7nf6eauz52usezffwrxc0mz4pywr",
			StartTime:     genParams.GenesisTime,
			EndTime:       genParams.GenesisTime.AddDate(0, 1, 0),
		},
	}

	filePath := ""
	results, err := readCSVFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("failed to read %s", filePath))
	}

	records := []claimtypes.ClaimRecord{}
	for _, r := range results {
		recipient := r[0]
		claimableAmtStr := r[1]
		claimableAmt, _ := sdk.NewIntFromString(claimableAmtStr)

		records = append(records, types.ClaimRecord{
			AirdropId:             1,
			Recipient:             recipient,
			InitialClaimableCoins: sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, claimableAmt)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewCoin(genParams.AirdropSupply.Denom, claimableAmt)),
		})
	}
	genParams.ClaimGenesisState.ClaimRecords = records

	return genParams
}

// parseGenesisParams returns GenesisParams based on the network type.
func parseGenesisParams(networkType string) *GenesisParams {
	switch strings.ToLower(networkType) {
	case "t", "testnet":
		return TestnetGenesisParams()
	case "m", "mainnet":
		return MainnetGenesisParams()
	default:
		return nil
	}
}

// readCSVFile reads csv file and returns all the records.
func readCSVFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}