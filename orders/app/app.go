package app

import (
	common "github.com/srinandan/sample-apps/common"
	odr "github.com/srinandan/sample-apps/orders/odr"
)

//Initialize logging, context, sec mgr and kms
func Initialize() {
	//init logging
	common.InitLog()
	//ReadOrdersFile
	if err := odr.ReadOrdersFile(); err != nil {
		common.Error.Fatalln("error reading orders file ", err)
	}
}
