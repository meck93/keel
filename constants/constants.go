package constants

// DefaultDockerRegistry - default docker registry
const DefaultDockerRegistry = "https://index.docker.io"

// DefaultNamespace - default namespace to initialise configmaps based kv
const DefaultNamespace = "kube-system"

// WebhookEndpointEnv if set - enables webhook notifications
const WebhookEndpointEnv = "WEBHOOK_ENDPOINT"

// slack bot/token
const (
	EnvSlackBotToken         = "SLACK_BOT_TOKEN"
	EnvSlackAppToken         = "SLACK_APP_TOKEN"
	EnvSlackBotName          = "SLACK_BOT_NAME"
	EnvSlackChannels         = "SLACK_CHANNELS"
	EnvSlackApprovalsChannel = "SLACK_APPROVALS_CHANNEL"

	// Mail notification settings
	EnvMailTo         = "MAIL_TO"
	EnvMailFrom       = "MAIL_FROM"
	EnvMailSmtpServer = "MAIL_SMTP_SERVER"
	EnvMailSmtpPort   = "MAIL_SMTP_PORT"
	EnvMailSmtpUser   = "MAIL_SMTP_USER"
	EnvMailSmtpPass   = "MAIL_SMTP_PASS"
)

// EnvNotificationLevel - minimum level for notifications, defaults to info
const EnvNotificationLevel = "NOTIFICATION_LEVEL"

// Basic Auth - User / Password
const EnvBasicAuthUser = "BASIC_AUTH_USER"
const EnvBasicAuthPassword = "BASIC_AUTH_PASSWORD"
const EnvAuthenticatedWebhooks = "AUTHENTICATED_WEBHOOKS"
const EnvTokenSecret = "TOKEN_SECRET"

// KeelLogoURL - is a logo URL for bot icon
const KeelLogoURL = "https://keel.sh/img/logo.png"

// Env var to define a namespace that keel will scan - avoid scan over all the cluster -
const EnvRestrictedNamespace = "RESTRICTED_NAMESPACE"
