#!/bin/bash

export GOPROXY=https://goproxy.cn # 设置代理
go mod tidy # 下载依赖
go build -o twitter-space # 打包成可执行文件
./twitter-space # 运行

name="世界和平"

readonly name

read age

echo "hi $name, your age is $age."
echo "hi ${name}, your age is ${age}."

read level

case $level in
  1)
    echo "1"
  ;;
  2)
    echo "2"
  ;;
  3)
    echo "2"
  ;;
esac

if [[ $level == 1 ]]; then
  echo "win"
elif [[ $level == 2 ]]; then
  echo "failed"
fi

unset level # 删除变量

echo ${#name} # 字符串长度


a=1
b=2
if (( a > b));
then
  echo "a>b"
elif (( a== b ))
then
  echo "a=b"
else
  echo "a<b"
fi

if [ $a -ge $b ];
then
  echo "hi"
fi

function foo() {
    echo "这是我的第一个 shell 函数!"
}
foo # 执行函数


function bar() {
    echo "这是我的第二个 shell 函数!"
    echo "第一个参数: $1"
    echo "第二个参数: $2"
    echo "第三个参数: $3"

    echo "第Ⅹ个参数: $10"
    echo "第Ⅹ个参数: ${10}"
}
bar 1 2 3 # 执行函数 - 携带参数


pwd1=$(pwd)
echo pwd1
pwd2=`pwd`
echo pwd2