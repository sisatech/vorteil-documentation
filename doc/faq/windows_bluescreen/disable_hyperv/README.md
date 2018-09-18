# Disable Hyper Virtualization

**Option 1 - Disable Hyper-V in Windows Features**

1.  Open windows features.
2.  Expand the Hyper-v sub-section.
3.  Uncheck Hyper-v and hit Ok.
4.  Restart your windows machine.

**Option 2 - Disable Hyper-V using Powershell**

1.  Open an elevated PowerShell
2.  Copy and paste this command into the prompt
  `Disable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V-All`

  3. Restart your windows machine. 