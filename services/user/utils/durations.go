package utils

import "time"

const (
	NormalActionDuration time.Duration = 3 * time.Hour
	RefreshTokenDuration time.Duration = 24 * 7 * time.Hour
	AdminLockDuration    time.Duration = time.Minute * 15
	StaffLockDuration    time.Duration = time.Minute * 30
)
