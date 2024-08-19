package constants

import "fmt"

const (
	HFollowed = "HFollowed"
)

const (
	FollowedKey    = "followed:%d"
	MovingListKey  = "moving:list:%d"
	MovingViewsKey = "moving:views:%d"
)

// GetFollowedKey 获取被关注者key
func GetFollowedKey(uid uint) string {
	return fmt.Sprintf(FollowedKey, uid)
}

// GetMovingListKey 获取动态列表key
func GetMovingListKey(uid uint) string {
	return fmt.Sprintf(MovingListKey, uid)
}

// GetMovingViewsKey 获取收件动态列表key
func GetMovingViewsKey(uid uint) string {
	return fmt.Sprintf(MovingViewsKey, uid)
}
