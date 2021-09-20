## <center> Personal Work – 中小学数学卷子自动生成  </center>

<p style="float:right;" >Personal Project – 个人项目  软件工程导论 </p>

### 技术文档

#### 技术栈

&ensp;&ensp;&ensp;&ensp;本次开发选择性能优良的Go语言，便于后期优化以及前后端拆分，利用Go语言优秀的解析处理能力以及超高的效率，实现所需要的功能。

&ensp;&ensp;&ensp;&ensp;本次为保证数据的封装性和良好的迁移性，使用`Gorm`数据库映射，选用嵌入式良好的`SQLite3`数据库作为用户存储，便于后期维护以及新功能特性的开发。

> PS: 数据库已开启自动同步，已经实现自动创建、自动初始化功能。

#### 目录结构分析

``` shell
ProjectFile # 项目文件夹
│  .gitignore # git 提交忽略项
│  application.yml # 配置文件
│  go.mod # go mod 集成 依赖详情
│  go.sum
│  main.go # 程序运行入口
│  README.md
│  test.db # sqlite3 数据库 不存在自动生成  
├─app # 程序文件夹
│  │  app.go # app 方法定义：数据库初始化、运行
│  │  gorm.go # gorm 数据库链接初始化
│  │  wire.go # wire 依赖注入
│  │  wire_gen.go
│  ├─config # 配置文件读取包
│  │      config.go # 配置文件读取操作
│  │      type.go  # 配置文件类型定义
│  ├─dao # 数据库操作
│  │      dao.go # dao层控制DI
│  │      user.go # user 操作    
│  ├─entity # 实例映射
│  │      role.go # 角色表
│  │      user.go # 用户表
│  ├─model # 操作处理层
│  │      base.go # Base 生成方法
│  │      common.go # common：随机种子
│  │      high.go # 高中特有生成方法
│  │      junior.go # 初中特有生成方法
│  │      login.go # model层登录操作
│  │      model.go # model控制DI
│  │      operator.go # 操作符定义
│  │      userGenerate.go # 接口定义实现用户生成   
│  ├─schema # 参数拆分
│  │      login.go # Login 相关参数
│  │      paper.go # PaperGenerate 相关参数
│  ├─service # 服务Provider
│  │      customer.go # currentUser 定义
│  │      login.go # 登录入口
│  │      paper.go # PaperGenerate 入口
│  │      service.go # service 控制DI
│  └─util # 抽象工具
│          input.go # 输入封装
│          judge.go # 数组去重
│          match.go # 输入验证
│          saveTXT.go # 数据保存
└─paperResult # 生成结果保存
    ├─user # 用户名
    │      *.txt # 生成文件  时间命名
    ...
```

#### 配置文件说明

``` yaml
# 运行模式
Mode: prod

GORM:
  # 是否开启调试模式
  Debug: false
  # 设置连接可以重用的最长时间(单位：秒)
  MaxLifetime: 7200
  # 设置数据库的最大打开连接数
  MaxOpenConns: 150
  # 设置空闲连接池中的最大连接数
  MaxIdleConns: 50
  # 是否启用自动映射数据库表结构
  EnableAutoMigrate: true

DB:
  # 数据库
  DBName: test.db
```

### 构建须知

#### 项目下拉

``` shell
git clone https://gitee.com/itse/personal-work.git
```

#### 依赖下载

进入项目根目录，执行下列命令:

```shell
go mod download
```

> PS: `PATH` 系统环境变量中必须存在 `GOROOT`

#### 运行

```shell
go run main.go
```

> PS: 项目根目录下会创建 `test.db`，在执行生成后，会创建 `paperResult` 文件夹

#### 构建

```shell
go build main.go
```

> PS: golang 支持交叉编译，通过环境变量设置实现。

#####

`build` 之后生成的可执行文件不需要任何依赖，可直接运行，仅需要将配置文件`application.yml`一起打包即可。

### 使用文档

将程序使用上面两种方法之一运行即可。

#### 登录

在一行中输入用户名和密码，以空格相隔

* 登录失败
    * ![image-20210921045233770](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210452356.png)
* 登录成功
    * ![image-20210921045403014](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210454121.png)

#### 试题生成

在一行中输入一个在10到30之间的数字(包括10和30)，若在范围内，出现下一次邀请生成；若不在，则提示有效输入

* 范围正确
    * ![image-20210921045952702](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210459863.png)
* 范围错误
    * ![image-20210921045904654](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210459905.png)

若生成成功，则可以在`paperResult`文件夹下对应的用户文件夹下找到以时间命名的`TXT`文件

![image-20210921050157396](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210501457.png)

#### 模式切换

输入`切换为小学 切换为初中 切换为高中`三句中任意一句时，切换至对应出题模式，当输入中包含"切换为"三个字时，若后面不符合，则出现提示

* 切换失败
    * ![image-20210921050702712](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210507848.png)
* 切换成功
    * ![image-20210921050730171](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210507405.png)

#### 用户切换

输入 `-1` 进行用户退出

![image-20210921050914129](https://gitee.com/Monkeyman520/MonkeyImgURL/raw/master/img/202109210509198.png)
