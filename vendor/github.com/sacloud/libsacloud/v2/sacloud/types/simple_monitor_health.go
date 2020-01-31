// Copyright 2016-2020 The Libsacloud Authors
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

package types

// ESimpleMonitorHealth シンプル監視ステータス
type ESimpleMonitorHealth string

// SimpleMonitorHealth シンプル監視ステータス
var SimpleMonitorHealth = struct {
	// Up アップ
	Up ESimpleMonitorHealth
	// Down ダウン
	Down ESimpleMonitorHealth
}{
	Up:   ESimpleMonitorHealth("UP"),
	Down: ESimpleMonitorHealth("DOWN"),
}

// IsUp アップ
func (e ESimpleMonitorHealth) IsUp() bool {
	return e == SimpleMonitorHealth.Up
}

// IsDown ダウン
func (e ESimpleMonitorHealth) IsDown() bool {
	return e == SimpleMonitorHealth.Down
}
