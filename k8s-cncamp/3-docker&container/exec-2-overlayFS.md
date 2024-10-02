### Create test lab

```shell
mkdir testlab && cd $_

mkdir upper lower merged work

echo "from lower" > lower/in_lower.txt
echo "from upper" > upper/in_upper.txt
echo "from lower" > lower/in_both.txt
echo "from upper" > upper/in_both.txt
```

### Check it

```log
root@docker:/home/pp/testlab# tree .
.
├── lower
│   ├── in_both.txt
│   └── in_lower.txt
├── merged
├── upper
│   ├── in_both.txt
│   └── in_upper.txt
└── work

4 directories, 4 files
```

### Combine

```shell
sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
```

### Check the merged file

```shell
cat merged/in_both.txt

cat merged/in_lower.txt

cat merged/in_upper.txt
```

```shell
df -h | grep testlab

overlay                             29G  9.7G   18G  36% /root/testlab/merged
```

### modify
```shell
echo "modifyupper" > merged/in_upper.txt
echo "modifylower" > merged/in_lower.txt
echo "modifyboth" > merged/in_both.txt
```
```shell
cat upper/in_both.txt
cat upper/in_upper.txt
cat lower/in_both.txt
cat lower/in_lower.txt
```

```log
root@docker:/home/pp/testlab# cat upper/in_both.txt
modifyboth
root@docker:/home/pp/testlab# cat upper/in_upper.txt
modifyupper
root@docker:/home/pp/testlab# cat lower/in_both.txt
from lower
root@docker:/home/pp/testlab# cat lower/in_lower.txt
from lower
```

### Delete merged file

```shell
sudo umount `pwd`/merged
```
