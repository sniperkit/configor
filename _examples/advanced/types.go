package main

type AppConfig struct {
	Name     string `json:"name" yaml:"name" toml:"name"`
	Timezone string `default:"Europe/Paris" json:"timezone" yaml:"timezone" toml:"timezone"`
	Version  string `env:"AppVersion" json:"version" yaml:"version" toml:"version"`
	Sandbox  bool   `env:"AppSandbox" default:"false" json:"sandbox_mode" yaml:"sandbox_mode" toml:"amazonPaySandbox"`
}

type ServerConfig struct {
	Host uint `default:"localhost" env:"HOST" json:"host" yaml:"host" toml:"host"`
	Port uint `default:"7000" env:"PORT" json:"port" yaml:"port" toml:"port"`
	TLS  bool `default:"false" env:"TLS" json:"tls" yaml:"tls" toml:"tls"`
}

type SMTPConfig struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	User     string `json:"user" yaml:"user" toml:"user"`
	Password string `json:"password" yaml:"password" toml:"password"`
}

type DBConfig struct {
	Name     string `env:"DBName" default:"qor_example" json:"db_name" yaml:"db_name" toml:"dbName"`
	Adapter  string `env:"DBAdapter" default:"mysql" json:"adapter" yaml:"adapter" toml:"adapter"`
	Host     string `env:"DBHost" default:"localhost" json:"host" yaml:"host" toml:"host"`
	Port     string `env:"DBPort" default:"3306" json:"port" yaml:"port" toml:"port"`
	User     string `env:"DBUser" json:"user" yaml:"user" toml:"user"`
	Password string `env:"DBPassword" json:"password" yaml:"password" toml:"password"`
}

type S3Config struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" json:"access_key" yaml:"access_key" toml:"accessKey"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" json:"secret_key" yaml:"secret_key" toml:"secretKey"`
	Region          string `env:"AWS_Region" json:"region" yaml:"region" toml:"region"`
	S3Bucket        string `env:"AWS_Bucket" json:"bucket" yaml:"bucket" toml:"bucket"`
}

type AmazonPayConfig struct {
	MerchantID   string `env:"AmazonPayMerchantID" json:"merchant_id" yaml:"merchant_id" toml:"amazonPayMerchantID"`
	AccessKey    string `env:"AmazonPayAccessKey" json:"access_key" yaml:"access_key" toml:"amazonPayAccessKey"`
	SecretKey    string `env:"AmazonPaySecretKey" json:"secret_key" yaml:"secret_key" toml:"amazonPaySecretKey"`
	ClientID     string `env:"AmazonPayClientID" json:"client_id" yaml:"client_id" toml:"amazonPayClientID"`
	ClientSecret string `env:"AmazonPayClientSecret" json:"client_secret" yaml:"client_secret" toml:"amazonPayClientSecret"`
	CurrencyCode string `env:"AmazonPayCurrencyCode" default:"EUR" json:"currency_code" yaml:"currency_code" toml:"amazonPayCurrencyCode"`
	Sandbox      bool   `env:"AmazonPaySandbox" default:"false" json:"sandbox_mode" yaml:"sandbox_mode" toml:"amazonPaySandbox"`
}

type AWSConfig struct {
	AmazonS3  S3Config        `json:"amazon_s3" yaml:"amazon_s3" toml:"amazonS3"`
	AmazonPay AmazonPayConfig `json:"amazon_pay" yaml:"amazon_pay" toml:"amazonPay"`
}

type ServiceConfig struct {
	AWS   AWSConfig          `json:"amazon_aws" yaml:"amazon_aws" toml:"amazon_aws"`
	SMTP  SMTPConfig         `json:"smtp" yaml:"smtp" toml:"smtp"`
	API   APIProvidersConfig `json:"api" yaml:"api" toml:"api"`
	OAuth OAuthConfig        `json:"oauth" yaml:"oauth" toml:"oauth"`
}

type OAuthConfig struct {
	Provider OAuthProvidersConfig `json:"provider" yaml:"provider" toml:"provider"`
}

type OAuthProvidersConfig struct {
	Github   OAuthProviderConfig `json:"github" yaml:"github" toml:"github"`
	Google   OAuthProviderConfig `json:"google" yaml:"google" toml:"google"`
	Facebook OAuthProviderConfig `json:"facebook" yaml:"facebook" toml:"facebook"`
	Twitter  OAuthProviderConfig `json:"twitter" yaml:"twitter" toml:"twitter"`
}

type PaypalConfig struct {
	Email   string `json:"email" yaml:"email" toml:"email"`
	Sandbox bool   `env:"PaypalSandbox" default:"false" json:"sandbox_mode" yaml:"sandbox_mode" toml:"paypalSandbox"`
}

type ContactConfig struct {
	Email string `json:"client_id" yaml:"client_id" toml:"clientID"`
}

// Config github Config
type OAuthProviderConfig struct {
	ClientID     string   `json:"client_id" yaml:"client_id" toml:"clientID"`
	ClientSecret string   `json:"client_secret" yaml:"client_secret" toml:"clientSecret"`
	AuthorizeURL string   `json:"-" yaml:"-" toml:"-"` // `json:"authorize_url" yaml:"authorize_url" toml:"twitter"`
	TokenURL     string   `json:"-" yaml:"-" toml:"-"` // `json:"token_url" yaml:"token_url" toml:"tokenURL"`
	RedirectURL  string   `json:"-" yaml:"-" toml:"-"` // `json:"redirect_url" yaml:"redirect_url" toml:"redirectURL"`
	Scopes       []string `json:"-" yaml:"-" toml:"-"` // !!! Important: yaml decoder cannot parse string slices
}

type APIProvidersConfig struct {
	GoogleAPIKey string `env:"GoogleAPIKey" json:"google_api_key" yaml:"google_api_key" toml:"googleAPIKey"`
	BaiduAPIKey  string `env:"BaiduAPIKey" json:"baidu_api_key" yaml:"baidu_api_key" toml:"baiduAPIKey"`
}
