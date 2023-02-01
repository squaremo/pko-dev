// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LimitRangeList is a list of LimitRange items.
type LimitRangeList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Items LimitRangeTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewLimitRangeList registers a new resource with the given unique name, arguments, and options.
func NewLimitRangeList(ctx *pulumi.Context,
	name string, args *LimitRangeListArgs, opts ...pulumi.ResourceOption) (*LimitRangeList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("LimitRangeList")
	var resource LimitRangeList
	err := ctx.RegisterResource("kubernetes:core/v1:LimitRangeList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLimitRangeList gets an existing LimitRangeList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLimitRangeList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LimitRangeListState, opts ...pulumi.ResourceOption) (*LimitRangeList, error) {
	var resource LimitRangeList
	err := ctx.ReadResource("kubernetes:core/v1:LimitRangeList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LimitRangeList resources.
type limitRangeListState struct {
}

type LimitRangeListState struct {
}

func (LimitRangeListState) ElementType() reflect.Type {
	return reflect.TypeOf((*limitRangeListState)(nil)).Elem()
}

type limitRangeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Items []LimitRangeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a LimitRangeList resource.
type LimitRangeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Items LimitRangeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (LimitRangeListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*limitRangeListArgs)(nil)).Elem()
}

type LimitRangeListInput interface {
	pulumi.Input

	ToLimitRangeListOutput() LimitRangeListOutput
	ToLimitRangeListOutputWithContext(ctx context.Context) LimitRangeListOutput
}

func (*LimitRangeList) ElementType() reflect.Type {
	return reflect.TypeOf((**LimitRangeList)(nil)).Elem()
}

func (i *LimitRangeList) ToLimitRangeListOutput() LimitRangeListOutput {
	return i.ToLimitRangeListOutputWithContext(context.Background())
}

func (i *LimitRangeList) ToLimitRangeListOutputWithContext(ctx context.Context) LimitRangeListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeListOutput)
}

// LimitRangeListArrayInput is an input type that accepts LimitRangeListArray and LimitRangeListArrayOutput values.
// You can construct a concrete instance of `LimitRangeListArrayInput` via:
//
//	LimitRangeListArray{ LimitRangeListArgs{...} }
type LimitRangeListArrayInput interface {
	pulumi.Input

	ToLimitRangeListArrayOutput() LimitRangeListArrayOutput
	ToLimitRangeListArrayOutputWithContext(context.Context) LimitRangeListArrayOutput
}

type LimitRangeListArray []LimitRangeListInput

func (LimitRangeListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LimitRangeList)(nil)).Elem()
}

func (i LimitRangeListArray) ToLimitRangeListArrayOutput() LimitRangeListArrayOutput {
	return i.ToLimitRangeListArrayOutputWithContext(context.Background())
}

func (i LimitRangeListArray) ToLimitRangeListArrayOutputWithContext(ctx context.Context) LimitRangeListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeListArrayOutput)
}

// LimitRangeListMapInput is an input type that accepts LimitRangeListMap and LimitRangeListMapOutput values.
// You can construct a concrete instance of `LimitRangeListMapInput` via:
//
//	LimitRangeListMap{ "key": LimitRangeListArgs{...} }
type LimitRangeListMapInput interface {
	pulumi.Input

	ToLimitRangeListMapOutput() LimitRangeListMapOutput
	ToLimitRangeListMapOutputWithContext(context.Context) LimitRangeListMapOutput
}

type LimitRangeListMap map[string]LimitRangeListInput

func (LimitRangeListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LimitRangeList)(nil)).Elem()
}

func (i LimitRangeListMap) ToLimitRangeListMapOutput() LimitRangeListMapOutput {
	return i.ToLimitRangeListMapOutputWithContext(context.Background())
}

func (i LimitRangeListMap) ToLimitRangeListMapOutputWithContext(ctx context.Context) LimitRangeListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitRangeListMapOutput)
}

type LimitRangeListOutput struct{ *pulumi.OutputState }

func (LimitRangeListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LimitRangeList)(nil)).Elem()
}

func (o LimitRangeListOutput) ToLimitRangeListOutput() LimitRangeListOutput {
	return o
}

func (o LimitRangeListOutput) ToLimitRangeListOutputWithContext(ctx context.Context) LimitRangeListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LimitRangeListOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LimitRangeList) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Items is a list of LimitRange objects. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
func (o LimitRangeListOutput) Items() LimitRangeTypeArrayOutput {
	return o.ApplyT(func(v *LimitRangeList) LimitRangeTypeArrayOutput { return v.Items }).(LimitRangeTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LimitRangeListOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LimitRangeList) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LimitRangeListOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v *LimitRangeList) metav1.ListMetaPtrOutput { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

type LimitRangeListArrayOutput struct{ *pulumi.OutputState }

func (LimitRangeListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LimitRangeList)(nil)).Elem()
}

func (o LimitRangeListArrayOutput) ToLimitRangeListArrayOutput() LimitRangeListArrayOutput {
	return o
}

func (o LimitRangeListArrayOutput) ToLimitRangeListArrayOutputWithContext(ctx context.Context) LimitRangeListArrayOutput {
	return o
}

func (o LimitRangeListArrayOutput) Index(i pulumi.IntInput) LimitRangeListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LimitRangeList {
		return vs[0].([]*LimitRangeList)[vs[1].(int)]
	}).(LimitRangeListOutput)
}

type LimitRangeListMapOutput struct{ *pulumi.OutputState }

func (LimitRangeListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LimitRangeList)(nil)).Elem()
}

func (o LimitRangeListMapOutput) ToLimitRangeListMapOutput() LimitRangeListMapOutput {
	return o
}

func (o LimitRangeListMapOutput) ToLimitRangeListMapOutputWithContext(ctx context.Context) LimitRangeListMapOutput {
	return o
}

func (o LimitRangeListMapOutput) MapIndex(k pulumi.StringInput) LimitRangeListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LimitRangeList {
		return vs[0].(map[string]*LimitRangeList)[vs[1].(string)]
	}).(LimitRangeListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LimitRangeListInput)(nil)).Elem(), &LimitRangeList{})
	pulumi.RegisterInputType(reflect.TypeOf((*LimitRangeListArrayInput)(nil)).Elem(), LimitRangeListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LimitRangeListMapInput)(nil)).Elem(), LimitRangeListMap{})
	pulumi.RegisterOutputType(LimitRangeListOutput{})
	pulumi.RegisterOutputType(LimitRangeListArrayOutput{})
	pulumi.RegisterOutputType(LimitRangeListMapOutput{})
}
