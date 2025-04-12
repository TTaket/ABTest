package handlers

import (
	"context"
	"net/http"
	"strconv"

	client "ABTest/apps/Micro/userb/client"
	pb "ABTest/pkgs/proto/pb_userb"
	"github.com/gin-gonic/gin"
)

// UserbHandler 处理用户相关的HTTP请求
type UserbHandler struct {
	userbClient client.UserbClient
}

// NewUserbHandler 创建一个新的用户处理器
func NewUserbHandler() *UserbHandler {
	return &UserbHandler{
		userbClient: client.NewUserbClient(),
	}
}

// GetUserInfo godoc
// @Summary 获取用户信息
// @Description 获取用户详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body pb.GetUserInfoRequest true "用户查询请求"
// @Success 200 {object} pb.GetUserInfoResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/info [post]
func (h *UserbHandler) GetUserInfo(c *gin.Context) {
	var req pb.GetUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userbClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddUser godoc
// @Summary 添加用户
// @Description 添加新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body pb.AddUserRequest true "用户信息"
// @Success 200 {object} pb.AddUserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [post]
func (h *UserbHandler) AddUser(c *gin.Context) {
	var req pb.AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userbClient.AddUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} pb.DeleteUserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [delete]
func (h *UserbHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resp, err := h.userbClient.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		UserId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary 更新用户信息
// @Description 更新用户详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body pb.UpdateUserRequest true "更新用户请求"
// @Success 200 {object} pb.UpdateUserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [put]
func (h *UserbHandler) UpdateUser(c *gin.Context) {
	var req pb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userbClient.UpdateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}