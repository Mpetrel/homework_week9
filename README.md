### Socket 粘包的解包方式总结
#### 1. fix length
通过发送固定长度的数据，接受方根据消息长度来判断是否可作为一个消息处理
#### 2. delimiter based
包尾加上分隔符来标记，如`FTP`的`\r\n`，但数据中包含同样的分隔符则容易误判
#### 3. length field based frame decoder
通过在包头上加入包体长度，接收方获取包体长度后，再根据该长度来接受包体，如`HTTP`协议header中的`Content-Length`
   