package cqueue

import (
	"CraneFrontEnd/generated/protos"
	"CraneFrontEnd/internal/util"
	"os/user"

	log "github.com/sirupsen/logrus"
)

type FilterProcessor interface {
	Process(req *protos.QueryTasksInfoRequest) error
}

// FlagFilterStates
type StatesFilterProcessor struct{}

func (p *StatesFilterProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterStates == "" {
		return nil
	}
	stateList, err := util.ParseInRamTaskStatusList(FlagFilterStates)
	if err != nil {
		log.Errorf("%s", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterStates = stateList
	return nil
}

// FlagSelf
type SelfProcessor struct{}

func (p *SelfProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if !FlagSelf {
		return nil
	}
	cu, err := user.Current()
	if err != nil {
		log.Errorf("Failed to get current username: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterUsers = []string{cu.Username}
	return nil
}

// FlagFilterJobNames
type JobNamesProcessor struct{}

func (p *JobNamesProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterJobNames == "" {
		return nil
	}
	filterJobNameList, err := util.ParseStringParamList(FlagFilterJobNames, ",")
	if err != nil {
		log.Errorf("Invalid job name list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterTaskNames = filterJobNameList
	return nil
}

// FlagFilterUsers
type UserFilterProcessor struct{}

func (p *UserFilterProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterUsers == "" {
		return nil
	}
	filterUserList, err := util.ParseStringParamList(FlagFilterUsers, ",")
	if err != nil {
		log.Errorf("Invalid user list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterUsers = filterUserList
	return nil
}

// FlagFilterQos
type QosProcessor struct{}

func (p *QosProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterQos == "" {
		return nil
	}
	filterJobQosList, err := util.ParseStringParamList(FlagFilterQos, ",")
	if err != nil {
		log.Errorf("Invalid Qos list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterQos = filterJobQosList
	return nil
}

// FlagFilterAccounts
type AccountProcessor struct{}

func (p *AccountProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterAccounts == "" {
		return nil
	}
	filterAccountList, err := util.ParseStringParamList(FlagFilterAccounts, ",")
	if err != nil {
		log.Errorf("Invalid account list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterAccounts = filterAccountList
	return nil
}

// FlagFilterPartitions
type PartitionsProcessor struct{}

func (p *PartitionsProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterPartitions == "" {
		return nil
	}
	filterPartitionList, err := util.ParseStringParamList(FlagFilterPartitions, ",")
	if err != nil {
		log.Errorf("Invalid partition list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterPartitions = filterPartitionList
	return nil
}

// FlagFilterJobIDs
type JobIDsProcessor struct{}

func (p *JobIDsProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterJobIDs == "" {
		return nil
	}
	filterJobIdList, err := util.ParseJobIdList(FlagFilterJobIDs, ",")
	if err != nil {
		log.Errorf("Invalid job list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	idFilter := make(map[uint32]*protos.JobStepIds)
	for _, jobId := range filterJobIdList {
		idFilter[jobId] = nil
	}
	req.FilterIds = idFilter
	if !FlagStep {
		req.NumLimit = uint32(len(filterJobIdList))
	} else {
		req.NumLimit = 0
	}
	return nil
}

type StepIDsProcessor struct{}

func (p *StepIDsProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterStepIDs == "" {
		return nil
	}
	filterStepList, err := util.ParseStepIdList(FlagFilterStepIDs, ",")
	if err != nil {
		log.Errorf("Invalid step list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	if len(filterStepList) == 0 {
		req.FilterIds = filterStepList
	} else {
		if req.FilterIds == nil {
			req.FilterIds = make(map[uint32]*protos.JobStepIds)
		}
		for jobId, steps := range filterStepList {
			if _, exist := req.FilterIds[jobId]; exist {
				steps.Steps = append(steps.Steps, req.FilterIds[jobId].Steps...)
			} else {
				req.FilterIds[jobId] = steps
			}
		}
	}
	req.NumLimit = 0
	return nil
}

// FlagFilterTaskTypes
type TaskTypesProcessor struct{}

func (p *TaskTypesProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterTaskTypes == "" {
		return nil
	}
	filterTaskTypeList, err := util.ParseTaskTypeList(FlagFilterTaskTypes)
	if err != nil {
		log.Errorf("Invalid task type list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterTaskTypes = filterTaskTypeList
	return nil
}

// FlagFilterNodeNames
type NodeNamesProcessor struct{}

func (p *NodeNamesProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterNodeNames == "" {
		return nil
	}
	filterNodenameList, ok := util.ParseHostList(FlagFilterNodeNames)
	if !ok {
		log.Errorf("Invalid node pattern: %s.", FlagFilterNodeNames)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterNodenameList = filterNodenameList
	return nil
}

// FlagFilterLicenses
type LicensesProcessor struct{}

func (p *LicensesProcessor) Process(req *protos.QueryTasksInfoRequest) error {
	if FlagFilterLicenses == "" {
		return nil
	}
	filterLicensesList, err := util.ParseStringParamList(FlagFilterLicenses, ",")
	if err != nil {
		log.Errorf("Invalid licenses list specified: %s.", err)
		return &util.CraneError{Code: util.ErrorCmdArg}
	}
	req.FilterLicenses = filterLicensesList
	return nil
}
