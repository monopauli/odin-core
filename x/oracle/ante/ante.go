package ante

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ODIN-PROTOCOL/odin-core/x/oracle/keeper"
	"github.com/ODIN-PROTOCOL/odin-core/x/oracle/types"
)

var (
	repTxCount       *lru.Cache
	nextRepOnlyBlock int64
)

func init() {
	var err error
	repTxCount, err = lru.New(20000)
	if err != nil {
		panic(err)
	}
}

func checkValidReportMsg(ctx sdk.Context, oracleKeeper oraclekeeper.Keeper, rep *types.MsgReportData) bool {
	validator, _ := sdk.ValAddressFromBech32(rep.Validator)
	reporter, _ := sdk.AccAddressFromBech32(rep.Reporter)
	if !oracleKeeper.IsReporter(ctx, validator, reporter) {
		return false
	}
	if rep.RequestID <= oracleKeeper.GetRequestLastExpired(ctx) {
		return false
	}

	req, err := oracleKeeper.GetRequest(ctx, rep.RequestID)
	if err != nil {
		return false
	}

	reqVals := make([]sdk.ValAddress, len(req.RequestedValidators))
	for idx, reqVal := range req.RequestedValidators {
		val, _ := sdk.ValAddressFromBech32(reqVal)
		reqVals[idx] = val
	}

	if !oraclekeeper.ContainsVal(reqVals, validator) {
		return false
	}
	if len(rep.RawReports) != len(req.RawRequests) {
		return false
	}
	for _, report := range rep.RawReports {
		if !oraclekeeper.ContainsEID(req.RawRequests, report.ExternalID) {
			return false
		}
	}
	return true
}

// NewFeelessReportsAnteHandler returns a new ante handler that waives minimum gas price
// requirement if the incoming tx is a valid report transaction.
func NewFeelessReportsAnteHandler(ante sdk.AnteHandler, oracleKeeper oraclekeeper.Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		if ctx.IsCheckTx() && !simulate {
			// TODO: Move this out of "FeelessReports" ante handler.
			isRepOnlyBlock := ctx.BlockHeight() == nextRepOnlyBlock
			isValidReportTx := true
			for _, msg := range tx.GetMsgs() {
				rep, ok := msg.(*types.MsgReportData)
				if !ok || !checkValidReportMsg(ctx, oracleKeeper, rep) {
					isValidReportTx = false
					break
				}
				if !isRepOnlyBlock {
					key := fmt.Sprintf("%s:%d", rep.Validator, rep.RequestID)
					val, ok := repTxCount.Get(key)
					nextVal := 1
					if ok {
						nextVal = val.(int) + 1
					}
					repTxCount.Add(key, nextVal)
					if nextVal > 20 {
						nextRepOnlyBlock = ctx.BlockHeight() + 1
					}
				}
			}
			if isRepOnlyBlock && !isValidReportTx {
				return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Block reserved for report txs")
			}
			if isValidReportTx {
				minGas := ctx.MinGasPrices()
				newCtx, err := ante(ctx.WithMinGasPrices(sdk.DecCoins{}), tx, simulate)
				// Set minimum gas price context and return context to caller.
				return newCtx.WithMinGasPrices(minGas), err
			}
		}
		return ante(ctx, tx, simulate)
	}
}
