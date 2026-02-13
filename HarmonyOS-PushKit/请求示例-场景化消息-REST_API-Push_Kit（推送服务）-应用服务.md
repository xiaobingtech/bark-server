请求示例
更新时间: 2026-01-14 17:55
通知消息
// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:0
 
// Request Body
{ 
  "payload": { 
    "notification": { 
      "category": "xxxxxx", // category替换为实际通知消息类型
      "title": "普通标题", 
      "body": "普通内容", 
      "clickAction": { 
        "actionType": 0
      },
      "style": 0, 
      "image":"https://lf*******246.png" 
    } 
  }, 
  "target": { 
    "token": ["MAAALgE4G98BAAAAst*******jg"] 
  } 
}
应用在前台时接收通知消息
应用只在后台展示通知消息；应用在前台时，通知消息将不会展示，但可以接收通知消息后自行完成业务处理，详情请参见应用在前台时处理通知消息。

// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type: 0

// Request Body
{
  "payload": {
    "notification": {
      "category": "MARKETING",
      "title": "普通通知标题",
      "body": "普通通知内容",
      "clickAction": {
        "actionType": 0
      },
      "foregroundShow": false  // 设置为false则应用在前台时不会展示通知消息，默认为true表示前后台都展示
    }
  },
  "target": {
    "token": ["MAMzLg**********lPW"]
  },
  "pushOptions": {
    "testMessage": true,
    "ttl": 86400
  }
}
卡片刷新消息
// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:1
 
// Request Body
{ 
  "payload": { 
    "formData": { 
      "123": 96, 
      "class": "123" 
    }, 
    "version": 922337203, 
    "images": [ 
      { 
        "keyName": "hello", 
        "url": "https://xxx.png", 
        "require": 1 
      } 
    ], 
    "formId": 0, 
    "moduleName": "testName", 
    "formName": "testFormName", 
    "abilityName": "testAbilityName"      
  }, 
  "pushOptions": { 
    "biTag": "this is bi", 
    "ttl": 666 
  }, 
  "target": { 
    "token": [ 
      "MAAALgE4G98BAAAAst************ttQd4Tw" 
    ]   
  } 
}
语音播报消息
// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:2
 
// Request Body
{ 
  "payload": { 
    "extraData": "Extension extra data", 
    "notification": { 
      "category": "PLAY_VOICE", 
      "title": "普通标题", 
      "body": "普通内容", 
      "clickAction": { 
        "actionType": 0
      }, 
      "style": 0, 
      "image":"https://lf*******246.png" 
    } 
  }, 
  "target": { 
    "token": ["MAAALgE4G98BAAAAst*******jg"] 
  } 
}
后台消息
// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:6
 
// Request Body
{ 
  "payload": { 
    "extraData": "携带的数据" 
  }, 
  "target": { 
    "token": ["MAAALgE4G98BAAAAst************jq"] 
  } 
}
创建实况窗消息
航班场景（event为FLIGHT）
计划出发，使用左右文本模板。

// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"
 
// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---**** 
push-type:7 

// Request Body
{
  "pushOptions": {
    "ttl": 1000,
    "biTag": "biTag"
  },
  "payload": {
    "activityId": 1,
    "operation": 0,
    "event": "FLIGHT",
    "status": "DEPART", // 计划出发
    "activityData": {
      "notificationData": {
        "keywords": {
          "flightNo": "MU1471"
        },
        "type": 5,
        "contentTitle": "航班{{status}}", // 航班计划出发
        "contentText": [
          {
            "text": "航班号："
          },
          {
            "text": "{{flightNo}}", // MU1471
            "foregroundColor": "#FF317AF7"
          }
        ],
        "clickAction": {
          "actionType": 0
        },
        "firstTextBlock": {
          "firstLine": "12:00",
          "secondLine": "上海虹桥"
        },
        "lastTextBlock": {
          "firstLine": "14:20",
          "secondLine": "成都天府"
        },
        "displayHorizontalLine": true,
        "spaceIcon": "space.png", // 取值为“/resources/rawfile”路径下的文件名
        "extend": {
          "type": 3,
          "pic": "flight.png", // 取值为“/resources/rawfile”路径下的文件名
          "clickAction": {
            "actionType": 1,
            "action": "xxxxxxx"
          }
        }
      },
      "capsuleData": {
        "type": 1,
        "status": 1,
        "icon": "icon.svg", // 取值为“/resources/rawfile”路径下的文件名
        "bgColor": "#FF317AF7",
        "remind": "EXPAND",
        "title": "即将出发",
        "content": "请尽快前往机场"
      }
    }
  },
  "target": {
    "token": [
      "MAAALgE4G98BAAAAst************jq"
    ]
  }
}
出行打车场景（event为TAXI）
司机正在赶来，使用进度可视化模板。

// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:7

// Request Body
{
  "pushOptions": {
    "ttl": 1000,
    "biTag": "biTag"
  },
  "payload": {
    "activityId": 2,
    "operation": 0,
    "event": "TAXI",
    "status": "DRIVER_ON_THE_WAY", // 司机正在赶来
    "activityData": {
      "notificationData": {
        "type": 3,       
        "contentTitle": "{{status}}", // 司机正在赶来
        "contentText": [
          {
            "text": "距您"
          },
          {
            "text": "1.2公里",
            "foregroundColor": "#FF317AF7"
          },
          {
            "text": " | "
          },
          {
            "text": "5分钟",
            "foregroundColor": "#FF317AF7"
          }
        ],
        "clickAction": {
          "actionType": 1, // 打开应用自定义页面
          "action": "xxxxxx" // 应用内置页面ability对应的action
        },
        "richProgress": {
          "type": 0,
          "nodeIcons": ["icon1.png", "icon2.png", "icon3.png"], // 取值为“/resources/rawfile”路径下的文件名
          "indicatorIcon": "taxi.png", // 取值为“/resources/rawfile”路径下的文件名
          "progress": 40,
          "indicatorType": 1,
          "color": "#FF317AF7",
          "bgColor": "#19000000"
        },
        "extend": {
          "type": 3,
          "pic": "phone.png", // 取值为“/resources/rawfile”路径下的文件名
          "clickAction": {
            "actionType": 5, // 打开拨号界面
            "data": {
              "tel": "138xxxxxxxx" // 通过tel字段携带电话号码
            }
          }
        }
      },
      "capsuleData": {
        "type": 1,
        "status": 1,
        "icon": "icon.svg", // 取值为“/resources/rawfile”路径下的文件名
        "bgColor": "#FF317AF7",
        "remind": "EXPAND",
        "title": "接驾中",
        "content": "预计5分钟"
      }
    }
  },
  "target": {
    "token": [
      "MAAALgE4G98BAAAAst************jq"
    ]
  }
}
高铁/火车场景（event为TRAIN）
计划出发，使用左右文本模板。

// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"

// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---****
push-type:7

// Request Body
{
  "pushOptions": {
    "ttl": 1000,
    "biTag": "biTag"
  },
  "payload": {
    "activityId": 3,
    "operation": 0,
    "event": "TRAIN",
    "status": "DEPART", // 计划出发
    "title": "列车即将出发",
    "content": "请尽快去高铁站",
    "activityData": {
      "notificationData": {
        "keywords": {
          "trainNo": "G1406"
        },
        "type": 5,        
        "contentTitle": "列车{{status}}", // 列车计划出发
        "contentText": [
          {
            "text": "车次："
          },
          {
            "text": "{{trainNo}}", // G1406
            "foregroundColor": "#FF317AF7"
          }
        ],
        "clickAction": {
          "actionType": 0
        },
        "firstTextBlock": {
          "firstLine": "13:00",
          "secondLine": "上海虹桥"
        },
        "lastTextBlock": {
          "firstLine": "14:20",
          "secondLine": "南京南"
        },
        "displayHorizontalLine": true,
        "spaceIcon": "space.png", // 取值为“/resources/rawfile”路径下的文件名
        "extend": {
          "type": 3,
          "pic": "train.png", // 取值为“/resources/rawfile”路径下的文件名
          "clickAction": {
            "actionType": 1,
            "action": "xxxxxxx"
          }
        }
      },
      "capsuleData": {
        "type": 1,
        "status": 1,
        "icon": "icon.svg", // 取值为“/resources/rawfile”路径下的文件名
        "bgColor": "#FF317AF7",
        "remind": "EXPAND",
        "title": "即将出发",
        "content": "请尽快去高铁站"
      }
    }
  },
  "target": {
    "token": [
      "MAAALgE4G98BAAAAst************jq"
    ]
  }
}
更新实况窗消息
航班场景（event为FLIGHT）
已值机，使用左右文本模板。

// Request URL
POST "https://push-api.cloud.huawei.com/v3/3158882***52863/messages:send"
 
// Request Header
Content-Type: application/json
Authorization: Bearer eyJr*****OiIx---****.eyJh*****iJodHR--***.QRod*****4Gp---**** 
push-type:7 

// Request Body
{
  "pushOptions": {
    "ttl": 1000,
    "biTag": "biTag"
  },
  "payload": {
    "activityId": 1,
    "operation": 1,
    "event": "FLIGHT",
    "status": "CHECKED_IN", // 已值机
    "version": 1,
    "activityData": {
      "notificationData": {
        "keywords": {
          "flightNo": "MU1471"
        },
        "type": 5,
        "contentTitle": "登机口88",
        "contentText": [
          {
            "text": "{{status}} | " // 已值机
          },
          {
            "text": "{{flightNo}}", // MU1471
            "foregroundColor": "#FF317AF7"
          }
        ],
        "clickAction": {
          "actionType": 0
        },
        "firstTextBlock": {
          "firstLine": "12:00",
          "secondLine": "上海虹桥"
        },
        "lastTextBlock": {
          "firstLine": "14:20",
          "secondLine": "成都天府"
        },
        "displayHorizontalLine": true,
        "spaceIcon": "space.png", // 取值为“/resources/rawfile”路径下的文件名
        "extend": {
          "type": 3,
          "pic": "flight.png", // 取值为“/resources/rawfile”路径下的文件名
          "clickAction": {
            "actionType": 1,
            "action": "xxxxxxx"
          }
        }
      },
      "capsuleData": {
        "type": 1,
        "status": 1,
        "icon": "icon.svg", // 取值为“/resources/rawfile”路径下的文件名
        "bgColor": "#FF317AF7",
        "remind": "EXPAND",
        "title": "登机口88",
        "content": "请尽快完成安检"
      }
   