# qb2tr

qb2tr 是使用Golang语言编写的qBittorrent高版本转换Transmission种子程序，由于qBittorrent4.4版本后，保存的种子文件，都将Tracker从种子中删除，然后保存到了fastresume文件中。

所以导致如果要从qBittorrent的高版本中，把种子转移到Transmission里，会出现转移后果没有Tracker的尴尬情况，尤其是涉及到需要保种的情况，如果涉及到成百上千，那恐怕会很崩溃。

## 使用方法

1. 下载代码

```shell
git clone https://github.com/ylqjgm/qb2tr.git
cd qb2tr
go get .
```

2. 修改配置

使用任意编辑工具修改`main.go`中的`host`、`user`、`pass`三个数值为自己qBittorrent的连接信息。

3. 拷贝种子

将qBittorrent的`BT_backup`目录拷贝到`qb2tr`下，与`main.go`同层级，并在同级新建一个`export`目录。

4. 执行转换

确保qBittorrent当前已启动并可正常访问，在命令行执行：

```shell
go run main.go
```

完成后到`export`目录下找到转换好的种子文件，直接通过Transmission添加即可。
