# User Guide

## WebHook 说明
创建webhook后您的系统可以收到UCloud告警信息。当有告警时，UCloud的系统会将告警信息以HTTP POST方式发送到指定URL，您可以对收到信息进行处理。

### 先决条件
用户需要提供接收POST请求的HTTP服务，以处理UCloud发送的POST请求，并将该服务的URL注册到UCloud的告警系统中。

### JSON Body Example
  
#### 告警
    {
        SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
        Region: "cn-north-03",
        ResourceType: "uhost",
        ResourceId: "uhost-xxxx",
        MetricName: "MemUsage",
        AlarmTime: 1458733318,
        RecoveryTime: 0
    }
  
#### 恢复
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

### Response
我们这边需要收到这样的response， 表明用户成功接收推送信息，否则会再重试2次：

    {
        SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
        RetCode: 0
    }


## WebHook Demo 说明
本Demo提供实现支持WebHook服务的一种方式，基于Golang实现，使用MySQL存储收到的报警数据，Demo仅供参考。

### 依赖
WebHook Demo依赖3个第三方包，编译前需要使用go get 命令安装到本地

1. github.com/gorilla/mux
2. github.com/google/uuid
3. github.com/go-sql-driver/mysql

### 编译
1. 在目录warn-webhook中创建名为src指向webhook-demo-go的软连接
2. 将目录warn-webhook添加到GOPATH；
3. 在目录warn-webhook中执行：*go build -a . *

### 执行
1. 初始化MySQL数据库及表，表Schema可参考如下：

        CREATE TABLE `warn_message` (
        `session_id` char(36) NOT NULL,
        `region` varchar(20) NOT NULL,
        `resource_type` varchar(45) NOT NULL,
        `resource_id` varchar(45) NOT NULL,
        `metric_name` varchar(50) NOT NULL,
        `alarm_time` int(11) DEFAULT NULL,
        `recovery_time` int(11) DEFAULT '0',
        PRIMARY KEY (`session_id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

2. 根据数据库表相关信息，修改webhook-demo-go/etc/conf.ini文件中相关内容
3. 执行编译生成的可执行程序，默认加载etc/conf.ini文件，也可以指定其他位置配置文件

### 资源路径及接口
服务器启动时默认端口为80

1. / -> User Guide
2. /add -> 添加报警信息，即实际WebHook调用接口
3. /get -> 获取当前已经添加的报警信息

### 客户端说明
仅测试添加报警信息，实际数据请以接口实际返回为准。

