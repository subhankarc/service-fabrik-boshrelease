// +build !ignore_autogenerated

//TODO copyright header

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
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
		Convert_v1alpha1_DefaultBackup_To_backup_DefaultBackup,
		Convert_backup_DefaultBackup_To_v1alpha1_DefaultBackup,
		Convert_v1alpha1_DefaultBackupList_To_backup_DefaultBackupList,
		Convert_backup_DefaultBackupList_To_v1alpha1_DefaultBackupList,
		Convert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec,
		Convert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec,
		Convert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus,
		Convert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus,
		Convert_v1alpha1_DefaultBackupStatusStrategy_To_backup_DefaultBackupStatusStrategy,
		Convert_backup_DefaultBackupStatusStrategy_To_v1alpha1_DefaultBackupStatusStrategy,
		Convert_v1alpha1_DefaultBackupStrategy_To_backup_DefaultBackupStrategy,
		Convert_backup_DefaultBackupStrategy_To_v1alpha1_DefaultBackupStrategy,
	)
}

func autoConvert_v1alpha1_DefaultBackup_To_backup_DefaultBackup(in *DefaultBackup, out *backup.DefaultBackup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_DefaultBackup_To_backup_DefaultBackup is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackup_To_backup_DefaultBackup(in *DefaultBackup, out *backup.DefaultBackup, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackup_To_backup_DefaultBackup(in, out, s)
}

func autoConvert_backup_DefaultBackup_To_v1alpha1_DefaultBackup(in *backup.DefaultBackup, out *DefaultBackup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_backup_DefaultBackup_To_v1alpha1_DefaultBackup is an autogenerated conversion function.
func Convert_backup_DefaultBackup_To_v1alpha1_DefaultBackup(in *backup.DefaultBackup, out *DefaultBackup, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackup_To_v1alpha1_DefaultBackup(in, out, s)
}

func autoConvert_v1alpha1_DefaultBackupList_To_backup_DefaultBackupList(in *DefaultBackupList, out *backup.DefaultBackupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]backup.DefaultBackup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_DefaultBackupList_To_backup_DefaultBackupList is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackupList_To_backup_DefaultBackupList(in *DefaultBackupList, out *backup.DefaultBackupList, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackupList_To_backup_DefaultBackupList(in, out, s)
}

func autoConvert_backup_DefaultBackupList_To_v1alpha1_DefaultBackupList(in *backup.DefaultBackupList, out *DefaultBackupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]DefaultBackup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_backup_DefaultBackupList_To_v1alpha1_DefaultBackupList is an autogenerated conversion function.
func Convert_backup_DefaultBackupList_To_v1alpha1_DefaultBackupList(in *backup.DefaultBackupList, out *DefaultBackupList, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackupList_To_v1alpha1_DefaultBackupList(in, out, s)
}

func autoConvert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec(in *DefaultBackupSpec, out *backup.DefaultBackupSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec(in *DefaultBackupSpec, out *backup.DefaultBackupSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackupSpec_To_backup_DefaultBackupSpec(in, out, s)
}

func autoConvert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec(in *backup.DefaultBackupSpec, out *DefaultBackupSpec, s conversion.Scope) error {
	out.Options = in.Options
	return nil
}

// Convert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec is an autogenerated conversion function.
func Convert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec(in *backup.DefaultBackupSpec, out *DefaultBackupSpec, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackupSpec_To_v1alpha1_DefaultBackupSpec(in, out, s)
}

func autoConvert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus(in *DefaultBackupStatus, out *backup.DefaultBackupStatus, s conversion.Scope) error {
	out.State = in.State
	out.Error = in.Error
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus(in *DefaultBackupStatus, out *backup.DefaultBackupStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackupStatus_To_backup_DefaultBackupStatus(in, out, s)
}

func autoConvert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus(in *backup.DefaultBackupStatus, out *DefaultBackupStatus, s conversion.Scope) error {
	out.State = in.State
	out.Error = in.Error
	out.LastOperation = in.LastOperation
	out.Response = in.Response
	return nil
}

// Convert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus is an autogenerated conversion function.
func Convert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus(in *backup.DefaultBackupStatus, out *DefaultBackupStatus, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackupStatus_To_v1alpha1_DefaultBackupStatus(in, out, s)
}

func autoConvert_v1alpha1_DefaultBackupStatusStrategy_To_backup_DefaultBackupStatusStrategy(in *DefaultBackupStatusStrategy, out *backup.DefaultBackupStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_DefaultBackupStatusStrategy_To_backup_DefaultBackupStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackupStatusStrategy_To_backup_DefaultBackupStatusStrategy(in *DefaultBackupStatusStrategy, out *backup.DefaultBackupStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackupStatusStrategy_To_backup_DefaultBackupStatusStrategy(in, out, s)
}

func autoConvert_backup_DefaultBackupStatusStrategy_To_v1alpha1_DefaultBackupStatusStrategy(in *backup.DefaultBackupStatusStrategy, out *DefaultBackupStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_backup_DefaultBackupStatusStrategy_To_v1alpha1_DefaultBackupStatusStrategy is an autogenerated conversion function.
func Convert_backup_DefaultBackupStatusStrategy_To_v1alpha1_DefaultBackupStatusStrategy(in *backup.DefaultBackupStatusStrategy, out *DefaultBackupStatusStrategy, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackupStatusStrategy_To_v1alpha1_DefaultBackupStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_DefaultBackupStrategy_To_backup_DefaultBackupStrategy(in *DefaultBackupStrategy, out *backup.DefaultBackupStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_DefaultBackupStrategy_To_backup_DefaultBackupStrategy is an autogenerated conversion function.
func Convert_v1alpha1_DefaultBackupStrategy_To_backup_DefaultBackupStrategy(in *DefaultBackupStrategy, out *backup.DefaultBackupStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_DefaultBackupStrategy_To_backup_DefaultBackupStrategy(in, out, s)
}

func autoConvert_backup_DefaultBackupStrategy_To_v1alpha1_DefaultBackupStrategy(in *backup.DefaultBackupStrategy, out *DefaultBackupStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_backup_DefaultBackupStrategy_To_v1alpha1_DefaultBackupStrategy is an autogenerated conversion function.
func Convert_backup_DefaultBackupStrategy_To_v1alpha1_DefaultBackupStrategy(in *backup.DefaultBackupStrategy, out *DefaultBackupStrategy, s conversion.Scope) error {
	return autoConvert_backup_DefaultBackupStrategy_To_v1alpha1_DefaultBackupStrategy(in, out, s)
}
