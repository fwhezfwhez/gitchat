## project
   项目为某游戏业务内可插拔的http后端活动，提供了<邀请活动>的实现与集成
   title: 为游戏添加<邀请有礼>服务
   description:
      开发约束:
         XX公司研制的xx游戏上架了，目前的业务功能简单，需要为该游戏添加各类活动，要求组员以插件化的形式开发。
         activities
             | ---- invite-activity
             | ---- login-activity
             ...

      功能约束:
         1. 每邀请1人，调取赠送奖品回调，最高奖励次数20次.
   需求:
       1. 活动插件接入app的jwt鉴权,即仅仅允许app客户端调取接口。
       2. 要求用户进度读取响应极快,不抢占数据库连接.
       3. 奖品回调仅允许活动插件调用,即app回调不允许给app客户端使用

## gin
   gin的基本用法

## gorm-postgres
   go 使用gorm连接postgres的基本用法，并存储用户数据

## redis
   go 客户端连接redis的用法，并缓存用户记录

## mongodb
  go 客户端连接mongodb，并存储访问记录
