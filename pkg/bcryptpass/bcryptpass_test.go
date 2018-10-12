//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package bcryptpass

import "testing"

func TestEncryptThenVerify(t *testing.T) {
	testcases := []struct {
		name string
		pass string
	}{
		{name: "pass", pass: "pass"},
		{name: "password", pass: "password"},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			hashedPass, err := Encrypt(tc.pass)
			if err != nil {
				t.Fatalf("encryption error | %s", err)
			}

			same, err := Verify(hashedPass, tc.pass)
			if err != nil {
				t.Fatalf("verification error | %s", err)
			}

			if !same {
				t.Fatal("hashed pass is not the same as pass")
			}
		})
	}
}
