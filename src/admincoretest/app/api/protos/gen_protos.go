package protos

// Generated by niuhe.idl

type NoneReq struct {
}

type NoneRsp struct {
}

type SystemTestReq struct {
	World string `json:"world" zpf_name:"world" zpf_reqd:"true"` //	hello world
}
