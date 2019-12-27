package app

import (
	common "github.com/srinandan/sample-apps/common"
	items "github.com/srinandan/sample-apps/inventory/items"
)

//Initialize logging, context, sec mgr and kms
func Initialize() {
	//init logging
	common.InitLog()
	//ReadInventoryFile
	if err := items.ReadInventoryFile(); err != nil {
		common.Error.Fatalln("error reading inventory file ", err)
	}
}
