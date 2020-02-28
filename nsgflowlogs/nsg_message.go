package nsgflowlogs

import "time"

type NsgMessage struct {
	Time          time.Time            `json:"time"`
	SystemID      string               `json:"systemId"`
	MacAddress    string               `json:"macAddress"`
	Category      string               `json:"category"`
	ResourceID    string               `json:"resourceId"`
	OperationName string               `json:"operationName"`
	Properties    NsgMessageProperties `json:"properties"`
}

type NsgMessageProperties struct {
	Version int              `json:"Version"`
	Flows   []NsgMessageFlow `json:"flows"`
}

type NsgMessageFlow struct {
	Rule  string        `json:"rule"`
	Flows []NsgRuleFlow `json:"flows"`
}

type NsgRuleFlow struct {
	Mac        string   `json:"mac"`
	FlowTuples []string `json:"flowTuples"`
}
