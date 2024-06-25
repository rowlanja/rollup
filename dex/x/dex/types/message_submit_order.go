package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitOrder{}

func NewMsgSubmitOrder(creator string, amount string, token string) *MsgSubmitOrder {
	return &MsgSubmitOrder{
		Creator: creator,
		Amount:  amount,
		Token:   token,
	}
}

func (msg *MsgSubmitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
