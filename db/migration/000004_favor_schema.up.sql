create table if not EXISTS favor (
  userId BIGINT not null comment '用户 id',
  gameId VARCHAR(100) not null comment '比赛 id, 来自官网',
  favorLevel TINYINT not null comment '喜欢 1',
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
  primary key (userId, gameId)
)