# User Guide

创建webhook后您的系统可以收到UCloud告警信息。当有告警时，UCloud的系统会将告警信息以HTTP POST方式发送到指定URL，您可以对收到信息进行处理。

## 先决条件
用户需要提供接收POST请求的HTTP服务，以处理UCloud发送的POST请求，并将该服务的URL注册到UCloud的告警系统中。

## JSON Body Example
  
### 告警
    {
        SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
        Region: "cn-north-03",
        ResourceType: "uhost",
        ResourceId: "uhost-xxxx",
        MetricName: "MemUsage",
        AlarmTime: 1458733318,
        RecoveryTime: 0
    }
  
### 恢复
    {
        SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
        Region: "cn-north-03",
        ResourceType: "uhost",
        ResourceId: "uhost-xxxx",
        MetricName: "MemUsage",
        AlarmTime: 0,
        RecoveryTime: 1458733318
    }
  
**Field Explaination**
  

<table>
      <tr>
            <td>Field</td><td>Explaination</td>
      </tr>
      <tr>
            <td>SessionID</td><td>Session ID for this message</td>
      </tr>
      <tr>
            <td>Region</td><td>Region Name</td>
      </tr>
      <tr>
            <td>ResourceType</td><td>Resource Type</td>
      </tr>
      <tr>
            <td>ResourceId</td><td>short id</td>
      </tr>
      <tr>
            <td>MetricName</td><td>Metric Name for current warning</td>
      </tr>
      <tr>
            <td>AlarmTime</td><td>Alarm time</td>
      </tr>
      <tr>
            <td>RecoveryTime</td><td>Restory time</td>
      </tr>
      <tr>
            <td>Content</td><td>Warning content</td>
      </tr>
</table>

## Response
我们这边需要收到这样的response， 表明用户成功接收推送信息，否则会再重试2次：

    {
        SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
        RetCode: 0
    }

