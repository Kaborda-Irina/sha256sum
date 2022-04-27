# sha256sum - cli-app hash sum on Golang
sha256sum is a console application for getting the hash sum of files with different algorithms( **MD5, SHA256, SHA1, SHA224, SHA384, SHA512**).

+ we used standard libraries "crypto/sha1", "crypto/sha256","crypto/sha512"
+ you can see https://pkg.go.dev/crypto

## :hammer: Installation
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
### Clone repository and install dependencies
```
$ git clone https://github.com/Kaborda-Irina/sha256sum.git
$ cd path/to/install
```
### How to build app
1. You need to go to file ".env" and set "VOLUME_PATH" - this is the relative path to the folder you want to check,
for example:
```
VOLUME_PATH=/home/dir/test
VOLUME_PATH=C:\ProgramData\
```
2. Then you need to build a docker image and database container
```
 docker-compose build
 docker-compose up -d postgres
 ```

### You can use the app:
+ **in container**
```
 docker-compose run web -d /local_path
 or
 docker-compose run web -d=/local_path
```
Where `/local_path = VOLUME_PATH = /home/dir/test` you get the hash sum  of all files in the directory /home/dir/test , but you can set a more detailed path,
for example:
```
 docker-compose run web -d /local_path/dir/test.txt
 or  
 docker-compose run web -d=/local_path/dir/test.txt
```
And you can use flags (-d, -h, -c, -a)

```
 docker-compose run web -d /local_path/example/test.txt -a SHA512
 docker-compose run web -d=/local_path/example/test.txt -a=SHA512
 or
 docker-compose run web -с /local_path/example/test.txt
 docker-compose run web -с=/local_path/example/test.txt
 or
 docker-compose run web -h
 docker-compose run web -h
```

+ **locally in console**
You need to set in file `config.yaml` "localhost" in place of "postgres"
```go
#Database credentials
postgres:
  uri: "postgres://user:1234@localhost:5432/hashdb?sslmode=disable"
```
Then start in console
```
go run cmd/main.go -d ../dir
or
go run cmd/main.go -d ../dir/file_name.txt -a MD5
```
____
## :house: Usage

:one: if you need the hash sum of a single file or all files in a specific directory, use the **`flag -d`** (-d then directory path)
```
go run cmd/main.go -d ../name_dir/example.txt
go run cmd/main.go -d ../name_dir
or 
docker-compose run web -d /local_path/example/test.txt
docker-compose run web -d /local_path
```

:two: if you need to set algorithm other than SHA256,use the **`flag -a`** (-a MD5, SHA1, SHA224, SHA384, SHA512)
```
go run cmd/main.go -a MD5 -d ../name_dir/example.txt 
go run cmd/main.go -a SHA512 -d ../name_dir

go run cmd/main.go -d ../name_dir/example.txt -a MD5
go run cmd/main.go -d ../name_dir -a SHA512

or 

docker-compose run web -a SHA512 -d /local_path/example/test.txt
docker-compose run web -d /local_path -a SHA512
```

:three: if you need to check for changes in directory or file use **`flag -c`** (-c then directory path)
```
go run cmd/main.go -c ../name_dir/example.txt 
go run cmd/main.go -c ../name_dir

or

docker-compose run web -a SHA512 m-d /local_path/example/test.txt
docker-compose run web -d /local_path -a SHA512

```
:four: if you need to check for changes in directory or file with a specific algorithm, use the **`flag -c and flag -a`** (-c then directory path -a algorithm)
```
go run cmd/main.go -c ../name_dir/example.txt -a MD5
go run cmd/main.go -a SHA512 -c ../name_dir -a SHA512
```
___________________________
### :notebook_with_decorative_cover: Godoc extracts and generates documentation for Go programs
#### Presents the documentation as a web page.
```go
godoc -http=:6060/sha256sum
go doc packge.function_name
```
for example
```go
go doc pkg/api.Result
```

### :mag: Running tests

You need to go to the folder where the file is located *_test.go and run the following command:
```go
go test -v
```

for example
```go
cd ../pkg/api
go test -v
```