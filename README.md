#sha256sum - cli-app hash sum on Golang
sha256sum is a console application for getting the hash sum of files with different algorithms( **MD5, SHA256, SHA1, SHA224, SHA384, SHA512**).

+ we used standard libraries "crypto/sha1", "crypto/sha256","crypto/sha512"
+ you can see https://pkg.go.dev/crypto

##:hammer: Installation
```go
import (
	"context"
	"flag"
	"fmt"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/initialize"
	"os"
	"os/signal"
	"time"
)

var dirPath string
var doHelp bool
var algorithm string
var checkHashSumFile string

//initializes the binding of the flag to a variable that must run before the main() function
func init() {
	flag.StringVar(&dirPath, "d", "", "directory path")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&algorithm, "a", "SHA256", "algorithm MD5, SHA1, SHA224, SHA256, SHA384, SHA512, default: SHA256")
	flag.StringVar(&checkHashSumFile, "c", "", "check hash sum files in directory")
}

func main() {
	flag.Parse()
	
	//Initialize config
	cfg, logger, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Error during loading from config file", err)
	}

	ctx := context.Background()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		signal.Stop(sig)
		cancel()
	}()

	initialize.Initialize(ctx, cfg, logger, sig, doHelp, dirPath, algorithm, checkHashSumFile)
}
```
###Clone repository and install dependencies
```
$ git clone https://github.com/Kaborda-Irina/sha256sum.git
$ cd path/to/install
```
###Сonnect database
```
sudo docker-compose build
sudo docker-compose up -d
```

###Run code in console
```
go run cmd/main.go -d ../dir
or
go run cmd/feature_fifth/main.go -d ../dir/file_name.txt -a MD5
```
____
##:house: Usage

:one: if you need the hash sum of a single file or all files in a specific directory, use the **`flag -d`** (-d then directory path)
```
go run cmd/main.go -d ../name_dir/example.txt
go run cmd/main.go -d ../name_dir
```

:two: if you need to set algorithm other than SHA256,use the **`flag -a`** (-a MD5, SHA1, SHA224, SHA384, SHA512)
```
go run cmd/main.go -a md5 -d ../name_dir/example.txt 
go run cmd/main.go -a SHA512 -d ../name_dir

go run cmd/main.go -d ../name_dir/example.txt -a md5
go run cmd/main.go -d ../name_dir -a SHA512
```

:three: if you need to check for changes in directory or file use **`flag -c`** (-c then directory path)
```
go run cmd/main.go -c ../name_dir/example.txt 
go run cmd/main.go -c ../name_dir
```
:four: if you need to check for changes in directory or file with a specific algorithm, use the **`flag -c and flag -a`** (-c then directory path -a algorithm)
```
go run cmd/main.go -c ../name_dir/example.txt -a MD5
go run cmd/main.go -a SHA512 -c ../name_dir -a SHA512
```
___________________________
### :notebook_with_decorative_cover: Godoc extracts and generates documentation for Go programs
####Presents the documentation as a web page.
```go
godoc -http=:6060/sha256sum
go doc packge.function_name
```
for example
```go
go doc pkg/api.Result
```

###:mag: Running tests

You need to go to the folder where the file is located *_test.go and run the following command:
```go
go test -v
```

for example
```go
cd ../pkg/api
go test -v
```
