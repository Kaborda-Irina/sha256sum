# sha256sum on Golang

###Getting the hash sum of a file or all files in a directory in Golang:
_____
## Usage

### 1) if you need the hash sum of a single file sha256, use the flag -d (-d then directory path)
go run cmd/feature_third/main.go -d ../name_dir/example.txt

### 2) if you need the hash sum of all files in a specific directory sha256, use the flag -d (-d then directory path)
go run cmd/feature_third/main.go -d ../name_dir

### 3) if you need the hash sum  of multiple directories sha256, use the flag -d (-d then directory path, directory path)
go run cmd/feature_third/main.go -d ../name_dir ../dir/dir_name

### 4) if you need to set algorithm other than sha256,use the flag -a (-a md5, 1, 224, 256, 384, 512)
go run cmd/feature_third/main.go -a md5 -d ../name_dir/example.txt
go run cmd/feature_third/main.go -a md5 -d ../name_dir
