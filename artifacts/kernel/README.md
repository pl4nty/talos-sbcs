# Kernel

## Updating kernel config

When updating kernel to the new version, import proper defaults with:

```sh
make kernel-olddefconfig-board
```

If you want to update for a specific architecture only, use:

```sh
make kernel-olddefconfig-board PLATFORM=linux/arm64
```

## Customizing the kernel

Run another target to get into `menuconfig`:

```sh
make kernel-menuconfig-board
```
