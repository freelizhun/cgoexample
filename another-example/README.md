###对于c与go混合编译的解释（对于大型工程项目中也是如此，形如main.go的代码放入对应的包目录即可）：
第一步：先将openimage.c文件编译成openimage.so文件：
```bash
$ gcc -shared -fPIC -nostartfiles openimage.c -o openimage.so
$ cp openimage.so /usr/lib
```
第二部：则从main.go中引入openimage.so文件，参照main.go
```bash
$ go build main.go
$ ./main
```
输出：
the foyouage need to tagtagtagtaged is busybox:1.28.4
opened with handle 100
