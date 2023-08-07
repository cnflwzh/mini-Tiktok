package jwt

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var token string

func TestGenerateAndParseToken(t *testing.T) {
	username := "testUser"
	userID := int64(456)

	// 生成Token
	token, err := GenerateToken(username, userID)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
		return
	}

	// 解析Token
	parsedUserID, err := ParseToken(token)
	hlog.Info("parsedUserID", parsedUserID)
	if err != nil {
		t.Errorf("Error parsing token: %v", err)
		return
	}

	// 验证生成的Token是否正确
	if parsedUserID != userID {
		t.Errorf("Expected userID %d, but got %d", userID, parsedUserID)
		return
	}
}

func TestParseInvalidToken(t *testing.T) {
	invalidToken := token + "123"

	// 尝试解析无效的Token
	_, err := ParseToken(invalidToken)
	if err == nil {
		t.Error("Expected error when parsing invalid token, but got nil")
		return
	}
}
