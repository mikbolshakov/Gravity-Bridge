package types

import (
	fmt "fmt"
	"strings"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeUnhaltBridge = "UnhaltBridge"
	ProposalTypeAirdrop      = "Airdrop"
)

func (p *UnhaltBridgeProposal) GetTitle() string { return p.Title }

func (p *UnhaltBridgeProposal) GetDescription() string { return p.Description }

func (p *UnhaltBridgeProposal) ProposalRoute() string { return RouterKey }

func (p *UnhaltBridgeProposal) ProposalType() string {
	return ProposalTypeUnhaltBridge
}

func (p *UnhaltBridgeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	return nil
}

func (p UnhaltBridgeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Unhalt Bridge Proposal:
  Title:          %s
  Description:    %s
  target_nonce:   %d
`, p.Title, p.Description, p.TargetNonce))
	return b.String()
}

func (p *AirdropProposal) GetTitle() string { return p.Title }

func (p *AirdropProposal) GetDescription() string { return p.Description }

func (p *AirdropProposal) ProposalRoute() string { return RouterKey }

func (p *AirdropProposal) ProposalType() string {
	return ProposalTypeAirdrop
}

func (p *AirdropProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	return nil
}

func (p AirdropProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Airdrop Proposal:
  Title:          %s
  Description:    %s
  Amount:         %d%s
  Recipients:     %s
`, p.Title, p.Description, p.Amount.Amount.Int64(), p.Amount.Denom, p.Recipients))
	return b.String()
}