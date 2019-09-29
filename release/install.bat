@echo off

reg add "HKEY_CLASSES_ROOT\*\Shell\base64Encode"  /d "" /f

reg add "HKEY_CLASSES_ROOT\*\Shell\base64Encode\command"  /d "\"%~dp0base64.exe\"  \"%%1\"" /f
