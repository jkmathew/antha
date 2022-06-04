if not exist "%TEMP%\7z920-x64.msi" (
    powershell -Command "(New-Object System.Net.WebClient).DownloadFile('http://downloads.sourceforge.net/sevenzip/7z920-x64.msi', '%TEMP%\7z920-x64.msi')" <NUL
)

msiexec /qb /i "%TEMP%\7z920-x64.msi"

pushd "%TEMP%"
rem Make sure the '%' in URL are escaped when copy-and-pasting in
if not exist "%TEMP%\mingw.7z" (
  powershell -Command "(New-Object System.Net.WebClient).DownloadFile('http://cznic.dl.sourceforge.net/project/mingw-w64/Toolchains%%20targetting%%20Win64/Personal%%20Builds/mingw-builds/5.1.0/threads-posix/seh/x86_64-5.1.0-release-posix-seh-rt_v4-rev0.7z', 'mingw.7z')" <NUL
)
cmd /c ""%PROGRAMFILES%\7-Zip\7z.exe" x mingw.7z"
mkdir "C:\MinGW"
xcopy /Y /E mingw64 "C:\MinGW"
rmdir /s /q mingw64
del mingw.7z
popd

msiexec /qb /x "%TEMP%\7z920-x64.msi"

%SystemRoot%\System32\setx.exe C_INCLUDE_PATH C:\MinGW\include /M
%SystemRoot%\System32\setx.exe LIBRARY_PATH C:\MinGW\lib /M
%SystemRoot%\System32\setx.exe PATH "%PATH%;C:\MinGW\bin" /M
