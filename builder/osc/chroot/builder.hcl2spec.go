// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package chroot

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/outscale/packer-plugin-outscale/builder/osc/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName         *string                      `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType       *string                      `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion       *string                      `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug             *bool                        `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce             *bool                        `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError           *string                      `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars          map[string]string            `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars     []string                     `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	OMIMappings             []common.FlatBlockDevice     `mapstructure:"omi_block_device_mappings" cty:"omi_block_device_mappings" hcl:"omi_block_device_mappings"`
	OMIName                 *string                      `mapstructure:"omi_name" cty:"omi_name" hcl:"omi_name"`
	OMIDescription          *string                      `mapstructure:"omi_description" cty:"omi_description" hcl:"omi_description"`
	OMIAccountIDs           []string                     `mapstructure:"omi_account_ids" cty:"omi_account_ids" hcl:"omi_account_ids"`
	OMIGroups               []string                     `mapstructure:"omi_groups" cty:"omi_groups" hcl:"omi_groups"`
	OMIProductCodes         []string                     `mapstructure:"omi_product_codes" cty:"omi_product_codes" hcl:"omi_product_codes"`
	OMIRegions              []string                     `mapstructure:"omi_regions" cty:"omi_regions" hcl:"omi_regions"`
	OMISkipRegionValidation *bool                        `mapstructure:"skip_region_validation" cty:"skip_region_validation" hcl:"skip_region_validation"`
	OMITags                 common.TagMap                `mapstructure:"tags" cty:"tags" hcl:"tags"`
	OMIForceDeregister      *bool                        `mapstructure:"force_deregister" cty:"force_deregister" hcl:"force_deregister"`
	OMIForceDeleteSnapshot  *bool                        `mapstructure:"force_delete_snapshot" cty:"force_delete_snapshot" hcl:"force_delete_snapshot"`
	SnapshotTags            common.TagMap                `mapstructure:"snapshot_tags" cty:"snapshot_tags" hcl:"snapshot_tags"`
	SnapshotAccountIDs      []string                     `mapstructure:"snapshot_account_ids" cty:"snapshot_account_ids" hcl:"snapshot_account_ids"`
	SnapshotGroups          []string                     `mapstructure:"snapshot_groups" cty:"snapshot_groups" hcl:"snapshot_groups"`
	AccessKey               *string                      `mapstructure:"access_key" cty:"access_key" hcl:"access_key"`
	CustomEndpointOAPI      *string                      `mapstructure:"custom_endpoint_oapi" cty:"custom_endpoint_oapi" hcl:"custom_endpoint_oapi"`
	InsecureSkipTLSVerify   *bool                        `mapstructure:"insecure_skip_tls_verify" cty:"insecure_skip_tls_verify" hcl:"insecure_skip_tls_verify"`
	MFACode                 *string                      `mapstructure:"mfa_code" cty:"mfa_code" hcl:"mfa_code"`
	ProfileName             *string                      `mapstructure:"profile" cty:"profile" hcl:"profile"`
	RawRegion               *string                      `mapstructure:"region" cty:"region" hcl:"region"`
	SecretKey               *string                      `mapstructure:"secret_key" cty:"secret_key" hcl:"secret_key"`
	SkipMetadataApiCheck    *bool                        `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check" hcl:"skip_metadata_api_check"`
	Token                   *string                      `mapstructure:"token" cty:"token" hcl:"token"`
	X509certPath            *string                      `mapstructure:"x509_cert_path" cty:"x509_cert_path" hcl:"x509_cert_path"`
	X509keyPath             *string                      `mapstructure:"x509_key_path" cty:"x509_key_path" hcl:"x509_key_path"`
	ChrootMounts            [][]string                   `mapstructure:"chroot_mounts" cty:"chroot_mounts" hcl:"chroot_mounts"`
	CommandWrapper          *string                      `mapstructure:"command_wrapper" cty:"command_wrapper" hcl:"command_wrapper"`
	CopyFiles               []string                     `mapstructure:"copy_files" cty:"copy_files" hcl:"copy_files"`
	DevicePath              *string                      `mapstructure:"device_path" cty:"device_path" hcl:"device_path"`
	NVMEDevicePath          *string                      `mapstructure:"nvme_device_path" cty:"nvme_device_path" hcl:"nvme_device_path"`
	FromScratch             *bool                        `mapstructure:"from_scratch" cty:"from_scratch" hcl:"from_scratch"`
	MountOptions            []string                     `mapstructure:"mount_options" cty:"mount_options" hcl:"mount_options"`
	MountPartition          *string                      `mapstructure:"mount_partition" cty:"mount_partition" hcl:"mount_partition"`
	MountPath               *string                      `mapstructure:"mount_path" cty:"mount_path" hcl:"mount_path"`
	PostMountCommands       []string                     `mapstructure:"post_mount_commands" cty:"post_mount_commands" hcl:"post_mount_commands"`
	PreMountCommands        []string                     `mapstructure:"pre_mount_commands" cty:"pre_mount_commands" hcl:"pre_mount_commands"`
	RootDeviceName          *string                      `mapstructure:"root_device_name" cty:"root_device_name" hcl:"root_device_name"`
	RootVolumeSize          *int64                       `mapstructure:"root_volume_size" cty:"root_volume_size" hcl:"root_volume_size"`
	RootVolumeType          *string                      `mapstructure:"root_volume_type" cty:"root_volume_type" hcl:"root_volume_type"`
	SourceOMI               *string                      `mapstructure:"source_omi" cty:"source_omi" hcl:"source_omi"`
	SourceOMIFilter         *common.FlatOmiFilterOptions `mapstructure:"source_omi_filter" cty:"source_omi_filter" hcl:"source_omi_filter"`
	RootVolumeTags          common.TagMap                `mapstructure:"root_volume_tags" cty:"root_volume_tags" hcl:"root_volume_tags"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"omi_block_device_mappings":  &hcldec.BlockListSpec{TypeName: "omi_block_device_mappings", Nested: hcldec.ObjectSpec((*common.FlatBlockDevice)(nil).HCL2Spec())},
		"omi_name":                   &hcldec.AttrSpec{Name: "omi_name", Type: cty.String, Required: false},
		"omi_description":            &hcldec.AttrSpec{Name: "omi_description", Type: cty.String, Required: false},
		"omi_account_ids":            &hcldec.AttrSpec{Name: "omi_account_ids", Type: cty.List(cty.String), Required: false},
		"omi_groups":                 &hcldec.AttrSpec{Name: "omi_groups", Type: cty.List(cty.String), Required: false},
		"omi_product_codes":          &hcldec.AttrSpec{Name: "omi_product_codes", Type: cty.List(cty.String), Required: false},
		"omi_regions":                &hcldec.AttrSpec{Name: "omi_regions", Type: cty.List(cty.String), Required: false},
		"skip_region_validation":     &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"tags":                       &hcldec.AttrSpec{Name: "tags", Type: cty.Map(cty.String), Required: false},
		"force_deregister":           &hcldec.AttrSpec{Name: "force_deregister", Type: cty.Bool, Required: false},
		"force_delete_snapshot":      &hcldec.AttrSpec{Name: "force_delete_snapshot", Type: cty.Bool, Required: false},
		"snapshot_tags":              &hcldec.AttrSpec{Name: "snapshot_tags", Type: cty.Map(cty.String), Required: false},
		"snapshot_account_ids":       &hcldec.AttrSpec{Name: "snapshot_account_ids", Type: cty.List(cty.String), Required: false},
		"snapshot_groups":            &hcldec.AttrSpec{Name: "snapshot_groups", Type: cty.List(cty.String), Required: false},
		"access_key":                 &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"custom_endpoint_oapi":       &hcldec.AttrSpec{Name: "custom_endpoint_oapi", Type: cty.String, Required: false},
		"insecure_skip_tls_verify":   &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"mfa_code":                   &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                    &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                     &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                 &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_metadata_api_check":    &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"token":                      &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"x509_cert_path":             &hcldec.AttrSpec{Name: "x509_cert_path", Type: cty.String, Required: false},
		"x509_key_path":              &hcldec.AttrSpec{Name: "x509_key_path", Type: cty.String, Required: false},
		"chroot_mounts":              &hcldec.AttrSpec{Name: "chroot_mounts", Type: cty.List(cty.List(cty.String)), Required: false},
		"command_wrapper":            &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"copy_files":                 &hcldec.AttrSpec{Name: "copy_files", Type: cty.List(cty.String), Required: false},
		"device_path":                &hcldec.AttrSpec{Name: "device_path", Type: cty.String, Required: false},
		"nvme_device_path":           &hcldec.AttrSpec{Name: "nvme_device_path", Type: cty.String, Required: false},
		"from_scratch":               &hcldec.AttrSpec{Name: "from_scratch", Type: cty.Bool, Required: false},
		"mount_options":              &hcldec.AttrSpec{Name: "mount_options", Type: cty.List(cty.String), Required: false},
		"mount_partition":            &hcldec.AttrSpec{Name: "mount_partition", Type: cty.String, Required: false},
		"mount_path":                 &hcldec.AttrSpec{Name: "mount_path", Type: cty.String, Required: false},
		"post_mount_commands":        &hcldec.AttrSpec{Name: "post_mount_commands", Type: cty.List(cty.String), Required: false},
		"pre_mount_commands":         &hcldec.AttrSpec{Name: "pre_mount_commands", Type: cty.List(cty.String), Required: false},
		"root_device_name":           &hcldec.AttrSpec{Name: "root_device_name", Type: cty.String, Required: false},
		"root_volume_size":           &hcldec.AttrSpec{Name: "root_volume_size", Type: cty.Number, Required: false},
		"root_volume_type":           &hcldec.AttrSpec{Name: "root_volume_type", Type: cty.String, Required: false},
		"source_omi":                 &hcldec.AttrSpec{Name: "source_omi", Type: cty.String, Required: false},
		"source_omi_filter":          &hcldec.BlockSpec{TypeName: "source_omi_filter", Nested: hcldec.ObjectSpec((*common.FlatOmiFilterOptions)(nil).HCL2Spec())},
		"root_volume_tags":           &hcldec.AttrSpec{Name: "root_volume_tags", Type: cty.Map(cty.String), Required: false},
	}
	return s
}
