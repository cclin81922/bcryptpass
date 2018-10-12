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

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt ...
func Encrypt(plainPass string) (hashedPass string, e error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)

	if err != nil {
		e = err
		return
	}

	hashedPass = string(hash)
	return
}

// Verify ...
func Verify(hashedPass, plainPass string) (same bool, e error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))

	if err != nil {
		e = err
		return
	}

	same = true
	return
}
