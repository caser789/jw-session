package agents

//go:generate mockgen -destination=mock_account_agent.go -package=agents github.com/caser789/jw-session/agents AccountAgent
type AccountAgent interface {
	VerifyToken(token string) (uint64, error)
}
