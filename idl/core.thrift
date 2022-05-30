namespace go core

// 公共返回值（减少重复代码）
struct BaseResp {
    1: i32 status_code  // 状态码，0-成功，其它-失败
    2: optional string status_msg  // 返回状态描述
}

// 用户信息
struct User {
    1: i64 id   // 用户id
    2: string name  // 用户名称
    3: optional i64 follow_count  // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: bool is_follow  // true-已关注，false-未关注
}

struct douyin_user_request {
    1: i64 user_id   // 用户id
    2: string token   // 用户鉴权token
}

struct douyin_user_response {
    1: BaseResp base_resp
    3: User user  // 用户信息
}

// 视频信息
struct Video {
    1: i64 id   // 视频唯一标识
    2: User author  // 视频作者信息
    3: string play_url  // 视频播放地址
    4: string cover_url // 视频封面地址
    5: i64 favorite_count // 视频的点赞总数
    6: i64 comment_count  // 视频的评论总数
    7: bool is_favorite  // true-已点赞，false-未点赞
    8: string title  // 视频标题
}

// 用户注册请求
struct douyin_user_register_request {
    1: string username  // 登录用户名
    2: string password  // 密码
}

struct douyin_user_register_response {
    1: BaseResp base_resp
    3: i64 user_id  // 用户id
    4: string token   // 用户鉴权token
}

// 用户登录请求
struct douyin_user_login_request {
    1: string username  // 登录用户名
    2: string password  // 密码
}

struct douyin_user_login_response {
    1: BaseResp base_resp
    3: i64 user_id  // 用户id
    4: string token  // 用户鉴权token
}

// 视频流请求
struct douyin_feed_request {
    1: optional i64 latest_time  // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token   // 登录用户设置
}

struct douyin_feed_response {
    1: BaseResp base_resp
    3: list<Video> video_list  // 视频列表
    4: optional i64 next_time  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// 视频投稿请求
struct douyin_publish_action_request {
    1: string token   // 用户鉴权token
    2: binary data   // 视频数据（字节数组）
    3: string title  // 视频标题
}

struct douyin_publish_action_response {
    1: BaseResp base_resp
}

// 查看发布列表请求
struct douyin_publish_list_request {
    1: i64 user_id   // 用户id
    2: string token   // 用户鉴权token
}

struct douyin_publish_list_response {
    1: BaseResp base_resp
    3: list<Video> video_list   // 用户发布的视频列表
}

service CoreService {   // 用户服务
    douyin_user_register_response Register(1:douyin_user_register_request req)  // 注册
    douyin_user_login_response Login(1:douyin_user_login_request req)  // 登录
    douyin_user_response GetUser(1:douyin_user_request req)   // 获取用户信息
    douyin_feed_response GetVideoFeed(1:douyin_feed_request req)  // 获取视频流推送
    douyin_publish_action_response PublishVideo(1:douyin_publish_action_request req)  // 视频投稿
    douyin_publish_list_response GetPublishedList(1:douyin_publish_list_request req)  // 已发布视频列表
}