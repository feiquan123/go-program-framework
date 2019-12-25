# Linux Script
echo start [git push]
git add .  -A
if [ "$1" = "" ]; then
    git commit -m "debug"
else
    git commit -m "$1"
fi

git push origin master
echo [ git push ] Success!
exit


rem Windows Script
cls
@echo off

echo start [git push]
git add . -A
if not "%1" == "" (
    git commit -m "%1"
) else (
    git commit -m "debug"
)
git push origin master
echo [git push ] Success!