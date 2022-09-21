# usage case
# ./gcs.sh 'd' 'update shell code' 'dev'

DATE_FORMAT=`date "+%Y-%m-%d %H:%M:%S"`
logInfo() {
 local infoMsg="$*"
 echo  "\033[32m [Info] $DATE_FORMAT $infoMsg\033[0m"
}

logWarn() {
 local infoMsg="$*"
 echo  "\033[31m [Warn]  $DATE_FORMAT $infoMsg\033[0m"
}

logError() {
 local infoMsg="$*"
 echo  "\033[31m [Error]  $DATE_FORMAT $infoMsg\033[0m"
 exit 1
}

if [ ! $1 ]
then
  selectProject=" 请选择项目!!! "
  logWarn "$selectProject"
  exit;
fi
 
if [ ! $2 ]
then
  inputCommitValue=" 请输入commit值!!! "
  logWarn "$inputCommitValue"
  exit;
fi
 
#给出一个默认的项目路径
path="/Users/kubrick/go/src/kubrick/july"
 
case "$1" in
  "d")
  logInfo "moppo-ocr"
  path="/Users/kubrick/go/src/kubrick/july"
  ;;
esac
 
#先进入项目当中
cd $path
 
ls
 
logInfo "开始执行命令"

git add .

git status
 
#写个sleep 1s 是为了解决并发导致卡壳
sleep 1s
 
logInfo " 添加文件!"
 
git commit -m "$2"
 
logInfo " 提交完成!"
 
sleep 1s
 
logInfo " 开始推送!"
 
if [ ! $3 ]
then
  logInfo " 请输入自己提交代码的分支!"
  exit;
fi
 
git push origin "$3"
 
logInfo " 推送成功!"


