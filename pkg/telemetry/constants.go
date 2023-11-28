package telemetry

const (
	// SegmentIO constants
	// SegmentIOWriteKey The write key is the unique identifier for a source that tells Segment which source data comes
	// from, to which workspace the data belongs, and which destinations should receive the data.
	SegmentIOWriteKey = "0gAYkX5RV3vt7s4pqCOOsDb6WHPLT30M"

	// Heartbeat
	KubefirstHeartbeat = "kubefirst.heartbeat"

	// Install
	KubefirstInstalled = "kubefirst.installed"

	// Init
	InitStarted   = "kubefirst.init.started"
	InitCompleted = "kubefirst.init.completed"

	CloudCredentialsCheckStarted   = "kubefirst.init.cloud_credentials_check.started"
	CloudCredentialsCheckCompleted = "kubefirst.init.cloud_credentials_check.completed"
	CloudCredentialsCheckFailed    = "kubefirst.init.cloud_credentials_check.failed"

	DomainLivenessStarted   = "kubefirst.init.domain_liveness.started"
	DomainLivenessCompleted = "kubefirst.init.domain_liveness.completed"
	DomainLivenessFailed    = "kubefirst.init.domain_liveness.failed"

	StateStoreCreateStarted   = "kubefirst.init.state_store_create.started"
	StateStoreCreateCompleted = "kubefirst.init.state_store_create.completed"
	StateStoreCreateFailed    = "kubefirst.init.state_store_create.failed"

	StateStoreCredentialsCreateStarted   = "kubefirst.init.state_store_credentials_create.started"
	StateStoreCredentialsCreateCompleted = "kubefirst.init.state_store_credentials_create.completed"
	StateStoreCredentialsCreateFailed    = "kubefirst.init.state_store_credentials_create.failed"

	GitCredentialsCheckStarted   = "kubefirst.init.git_credentials_check.started"
	GitCredentialsCheckCompleted = "kubefirst.init.git_credentials_check.completed"
	GitCredentialsCheckFailed    = "kubefirst.init.git_credentials_check.failed"

	KbotSetupStarted   = "kubefirst.init.kbot_setup.started"
	KbotSetupCompleted = "kubefirst.init.kbot_setup.completed"
	KbotSetupFailed    = "kubefirst.init.kbot_setup.failed"

	// Create
	// note: as of kubefirst 2.3.6, there will no longer a kubefirst.cluster_install.started event
	// as all installs are conducted through helm
	// ClusterInstallStarted   = "kubefirst.cluster_install.started" # permanently removed
	ClusterInstallCompleted = "kubefirst.cluster_install.completed"
	ClusterInstallFailed    = "kubefirst.cluster_install.failed"

	GitTerraformApplyStarted   = "kubefirst.git_terraform_apply.started"
	GitTerraformApplyCompleted = "kubefirst.git_terraform_apply.completed"
	GitTerraformApplyFailed    = "kubefirst.git_terraform_apply.failed"

	GitopsRepoPushStarted   = "kubefirst.gitops_repo_push.started"
	GitopsRepoPushCompleted = "kubefirst.gitops_repo_push.completed"
	GitopsRepoPushFailed    = "kubefirst.gitops_repo_push.failed"

	CloudTerraformApplyStarted   = "kubefirst.cloud_terraform_apply.started"
	CloudTerraformApplyCompleted = "kubefirst.cloud_terraform_apply.completed"
	CloudTerraformApplyFailed    = "kubefirst.cloud_terraform_apply.failed"

	ArgoCDInstallStarted   = "kubefirst.argocd_install.started"
	ArgoCDInstallCompleted = "kubefirst.argocd_install.completed"
	ArgoCDInstallFailed    = "kubefirst.argocd_install.failed"

	CreateRegistryStarted   = "kubefirst.create_registry.started"
	CreateRegistryCompleted = "kubefirst.create_registry.completed"
	CreateRegistryFailed    = "kubefirst.create_registry.failed"

	VaultInitializationStarted   = "kubefirst.vault_initialization.started"
	VaultInitializationCompleted = "kubefirst.vault_initialization.completed"
	VaultInitializationFailed    = "kubefirst.vault_initialization.failed"

	VaultTerraformApplyStarted   = "kubefirst.vault_terraform_apply.started"
	VaultTerraformApplyCompleted = "kubefirst.vault_terraform_apply.completed"
	VaultTerraformApplyFailed    = "kubefirst.vault_terraform_apply.failed"

	UsersTerraformApplyStarted   = "kubefirst.users_terraform_apply.started"
	UsersTerraformApplyCompleted = "kubefirst.users_terraform_apply.completed"
	UsersTerraformApplyFailed    = "kubefirst.users_terraform_apply.failed"

	// Delete
	ClusterDeleteStarted   = "kubefirst.cluster_delete.started"
	ClusterDeleteCompleted = "kubefirst.cluster_delete.completed"
)
