-- postgres
-- 游戏道具
create table prop (
  id serial primary key,                           -- prop_id, 主键唯一
  prop_name varchar not null unique,               -- 道具名，唯一
  can_present smallint not null default 2,         -- 1-可赠送,2-不可赠送
  can_destroy smallint not null default 1,         -- 1-可销毁(丢弃), 2-不可销毁(丢弃)
  auto_used smallint not null default 2,           -- 1-自动使用， 2-不自动使用
  image_url varchar not null default ''            -- 图片cdn地址
);

-- 用户背包
create table user_prop(
  id serial primary key,
  user_id integer not null default 0,            -- 用户id
  prop_id integer not null default 0,            -- 道具id
  expire_in integer not null default -1,         -- 失效于
  prop_num integer not null default 1,           -- 道具数量
  prop_title varchar not null default ''         -- 道具标题
);

-- 活动配置
create table activity_config(
    id serial primary key,                       -- activity_id
    state integer not null default 1,            -- 1-关闭，2-开启
    open_config jsonb,                           -- 开启设置
    reward_config jsonb                          -- 奖品设置
);

-- 用户进度
create table user_activity_process(
   id serial primary key,
   user_id integer not null,              -- 用户id
   activity_id integer not null,          -- 活动id
   date_time varchar not null,            -- 生成时间 '2019-06-11 18:11:11'
   joint_config jsonb,                    -- 游戏进度
   unique(user_id, activity_id)
);

create table backend_user(
  id serial primary key,              -- user_id
  username varchar not null unique,   --用户名，唯一
  password_hash varchar not null,     -- 加盐密码
  role_id integer not null default 1, -- 1-普通用户，2-管理员
  content_control jsonb               -- 存放了对应后台的权限控制配置
)
