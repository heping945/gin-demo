package global

import "time"

type Server struct {
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	AliYunOSS AliYunOSS `mapstructure:"aliyun_oss" json:"aliyun_oss" yaml:"aliyun_oss"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	UseMultipoint bool          `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Env           string        `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int           `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string        `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	ReadTimeout   time.Duration `mapstructure:"read-timeout" json:"ReadTimeout" yaml:"read-timeout"`
	WriteTimeout  time.Duration `mapstructure:"write-timeout" json:"WriteTimeout" yaml:"write-timeout"`
}

type AliYunOSS struct {
	Endpoint        string `json:"endpoint" mapstructure:"endpoint" `
	AccessKeyID     string `json:"access_key_id" mapstructure:"access_key_id" `
	Domain          string `json:"domain" mapstructure:"domain" `
	AccessKeySecret string `json:"access_key_secret" mapstructure:"access_key_secret" `
	BucketName      string `json:"bucket_name" mapstructure:"bucket_name" `
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Casbin struct {
	ModelPath string `mapstructure:"dao-path" json:"modelPath" yaml:"dao-path"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

// 日志输出前缀，输出文件，控制台输出
type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

// mapstructure这个属性和.yaml文件的字段会做映射
