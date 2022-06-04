if not exist "%TEMP%\go.msi" (
  powershell -Command "(New-Object System.Net.WebClient).DownloadFile('https://storage.googleapis.com/golang/go1.4.2.windows-amd64.msi', '%TEMP%\go.msi')" <NUL
)

msiexec /quiet /i "%TEMP%\go.msi

%SystemRoot%\System32\setx GOPATH "%USERPROFILE%\go"
mkdir "%GOPATH%"
%SystemRoot%\System32\setx.exe PATH "%PATH%;C:\Go\bin" /M
