package handlers

import (
	"context"
	"net/http"
	"strconv"

	client "ABTest/apps/Micro/experiment/client"
	pb "ABTest/pkgs/proto/pb_experiment"

	"github.com/gin-gonic/gin"
)

// ExperimentHandler 处理实验相关的HTTP请求
type ExperimentHandler struct {
	experimentClient client.ExperimentClient
}

// NewExperimentHandler 创建一个新的实验处理器
func NewExperimentHandler() *ExperimentHandler {
	return &ExperimentHandler{
		experimentClient: client.NewExperimentClient(),
	}
}

// GetExperiment godoc
// @Summary 获取实验信息
// @Description 根据实验ID获取实验详细信息
// @Tags 实验
// @Accept json
// @Produce json
// @Param id path int true "实验ID"
// @Success 200 {object} pb.GetExperimentResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/getinfo [get]
func (h *ExperimentHandler) GetExperiment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experiment ID"})
		return
	}

	resp, err := h.experimentClient.GetExperiment(context.Background(), &pb.GetExperimentRequest{
		ExperimentId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateExperiment godoc
// @Summary 创建新实验
// @Description 创建一个新的实验
// @Tags 实验
// @Accept json
// @Produce json
// @Param experiment body pb.CreateExperimentRequest true "实验信息"
// @Success 200 {object} pb.CreateExperimentResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/create-experiment [post]
func (h *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var req pb.CreateExperimentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.experimentClient.CreateExperiment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteExperiment godoc
// @Summary 删除实验
// @Description 根据实验ID删除实验
// @Tags 实验
// @Accept json
// @Produce json
// @Param id path int true "实验ID"
// @Success 200 {object} pb.DeleteExperimentResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/delete-experiment [delete]
func (h *ExperimentHandler) DeleteExperiment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experiment ID"})
		return
	}

	resp, err := h.experimentClient.DeleteExperiment(context.Background(), &pb.DeleteExperimentRequest{
		ExperimentId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateExperimentBaseInfo godoc
// @Summary 更新实验基本信息
// @Description 根据实验ID更新实验基本信息
// @Tags 实验
// @Accept json
// @Produce json
// @Param id path int true "实验ID"
// @Param info body pb.UpdateExperimentBaseInfoRequest true "实验基本信息"
// @Success 200 {object} pb.UpdateExperimentBaseInfoResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/baseinfo/update-info [put]
func (h *ExperimentHandler) UpdateExperimentBaseInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experiment ID"})
		return
	}

	var req pb.UpdateExperimentBaseInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 确保设置了正确的实验ID
	req.ExperimentInfo = &pb.ExperimentInfo{
		ExperimentId: id,
	}

	resp, err := h.experimentClient.UpdateExperimentBaseInfo(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddExperimentGroup godoc
// @Summary 添加实验组
// @Description 向现有实验添加一个新的实验组
// @Tags 实验
// @Accept json
// @Produce json
// @Param id path int true "实验ID"
// @Param group body pb.AddExperimentGroupRequest true "实验组信息"
// @Success 200 {object} pb.AddExperimentGroupResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/groups/create-groups [post]
func (h *ExperimentHandler) AddExperimentGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experiment ID"})
		return
	}

	var req pb.AddExperimentGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 确保设置了正确的实验ID
	req.ExperimentId = id

	resp, err := h.experimentClient.AddExperimentGroup(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteExperimentGroup godoc
// @Summary 删除实验组
// @Description 删除指定实验中的特定实验组
// @Tags 实验
// @Accept json
// @Produce json
// @Param experiment_id path int true "实验ID"
// @Param group_id path int true "实验组ID"
// @Success 200 {object} pb.DeleteExperimentGroupResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/experiments/delete-gropup [delete]
func (h *ExperimentHandler) DeleteExperimentGroup(c *gin.Context) {
	experimentIDStr := c.Param("experiment_id")
	experimentID, err := strconv.ParseUint(experimentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experiment ID"})
		return
	}

	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	resp, err := h.experimentClient.DeleteExperimentGroup(context.Background(), &pb.DeleteExperimentGroupRequest{
		ExperimentId: experimentID,
		GroupId:      groupID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
