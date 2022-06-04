if not exist "%TEMP%\git-install.exe" (
  powershell -Command "(New-Object System.Net.WebClient).DownloadFile('https://github.com/msysgit/msysgit/releases/download/Git-1.9.5-preview20150319/Git-1.9.5-preview20150319.exe', '%TEMP%\git-install.exe')" <NUL
)

cmd /c "%TEMP%\git-install.exe" /SILENT

%SystemRoot%\System32\setx.exe PATH "%PATH%;%PROGRAMFILES(x86)%\Git\bin" /M
