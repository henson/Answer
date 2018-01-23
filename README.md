[![Go Report Card](https://goreportcard.com/badge/github.com/henson/Answer)](https://goreportcard.com/report/github.com/henson/Answer)  [![Travis Status for henson/Answer](https://travis-ci.org/henson/Answer.svg?branch=master)](https://travis-ci.org/henson/Answer)  [![GitHub release](https://img.shields.io/github/release/henson/Answer.svg)](https://github.com/henson/Answer/releases/tag/v1.0)

# Answer
答题助手，适用于 百万英雄 / 芝士超人 / 冲顶大会 / 花椒百万赢家 等多个直播答题类 APP，支持 iOS、Android 手机和模拟器，3 秒出结果。

![](/doc/1.png)

**目录**

- [答题助手](#Answer)
  - [特点](#特点)
  - [更新日志](#更新日志)
  - [方法原理](#方法原理)
  - [文本关联相似度算法](#文本关联相似度算法)
  - [使用步骤](#使用步骤)
    -[Android](#Android)
    -[iOS](#iOS)
  - [项目参考](#项目参考)
- [TODO](#TODO)

## 特点

- 超快的响应速度
- 支持iOS、Android
- 支持真机测试和模拟器
- 全面覆盖百万英雄/芝士超人/冲顶大会/花椒百万赢家等多个直播答题类APP
- 优化搜索逻辑保证高正确率
- 多维度参考值，答案交叉验证

## 更新日志

- 2018.01.23
  - 把OCR识别方式分离（[issues/1](https://github.com/henson/Answer/issues/1)）
  - fix [issues/3](https://github.com/henson/Answer/issues/3)
- 2018.01.21
  - 加入搜狗汪仔答题助手的结果
- 2018.01.20
  - 改进 ABD 手机截屏获取方式，提高速度
  - 题目和选项一次截取识别
- 2018.01.18
  - 增加知识图谱结果
- 2018.01.15
  - 调整适应性，百万英雄/芝士超人/冲顶大会/花椒百万赢家等全适用
- 2018.01.10
  - 改进算法逻辑，提高正确率


## 方法原理

1. ADB 获取手机截屏

  - ~~命令行生成截屏图片，把图片传输到电脑上~~
```
adb shell screencap -p /sdcard/screenshot.png
adb pull /sdcard/screenshot.png .
```

  - 直接读取屏幕数据，速度更快（参考 [http://blog.csdn.net/wirelessqa/article/details/29187339](http://blog.csdn.net/wirelessqa/article/details/29187339)）
```
adb shell screencap -p
```

2. OCR 识别题目与选项文字   

  ![](/doc/cut.png)

  ​
  两个方法：

  - 谷歌 [Tesseract](https://github.com/tesseract-ocr/tesseract) ，安装软件即可，不同电脑配置运行效率不同
  - [百度 OCR](https://cloud.baidu.com/product/ocr) ，需要注册百度 API，每天调用次数有限
  
3. 申请百度OCR

* [注册百度云账号](http://ai.baidu.com/tech/ocr/general)，打开[管理中心](https://console.bce.baidu.com/iam)，然后按下图操作

![](/doc/baidu.gif)


4. 通过算法对搜索结果进行筛选、判断

    目前用到的算法：
  - 文本关联相似度算法
  - 结巴分词算法


## 文本关联相似度算法

（参考 [github.com/smileboywtu/MillionHeroAssistant](https://github.com/smileboywtu/MillionHeroAssistant)）

通过分别统计问题与三个答案的关联度来选择正确的答案，在集合相当大的情况下，关联度会呈现正相关。

假设题目是： 

*中国历史上著名的科举制度开始于那个朝代？*
- 汉朝
- 唐朝
- 隋朝

我们先用百度分别搜索`汉朝`，`唐朝`，`隋朝`，得到如下数据：

朝代 | 搜索出的数量（来自百度为您找到相关结果约）
---- | ------------------------------------------
汉朝 | 17900000
唐朝 | 30500000
隋朝 | 16600000

然后我们在用`题目` + `答案`的方式，搜索示例：

`中国历史上著名的科举制度开始于那个朝代？ 汉朝` 得到三次的搜索结果：

 关键字  | 搜索出的数量（来自百度为您找到相关结果约）
-------- | ------------------------------------------
Q + 汉朝 | 602000
Q + 唐朝 | 837000
Q + 隋朝 | 658000

关联度计算方式：

``` shell
K = count(Q & A) / (count(Q) * count(A))
```

关联度如下：

答案 | 关联度
---- | ------
汉朝 | 0.0336
唐朝 | 0.0274
隋朝 | 0.0396


## 使用步骤
### Android

#### 1. 安装 ADB

**windows**

下载地址：https://adb.clockworkmod.com/ ，并配置环境变量

**Mac**

使用 brew 进行安装 `brew cask install android-platform-tools`



安装完后插入安卓设备且安卓已打开 USB 调试模式，终端输入 `adb devices` ，显示设备号则表示成功。我手上的机子是坚果 pro1，第一次不成功,查看设备管理器有叹号，使用 [handshaker](https://www.smartisan.com/apps/handshaker) 加载驱动后成功，也可以使用豌豆荚之类的试试。

```
List of devices attached
6934dc33    device
```

若不成功，可以参考[Android 和 iOS 操作步骤](https://github.com/wangshub/wechat_jump_game/wiki/Android-%E5%92%8C-iOS-%E6%93%8D%E4%BD%9C%E6%AD%A5%E9%AA%A4)进行修改


#### 2.安装模拟器

**windows**

  - 安装[夜神模拟器](https://www.yeshen.com/)或[逍遥模拟器](http://www.xyaz.cn/)，然后在模拟器上安装西瓜视频、冲顶大会等答题APP
  - 把模拟器设置成竖屏（分辨率900*1440），打开 USB 调试模式
  - 确认端口号，夜神模拟器62001，逍遥模拟器21503
  - 连接测试，连接成功会有提示
  ```shell
  adb connect 127.0.0.1:62001
  ```

#### 3. 安装谷歌 Tesseract

Windows下链接：
*推荐使用安装版，在安装时选择增加中文简体语言包*
- 安装版：
  https://digi.bib.uni-mannheim.de/tesseract/tesseract-ocr-setup-3.05.01.exe
- 免安装版：
  https://github.com/parrot-office/tesseract/releases/download/3.5.1/tesseract-Win64.zip
  *免安装版需要下载[中文语言包](https://github.com/tesseract-ocr/tesseract/wiki/Data-Files)，放置到Tesseract的`tessdata`目录下*

其他系统：
https://github.com/tesseract-ocr/tesseract/wiki

### iOS

#### 1.安装WDA

iOS 真机如何安装 WebDriverAgent（参考 [https://testerhome.com/topics/7220](https://testerhome.com/topics/7220)）

#### 2.安装tesseract以及简体中文包

**Mac**

```shell
brew install tesseract
cd /usr/local/Cellar/tesseract/{version}/share/tessdata
wget https://github.com/tesseract-ocr/tessdata/raw/master/chi_sim.traineddata
```

### 配置 Golang 环境

#### 安装 Golang
```
https://golang.org/doc/
```

#### 从源码安装本项目

```
go get -u github.com/henson/Answer
```

#### 直接下载可执行文件

根据平台直接下载各版本运行文件： [releases](https://github.com/henson/Answer/releases)


### 运行

 1. 根据实际情况设置配置文件

**配置参数说明：**

```
# 是否开始调试模式
debug: true
# 游戏名称, xigua西瓜视频/cddh冲顶大会/huajiao花椒/zscr芝士超人
app: xigua
# 是否模拟自动答题
automatic: false
# 对应的设备类型：ios or android
device: android
# android 连接地址
adb_address: '127.0.0.1:62001'
# ios 设备连接wda的地址
wda_address: '127.0.0.1:8100'
# 西瓜视频裁剪区域：x,y,w,h
xg_q_x: 30
xg_q_y: 240
xg_q_w: 840
xg_q_h: 180
xg_a_x: 30
xg_a_y: 410
xg_a_w: 840
xg_a_h: 450
……
```

  2. 填入百度 OCR key

修改参数：

```
# 百度OCR API参数
Baidu_API_Key: "**************************"
Baidu_Secret_Key: "**************************"
```

  3. 程序运行：

```
./Answer
```
或
```
cd cmd
go run main.go
```

## 项目参考

  - [qanswer](https://github.com/silenceper/qanswer) (golang)
  - [MillionHeroAssistant](https://github.com/smileboywtu/MillionHeroAssistant) (python)
  - [TopSup](https://github.com/Skyexu/TopSup) (python)
  - [wenda-helper](https://github.com/rrdssfgcs/wenda-helper) (python)

本项目在开发过程中参考了以上开源项目，在此对开源作者表示感谢！

## TODO

- 对题目进行分类，建立题型库，针对不同题型采取不同的处理方法

