package utils

import (
	model "ABTest/apps/Dao/experiment/model"
	pb "ABTest/pkgs/proto/pb_experiment"
)

/*
	message ExperimentInfo {
		uint64 experiment_id = 1;
		string name = 2;
		string description = 3;
		repeated ExperimentGroup groups = 4;
		ExperimentStatus status = 5;

		google.protobuf.Timestamp create_time = 50;
		google.protobuf.Timestamp update_time = 51;
	}
*/
/*
	message ExperimentGroup {
	  uint64 group_id = 1;
	  string name = 2;
	  string description = 3;
	  float allocation = 4; // 实验组占实验比例
	  uint64 from_experiment_id = 5; // 实验组来源实验ID
	  uint64 whitelist_userpackageID = 6; // 白名单
	}
*/
func TranslateExperimentModelToProtoExperimentInfo(experimentInfo model.ExperimentInfoBasic, experimentGroup []model.ExperimentGroupBasic) *pb.ExperimentInfo {
	var groups []*pb.ExperimentGroup
	for _, group := range experimentGroup {
		groups = append(groups, &pb.ExperimentGroup{
			GroupId:                group.GroupID,
			Name:                   group.Name,
			Description:            group.Description,
			Allocation:             group.Allocation,
			FromExperimentId:       group.FromExperimentID,
			WhitelistUserpackageID: group.WhiteListUserPackageId,
		})
	}
	return &pb.ExperimentInfo{
		ExperimentId: experimentInfo.ExperimentID,
		Name:         experimentInfo.Name,
		Description:  experimentInfo.Description,
		Groups:       groups,
		CreateTime:   experimentInfo.CreatedAt.String(),
		UpdateTime:   experimentInfo.UpdatedAt.String(),
	}
}

func TranslateProtoExperimentInfoToExperimentModel(experimentInfo *pb.ExperimentInfo) (model.ExperimentInfoBasic, []model.ExperimentGroupBasic) {
	var groups []model.ExperimentGroupBasic
	for _, group := range experimentInfo.Groups {
		groups = append(groups, model.ExperimentGroupBasic{
			GroupID:                group.GroupId,
			Name:                   group.Name,
			Description:            group.Description,
			Allocation:             group.Allocation,
			FromExperimentID:       group.FromExperimentId,
			WhiteListUserPackageId: group.WhitelistUserpackageID,
		})
	}
	return model.ExperimentInfoBasic{
		ExperimentID: experimentInfo.ExperimentId,
		Name:         experimentInfo.Name,
		Description:  experimentInfo.Description,
		Status:       int(experimentInfo.Status),
	}, groups
}

// group
func TranslateExperimentGroupModelToProtoExperimentGroup(group model.ExperimentGroupBasic) *pb.ExperimentGroup {
	return &pb.ExperimentGroup{
		GroupId:                group.GroupID,
		Name:                   group.Name,
		Description:            group.Description,
		Allocation:             group.Allocation,
		FromExperimentId:       group.FromExperimentID,
		WhitelistUserpackageID: group.WhiteListUserPackageId,
	}
}

func TranslateProtoExperimentGroupToExperimentModel(group *pb.ExperimentGroup) model.ExperimentGroupBasic {
	return model.ExperimentGroupBasic{
		GroupID:                group.GroupId,
		Name:                   group.Name,
		Description:            group.Description,
		Allocation:             group.Allocation,
		FromExperimentID:       group.FromExperimentId,
		WhiteListUserPackageId: group.WhitelistUserpackageID,
	}
}
