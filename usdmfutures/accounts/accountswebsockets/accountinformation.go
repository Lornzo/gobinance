package accountswebsockets

type AccountInformation struct {
	FeeTier      int64  `json:"feeTier"`
	CanTrade     bool   `json:"canTrade"`
	CanDeposit   bool   `json:"canDeposit"`
	CanWithdraw  bool   `json:"canWithdraw"`
	AccountAlias string `json:"accountAlias"`
}
