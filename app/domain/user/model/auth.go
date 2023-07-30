package model

import (
	"encoding/hex"
	"time"
)

const (
	_newSessionBinByteLen = 16
)

// Cookie .
type Cookie struct {
	ID      int64  `json:"id"`
	Uid     int64  `json:"Uid"`
	Session []byte `json:"session"`
	CSRF    []byte `json:"csrf"`
	Type    int64  `json:"type"`
	Expires int64  `json:"expires"`
}

// Token .
type Token struct {
	ID      int64  `json:"id"`
	Uid     int64  `json:"Uid"`
	AppID   int64  `json:"appid"`
	Token   []byte `json:"token"`
	Expires int64  `json:"expires"`
	Type    int64  `json:"type"`
}

// Refresh refresh token
type Refresh struct {
	ID      int64  `json:"id"`
	Uid     int64  `json:"uid"`
	AppID   int64  `json:"appid"`
	Refresh []byte `json:"refresh"`
	Token   []byte `json:"token"`
	Expires int64  `json:"expires"`
}

// OldCookie old cookie
type OldCookie struct {
	Uid       int64  `json:"uid"`
	Session   string `json:"session_data"`
	CSRFToken string `json:"csrf_token"`
	Type      int64  `json:"type"`
	Expires   int64  `json:"expire_time"`
}

// OldToken old token
type OldToken struct {
	ID           int64  `json:"id"`
	Uid          int64  `json:"uid"`
	AppID        int64  `json:"appid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AppSubID     int64  `json:"app_subid"`
	CreateAt     int64  `json:"create_at"`
	Expires      int64  `json:"expires"`
	Type         int64  `json:"type"`
	CTime        string `json:"ctime"`
}

// ConvertToOld convert to old cookie
func (c *Cookie) ConvertToOld() *OldCookie {
	return &OldCookie{
		Uid:       c.Uid,
		Session:   encodeSession(c.Session),
		CSRFToken: hex.EncodeToString(c.CSRF),
		Type:      c.Type,
		Expires:   c.Expires,
	}
}

// ConvertToOld convert to old token
func (t *Token) ConvertToOld(refresh []byte, subID int64) *OldToken {
	return &OldToken{
		ID:           t.ID,
		Uid:          t.Uid,
		AppID:        t.AppID,
		AccessToken:  hex.EncodeToString(t.Token),
		RefreshToken: hex.EncodeToString(refresh),
		AppSubID:     subID,
		CreateAt:     time.Now().Unix(),
		Expires:      t.Expires,
		Type:         t.Type,
	}
}

// ConvertToProto convert to proto
func (c *Cookie) ConvertToProto() *CookieProto {
	return &CookieProto{
		Mid:     c.Uid,
		Session: encodeSession(c.Session),
		CSRF:    hex.EncodeToString(c.CSRF),
		Type:    c.Type,
		Expires: c.Expires,
	}
}

// ConvertToProto convert to proto
func (t *Token) ConvertToProto() *TokenProto {
	return &TokenProto{
		Mid:     t.Uid,
		AppID:   t.AppID,
		Token:   hex.EncodeToString(t.Token),
		Expires: t.Expires,
		Type:    t.Type,
	}
}

// ConvertToProto convert to proto
func (r *Refresh) ConvertToProto() *RefreshProto {
	return &RefreshProto{
		Mid:     r.Uid,
		AppID:   r.AppID,
		Refresh: hex.EncodeToString(r.Refresh),
		Token:   hex.EncodeToString(r.Token),
		Expires: r.Expires,
	}
}

// RefreshTokenResp refreshToken response
type RefreshTokenResp struct {
	Uid     int64  `json:"uid"`
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Expires int64  `json:"expires"`
}

func encodeSession(b []byte) (s string) {
	// format new
	if len(b) == _newSessionBinByteLen {
		return hex.EncodeToString(b)
	}
	// or format old
	return string(b)
}
