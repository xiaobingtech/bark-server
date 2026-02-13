type

是

Integer

布局类型：

+   3：进度可视化类型，适用于外卖配送、生鲜配送、车辆接驾进展等涉及进度节点显示的活动。
+   4：强调文本模板类型，适用于展示取餐码、取件码、车牌号等关键信息的活动。
+   5：左右文本模板类型，适用于高铁、火车、航班等涉及展示起点、终点的活动。
+   7：赛事类型，适用于体育赛事比分场景、游戏赛事比分场景等。

keywords

否

Map<String, String>

实况窗关键词，operation为0且event为如下场景时，必填。

+   event为FLIGHT时，仅有**flightNo**一个keyword，表示航班号，占位符格式：{{flightNo}}。
    
    示例：
    
    收起
    
    自动换行
    
    深色代码主题
    
    复制
    
+   event为TRAIN时，仅有**trainNo**一个keyword，表示火车车次，占位符格式：{{trainNo}}。
    
    示例：
    
    收起
    
    自动换行
    
    深色代码主题
    
    复制
    

消息体中占位符的使用，参见[支持携带占位符的字段](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section111894703518)。

additionalText

否

String

提示信息/免责声明。仅在NotificationData.type=5时可用。（注意消息体大小限制，详情参见[使用约束](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-intro#section114731924131418)）

keepTime

否

Long

实况窗通知存档期，在结束实况窗通知后，通知仍保留在通知中心的时长，**默认****0****不保留**，最多设置1小时，单位为秒（s）。

存档期时间以结束实况窗消息中携带的此字段数据为准，存档期期间不支持再次更新或结束通知。

contentTitle

否

String

通知标题，长度最大1024字符。

operation为0时必填，且不能为空字符串。

contentText

否

Array \[[RichText](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section33109121415) Object\]

通知内容，由多段富文本RichText组成，文本长度总和不超过1024字符，若设置文本颜色，只允许设置为同一种颜色。

operation为0时必填，且不能为空Array。

richProgress

否

[RichProgress](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section3869231875) Object

丰富进度信息，type为3时必填，具体字段请参见[RichProgress](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section3869231875)结构体。

singleTextBlock

否

[SingleTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section4853202219519) Object

强调文本模板样式中，强调的文本块，type为4时必填，默认占据左侧扩展区，具体字段请参见[SingleTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section4853202219519)结构体。

firstTextBlock

否

[FirstTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section194733293615) Object

多文本块布局中的左侧文本块，type为5时必填，详情可参见[FirstTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section194733293615)结构体。

lastTextBlock

否

[LastTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1714025575) Object

多文本块布局中的右侧文本块，type为5时必填，详情可参见[LastTextBlock](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1714025575)结构体。

displayHorizontalLine

否

Boolean

是否显示扩展区域的分割线，不设置默认显示分割线。

+   true：显示
+   false：不显示

说明

当type为5或7时才会显示分割线**。**

spaceIcon

否

String

间隔图标，本地资源，type为5时占据扩展区中间。

operation为0，type为5，spaceType未传或者spaceType为0时必填，且不能为空字符串。

取值为在指定路径下的文件名。

示例：图标文件“icon.png”存放在应用的“/resources/rawfile”路径下，则取值为“icon.png”。

spaceText

否

String

间隔文本，type为5时占据扩展区中间。

operation为0，type为5，spaceType为1时必填，且不能为空字符串。（注意消息体大小限制，详情参见[使用约束](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-intro#section114731924131418)）

style

否

Integer

左右文本样式类型

0：强调型

1：均衡型

说明

创建时未传style字段将使用强调型展示。

spaceType

否

Integer

间隔类型

0：使用图标

1：使用文本

说明

创建时未传spaceType字段将使用图标展示。

extend

否

[Extend](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section15169131522812) Object

辅助区样式，无更新时可不携带。具体字段请参见[Extend](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section15169131522812)结构体。

说明

+   更新type类型为新布局时，需重新携带本字段。
+   刷新实况窗通知内容时，**辅助区显示类型为图片且图片路径填写错误会导致刷新内容失败**。

game

否

[Game](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1134158205816) Object

赛事信息扩展区，type为7时必填，具体字段请参见[Game](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1134158205816)结构体。

descPic

否

String

扩展区域描述图片，默认不显示，当type为4时且传值会占据右侧扩展区。不携带时系统显示时采用上次刷新的图像。

operation为0且type为4时必填，且不能为空字符串。

取值为在指定路径下的文件名。

示例：图标文件“icon.png”存放在应用的“/resources/rawfile”路径下，则取值为“icon.png”

clickAction

是

[ClickAction](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section152462191216) Object

消息点击行为，具体字段请参见[ClickAction](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section152462191216)结构体。

lockScreen

否

[LiveViewLockScreen](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1786105519115) Object

锁屏沉浸实况窗相关字段，具体字段请参见[LiveViewLockScreen](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1786105519115)结构体。

weather

否

[Weather](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section831712212426) Object

传入天气信息结构体。需要同时传入天气类型、天气位置类型与最高最低温度参数，才会在卡片上展示天气。仅支持左右文本模板（即type为5）。

当传入天气类型为雨、雪特殊天气，且同时传入实况窗卡片的背景氛围类型参数backgroundType（合法值参见Live View Kit [BackgroundType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/liveview-liveviewmanager#section15834105743213)枚举值）为赏月航班或夕阳航班对应的值时，卡片上优先展示天气背景，其余非特殊天气在卡片上优先展示赏月航班或夕阳航班背景氛围。

backgroundType

否

Integer

表示实况窗卡片的背景氛围类型，仅支持左右文本模板（即type为5），合法值参见Live View Kit [BackgroundType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/liveview-liveviewmanager#section15834105743213)枚举值。

当传入实况窗卡片的背景氛围类型参数为赏月航班或夕阳航班对应的值时，且同时传入天气类型（[Weather](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section831712212426)）为雨、雪特殊天气，卡片上优先展示天气背景，其余非特殊天气在卡片上优先展示赏月航班或夕阳航班背景氛围。