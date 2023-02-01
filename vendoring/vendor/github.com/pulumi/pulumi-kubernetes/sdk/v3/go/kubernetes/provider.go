// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the kubernetes package.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.EnableConfigMapMutable) {
		args.EnableConfigMapMutable = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE").(bool))
	}
	if isZero(args.EnableDryRun) {
		args.EnableDryRun = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_ENABLE_DRY_RUN").(bool))
	}
	if isZero(args.EnableReplaceCRD) {
		args.EnableReplaceCRD = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_ENABLE_REPLACE_CRD").(bool))
	}
	if isZero(args.EnableServerSideApply) {
		args.EnableServerSideApply = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_ENABLE_SERVER_SIDE_APPLY").(bool))
	}
	if args.HelmReleaseSettings != nil {
		args.HelmReleaseSettings = args.HelmReleaseSettings.ToHelmReleaseSettingsPtrOutput().ApplyT(func(v *HelmReleaseSettings) *HelmReleaseSettings { return v.Defaults() }).(HelmReleaseSettingsPtrOutput)
	}
	if args.KubeClientSettings != nil {
		args.KubeClientSettings = args.KubeClientSettings.ToKubeClientSettingsPtrOutput().ApplyT(func(v *KubeClientSettings) *KubeClientSettings { return v.Defaults() }).(KubeClientSettingsPtrOutput)
	}
	if isZero(args.Kubeconfig) {
		args.Kubeconfig = pulumi.StringPtr(getEnvOrDefault("", nil, "KUBECONFIG").(string))
	}
	if isZero(args.SuppressDeprecationWarnings) {
		args.SuppressDeprecationWarnings = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS").(bool))
	}
	if isZero(args.SuppressHelmHookWarnings) {
		args.SuppressHelmHookWarnings = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster *string `pulumi:"cluster"`
	// If present, the name of the kubeconfig context to use.
	Context *string `pulumi:"context"`
	// If present and set to true, the provider will delete resources associated with an unreachable Kubernetes cluster from Pulumi state
	DeleteUnreachable *bool `pulumi:"deleteUnreachable"`
	// BETA FEATURE - If present and set to true, allow ConfigMaps to be mutated.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways using this precedence:
	// 1. This `enableConfigMapMutable` parameter.
	// 2. The `PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE` environment variable.
	EnableConfigMapMutable *bool `pulumi:"enableConfigMapMutable"`
	// Deprecated. If present and set to true, enable server-side diff calculations.
	//
	// Deprecated: This option has been replaced by `enableServerSideApply`.
	EnableDryRun *bool `pulumi:"enableDryRun"`
	// Obsolete. This option has no effect.
	//
	// Deprecated: This option is deprecated, and will be removed in a future release.
	EnableReplaceCRD *bool `pulumi:"enableReplaceCRD"`
	// BETA FEATURE - If present and set to true, enable Server-Side Apply mode.
	// See https://github.com/pulumi/pulumi-kubernetes/issues/2011 for additional details.
	// This feature is in developer preview, and is disabled by default.
	EnableServerSideApply *bool `pulumi:"enableServerSideApply"`
	// Options to configure the Helm Release resource.
	HelmReleaseSettings *HelmReleaseSettings `pulumi:"helmReleaseSettings"`
	// Options for tuning the Kubernetes client used by a Provider.
	KubeClientSettings *KubeClientSettings `pulumi:"kubeClientSettings"`
	// The contents of a kubeconfig file or the path to a kubeconfig file.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace *string `pulumi:"namespace"`
	// BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
	// be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
	// to the Pulumi program. This feature is in developer preview, and is disabled by default.
	//
	// Note that some computed Outputs such as status fields will not be populated
	// since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
	// and may result in an error if they are referenced by other resources. Also note that any secret values
	// used in these resources will be rendered in plaintext to the resulting YAML.
	RenderYamlToDirectory *string `pulumi:"renderYamlToDirectory"`
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	SuppressDeprecationWarnings *bool `pulumi:"suppressDeprecationWarnings"`
	// If present and set to true, suppress unsupported Helm hook warnings from the CLI.
	SuppressHelmHookWarnings *bool `pulumi:"suppressHelmHookWarnings"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster pulumi.StringPtrInput
	// If present, the name of the kubeconfig context to use.
	Context pulumi.StringPtrInput
	// If present and set to true, the provider will delete resources associated with an unreachable Kubernetes cluster from Pulumi state
	DeleteUnreachable pulumi.BoolPtrInput
	// BETA FEATURE - If present and set to true, allow ConfigMaps to be mutated.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways using this precedence:
	// 1. This `enableConfigMapMutable` parameter.
	// 2. The `PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE` environment variable.
	EnableConfigMapMutable pulumi.BoolPtrInput
	// Deprecated. If present and set to true, enable server-side diff calculations.
	//
	// Deprecated: This option has been replaced by `enableServerSideApply`.
	EnableDryRun pulumi.BoolPtrInput
	// Obsolete. This option has no effect.
	//
	// Deprecated: This option is deprecated, and will be removed in a future release.
	EnableReplaceCRD pulumi.BoolPtrInput
	// BETA FEATURE - If present and set to true, enable Server-Side Apply mode.
	// See https://github.com/pulumi/pulumi-kubernetes/issues/2011 for additional details.
	// This feature is in developer preview, and is disabled by default.
	EnableServerSideApply pulumi.BoolPtrInput
	// Options to configure the Helm Release resource.
	HelmReleaseSettings HelmReleaseSettingsPtrInput
	// Options for tuning the Kubernetes client used by a Provider.
	KubeClientSettings KubeClientSettingsPtrInput
	// The contents of a kubeconfig file or the path to a kubeconfig file.
	Kubeconfig pulumi.StringPtrInput
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace pulumi.StringPtrInput
	// BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
	// be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
	// to the Pulumi program. This feature is in developer preview, and is disabled by default.
	//
	// Note that some computed Outputs such as status fields will not be populated
	// since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
	// and may result in an error if they are referenced by other resources. Also note that any secret values
	// used in these resources will be rendered in plaintext to the resulting YAML.
	RenderYamlToDirectory pulumi.StringPtrInput
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	SuppressDeprecationWarnings pulumi.BoolPtrInput
	// If present and set to true, suppress unsupported Helm hook warnings from the CLI.
	SuppressHelmHookWarnings pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}