// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x6D6.proto

package oidb

// Group File Upload
type OidbSvcTrpcTcp0X6D6 struct {
	File     *OidbSvcTrpcTcp0X6D6Upload   `protobuf:"bytes,1,opt"`
	Download *OidbSvcTrpcTcp0X6D6Download `protobuf:"bytes,3,opt"`
	Delete   *OidbSvcTrpcTcp0X6D6Delete   `protobuf:"bytes,4,opt"`
	Rename   *OidbSvcTrpcTcp0X6D6Rename   `protobuf:"bytes,5,opt"`
	Move     *OidbSvcTrpcTcp0X6D6Move     `protobuf:"bytes,6,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D6Upload struct {
	GroupUin        uint32 `protobuf:"varint,1,opt"`
	AppId           uint32 `protobuf:"varint,2,opt"` // 7
	BusId           uint32 `protobuf:"varint,3,opt"` // 102
	Entrance        uint32 `protobuf:"varint,4,opt"` // 6
	TargetDirectory string `protobuf:"bytes,5,opt"`
	FileName        string `protobuf:"bytes,6,opt"`
	LocalDirectory  string `protobuf:"bytes,7,opt"`
	FileSize        uint64 `protobuf:"varint,8,opt"`
	FileSha1        []byte `protobuf:"bytes,9,opt"`
	FileSha3        []byte `protobuf:"bytes,10,opt"`
	FileMd5         []byte `protobuf:"bytes,11,opt"`
	Field15         bool   `protobuf:"varint,15,opt"`
}

type OidbSvcTrpcTcp0X6D6Download struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	AppID    uint32 `protobuf:"varint,2,opt"` // 7
	BusID    uint32 `protobuf:"varint,3,opt"` // 102
	FileId   string `protobuf:"bytes,4,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D6Delete struct {
	GroupUin uint32 `protobuf:"varint,1,opt"`
	BusId    uint32 `protobuf:"varint,3,opt"`
	FileId   string `protobuf:"bytes,5,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D6Rename struct {
	GroupUin     uint32 `protobuf:"varint,1,opt"`
	BusId        uint32 `protobuf:"varint,3,opt"` // 102
	FileId       string `protobuf:"bytes,4,opt"`
	ParentFolder string `protobuf:"bytes,5,opt"`
	NewFileName  string `protobuf:"bytes,6,opt"`
	_            [0]func()
}

type OidbSvcTrpcTcp0X6D6Move struct {
	GroupUin        uint32 `protobuf:"varint,1,opt"`
	AppId           uint32 `protobuf:"varint,2,opt"` // 7
	BusId           uint32 `protobuf:"varint,3,opt"` // 102
	FileId          string `protobuf:"bytes,4,opt"`
	ParentDirectory string `protobuf:"bytes,5,opt"`
	TargetDirectory string `protobuf:"bytes,6,opt"`
	_               [0]func()
}

type OidbSvcTrpcTcp0X6D6Response struct {
	Upload   *OidbSvcTrpcTcp0X6D6_0Response     `protobuf:"bytes,1,opt"`
	Download *OidbSvcTrpcTcp0X6D6_2Response     `protobuf:"bytes,3,opt"`
	Delete   *OidbSvcTrpcTcp0X6D6_3_4_5Response `protobuf:"bytes,4,opt"`
	Rename   *OidbSvcTrpcTcp0X6D6_3_4_5Response `protobuf:"bytes,5,opt"`
	Move     *OidbSvcTrpcTcp0X6D6_3_4_5Response `protobuf:"bytes,6,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D6_0Response struct {
	RetCode       int32    `protobuf:"varint,1,opt"`
	RetMsg        string   `protobuf:"bytes,2,opt"`
	ClientWording string   `protobuf:"bytes,3,opt"`
	UploadIp      string   `protobuf:"bytes,4,opt"`
	ServerDns     string   `protobuf:"bytes,5,opt"`
	BusId         int32    `protobuf:"varint,6,opt"`
	FileId        string   `protobuf:"bytes,7,opt"`
	CheckKey      []byte   `protobuf:"bytes,8,opt"`
	FileKey       []byte   `protobuf:"bytes,9,opt"`
	BoolFileExist bool     `protobuf:"varint,10,opt"`
	UploadIpLanV4 []string `protobuf:"bytes,12,rep"`
	UploadIpLanV6 []string `protobuf:"bytes,13,rep"`
	UploadPort    uint32   `protobuf:"varint,14,opt"`
}

type OidbSvcTrpcTcp0X6D6_2Response struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
	DownloadIp    string `protobuf:"bytes,4,opt"`
	DownloadDns   string `protobuf:"bytes,5,opt"`
	DownloadUrl   []byte `protobuf:"bytes,6,opt"`
	FileSha1      []byte `protobuf:"bytes,7,opt"`
	FileSha3      []byte `protobuf:"bytes,8,opt"`
	FileMd5       []byte `protobuf:"bytes,9,opt"`
	CookieVal     []byte `protobuf:"bytes,10,opt"`
	SaveFileName  string `protobuf:"bytes,11,opt"`
	PreviewPort   uint32 `protobuf:"varint,12,opt"`
}

type OidbSvcTrpcTcp0X6D6_3_4_5Response struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
	_             [0]func()
}
