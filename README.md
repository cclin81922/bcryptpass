# Installation

```
go get -u github.com/cclin81922/bcryptpass/cmd/bcryptpass
export PATH=$PATH:~/go/bin
```

# Command Line Usage

Encrypt a plain password

```
bcryptpass encrypt <plain-pass>
```

Verify a plain password

```
bcryptpass verify <hashed-pass> <plain-pass>
```

# Package Usage

```
import "github.com/cclin81922/bcryptpass/pkg/bcryptpass"

func demo(pass string) {
    hashedPass := bcryptpass.Encrypt(pass)
    same := bcryptpass.Verify(hashedPass, pass)
}
```

# For Developer

Run all tests

```
go test github.com/cclin81922/bcryptpass/pkg/bcryptpass
```
