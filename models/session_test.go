package models

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/caser789/jw-session/agents"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSession(t *testing.T) {
	Convey("test session", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := NewMockSessionStorage(ctrl)
		a := agents.NewMockAccountAgent(ctrl)
		content := SessionContent{
			UserID:         1,
			LastVerifyTime: time.Now().Add(-time.Hour).Unix(),
		}
		data, _ := json.Marshal(content)
		Convey("session exist, return the same session id", func() {
			m.EXPECT().Get("session-id1").Return(string(data))
			s := NewSession(m, a)
			sessionID, userID, err := s.CheckSessionStatus("session-id1", "token")
			So(sessionID, ShouldEqual, "session-id1")
			So(userID, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("session not exist, return the new session id", func() {
			m.EXPECT().Get("session-id1").Return("")
			m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(1)
			m.EXPECT().SAdd(gomock.Any(), gomock.Any()).Return(1)
			m.EXPECT().Get(gomock.Not("session-id1")).Return(string(data))
			a.EXPECT().VerifyToken(gomock.Any()).Return(uint64(1), nil)
			s := NewSession(m, a)
			sessionID, userID, err := s.CheckSessionStatus("session-id1", "token")
			So(sessionID, ShouldNotEqual, "session-id1")
			So(len(sessionID), ShouldEqual, 32)
			So(userID, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("delete session by userid", func() {
			m.EXPECT().Del("1").Return(1)
			s := NewSession(m, a)
			err := s.DeleteSessionsBy(1)
			So(err, ShouldBeNil)
		})
	})
}
