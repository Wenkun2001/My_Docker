package subsystem

type ResourceConfig struct {
	// 内存限制
	MemoryLimit string
	// CPU时间片权重
	CpuShare string
	// CPU核数
	CpuSet string
}

// Subsystem 将cgroup抽象成path，因为在hierarchy中，cgroup便是虚拟的路径地址
type Subsystem interface {
	// Name 返回subsystem名字，如cpu，memory
	Name() string
	// Set 设置cgroup在这个subSystem中的资源限制
	Set(cgroupPath string, res *ResourceConfig) error
	// Remove 移除这个cgroup资源限制
	Remove(cgroupPath string) error
	// Apply 将某个进程添加到cgroup中
	Apply(cgroupPath string, pid int) error
}

type CpuSubSystem struct {
}

func (c CpuSubSystem) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c CpuSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	//TODO implement me
	panic("implement me")
}

func (c CpuSubSystem) Remove(cgroupPath string) error {
	//TODO implement me
	panic("implement me")
}

func (c CpuSubSystem) Apply(cgroupPath string, pid int) error {
	//TODO implement me
	panic("implement me")
}

type CpuSetSubSystem struct {
}

func (c CpuSetSubSystem) Name() string {
	//TODO implement me
	panic("implement me")
}

func (c CpuSetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	//TODO implement me
	panic("implement me")
}

func (c CpuSetSubSystem) Remove(cgroupPath string) error {
	//TODO implement me
	panic("implement me")
}

func (c CpuSetSubSystem) Apply(cgroupPath string, pid int) error {
	//TODO implement me
	panic("implement me")
}

var (
	Subsystems = []Subsystem{
		&MemorySubSystem{},
		&CpuSubSystem{},
		&CpuSetSubSystem{},
	}
)
