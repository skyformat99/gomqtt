// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import "time"

type Future interface {
	// Timeout will set a timeout for the future.
	Timeout(time.Duration) Future

	// Wait will wait until the future is completed. The return value may be false
	// if the call times out.
	Wait() bool

	// Wait will call the callback when the future is completed. The first argument
	// may be false id the call times out.
	Callback(func(bool))

	Error() error

	// Returns true if the future is already completed.
	Complete() bool
}
