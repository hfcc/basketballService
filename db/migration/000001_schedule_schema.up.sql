create table if not EXISTS schedule (
  id INT primary key comment '主键',
  seasonName VARCHAR(20) not null comment '赛季名称 例如:2021-2022 赛季',
  gameId VARCHAR(100) not null comment '比赛 id, 来自官网',
  seasonStageId TINYINT not null comment '比赛所处的赛季阶段, 1常规赛 2决赛',
  startTimeUTC TIMESTAMP comment '比赛时间',
  hTeamId VARCHAR(100) not null comment '主队 id',
  hScore TINYINT comment '主队分数',
  vTeamId VARCHAR(100) not null comment '客队 id',
  vScore TINYINT comment '客队分数',
  createdAt TIMESTAMP default now() comment '创建时间',
  updatedAt TIMESTAMP default now() comment '最后一次修改时间'
)
