# ts
timestamp input

### install

```shell
go get github.com/ma6174/ts
```

### test

```shell
$ cat ./example/test.sh
#!/bin/bash
echo "start"
sleep 1.0
echo "after sleep 1s"
sleep 0.5
echo "done"
$ ./example/test.sh | ts
2015/02/14-17:01:27.046446 start
2015/02/14-17:01:28.049754 after sleep 1s
2015/02/14-17:01:28.556886 done
$ ./example/test.sh | ts -i
0.000021s start
1.006260s after sleep 1s
0.507335s done
$ ./example/test.sh | ts -s
0.000022s start
1.005838s after sleep 1s
1.508020s done
$
```
