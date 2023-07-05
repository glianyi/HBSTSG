| cmd | eg | desc
| ----------- | ----------- | ------------ |
| > | ls > test.txt | 输出重定向覆盖 |
| < | ls < test.txt | 输入重定向覆盖 |
| >> | ls >> test.txt | 以追加的方式重定向 |
| fd >  | 0 > test.txt | fd描述符对应的文件重定向 |
| fd >>  | 0 >> test.txt | fd描述符对应的文件以追加方式重定向 |
| fd >& fd| 0 >& 1| |
| fd <& fd| 0 <& 1| |
| << tag | << EOF | here doc echo << EOF  nihaoshijie  EOF  |
| /dev/null | cat test.txt > /dev/null 2>&1 | | 