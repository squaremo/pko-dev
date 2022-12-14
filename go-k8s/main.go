package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	// arbitrary stuff
	_ "github.com/hashicorp/go-retryablehttp"
	_ "github.com/hashicorp/go-tfe"
	_ "github.com/hashicorp/vault/api"
	_ "github.com/hashicorp/vault/api/auth/kubernetes"
	_ "github.com/imdario/mergo"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	_ "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	_ "github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
	_ "k8s.io/apimachinery"
	_ "k8s.io/client-go"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		appLabels := pulumi.StringMap{
			"app": pulumi.String("nginx"),
		}
		deployment, err := appsv1.NewDeployment(ctx, "app-dep", &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String("www"),
								Image: pulumi.String("nginx"),
							}},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("name", deployment.Metadata.Elem().Name())

		return nil
	})
}
