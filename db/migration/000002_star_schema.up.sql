create table if not EXISTS star (
  id BIGINT primary key comment '主键' AUTO_INCREMENT,
  userId BIGINT not null comment '用户 id',
  gameId VARCHAR(100) not null comment '比赛 id, 来自官网',
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
)