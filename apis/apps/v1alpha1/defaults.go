package v1alpha1

// 这个方法把defaulter注册进入scheme
func init() {
	localSchemeBuilder.Register(RegisterDefaults)
}

// 这个方法的存在将影响code_generator生成的defaulter代码
func SetDefaults_GuestBookSpec(obj *GuestBookSpec) {
	if obj.InstanceAmount <= 0 {
		obj.InstanceAmount = 1
	}

	if obj.InstanceCpu <= 0 {
		obj.InstanceCpu = 1
	}
}
