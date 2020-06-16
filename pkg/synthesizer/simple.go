// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package synthesizer

import (
	"fmt"

	"chaos-mesh/matrix/pkg/node/data"
	"chaos-mesh/matrix/pkg/random"
)

func SimpleRecGen(hollow interface{}) interface{} {
	switch hollow.(type) {
	case data.HollowBool:
		b := hollow.(data.HollowBool).Value
		if b != nil {
			var boolValue bool = *b
			return boolValue
		}
		return random.RandChoose([]interface{}{true, false})
	case data.HollowInt:
		return random.RandInt(hollow.(data.HollowInt).RangeStart, hollow.(data.HollowInt).RangeEnd)
	case data.HollowFloat:
		return random.RandFloat(hollow.(data.HollowFloat).RangeStart, hollow.(data.HollowFloat).RangeEnd)
	case data.HollowTime:
		return random.RandTime(hollow.(data.HollowTime).RangeStart, hollow.(data.HollowTime).RangeEnd)
	case data.HollowSize:
		return random.RandSize(hollow.(data.HollowSize).RangeStart, hollow.(data.HollowSize).RangeEnd)
	case data.HollowString:
		hollowString := hollow.(data.HollowString)
		if hollowString.Value != "" {
			return hollowString.Value
		} else {
			return "rand-string"
		}
	case data.HollowMap:
		res := make(map[string]interface{})
		for k, v := range hollow.(data.HollowMap).Map {
			res[k] = SimpleRecGen(v)
		}
		return res
	case data.HollowList:
		var res []interface{}
		for _, v := range hollow.(data.HollowList).List {
			res = append(res, SimpleRecGen(v))
		}
		return res
	case data.HollowChoice:
		return SimpleRecGen(random.RandChoose(hollow.(data.HollowChoice).List))
	default:
		return fmt.Sprintf("%s", hollow)
	}
}
