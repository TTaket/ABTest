package handlers

import (
	"net/http"

	"ABTest/apps/Micro/dynamic/client"
	pb "ABTest/pkgs/proto/pb_dynamic"

	"github.com/gin-gonic/gin"
)

// DynamicHandler 处理 dynamic 相关的请求
type DynamicHandler struct {
	Client client.DynamicClient
}

// NewDynamicHandler 创建一个新的 DynamicHandler 实例
func NewDynamicHandler() *DynamicHandler {
	return &DynamicHandler{
		Client: client.NewDynamicClient(),
	}
}

// CreateLayerUserBinding godoc
// @Summary 创建Layer与User绑定关系
// @Description 创建一个Layer与User的绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.CreateLayerUserBindingRequest true "绑定信息"
// @Success 200 {object} pb.CreateLayerUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/layer-user-bindings/create-bind [post]
func (h *DynamicHandler) CreateLayerUserBinding(c *gin.Context) {
	var req pb.CreateLayerUserBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.CreateLayerUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLayerUserBinding godoc
// @Summary 获取Layer与User绑定关系
// @Description 获取特定的Layer与User绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request query pb.GetLayerUserBindingRequest true "查询参数"
// @Success 200 {object} pb.GetLayerUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/layer-user-bindings/get-bind [get]
func (h *DynamicHandler) GetLayerUserBinding(c *gin.Context) {
	var req pb.GetLayerUserBindingRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.GetLayerUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteLayerUserBinding godoc
// @Summary 删除Layer与User绑定关系
// @Description 删除特定的Layer与User绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.DeleteLayerUserBindingRequest true "删除请求"
// @Success 200 {object} pb.DeleteLayerUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/layer-user-bindings/delete-bind [delete]
func (h *DynamicHandler) DeleteLayerUserBinding(c *gin.Context) {
	var req pb.DeleteLayerUserBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.DeleteLayerUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateExperimentUserBinding godoc
// @Summary 创建Experiment与User绑定关系
// @Description 创建一个Experiment与User的绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.CreateExperimentUserBindingRequest true "绑定信息"
// @Success 200 {object} pb.CreateExperimentUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-user-bindings/create-bind [post]
func (h *DynamicHandler) CreateExperimentUserBinding(c *gin.Context) {
	var req pb.CreateExperimentUserBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.CreateExperimentUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetExperimentUserBinding godoc
// @Summary 获取Experiment与User绑定关系
// @Description 获取特定的Experiment与User绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request query pb.GetExperimentUserBindingRequest true "查询参数"
// @Success 200 {object} pb.GetExperimentUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-user-bindings/get-bind  [get]
func (h *DynamicHandler) GetExperimentUserBinding(c *gin.Context) {
	var req pb.GetExperimentUserBindingRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.GetExperimentUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteExperimentUserBinding godoc
// @Summary 删除Experiment与User绑定关系
// @Description 删除特定的Experiment与User绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.DeleteExperimentUserBindingRequest true "删除请求"
// @Success 200 {object} pb.DeleteExperimentUserBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-user-bindings/delete-bind  [delete]
func (h *DynamicHandler) DeleteExperimentUserBinding(c *gin.Context) {
	var req pb.DeleteExperimentUserBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.DeleteExperimentUserBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateExperimentLayerBinding godoc
// @Summary 创建Experiment与Layer绑定关系
// @Description 创建一个Experiment与Layer的绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.CreateExperimentLayerBindingRequest true "绑定信息"
// @Success 200 {object} pb.CreateExperimentLayerBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-layer-bindings/create-bind [post]
func (h *DynamicHandler) CreateExperimentLayerBinding(c *gin.Context) {
	var req pb.CreateExperimentLayerBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.CreateExperimentLayerBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetExperimentLayerBinding godoc
// @Summary 获取Experiment与Layer绑定关系
// @Description 获取特定的Experiment与Layer绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request query pb.GetExperimentLayerBindingRequest true "查询参数"
// @Success 200 {object} pb.GetExperimentLayerBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-layer-bindings/get-bind [get]
func (h *DynamicHandler) GetExperimentLayerBinding(c *gin.Context) {
	var req pb.GetExperimentLayerBindingRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.GetExperimentLayerBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLayerExperiments godoc
// @Summary 获取Layer关联的所有实验
// @Description 获取指定Layer关联的所有实验信息
// @Tags 实验
// @Accept json
// @Produce json
// @Param request query pb.GetLayerExperimentsRequest true "查询参数"
// @Success 200 {object} pb.GetLayerExperimentsResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/layer-experiments/get-all-bylayer [get]
func (h *DynamicHandler) GetLayerExperiments(c *gin.Context) {
	var req pb.GetLayerExperimentsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.GetLayerExperiments(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteExperimentLayerBinding godoc
// @Summary 删除Experiment与Layer绑定关系
// @Description 删除特定的Experiment与Layer绑定关系
// @Tags 绑定关系
// @Accept json
// @Produce json
// @Param request body pb.DeleteExperimentLayerBindingRequest true "删除请求"
// @Success 200 {object} pb.DeleteExperimentLayerBindingResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/dynamic/experiment-layer-bindings/delete-bind [delete]
func (h *DynamicHandler) DeleteExperimentLayerBinding(c *gin.Context) {
	var req pb.DeleteExperimentLayerBindingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.DeleteExperimentLayerBinding(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
