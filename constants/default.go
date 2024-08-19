package constants

const (
	defaultFollower = 20000 // 2000粉丝以上为大up
)

// GetDefaultFollower 获取默认大up粉丝数量
func GetDefaultFollower() int {
	return defaultFollower
}
