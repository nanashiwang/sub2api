package handler

import (
	"github.com/Wei-Shaw/sub2api/internal/pkg/ip"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
)

func buildGatewaySessionContext(c *gin.Context, apiKey *service.APIKey) *service.SessionContext {
	var apiKeyID, groupID int64
	if apiKey != nil {
		apiKeyID = apiKey.ID
		if apiKey.GroupID != nil {
			groupID = *apiKey.GroupID
		}
	}
	return &service.SessionContext{
		ClientIP:          ip.GetClientIP(c),
		UserAgent:         c.GetHeader("User-Agent"),
		APIKeyID:          apiKeyID,
		GroupID:           groupID,
		ExplicitSessionID: service.ExtractGatewaySessionIDFromHeaders(c.Request.Header),
	}
}
