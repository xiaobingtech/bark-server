## 接口原型

<table><tbody><tr><th class="firstcol" id="mcps1.3.1.2.1.3.1.1"><p>承载协议</p></th><td class="cellrowborder"><p>HTTPS POST</p></td></tr><tr><th class="firstcol" id="mcps1.3.1.2.1.3.2.1"><p>接口方向</p></th><td class="cellrowborder"><p>开发者服务器 -&gt; 华为Push服务器</p></td></tr><tr><th class="firstcol" id="mcps1.3.1.2.1.3.3.1"><p>接口URL</p></th><td class="cellrowborder"><p>https://push-api.cloud.huawei.com/v3/<strong>[projectId]</strong>/messages:send</p><div class="hw-editor-tip info"><p>说明</p><p><strong>[projectId]</strong>：项目ID，登录<a href="https://developer.huawei.com/consumer/cn/service/josp/agc/index.html" target="_blank">AppGallery Connect</a>网站，选择“开发与服务”，在项目列表中选择对应的项目，左侧导航栏选择“项目设置”，在该页面中获取。</p></div></td></tr><tr><th class="firstcol" id="mcps1.3.1.2.1.3.4.1"><p>数据格式</p></th><td class="cellrowborder"><p>Content-Type: application/json</p></td></tr></tbody></table>

## Request Header

展开

| 
参数

 | 

取值描述

 | 

样例

 |
| --- | --- | --- |
| 

Authorization

 | 

鉴权方式：**JWT方式**。

注意

HarmonyOS 5及以上系统版本推送不再支持OAuth 2.0开放鉴权方式。

详情参见[基于服务账号生成鉴权令牌](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-jwt-token)。

说明

建议JWT令牌过期时间设置为3600秒，有效期内可以复用。

Bearer后面拼接空格，再拼接获取的鉴权信息。





 | 

Bearer eyJr\*\*\*\*\*OiIx---\*\*\*\*.eyJh\*\*\*\*\*iJodHR--\*\*\*.QRod\*\*\*\*\*4Gp---\*\*\*\*

 |
| 

push-type

 | 

消息类型，取值如下：

0：Alert消息（通知消息）

1：卡片刷新消息

2：语音播报消息

6：后台消息

7：实况窗消息

10：应用内通话消息

 | 

0

 |

## Request Body

展开

| 
参数

 | 

是否必选

 | 

参数类型

 | 

描述

 |
| --- | --- | --- | --- |
| 

payload

 | 

是

 | 

Object

 | 

推送消息结构体，不同的push-type场景拥有不同的payload定义：

+   0：[AlertPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1152516418157) 通知消息
+   1：[FormUpdatePayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section6471174713249) 卡片刷新消息
+   2：[ExtensionPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section161192514234) 语音播报消息
+   6：[BackgroundPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section056120245274) 后台消息
+   7：[LiveViewPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section66881469306) 实况窗消息
+   10：[VoIPCallPayload](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section146341550144011) 应用内通话消息

 |
| 

pushOptions

 | 

否

 | 

Object

 | 

发送控制参数，详情请参见[pushOptions](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section418321011212)的定义。

 |
| 

target

 | 

是

 | 

Object

 | 

发送目标，详情请参见[target](https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section1344624991218)的定义。

 |

相关推荐

意见反馈

以上内容对您是否有帮助？

意见反馈

如果您有其他疑问，您也可以通过开发者社区问答频道来和我们联系探讨。