create table if not EXISTS star (
  id INT primary key comment '主键',
  userId INT not null comment '用户 id',
  gameId TEXT not null comment '比赛 id, 来自官网',
  createdAt TIMESTAMP default now() comment '创建时间'
)