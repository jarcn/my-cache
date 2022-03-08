package mongo

type ImMessage struct {
	_id           string `json:"name,_id" bson:"name,_id"`
	SignTimestamp int64  `json:"name,signTimestamp" bson:"name,signTimestamp"`
	Nonce         int64  `json:"name,nonce" bson:"name,nonce"`
	Signature     string `json:"name,signature" bson:"name,signature"`
	FromUserId    int64  `json:"name,fromUserId" bson:"name,fromUserId"`
	ToUserId      int64  `json:"name,toUserId" bson:"name,toUserId"`
	ObjectName    string `json:"name,objectName" bson:"name,objectName"`
	Content       string `json:"name,content" bson:"name,content"`
	ChannelType   string `json:"name,channelType" bson:"name,channelType"`
	MsgTimestamp  int64  `json:"name,msgTimestamp" bson:"name,msgTimestamp"`
	MsgUID        string `json:"name,msgUID" bson:"name,msgUID"`
	SensitiveType string `json:"name,sensitiveType" bson:"name,sensitiveType"`
	Source        string `json:"name,source" bson:"name,source"`
	_class        string `json:"name,_class" bson:"name,_class"`
	CreateTime    int64  `json:"name,createTime" bson:"name,createTime"`
	EmployeeId    int64  `json:"name,employeeId" bson:"name,employeeId"`
	EmployerId    int64  `json:"name,employerId" bson:"name,employerId"`
}
