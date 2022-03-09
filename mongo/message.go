package mongo

type ImMessage struct {
	Id            string `json:"_id" bson:"_id"`
	SignTimestamp int64  `json:"signTimestamp" bson:"signTimestamp"`
	Nonce         int64  `json:"nonce" bson:"nonce"`
	Signature     string `json:"signature" bson:"signature"`
	FromUserId    int64  `json:"fromUserId" bson:"fromUserId"`
	ToUserId      int64  `json:"toUserId" bson:"toUserId"`
	ObjectName    string `json:"objectName" bson:"objectName"`
	Content       string `json:"content" bson:"content"`
	ChannelType   string `json:"channelType" bson:"channelType"`
	MsgTimestamp  int64  `json:"msgTimestamp" bson:"msgTimestamp"`
	MsgUID        string `json:"msgUID" bson:"msgUID"`
	SensitiveType string `json:"sensitiveType" bson:"sensitiveType"`
	Source        string `json:"source" bson:"source"`
	Class         string `json:"_class" bson:"_class"`
	CreateTime    int64  `json:"createTime" bson:"createTime"`
	EmployeeId    int64  `json:"employeeId" bson:"employeeId"`
	EmployerId    int64  `json:"employerId" bson:"employerId"`
}
