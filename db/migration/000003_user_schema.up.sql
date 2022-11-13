create table if not EXISTS user (
  id BIGINT primary key comment '主键' AUTO_INCREMENT,
  openId TEXT not null comment '比赛 id, 来自官网',
  nickName TEXT comment '昵称',
  avatarFileId TEXT comment '头像 fileId',
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '最后一次修改时间'
)