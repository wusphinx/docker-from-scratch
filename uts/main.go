package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
	以下两种情况，会导致uidMappings/GidMappings中设置了非当前进程所属UID和GID的相关数值：
	1. HostID非本进程所有（与Getuid()和Getgid()不等）
	2. Size大于1 （则肯定包含非当前进程的UID和GID）
	则需要Host机使用Root权限才能正常执行此段代码。
*/
// 以root扫行后，当前id是10086，而宿主机的id并不是10086，以此可以说明user namespace生效了
var uidMappings = []syscall.SysProcIDMap{
	{
		ContainerID: 10086,
		HostID:      syscall.Getuid(),
		Size:        1,
	},
	{
		ContainerID: 10010,
		HostID:      syscall.Getgid() + 1,
		Size:        1,
	},
}

var gidMappings = []syscall.SysProcIDMap{
	{
		ContainerID: 10086,
		HostID:      syscall.Getgid(),
		Size:        1,
	},
	{
		ContainerID: 10010,
		HostID:      syscall.Getgid() + 1,
		Size:        1,
	},
}

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			// 实际上是一个位操作是: 67108864|134217728
			// refer: https://man7.org/linux/man-pages/man2/clone.2.html
			syscall.CLONE_NEWIPC |

			// 运行后看到当前pid为1
			// # echo $$
			// 1
			syscall.CLONE_NEWPID |

			// mount namespace
			// # mount -t proc proc /pro
			syscall.CLONE_NEWNS |

			syscall.CLONE_NEWUSER |

			// 运行后发现没有任何网络设备
			// # go run main.go
			// $ ifconfig
			// $
			syscall.CLONE_NEWNET,

		// refer: https://github.com/xianlubird/mydocker/issues/3
		UidMappings: uidMappings,
		GidMappings: gidMappings,
	}

	// 以下代码会报：fork/exec /usr/bin/sh: operation not permitted exit status 1
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid:uint32(1), Gid:uint32(1)}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
