# gin-gen

### 介绍
use bash to generate directory for gin project 
生成gin api项目结构。
> https://github.com/unionline/gin-gen.git

#### 安装
```shell script
go get github.com/unionline/gin-gen
```

#### 生成项目
```shell script
cd YourGOPath/src/github.com/unionline/gin-gen
go build main.go
./gin-gen new <project-name>
mv  <project-name> YourGOPath/src/
```
```├─ Project Name
    │  ├─ config          //配置文件
    │     ├── ...
    │  ├─ controllers     //控制器层
    │     ├── ...
    │  ├─ entities        //实体
    │     ├── ...
    │  ├─ initializers    //初始化
    │     ├── ...
    │  ├─ repositories    //数据库操作层
    │     ├── ...
    │  ├─ models          //数据库ORM
    │     ├── ...
    │  ├─ services        //业务层
    │     ├── ...
    │  ├─ proto           //proto文件
    │     ├── ...
    │  ├─ routers         //路由
    │     ├── middleware  //路由中间件
    │         ├── ...
    │     ├── ...
    │  ├─ utils           //工具类
    │     ├── ...
    │  ├─ vendor          //扩展包
    │     ├── ...
    │  ├─ views           //View
        │     ├── ...
    │  ├─ main.go         //入口文件
```

参考1：https://blog.csdn.net/XinLiangTalk/article/details/99670139

参考2：www.gitee.com/zhouxiaozhu/gin-gen
