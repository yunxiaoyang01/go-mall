
- [1. 项目文档](#1-项目文档)
- [2. 重要提示](#2-重要提示)
  - [2.1. 基本介绍](#21-基本介绍)
    - [2.1.1. 2.1.1 项目介绍](#211-211-项目介绍)
    - [2.1.2. 组织结构](#212-组织结构)
    - [2.1.3. 技术选型](#213-技术选型)
      - [2.1.3.1. 后端技术](#2131-后端技术)
      - [2.1.3.2. 前端技术](#2132-前端技术)
    - [2.1.4. 贡献指南](#214-贡献指南)
      - [2.1.4.1. Issue 规范](#2141-issue-规范)
      - [2.1.4.2. Pull Request 规范](#2142-pull-request-规范)
  - [2.2. 使用说明](#22-使用说明)
## 0.1. 前言

go-mall项目致力于使用golang语言打造一个完整的电商系统，采用现阶段流行技术实现。

# 1. 项目文档


# 2. 重要提示



## 2.1. 基本介绍

### 2.1.1. 2.1.1 项目介绍
go-mall项目是一套电商系统，包括前台商城系统及后台管理系统，基于gin+gorm实现，采用Docker容器化部署。前台商城系统包含首页门户、商品推荐、商品搜索、商品展示、购物车、订单流程、会员中心、客户服务、帮助中心等模块。后台管理系统包含商品管理、订单管理、会员管理、促销管理、运营管理、内容管理、统计报表、财务管理、权限管理、设置等模块。

### 2.1.2. 组织结构

``` lua
go-mall
├── mall-admin -- 商城后台接口
├── mall-web-admin -- 商城管理后台代码
├── mall-web-app -- 前台商城(敬请期待)
```
### 2.1.3. 技术选型

#### 2.1.3.1. 后端技术

| 技术          | 说明               | 官网                                     |
| ------------- | ------------------ | ---------------------------------------- |
| gin           | golangWeb框架      | https://github.com/gin-gonic/gin         |
| JWT           | JWT登录支持        | https://github.com/jwtk/jjwt             |
| Gorm          | ORM框架            | https://gorm.io/index.html               |
| Elasticsearch | 搜索引擎           | https://github.com/elastic/elasticsearch | ß |
| RabbitMQ      | 消息队列           | https://www.rabbitmq.com/                | ß |
| Redis         | 分布式缓存         | https://redis.io/                        |
| MongoDB       | NoSql数据库        | https://www.mongodb.com                  |
| LogStash      | 日志收集工具       | https://github.com/elastic/logstash      |
| Kibana        | 日志可视化查看工具 | https://github.com/elastic/kibana        |
| Nginx         | 静态资源服务器     | https://www.nginx.com/                   |
| Docker        | 应用容器引擎       | https://www.docker.com                   |
#### 2.1.3.2. 前端技术

| 技术       | 说明                  | 官网                                   |
| ---------- | --------------------- | -------------------------------------- |
| Vue        | 前端框架              | https://vuejs.org/                     |
| Vue-router | 路由框架              | https://router.vuejs.org/              |
| Vuex       | 全局状态管理框架      | https://vuex.vuejs.org/                |
| Element    | 前端UI框架            | https://element.eleme.io               |
| Axios      | 前端HTTP框架          | https://github.com/axios/axios         |
| v-charts   | 基于Echarts的图表框架 | https://v-charts.js.org/               |
| Js-cookie  | cookie管理工具        | https://github.com/js-cookie/js-cookie |
| nprogress  | 进度条控件            | https://github.com/rstacruz/nprogress  |
###  2.1.4. 贡献指南


#### 2.1.4.1. Issue 规范

#### 2.1.4.2. Pull Request 规范


## 2.2. 使用说明




