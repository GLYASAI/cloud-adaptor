// RAINBOND, Application Management Platform
// Copyright (C) 2020-2021 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package mapper

import (
	"github.com/devfeel/mapper"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

// Mapper mapper and set value from struct fromObj to toObj
// not support auto register struct
func Mapper(fromObj, toObj interface{}) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Warningf("mapper panic: %v", r)
			debug.PrintStack()
		}
	}()
	err := mapper.Mapper(fromObj, toObj)
	if err != nil {
		logrus.Warningf("mapper error: %v", err)
	}
}
