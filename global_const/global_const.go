package global_const

var (
	TimeFormat = "2024-05-15 09:04:40.224702"
)

type UpdateType string

const (
	UpdateTypeDeposit  UpdateType = "deposit"
	UpdateTypeLock     UpdateType = "lock"
	UpdateTypeWithdraw UpdateType = "withdraw"
	UpdateTypeRelease  UpdateType = "release"
)
