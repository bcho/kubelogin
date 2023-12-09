package kubelogin

import "github.com/Azure/kubelogin/internal/token"

// login method constants

const (
	DeviceCodeLogin = token.DeviceCodeLogin
	InteractiveLogin = token.InteractiveLogin
	ServicePrincipalLogin = token.ServicePrincipalLogin
	ROPCLogin = token.ROPCLogin
	MSILogin = token.MSILogin
	AzureCLILogin = token.AzureCLILogin
	WorkloadIdentityLogin = token.WorkloadIdentityLogin
)