// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"os"

	common "github.com/srinandan/sample-apps/common"
)

//Initialize logging, context, sec mgr and kms
func Initialize() {
	//init logging
	common.InitLog()
	//init tracing
	if os.Getenv("DISABLE_TRACING") == "" {
		common.Info.Println("Tracing enabled.")
		go common.InitTracing("tracking-client")
	} else {
		common.Info.Println("Tracing disabled.")
	}
}
