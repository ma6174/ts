# ts
timestamp input

### install

```console
go get -u github.com/ma6174/ts
```

### example

```console
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
```

### something interesting

- show timestamp after every command output

```console
~/tmp$ for i in {1..10};do touch $i;done
~/tmp$ ls
1   10  2   3   4   5   6   7   8   9
~/tmp$ bash | ts
~/tmp$ ls
2015/02/14-22:13:29.200884 1
2015/02/14-22:13:29.201127 10
2015/02/14-22:13:29.201130 2
2015/02/14-22:13:29.201133 3
2015/02/14-22:13:29.201136 4
2015/02/14-22:13:29.201139 5
2015/02/14-22:13:29.201141 6
2015/02/14-22:13:29.201143 7
2015/02/14-22:13:29.201145 8
2015/02/14-22:13:29.201146 9
~/tmp$ ls | grep 1
2015/02/14-22:13:39.579486 1
2015/02/14-22:13:39.579506 10
~/tmp$ exit
```

- show how log you have stay at this terminal

```console
~/tmp$ bash | ts -s
~/tmp$ ls
1.608955s 1
1.609007s 10
1.609013s 2
1.609018s 3
1.609022s 4
1.609026s 5
1.609030s 6
1.609034s 7
1.609038s 8
1.609043s 9
~/tmp$ ls | wc -l
5.184258s       10
~/tmp$ ls | wc -l
6.117924s       10
~/tmp$ ls | wc -l
6.869986s       10
~/tmp$ ls | wc -l
8.334861s       10
~/tmp$ ls | wc -l
9.079557s       10
~/tmp$ ls | wc -l
9.895030s       10
~/tmp$ exit
```
