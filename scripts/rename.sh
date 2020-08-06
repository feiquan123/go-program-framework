if [ -z "$1" ]; then
	echo input new program name is empty
else
	oldname=`head -1 go.mod | awk -F ' ' '{print $2}'`
	Oldname=`python ./scripts/convert.py -s $oldname`

	Newname=`python ./scripts/convert.py -s $1`
	echo $Oldname -\> $Newname
	echo $oldname -\> $1
	sed -i "" "s/$Oldname/$Newname/g" `grep "$oldname" -rl .`
	echo Success
fi