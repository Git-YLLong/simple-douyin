// 点赞操作请求
struct douyin_favorite_action_request {
    1: i64 user_id
    2: string token
    3: i64 video_id
    4: i32 action_type  // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1: i32 status_code
    2: optional string status_msg   
}

// 评论操作请求
struct douyin_comment_action_request {
    1: i64 user_id
    2: string token
    3: i64 video_id
    4: i32 action_type  // 1-发布评论，2-删除评论
    5: optional string comment_text  // 填写的评论内容，action_type=1时有用
    6: optional i64 comment_id // 删除的评论id，action_type=2时有用
}

struct douyin_comment_action_response {
    1: i32 status_code
    2: optional string status_msg 
    3: optional Comment comment  // 评论成功返回评论内容，不需要重新拉取整个列表
}

// 评论信息
struct Comment {
    1: i64 id
    2: User user  // 评论的用户信息
    3: string content
    4: string create_date  // 评论发布日期，格式mm-dd
}

// 用户信息
struct User {
    1: i64 id   // 用户id
    2: string name  // 用户名称
    3: optional i64 follow_count  // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: bool is_follow  // true-已关注，false-未关注
}

// 拉取评论列表请求
struct douyin_comment_list_request {
    1: string token
    2: i64 video_id
}

struct douyin_comment_list_response {
    1: i32 status_code
    2: optional string status_msg
    3: list<Comment> comment_list
}

// 关注操作
struct douyin_relation_action_request {
    1: i64 user_id
    2: string token
    3: i64 to_user_id   // 对方id
    4: i32 action_type  // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1: i32 status_code
    2: optional string status_msg
}

// 拉取关注列表操作
struct douyin_relation_follow_list_request {
    1: i64 user_id
    2: string token
}

struct douyin_relation_follow_list_response {
    1: i32 status_code
    2: optional string status_msg
    3: list<User> user_list
}

// 拉取粉丝列表操作
struct douyin_relation_follower_list_request {
    1: i64 user_id
    2: string token
}

struct douyin_relation_follower_list_response {
    1: i32 status_code
    2: optional string status_msg
    3: list<User> user_list
}

service ExtraService {   // 扩展接口
    douyin_favorite_action_response FavoriteAction(1:douyin_favorite_action_request req)  // 点赞
    douyin_comment_action_response CommentAction(1:douyin_comment_action_request req)  // 评论
    douyin_comment_list_response CommentList(1:douyin_comment_list_request req) // 评论列表
    douyin_relation_action_response RelationAction(1:douyin_relation_action_request req) // 关注
    douyin_relation_follow_list_response RelationFollowList(1:douyin_relation_follow_list_request req) // 关注列表
    douyin_relation_follower_list_response RelationFollowerList(1:douyin_relation_follower_list_request req) // 粉丝列表    
}