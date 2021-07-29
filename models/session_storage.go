package models

//go:generate mockgen -destination=mock_session_storage.go -package=models github.com/caser789/jw-session/models SessionStorage
type SessionStorage interface {
	Get(key string) string
	Set(key, val string) int
	Del(key string) int
	SAdd(key, value string) int
	SMembers(key string) []string
}
