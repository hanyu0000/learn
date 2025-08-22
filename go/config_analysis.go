package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 1. 日志配置结构体
type LogConfig struct {
	Level      string `yaml:"level"`       // 日志等级
	Format     string `yaml:"format"`      // 写入格式
	ShowLine   bool   `yaml:"show_line"`   // 是否显示调用行
	RootDir    string `yaml:"root_dir"`    // 日志根目录
	Filename   string `yaml:"filename"`    // 日志文件名称
	MaxSize    int    `yaml:"max_size"`    // 日志文件最大大小（MB）
	MaxBackups int    `yaml:"max_backups"` // 旧文件的最大个数
	MaxAge     int    `yaml:"max_age"`     // 旧文件的最大保留天数
	Compress   bool   `yaml:"compress"`    // 是否压缩
	EnableFile bool   `yaml:"enable_file"` // 是否启用日志文件
}

// 2. Redis配置结构体
type RedisConfig struct {
	Host     string `yaml:"host"`     // 服务地址
	Port     int    `yaml:"port"`     // 服务端口
	DB       int    `yaml:"db"`       // 库选择
	Password string `yaml:"password"` // 密码
}

// 3. PostgreSQL配置结构体
type PostgreSQLConfig struct {
	Host         string `yaml:"host"`           // 服务地址
	Port         int    `yaml:"port"`           // 端口号
	Username     string `yaml:"username"`       // 用户名
	Password     string `yaml:"password"`       // 密码
	Database     string `yaml:"database"`       // 数据库名称
	MaxIdleConns int    `yaml:"max_idle_conns"` // 空闲连接池中连接的最大数量
	MaxOpenConns int    `yaml:"max_open_conns"` // 打开数据库连接的最大数量
}

// 4. MinIO配置结构体
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`          // 服务地址
	AccessKeyID     string `yaml:"access_key_id"`     // 用户名
	SecretAccessKey string `yaml:"secret_access_key"` // 密码
	UseSSL          bool   `yaml:"use_ssl"`           // 是否启用加密
}

// 5. Etcd配置结构体
type EtcdConfig struct {
	Endpoint    string `yaml:"endpoint"`     // 服务地址
	DialTimeout int    `yaml:"dial_timeout"` // 连接超时时间
	TTL         int    `yaml:"ttl"`          // 租约时间
}

// 6. 代理配置结构体
type ProxyConfig struct {
	MinioNodeName string `yaml:"minio_node_name"` // minio 节点名称
	ProxyAddr     string `yaml:"proxy_addr"`      // 代理地址
}

// 7. 一致性哈希配置结构体
type ConsistentHashConfig struct {
	Replicas int      `yaml:"replicas"` // 虚拟节点数
	Nodes    []string `yaml:"nodes"`    // 节点列表
}

// 8. 主配置结构体
type Config struct {
	LogConfig            LogConfig            `yaml:"log"`
	RedisConfig          RedisConfig          `yaml:"redis"`
	PgDBConfig           PostgreSQLConfig     `yaml:"postgresql"`
	MinioConfig          MinioConfig          `yaml:"minio"`
	EtcdConfig           EtcdConfig           `yaml:"etcd"`
	ProxyConfig          ProxyConfig          `yaml:"proxy"`
	ConsistentHashConfig ConsistentHashConfig `yaml:"consistent_hash"`
}

// 配置加载函数
func loadConfig(filename string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("解析YAML失败: %v", err)
	}

	return &config, nil
}

// 配置验证函数
func validateConfig(config *Config) error {
	fmt.Println("=== 配置验证 ===")

	// 验证日志配置
	if config.LogConfig.Level == "" {
		return fmt.Errorf("日志等级不能为空")
	}
	fmt.Printf("日志配置: 等级=%s, 格式=%s, 文件=%s\n",
		config.LogConfig.Level, config.LogConfig.Format, config.LogConfig.Filename)

	// 验证Redis配置
	if config.RedisConfig.Host == "" {
		return fmt.Errorf("Redis主机地址不能为空")
	}
	fmt.Printf("Redis配置: 主机=%s, 端口=%d, 数据库=%d\n",
		config.RedisConfig.Host, config.RedisConfig.Port, config.RedisConfig.DB)

	// 验证PostgreSQL配置
	if config.PgDBConfig.Host == "" {
		return fmt.Errorf("PostgreSQL主机地址不能为空")
	}
	fmt.Printf("PostgreSQL配置: 主机=%s, 端口=%d, 数据库=%s\n",
		config.PgDBConfig.Host, config.PgDBConfig.Port, config.PgDBConfig.Database)

	// 验证MinIO配置
	if config.MinioConfig.Endpoint == "" {
		return fmt.Errorf("MinIO端点不能为空")
	}
	fmt.Printf("MinIO配置: 端点=%s, 用户=%s\n",
		config.MinioConfig.Endpoint, config.MinioConfig.AccessKeyID)

	// 验证Etcd配置
	if config.EtcdConfig.Endpoint == "" {
		return fmt.Errorf("Etcd端点不能为空")
	}
	fmt.Printf("Etcd配置: 端点=%s, 超时=%d秒\n",
		config.EtcdConfig.Endpoint, config.EtcdConfig.DialTimeout)

	// 验证代理配置
	if config.ProxyConfig.ProxyAddr == "" {
		return fmt.Errorf("代理地址不能为空")
	}
	fmt.Printf("代理配置: 节点=%s, 地址=%s\n",
		config.ProxyConfig.MinioNodeName, config.ProxyConfig.ProxyAddr)

	// 验证一致性哈希配置
	if len(config.ConsistentHashConfig.Nodes) == 0 {
		return fmt.Errorf("一致性哈希节点列表不能为空")
	}
	fmt.Printf("一致性哈希配置: 虚拟节点数=%d, 节点列表=%v\n",
		config.ConsistentHashConfig.Replicas, config.ConsistentHashConfig.Nodes)

	return nil
}

// 配置使用示例
func configUsageExamples(config *Config) {
	fmt.Println("\n=== 配置使用示例 ===")

	// 1. 构建Redis连接字符串
	redisConnStr := fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port)
	fmt.Printf("Redis连接字符串: %s\n", redisConnStr)

	// 2. 构建PostgreSQL连接字符串
	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PgDBConfig.Host, config.PgDBConfig.Port, config.PgDBConfig.Username,
		config.PgDBConfig.Password, config.PgDBConfig.Database)
	fmt.Printf("PostgreSQL连接字符串: %s\n", pgConnStr)

	// 3. 构建MinIO连接信息
	minioInfo := fmt.Sprintf("端点: %s, 用户: %s",
		config.MinioConfig.Endpoint, config.MinioConfig.AccessKeyID)
	fmt.Printf("MinIO连接信息: %s\n", minioInfo)

	// 4. 构建Etcd连接信息
	etcdInfo := fmt.Sprintf("端点: %s, 超时: %d秒",
		config.EtcdConfig.Endpoint, config.EtcdConfig.DialTimeout)
	fmt.Printf("Etcd连接信息: %s\n", etcdInfo)

	// 5. 显示代理信息
	proxyInfo := fmt.Sprintf("节点: %s, 地址: %s",
		config.ProxyConfig.MinioNodeName, config.ProxyConfig.ProxyAddr)
	fmt.Printf("代理信息: %s\n", proxyInfo)

	// 6. 显示一致性哈希节点
	fmt.Printf("一致性哈希节点数量: %d\n", len(config.ConsistentHashConfig.Nodes))
	for i, node := range config.ConsistentHashConfig.Nodes {
		fmt.Printf("  节点 %d: %s\n", i+1, node)
	}
}

// YAML结构体标签说明
func yamlTagExplanation() {
	fmt.Println("\n=== YAML结构体标签说明 ===")

	// 示例结构体
	type Example struct {
		Field1 string `yaml:"field1"`           // 字段映射到YAML中的field1
		Field2 int    `yaml:"field2,omitempty"` // omitempty表示空值时省略
		Field3 bool   `yaml:"-"`                // 减号表示不序列化此字段
	}

	fmt.Println("YAML标签说明:")
	fmt.Println("- `yaml:\"field1\"` - 将Go字段映射到YAML字段")
	fmt.Println("- `yaml:\"field2,omitempty\"` - 空值时省略该字段")
	fmt.Println("- `yaml:\"-\"` - 不序列化此字段")
	fmt.Println("- 没有标签时，使用字段名作为YAML键")
}

// 配置热重载示例
func configHotReloadExample() {
	fmt.Println("\n=== 配置热重载示例 ===")

	// 模拟配置文件变化
	config1 := &Config{
		LogConfig: LogConfig{
			Level:  "info",
			Format: "console",
		},
		RedisConfig: RedisConfig{
			Host: "127.0.0.1",
			Port: 6379,
		},
	}

	config2 := &Config{
		LogConfig: LogConfig{
			Level:  "debug", // 日志等级改变
			Format: "json",  // 格式改变
		},
		RedisConfig: RedisConfig{
			Host: "192.168.1.100", // 主机地址改变
			Port: 6380,            // 端口改变
		},
	}

	fmt.Println("配置1:")
	fmt.Printf("  日志等级: %s, 格式: %s\n", config1.LogConfig.Level, config1.LogConfig.Format)
	fmt.Printf("  Redis主机: %s, 端口: %d\n", config1.RedisConfig.Host, config1.RedisConfig.Port)

	fmt.Println("配置2:")
	fmt.Printf("  日志等级: %s, 格式: %s\n", config2.LogConfig.Level, config2.LogConfig.Format)
	fmt.Printf("  Redis主机: %s, 端口: %d\n", config2.RedisConfig.Host, config2.RedisConfig.Port)

	fmt.Println("配置变化检测完成")
}

func main() {
	fmt.Println("项目配置文件分析")
	fmt.Println("================")

	// 1. YAML标签说明
	yamlTagExplanation()

	// 2. 加载配置文件
	config, err := loadConfig("conf/config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 3. 验证配置
	err = validateConfig(config)
	if err != nil {
		log.Fatalf("配置验证失败: %v", err)
	}

	// 4. 配置使用示例
	configUsageExamples(config)

	// 5. 配置热重载示例
	configHotReloadExample()

	fmt.Println("\n================")
	fmt.Println("配置文件分析完成")
}
