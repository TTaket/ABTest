package utils

import (
	model "ABTest/apps/Dao/dynamic/model"
	pb "ABTest/pkgs/proto/pb_dynamic"
	"time"
)

// 转换 LayerUserBinding 模型至 proto 消息
func TranslateLayerUserBindingModelToProto(binding model.LayerUserBinding) *pb.LayerUserBinding {
	return &pb.LayerUserBinding{
		LayerId:       binding.LayerID,
		UserPackageId: binding.UserPackageID,
		CreatedAt:     binding.CreatedAt.Unix(),
		UpdatedAt:     binding.UpdatedAt.Unix(),
	}
}

// 转换 proto 消息至 LayerUserBinding 模型
func TranslateProtoToLayerUserBindingModel(binding *pb.LayerUserBinding) model.LayerUserBinding {
	return model.LayerUserBinding{
		LayerID:       binding.LayerId,
		UserPackageID: binding.UserPackageId,
		CreatedAt:     time.Unix(binding.CreatedAt, 0),
		UpdatedAt:     time.Unix(binding.UpdatedAt, 0),
	}
}

// 转换 ExperimentUserBinding 模型至 proto 消息
func TranslateExperimentUserBindingModelToProto(binding model.ExperimentUserBinding) *pb.ExperimentUserBinding {
	return &pb.ExperimentUserBinding{
		ExperimentGroupId: binding.ExperimentGroupID,
		UserPackageId:     binding.UserPackageID,
		CreatedAt:         binding.CreatedAt.Unix(),
		UpdatedAt:         binding.UpdatedAt.Unix(),
	}
}

// 转换 proto 消息至 ExperimentUserBinding 模型
func TranslateProtoToExperimentUserBindingModel(binding *pb.ExperimentUserBinding) model.ExperimentUserBinding {
	return model.ExperimentUserBinding{
		ExperimentGroupID: binding.ExperimentGroupId,
		UserPackageID:     binding.UserPackageId,
		CreatedAt:         time.Unix(binding.CreatedAt, 0),
		UpdatedAt:         time.Unix(binding.UpdatedAt, 0),
	}
}

// 转��� ExperimentLayerBinding 模型至 proto 消息
func TranslateExperimentLayerBindingModelToProto(binding model.ExperimentLayerBinding) *pb.ExperimentLayerBinding {
	return &pb.ExperimentLayerBinding{
		ExperimentId: binding.ExperimentID,
		LayerId:      binding.LayerID,
		Ratio:        binding.Ratio,
		CreatedAt:    binding.CreatedAt.Unix(),
		UpdatedAt:    binding.UpdatedAt.Unix(),
	}
}

// 转换 proto 消息至 ExperimentLayerBinding 模型
func TranslateProtoToExperimentLayerBindingModel(binding *pb.ExperimentLayerBinding) model.ExperimentLayerBinding {
	return model.ExperimentLayerBinding{
		ExperimentID: binding.ExperimentId,
		LayerID:      binding.LayerId,
		Ratio:        binding.Ratio,
		CreatedAt:    time.Unix(binding.CreatedAt, 0),
		UpdatedAt:    time.Unix(binding.UpdatedAt, 0),
	}
}
