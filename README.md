### go中引用c示例


Source files for https://www.thegoldfish.org/2019/04/using-c-libraries-from-go/

### 其中有一段编译的时候答主写错了，应该改为：
```bash
$ gcc hello.c -L. -lperson -o hello
然后将libperson.so文件复制到/usr/lib目录下，
在当前目录./hello则输出成功：
Hello C world: My name is tim, tim hughes.
```
### another-example可以参考：
https://rosettacode.org/wiki/Call_a_function_in_a_shared_library#Go

### another-example-2可以参考：
https://karthikkaranth.me/blog/calling-c-code-from-go/
