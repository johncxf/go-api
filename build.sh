echo "====== build ======"
# 文件
file_path="./bin/go-api-"$(date +"%Y%m%d")

# 判断文件是否存在
if [ -e "$file_path" ]; then
    # 如果文件存在，则删除
    rm "$file_path"
    echo "File deleted: $file_path"
fi

# 编译打包
build_cmd=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $file_path
# 进一步压缩
#build_cmd=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $file_path
eval $build_cmd

echo "file: "$file_path
echo "====== done ======"
