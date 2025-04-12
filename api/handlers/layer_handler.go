package handlers

import (
	"context"
	"net/http"
	"strconv"

	client "ABTest/apps/Micro/layer/client"
	pb "ABTest/pkgs/proto/pb_layer"

	"github.com/gin-gonic/gin"
)

// LayerHandler 处理层相关的HTTP请求
type LayerHandler struct {
	layerClient client.LayerClient
}

// NewLayerHandler 创建一个新的层处理器
func NewLayerHandler() *LayerHandler {
	return &LayerHandler{
		layerClient: client.NewLayerClient(),
	}
}

// GetLayer godoc
// @Summary 获取层信息
// @Description 根据层ID获取层详细信息
// @Tags 层
// @Accept json
// @Produce json
// @Param id path string true "层ID"
// @Success 200 {object} pb.GetLayerResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/layers/getinfo [get]
func (h *LayerHandler) GetLayer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Layer ID is required"})
		return
	}

	resp, err := h.layerClient.GetLayer(context.Background(), &pb.GetLayerRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateLayer godoc
// @Summary 创建新层
// @Description 创建一个新的层
// @Tags 层
// @Accept json
// @Produce json
// @Param layer body pb.CreateLayerRequest true "层信息"
// @Success 200 {object} pb.CreateLayerResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/layers/create [post]
func (h *LayerHandler) CreateLayer(c *gin.Context) {
	var req pb.CreateLayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.layerClient.CreateLayer(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateLayer godoc
// @Summary 更新层信息
// @Description 更新层详细信息
// @Tags 层
// @Accept json
// @Produce json
// @Param layer body pb.UpdateLayerRequest true "更新层请求"
// @Success 200 {object} pb.UpdateLayerResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/layers/updateinfo [put]
func (h *LayerHandler) UpdateLayer(c *gin.Context) {
	var req pb.UpdateLayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.layerClient.UpdateLayer(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteLayer godoc
// @Summary 删除层
// @Description 根据层ID删除层
// @Tags 层
// @Accept json
// @Produce json
// @Param id path string true "层ID"
// @Success 200 {object} pb.DeleteLayerResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/layers/delete-layer [delete]
func (h *LayerHandler) DeleteLayer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Layer ID is required"})
		return
	}

	resp, err := h.layerClient.DeleteLayer(context.Background(), &pb.DeleteLayerRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListLayers godoc
// @Summary 列出所有层
// @Description 获取所有层的列表
// @Tags 层
// @Accept json
// @Produce json
// @Param type query int false "层类型"
// @Param parent_id query string false "父层ID"
// @Param page query int false "页码"
// @Param page_size query int false "每页大小"
// @Success 200 {object} pb.ListLayersResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/layers/get-all [get]
func (h *LayerHandler) ListLayers(c *gin.Context) {
	var req pb.ListLayersRequest

	// 解析查询参数
	if typeStr := c.Query("type"); typeStr != "" {
		if typeStr == "0" {
			req.Type = pb.LayerType_LAYER
		} else if typeStr == "1" {
			req.Type = pb.LayerType_DOMAIN
		}
	}

	req.ParentId = c.Query("parent_id")

	// 解析分页参数
	if page := c.Query("page"); page != "" {
		if pageInt, err := strconv.ParseUint("page", 10, 64); err == nil {
			req.Page = int32(pageInt)
		}
	}

	if pageSize := c.Query("page_size"); pageSize != "" {
		if pageSizeInt, err := strconv.ParseUint("page_size", 10, 64); err == nil {
			req.PageSize = int32(pageSizeInt)
		}
	}

	resp, err := h.layerClient.ListLayers(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
