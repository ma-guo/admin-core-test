package views

// Generated by niuhe.idl
import (
	"github.com/ma-guo/niuhe"
)

var thisModule *niuhe.Module

func GetModule() *niuhe.Module {
	if thisModule == nil {
		thisModule = niuhe.NewModule("api")
	}
	return thisModule
}
