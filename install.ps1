# Installs the latest gcli release for Windows.
#
#   iwr -useb https://raw.githubusercontent.com/daflecardoso/gcli/main/install.ps1 | iex
#
# Install a specific version:
#
#   $env:GCLI_VERSION = "v0.1.0"; iwr -useb https://raw.githubusercontent.com/daflecardoso/gcli/main/install.ps1 | iex
$ErrorActionPreference = "Stop"

$Repo = "daflecardoso/gcli"
$InstallDir = if ($env:GCLI_INSTALL_DIR) { $env:GCLI_INSTALL_DIR } else { "$env:LOCALAPPDATA\gcli\bin" }
$Version = if ($env:GCLI_VERSION) { $env:GCLI_VERSION } else { "latest" }

$Arch = if ([System.Environment]::Is64BitOperatingSystem) {
  if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "amd64" }
} else {
  throw "gcli requires a 64-bit Windows"
}

$Asset = "gcli_windows_$Arch.zip"
$Url = if ($Version -eq "latest") {
  "https://github.com/$Repo/releases/latest/download/$Asset"
} else {
  "https://github.com/$Repo/releases/download/$Version/$Asset"
}

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
$TmpZip = New-TemporaryFile
$TmpZip = Rename-Item -Path $TmpZip -NewName "$($TmpZip.Name).zip" -PassThru

Write-Host "==> Downloading $Asset ($Version)..."
try {
  Invoke-WebRequest -Uri $Url -OutFile $TmpZip -UseBasicParsing
} catch {
  throw "download failed: $Url (no release built for windows/$Arch, or $Version doesn't exist?)"
}

Expand-Archive -Path $TmpZip -DestinationPath $InstallDir -Force
Remove-Item $TmpZip -Force

Write-Host "==> Installed to $InstallDir\gcli.exe"

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
  [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
  Write-Host "Added $InstallDir to your user PATH. Restart your terminal to use 'gcli'."
} else {
  Write-Host "Run 'gcli' inside a git repo with a gcli.json config to get started."
}
