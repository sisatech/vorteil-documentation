# KVM (Kernel-based Virtual Machine)

KVM is a supported hypervisor on Linux machines. Installating KVM is quick and easy.
To install KVM:

```bash
apt install qemu-kvm
```

You can verify that the installation was successful with the following:

```bash
$ kvm-ok
INFO: /dev/kvm exists
KVM acceleration can be used
```