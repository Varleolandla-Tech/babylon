package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// StakingKeeper defines the staking module interface contract needed by the
// epoching module.
type StakingKeeper interface {
	UnbondAllMatureValidators(ctx sdk.Context)
	DequeueAllMatureUBDQueue(ctx sdk.Context, currTime time.Time) (matureUnbonds []stakingtypes.DVPair)
	CompleteUnbonding(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
	DequeueAllMatureRedelegationQueue(ctx sdk.Context, currTime time.Time) (matureRedelegations []stakingtypes.DVVTriplet)
	CompleteRedelegation(ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) (sdk.Coins, error)
	ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error)
}

// Event Hooks
// These can be utilized to communicate between an epoching keeper and another
// keeper which must take particular actions when validators/delegators change
// state. The second keeper must implement this interface, which then the
// epoching keeper can call.

// EpochingHooks event hooks for epoching validator object (noalias)
type EpochingHooks interface {
	AfterEpochBegins(ctx sdk.Context, epoch sdk.Uint) error // Must be called after an epoch begins
	AfterEpochEnds(ctx sdk.Context, epoch sdk.Uint) error   // Must be called after an epoch ends
}

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) // Must be called right before a validator is slashed
}
