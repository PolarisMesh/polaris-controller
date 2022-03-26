package util

const (
	PolarisSync           = "polarismesh.cn/sync"
	PolarisEnableRegister = "polarismesh.cn/enableRegister"
	PolarisAliasNamespace = "polarismesh.cn/aliasNamespace"
	PolarisAliasService   = "polarismesh.cn/aliasService"

	PolarisMetadata     = "polarismesh.cn/metadata"
	PolarisWeight       = "polarismesh.cn/weight"
	PolarisHeartBeatTTL = "polarismesh.cn/ttl"

	WorkloadKind        = "polarismesh.cn/workloadKind"
	PolarisCustomWeight = "polarismesh.cn/customWeight"

	PolarisSidecarMode = "polaris-sidecar-mode"
)

const (
	PolarisClusterName = "clusterName"
	PolarisSource      = "source"
	PolarisVersion     = "version"
	PolarisProtocol    = "protocol"

	// PolarisOldSource 旧版本 controller 用来标志是 controller 同步的服务实例。
	// 已经废弃，项目中当前用来兼容存量的实例。
	PolarisOldSource = "platform"
)

const (
	SyncModeAll    = "all"
	SyncModeDemand = "demand"
	IsEnableSync   = "true"
	IsDisableSync  = "false"
)

// PolarisSystemMetaSet 由 polaris controller 决定的 meta，用户如果在 custom meta 中设置了，不会生效
var PolarisSystemMetaSet = map[string]struct{}{PolarisClusterName: {}, PolarisSource: {}}

// PolarisDefaultMetaSet 由 polaris controller 托管的 service ，注册的实例必定会带的 meta，
// 用于判断用户的 custom meta 是否发生了更新
var PolarisDefaultMetaSet = map[string]struct{}{
	PolarisClusterName: {},
	PolarisSource:      {},
	PolarisVersion:     {},
	PolarisProtocol:    {},
}

// ServiceChangeType 发升变更的类型
type ServiceChangeType string

const (
	ServicePolarisDelete         ServiceChangeType = "servicePolarisDelete" // 删除了北极星的服务
	ServiceNameSpacesChanged     ServiceChangeType = "serviceNameSpacesChanged"
	ServiceNameChanged           ServiceChangeType = "serviceNameChanged"
	ServiceWeightChanged         ServiceChangeType = "serviceWeightChanged"
	ServiceTokenChanged          ServiceChangeType = "serviceTokenChanged"
	ServiceTTLChanged            ServiceChangeType = "serviceTTLChanged"
	ServiceEnableRegisterChanged ServiceChangeType = "serviceEnableRegisterChanged"
	ServiceMetadataChanged       ServiceChangeType = "serviceMetadataChanged"
	ServiceCustomWeightChanged   ServiceChangeType = "serviceCustomWeightChanged"
)

const (
	PolarisGoConfigFileTpl string = "polaris-client-config-tpl"
	PolarisGoConfigFile string = "polaris-client-config"
)

type SidecarMode int

const (
	SidecarForUnknown SidecarMode = iota
	SidecarForMesh
	SidecarForDns

	SidecarMeshModeName string = "mesh"
	SidecarDnsModeName string = "dns"
)

func ParseSidecarMode(val string) SidecarMode {
	if val == SidecarMeshModeName {
		return SidecarForMesh
	}
	if val == SidecarDnsModeName {
		return SidecarForDns
	}
	return SidecarForMesh
}

func ParseSidecarModeName(mode SidecarMode) string {
	if mode == SidecarForMesh {
		return SidecarMeshModeName
	}
	if mode == SidecarForDns {
		return SidecarDnsModeName
	}
	return SidecarMeshModeName
}

// IndexPortMap 对应{"index-port":weight}
type IndexPortMap map[string]int
