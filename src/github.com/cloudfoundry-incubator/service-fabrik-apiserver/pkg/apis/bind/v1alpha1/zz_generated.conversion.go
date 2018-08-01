// +build !ignore_autogenerated

//TODO copyright header

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	bind "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_DirectorBind_To_bind_DirectorBind,
		Convert_bind_DirectorBind_To_v1alpha1_DirectorBind,
		Convert_v1alpha1_DirectorBindList_To_bind_DirectorBindList,
		Convert_bind_DirectorBindList_To_v1alpha1_DirectorBindList,
		Convert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec,
		Convert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec,
		Convert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus,
		Convert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus,
		Convert_v1alpha1_DirectorBindStatusStrategy_To_bind_DirectorBindStatusStrategy,
		Convert_bind_DirectorBindStatusStrategy_To_v1alpha1_DirectorBindStatusStrategy,
		Convert_v1alpha1_DirectorBindStrategy_To_bind_DirectorBindStrategy,
		Convert_bind_DirectorBindStrategy_To_v1alpha1_DirectorBindStrategy,
		Convert_v1alpha1_DockerBind_To_bind_DockerBind,
		Convert_bind_DockerBind_To_v1alpha1_DockerBind,
		Convert_v1alpha1_DockerBindList_To_bind_DockerBindList,
		Convert_bind_DockerBindList_To_v1alpha1_DockerBindList,
		Convert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec,
		Convert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec,
		Convert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus,
		Convert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus,
		Convert_v1alpha1_DockerBindStatusStrategy_To_bind_DockerBindStatusStrategy,
		Convert_bind_DockerBindStatusStrategy_To_v1alpha1_DockerBindStatusStrategy,
		Convert_v1alpha1_DockerBindStrategy_To_bind_DockerBindStrategy,
		Convert_bind_DockerBindStrategy_To_v1alpha1_DockerBindStrategy,
	)
}

func autoConvert_v1alpha1_DirectorBind_To_bind_DirectorBind(in *DirectorBind, out *bind.DirectorBind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_DirectorBind_To_bind_DirectorBind is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBind_To_bind_DirectorBind(in *DirectorBind, out *bind.DirectorBind, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBind_To_bind_DirectorBind(in, out, s)
}

func autoConvert_bind_DirectorBind_To_v1alpha1_DirectorBind(in *bind.DirectorBind, out *DirectorBind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_bind_DirectorBind_To_v1alpha1_DirectorBind is an autogenerated conversion function.
func Convert_bind_DirectorBind_To_v1alpha1_DirectorBind(in *bind.DirectorBind, out *DirectorBind, s conversion.Scope) error {
	return autoConvert_bind_DirectorBind_To_v1alpha1_DirectorBind(in, out, s)
}

func autoConvert_v1alpha1_DirectorBindList_To_bind_DirectorBindList(in *DirectorBindList, out *bind.DirectorBindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]bind.DirectorBind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_DirectorBindList_To_bind_DirectorBindList is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBindList_To_bind_DirectorBindList(in *DirectorBindList, out *bind.DirectorBindList, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBindList_To_bind_DirectorBindList(in, out, s)
}

func autoConvert_bind_DirectorBindList_To_v1alpha1_DirectorBindList(in *bind.DirectorBindList, out *DirectorBindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]DirectorBind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_bind_DirectorBindList_To_v1alpha1_DirectorBindList is an autogenerated conversion function.
func Convert_bind_DirectorBindList_To_v1alpha1_DirectorBindList(in *bind.DirectorBindList, out *DirectorBindList, s conversion.Scope) error {
	return autoConvert_bind_DirectorBindList_To_v1alpha1_DirectorBindList(in, out, s)
}

func autoConvert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec(in *DirectorBindSpec, out *bind.DirectorBindSpec, s conversion.Scope) error {
	out.Instance = in.Instance
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec(in *DirectorBindSpec, out *bind.DirectorBindSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBindSpec_To_bind_DirectorBindSpec(in, out, s)
}

func autoConvert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec(in *bind.DirectorBindSpec, out *DirectorBindSpec, s conversion.Scope) error {
	out.Instance = in.Instance
	out.Options = in.Options
	return nil
}

// Convert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec is an autogenerated conversion function.
func Convert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec(in *bind.DirectorBindSpec, out *DirectorBindSpec, s conversion.Scope) error {
	return autoConvert_bind_DirectorBindSpec_To_v1alpha1_DirectorBindSpec(in, out, s)
}

func autoConvert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus(in *DirectorBindStatus, out *bind.DirectorBindStatus, s conversion.Scope) error {
	out.State = in.State
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus(in *DirectorBindStatus, out *bind.DirectorBindStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBindStatus_To_bind_DirectorBindStatus(in, out, s)
}

func autoConvert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus(in *bind.DirectorBindStatus, out *DirectorBindStatus, s conversion.Scope) error {
	out.State = in.State
	out.Response = in.Response
	return nil
}

// Convert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus is an autogenerated conversion function.
func Convert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus(in *bind.DirectorBindStatus, out *DirectorBindStatus, s conversion.Scope) error {
	return autoConvert_bind_DirectorBindStatus_To_v1alpha1_DirectorBindStatus(in, out, s)
}

func autoConvert_v1alpha1_DirectorBindStatusStrategy_To_bind_DirectorBindStatusStrategy(in *DirectorBindStatusStrategy, out *bind.DirectorBindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_DirectorBindStatusStrategy_To_bind_DirectorBindStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBindStatusStrategy_To_bind_DirectorBindStatusStrategy(in *DirectorBindStatusStrategy, out *bind.DirectorBindStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBindStatusStrategy_To_bind_DirectorBindStatusStrategy(in, out, s)
}

func autoConvert_bind_DirectorBindStatusStrategy_To_v1alpha1_DirectorBindStatusStrategy(in *bind.DirectorBindStatusStrategy, out *DirectorBindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_bind_DirectorBindStatusStrategy_To_v1alpha1_DirectorBindStatusStrategy is an autogenerated conversion function.
func Convert_bind_DirectorBindStatusStrategy_To_v1alpha1_DirectorBindStatusStrategy(in *bind.DirectorBindStatusStrategy, out *DirectorBindStatusStrategy, s conversion.Scope) error {
	return autoConvert_bind_DirectorBindStatusStrategy_To_v1alpha1_DirectorBindStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_DirectorBindStrategy_To_bind_DirectorBindStrategy(in *DirectorBindStrategy, out *bind.DirectorBindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_DirectorBindStrategy_To_bind_DirectorBindStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DirectorBindStrategy_To_bind_DirectorBindStrategy(in *DirectorBindStrategy, out *bind.DirectorBindStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DirectorBindStrategy_To_bind_DirectorBindStrategy(in, out, s)
}

func autoConvert_bind_DirectorBindStrategy_To_v1alpha1_DirectorBindStrategy(in *bind.DirectorBindStrategy, out *DirectorBindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_bind_DirectorBindStrategy_To_v1alpha1_DirectorBindStrategy is an autogenerated conversion function.
func Convert_bind_DirectorBindStrategy_To_v1alpha1_DirectorBindStrategy(in *bind.DirectorBindStrategy, out *DirectorBindStrategy, s conversion.Scope) error {
	return autoConvert_bind_DirectorBindStrategy_To_v1alpha1_DirectorBindStrategy(in, out, s)
}

func autoConvert_v1alpha1_DockerBind_To_bind_DockerBind(in *DockerBind, out *bind.DockerBind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_DockerBind_To_bind_DockerBind is an autogenerated conversion function.
func Convert_v1alpha1_DockerBind_To_bind_DockerBind(in *DockerBind, out *bind.DockerBind, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBind_To_bind_DockerBind(in, out, s)
}

func autoConvert_bind_DockerBind_To_v1alpha1_DockerBind(in *bind.DockerBind, out *DockerBind, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_bind_DockerBind_To_v1alpha1_DockerBind is an autogenerated conversion function.
func Convert_bind_DockerBind_To_v1alpha1_DockerBind(in *bind.DockerBind, out *DockerBind, s conversion.Scope) error {
	return autoConvert_bind_DockerBind_To_v1alpha1_DockerBind(in, out, s)
}

func autoConvert_v1alpha1_DockerBindList_To_bind_DockerBindList(in *DockerBindList, out *bind.DockerBindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]bind.DockerBind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_DockerBindList_To_bind_DockerBindList is an autogenerated conversion function.
func Convert_v1alpha1_DockerBindList_To_bind_DockerBindList(in *DockerBindList, out *bind.DockerBindList, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBindList_To_bind_DockerBindList(in, out, s)
}

func autoConvert_bind_DockerBindList_To_v1alpha1_DockerBindList(in *bind.DockerBindList, out *DockerBindList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]DockerBind)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_bind_DockerBindList_To_v1alpha1_DockerBindList is an autogenerated conversion function.
func Convert_bind_DockerBindList_To_v1alpha1_DockerBindList(in *bind.DockerBindList, out *DockerBindList, s conversion.Scope) error {
	return autoConvert_bind_DockerBindList_To_v1alpha1_DockerBindList(in, out, s)
}

func autoConvert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec(in *DockerBindSpec, out *bind.DockerBindSpec, s conversion.Scope) error {
	out.Instance = in.Instance
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec is an autogenerated conversion function.
func Convert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec(in *DockerBindSpec, out *bind.DockerBindSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBindSpec_To_bind_DockerBindSpec(in, out, s)
}

func autoConvert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec(in *bind.DockerBindSpec, out *DockerBindSpec, s conversion.Scope) error {
	out.Instance = in.Instance
	out.Options = in.Options
	return nil
}

// Convert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec is an autogenerated conversion function.
func Convert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec(in *bind.DockerBindSpec, out *DockerBindSpec, s conversion.Scope) error {
	return autoConvert_bind_DockerBindSpec_To_v1alpha1_DockerBindSpec(in, out, s)
}

func autoConvert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus(in *DockerBindStatus, out *bind.DockerBindStatus, s conversion.Scope) error {
	out.State = in.State
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus is an autogenerated conversion function.
func Convert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus(in *DockerBindStatus, out *bind.DockerBindStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBindStatus_To_bind_DockerBindStatus(in, out, s)
}

func autoConvert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus(in *bind.DockerBindStatus, out *DockerBindStatus, s conversion.Scope) error {
	out.State = in.State
	out.Response = in.Response
	return nil
}

// Convert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus is an autogenerated conversion function.
func Convert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus(in *bind.DockerBindStatus, out *DockerBindStatus, s conversion.Scope) error {
	return autoConvert_bind_DockerBindStatus_To_v1alpha1_DockerBindStatus(in, out, s)
}

func autoConvert_v1alpha1_DockerBindStatusStrategy_To_bind_DockerBindStatusStrategy(in *DockerBindStatusStrategy, out *bind.DockerBindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_DockerBindStatusStrategy_To_bind_DockerBindStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DockerBindStatusStrategy_To_bind_DockerBindStatusStrategy(in *DockerBindStatusStrategy, out *bind.DockerBindStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBindStatusStrategy_To_bind_DockerBindStatusStrategy(in, out, s)
}

func autoConvert_bind_DockerBindStatusStrategy_To_v1alpha1_DockerBindStatusStrategy(in *bind.DockerBindStatusStrategy, out *DockerBindStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_bind_DockerBindStatusStrategy_To_v1alpha1_DockerBindStatusStrategy is an autogenerated conversion function.
func Convert_bind_DockerBindStatusStrategy_To_v1alpha1_DockerBindStatusStrategy(in *bind.DockerBindStatusStrategy, out *DockerBindStatusStrategy, s conversion.Scope) error {
	return autoConvert_bind_DockerBindStatusStrategy_To_v1alpha1_DockerBindStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_DockerBindStrategy_To_bind_DockerBindStrategy(in *DockerBindStrategy, out *bind.DockerBindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_DockerBindStrategy_To_bind_DockerBindStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DockerBindStrategy_To_bind_DockerBindStrategy(in *DockerBindStrategy, out *bind.DockerBindStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DockerBindStrategy_To_bind_DockerBindStrategy(in, out, s)
}

func autoConvert_bind_DockerBindStrategy_To_v1alpha1_DockerBindStrategy(in *bind.DockerBindStrategy, out *DockerBindStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_bind_DockerBindStrategy_To_v1alpha1_DockerBindStrategy is an autogenerated conversion function.
func Convert_bind_DockerBindStrategy_To_v1alpha1_DockerBindStrategy(in *bind.DockerBindStrategy, out *DockerBindStrategy, s conversion.Scope) error {
	return autoConvert_bind_DockerBindStrategy_To_v1alpha1_DockerBindStrategy(in, out, s)
}
