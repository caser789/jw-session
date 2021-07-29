package session

import (
	"github.com/caser789/jw-session/agents"
	"github.com/caser789/jw-session/constants"
)

type Session interface {
	CheckSessionStatus(sessionID, token string) (newSessionID string, err error)
	DeleteSessionBy(userID uint64) error
	GenerateNewSessionID(userID uint64) error
}

type SessionContent struct {
	UserID         uint64 `json:"user_id"`
	LastVerifyTime int64  `json:"last_verify_time"`
}

type session struct {
	storage SessionStorage
	agent   agents.AccountAgent
}

func NewSession(storage SessionStorage, agent agents.AccountAgent) *session {
	return &session{
		storage: storage,
		agent:   agent,
	}
}

// 1. get {user_id, v_time} from storage by key
// 2. verify using token if not exists
func (s *session) CheckSessionStatus(sessionID, token string) (newSessionID string, userID uint64, err error) {
	var (
		sessionContentStr string
		sessionContent    SessionContent
	)
	sessionContentStr = s.storage.Get(sessionID)
	if sessionContentStr == "" {
		// veriy token
		userID, err = s.agent.VerifyToken(token)
		if err != nil {
			return "", 0, fmt.Errorf("can not verify token")
		}
		sessionID, err = s.GenerateNewSessionID(userID)
		if err != nil {
			return "", 0, fmt.Errorf("can not generate session")
		}
		// Get session again
		sessionContentStr = s.storage.Get(sessionID)
	}
	json.Unmarshal([]byte(sessionContentStr), &sessionContent)
	// TODO: update sessionID when last verify time too long ago
	return sessionID, sessionContent.UserID, nil
}

// remove user id from set {user_id1, user_id2}
func (s *session) DeleteSessionsBy(userID uint64) error {
	if s.storage.Del(strconv.Itoa(int(userID))) != 1 {
		return fmt.Errorf("logout error")
	}
	return nil
}

// 1. generate session string key
// 2. save {key: '{user_id, v_time}'}
// 3. add user_id to set {user_id1, user_id2}
func (s *session) GenerateNewSessionID(userID uint64) (string, error) {
	key := generateRandomString(32, "")
	sessionContent := SessionContent{
		UserID:         userID,
		LastVerifyTime: time.Now().Unix(),
	}
	data, err := json.Marshal(sessionContent)
	if err != nil {
		return "", fmt.Errorf("fail to marshal session content")
	}
	if s.storage.Set(key, string(data)) != 1 {
		return "", fmt.Errorf("fail to generate new session")
	}
	if s.storage.SAdd(strconv.Itoa(int(userID)), key) != 1 {
		return "", fmt.Errorf("fail to add into the user's session set")
	}

	return key, nil
}

// generateRandomString with alphabet and length, by default used StringNumbersAlpha
func generateRandomString(length uint32, alphabet string) string {
	var buffer bytes.Buffer

	if alphabet == "" {
		alphabet = constants.StringNumbersAlpha
	}

	n := len(alphabet)
	for i := uint32(0); i < length; i++ {

		_, _ = buffer.WriteString(string(alphabet[generateRandomNumber(int64(n))]))
	}

	return buffer.String()
}

// generateRandomNumber -- To use the cypto rand first. If an error is returned, use the math rand
func generateRandomNumber(maxNum int64) int64 {
	randomNum, err := crand.Int(crand.Reader, big.NewInt(maxNum))
	if err != nil {
		// the int(maxNum) may reduce the space of random number
		return int64(rand.Intn(int(maxNum)))
	}
	return randomNum.Int64()
}
