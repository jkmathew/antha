
if not exist "%TEMP%\winglpk-4.55.zip" (
powershell -Command "(New-Object System.Net.WebClient).DownloadFile('http://downloads.sourceforge.net/project/winglpk/winglpk/GLPK-4.55/winglpk-4.55.zip?r=&ts=1434626279&use_mirror=cznic', '%TEMP%\winglpk-4.55.zip')" <NUL
)

pushd %TEMP%
unzip "%TEMP%\winglpk-4.55.zip"
copy glpk-4.55\w64\glpk_4_55.lib C:\MinGW\lib\
mklink C:\MinGW\lib\glpk.lib C:\MinGW\lib\glpk_4_55.lib
copy glpk-4.55\src\glpk.h C:\MinGW\include\
del winglpk-4.55.zip
rmdir /s /q glpk-4.55
popd
