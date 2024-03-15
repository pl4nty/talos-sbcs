package main

import (
	_ "embed"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

func main() {
	adapter.Execute(&BoardInstaller{})
}

type BoardInstaller struct{}

type boardExtraOptions struct{}

func (i *BoardInstaller) GetOptions(extra boardExtraOptions) (overlay.Options, error) {
	kernelArgs := []string{
		"console=tty0",
		"console=ttyS2,1500000n8",
		"talos.dashboard.disabled=1",
	}

	return overlay.Options{
		Name:       "soquartzmodela",
		KernelArgs: kernelArgs,
		PartitionOptions: overlay.PartitionOptions{
			Offset: 512 * 64,
		},
	}, nil
}

func (i *BoardInstaller) Install(options overlay.InstallOptions[boardExtraOptions]) error {
	// allows to copy a directory from the overlay to the target
	err := copy.Dir(filepath.Join(options.ArtifactsPath, "arm64/dtb"), filepath.Join(options.MountPrefix, "/boot/EFI/dtb"))
	if err != nil {
		return err
	}

	// allows to copy a file from the overlay to the target
	err = copy.File(filepath.Join(options.ArtifactsPath, "arm64/u-boot/soquartzmodela/u-boot-rockchip.bin"), filepath.Join(options.MountPrefix, "/boot/EFI/u-boot.bin"))
	if err != nil {
		return err
	}

	return nil
}
