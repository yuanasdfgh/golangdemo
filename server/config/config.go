package config

// 组合全部配置模型
type Config struct {
	Server       Server `mapstructure:"server"`
	Log          Log    `mapstructure:"log"`
	Mysql        Mysql  `mapstructure:"mysql"`
	Jwt          Jwt    `mapstructure:"jwt"`
	Cron         Cron   `mapstructure:"cron"`
	Code2Session `mapstructure:"code2Session"`
	Gpt          Gpt `mapstructure:"Jpt"`
}

// 服务端启动配置
type Server struct {
	Port string `mapstructure:"port"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Mysql struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

// 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}

// 定时任务配置
type Cron struct {
	Enable bool `mapstructure:"enable"`
}

// 微信小程序相关配置
type Code2Session struct {
	Code      string `mapstructure:"code"`
	AppId     string `mapstructure:"appId"`
	AppSecret string `mapstructure:"appSecret"`
}
type Gpt struct {
	Url string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}
