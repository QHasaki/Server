# 路由

## 配置模块

| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 配置 | /conf  | GET |  conf.proto --> Config |
| 用户角色 | /userRole  | GET |  conf.proto --> UserRole |
| 判题语言 | /getAllLanguage  | GET |  conf.proto --> JudgeLanguage |
| 判题结果 | /getJudgeResult  | GET |  conf.proto --> JudgeResults |
| 组卷算法 | /paperComposeAlgorithm  | GET |  conf.proto --> PaperComposeAlgorithm |

## 登陆模块


| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 注册 | /register  | POST |  login.proto --> RegisterReq & RegisterResp |
| 登陆 | /login  | POST |  login.proto --> LoginReq & LoginResp |
| 登出 | /logout | GET | login.proto --> LogoutResp |
 



## 管理员模块

### 用户管理

| 功能描述 | 路由 | HTTP方法 |protobuf |  
| ----- | ---- | ----- |  ----- |
| 获得用户 | /getUsers  | POST |  user_manage.proto --> GetUsersReq & GetUsersResp |
| 新增用户 | /addUsers | POST |  user_manage.proto --> AddUsersReq & AddUsersResp |
| 更新用户 | /updateUsers | POST |  user_manage.proto --> UpdateUsersReq & UpdateUsersResp |
| 删除用户 | /delUsers | POST |  user_manage.proto --> DelUsersReq & DelUsersResp |
| 提交记录 | /submitRecord | POST |  user_manage.proto --> GetSubmitRecordReq & GetSubmitRecordResp |

### 题目管理 （目前没有判题文件的部分）

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获得所有题目 （只有部分信息，包括题目标题，难度...） | /getProblems  | POST |  problem_manage.proto --> GetProblemsReq & GetProblemsResp |
| 获得题目具体信息 | /getProblemByID | POST | problem_manage.proto --> GetProblemByIDReq & GetProblemByIDResp |
| 新增题目 | /addProblem | POST | problem_manage.proto --> AddProblemReq & AddProblemResp |
| 编辑题目 | /editProblem | POST | problem_manage.proto --> EditProblemReq & EditProblemResp |


### 班级管理 

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获得所有班级 （只有部分信息，包括班级名称，创建时间，总人数...） | /getClasses  | POST | class_manage.proto --> GetClassesReq & GetClassesResp |
| 获得班级具体信息 | /getClassByID | POST | class_manage.proto --> GetClassByIDReq & GetClassByIDResp |
| 新增班级 | /addClass | POST | class_manage.proto --> AddClassReq & AddClassResp |
| 编辑班级 | /editClass | POST | class_manage.proto --> EditClassReq & EditClassResp |
| 班级成员管理 | /memberManage | POST | class_manage.proto --> MemberManageReq & MemberManageResp |

### 公告管理 

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获得全局公告 | /getAnnouncements  | POST | announcement.proto --> AnnouncementsReq & AnnouncementsResp |
| 获得公告具体信息 | /announcementDetail | POST | announcement.proto --> AnnouncementDetailReq & AnnouncementDetailResp |
| 新增公告 | /addAnnouncement | POST | announcement.proto --> AddAnnouncementReq & AddAnnouncementResp |
| 编辑公告 | /editAnnouncement | POST | announcement.proto --> EditAnnouncementReq & EditAnnouncementResp |
| 删除公告 | /delAnnouncement | POST | announcement.proto --> DelAnnouncementReq & DelAnnouncementResp |

### 比赛管理 

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 新建比赛(自动组卷) | /newMatch  | POST | match.proto --> NewMatchReq & NewMatchResp |
| 获得所有比赛信息 | /getMatches | POST | match.proto --> GetMatchesReq & GetMatchesResp |
| 获得比赛具体信息 | /getMatchByID | POST | match.proto --> GetPaperByIDReq & GetPaperByIDResp |
| 获取比赛试卷 | /getPaperByID | POST | match.proto --> GetPaperByIDReq & GetPaperByIDResp |
| 编辑比赛 | /editMatch | POST | match.proto --> EditMatchReq & EditMatchResp |

### Tag管理 

| 功能描述 | 路由 | HTTP方法 |  protobuf |  
| ----- | ---- | ----- | ----- |
| 获取所有tag信息 | /getTags  | POST | tag.proto --> GetTagsReq & GetTagsResp |
| 修改tag信息 | /updateTag | POST | tag.proto --> UpdateTagReq & UpdateTagResp |
| 新增他给 | /addTag | POST | tag.proto --> AddTagReq & AddTagResp |


## 排行榜模块

| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 排行榜信息 | /rank  | POST |  rank.proto --> RankListReq & RankListResp |


## 判题模块

| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 排行榜信息 | /judge  | POST |  judge.proto --> JudgeRequest & JudgeResponse |

## 用户做题情况模块

| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 根据难度 | /difficultyAnalysis  | POST |  analysis.proto --> AnalysisByDifficultyReq & AnalysisByDifficultyResp |
| 根据tags | /tagsAnalysis  | POST |  analysis.proto --> AnalysisByTagsReq & AnalysisByTagsResp |

## 学生进入班级
| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 搜索班级 | /searchClasses | POST | class_enter.proto --> SearchClassesReq & SearchClassesResp |
| 学生申请进入班级 | /enterClass | POST | class_enter.proto --> EnterClassReq & EnterClassResp |
| 学生申请退出班级 | /quitClass | POST | class_enter.proto --> QuitClassReq & QuitClassResp |
| 教师查看申请加入班级的请求 | /getMember | POST | class_enter.proto --> GetMemberReq & GetMemberResp |
| 教师批准/拒绝学生进入班级的申请 | /applyEnterRequest | POST | class_enter.proto --> ApplyEnterRequestReq & ApplyEnterRequestResp |
| 根据用户的id获取参加的班级 | /getClassesByUserID | GET | class_enter.proto --> GetClasserByUserID |

## 智能组卷
| 功能描述 | 路由 | HTTP方法 | protobuf |
| ----- | ---- | ----- | ---- |
| 智能组卷 | /newPaper | POST | paper.proto --> NewPaperReq & NewPaperResp |
| 增/删试卷的题目 | /modifyPaper | POST | paper.proto --> ManualModifyPaperReq & ManualModifyPaperResp |