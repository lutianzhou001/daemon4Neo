# Daemon Service



## APIs





### isActive

检查进程是否存活



#### 输入

{

  "jsonrpc": "2.0",

  "method": "isActive",

  "params": [],

  "id": 1

}



#### 输出

{

​    "jsonrpc": "2.0",

​    "id": 1,

​    "result": true

}



### restart

重启进程



#### 输入

{

  "jsonrpc": "2.0",

  "method": "restart",

  "params": [],

  "id": 1

}



#### 输出

{

​    "jsonrpc": "2.0",

​    "id": 1,

​    "result": true

}



### showLogs

#### 输入

{

  "jsonrpc": "2.0",

  "method": "showLogs",

  "params": [],

  "id": 1

}

#### 输出

{

​    "jsonrpc": "2.0",

​    "id": 1,

​    "result": [{id:1,message:"query data success",timestamp:1230313156},{id:2,message:"send data success",timestamp:13465465446},{id:3,message:"query data error",timestamp:123465465486}]

}





# Sync Service



### depositData

输入

{

  "jsonrpc": "2.0",

  "method": "depositData",

  "params": [{key: a, value:b},{key:c, value:d}],

  "id": 1

}

#### 输出

{

​    "jsonrpc": "2.0",

​    "id": 1,

​    "result": 2

}































