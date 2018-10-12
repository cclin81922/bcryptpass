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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cclin81922/bcryptpass/pkg/bcryptpass"
)

func main() {
	switch {
	case len(os.Args) == 3 && os.Args[1] == "encrypt":
		pass := os.Args[2]
		hashedPass, err := bcryptpass.Encrypt(pass)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(hashedPass)

	case len(os.Args) == 4 && os.Args[1] == "verify":
		hashedPass := os.Args[2]
		plainPass := os.Args[3]
		same, err := bcryptpass.Verify(hashedPass, plainPass)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(same)

	default:
		usage := `` +
			`Usage` + "\n" +
			"\n" +
			"\t" + `bcryptpass encrypt <plain-pass>` + "\n" +
			"\t" + `bcryptpass verify <hashed-pass> <plain-pass>` + "\n" +
			"\n"
		fmt.Printf("%s", usage)
		os.Exit(1)
	}
}
