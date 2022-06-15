# Gopus
## 介绍 Introduction
Gopus是一个基于Gin的Go WebService开发框架。  
本意是为向习惯了SpringMVC开发模式的WebService开发者提供MVC架构的支持。

## 架构 Architecture

### src
src目录包含了Gopus的代码以及相关的配置文件

### src/config
config目录包含了一些配置实体类以及配置文件

### src/main
main目录包含所有的业务性代码以及一些设计模式

### src/main/framework
framework下包含了若干设计模式类

### src/main/mvc
mvc即框架的基本主体

## 开发过程 Updates
2022-6-15: 项目创建与基本的README
2022-6-15: 集成Gorm，并通过模板方法模式(Template Method)抽象出数据库连接的初始化过程
2022-6-15: 通过模板方法模式(Template Method)抽象出WebService服务器的初始化过程

