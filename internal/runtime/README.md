# Runtime Server

動態控制、提取 Event Engine 的 History、Story。

## RESTful API

* 支援 HTTP 格式來獲取資料
* 支援的 url：

| status | Method | URL | 功能 |
| --- | --- | --- | --- |
| [x] | GET | `/health` | health check |
| [x] | GET | `/fetch`  | 獲得所有事件表 |
| [x] | GET | `/fetch/time` | 獲得編年史式的事件紀錄 |
| [x] | POST | `/insert` | runtime 塞入新的 event model |

## gRPC 

> `Work in progress` ...

* 使用 gRPC protocol 來獲取資料