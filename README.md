# sha256sum on Golang

Use sha256sum on Golang you should  enter the next operation:
_____
## if you need sha256sum specific file (-f then file path)
go run cmd/feature_third/main.go -f ../dir/test.txt

## if you need sha256sum multiple file (-f then file path, file path)
go run cmd/feature_third/main.go -f ../dir/test.txt ../dir/example.txt

## if you need sha256sum specific directory (-d then directory path)
go run cmd/feature_third/main.go -d ../name_dir

## if you need sha256sum multiple directory (-d then directory path, directory path)
go run cmd/feature_third/main.go -d ../name_dir ../dir/dir_name