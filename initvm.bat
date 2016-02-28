SET vername=su16vm-v1.0

vboxmanage unregistervm "%vername%" --delete

vboxmanage import "%vername%.ova"

sfmount "conf"
sfmount "go"
sfmount "static"

if not exist "log" mkdir "log"
if not exist "log\nginx" mkdir "log\nginx"

sfmount "log"

vboxmanage startvm "%vername%" --type headless

:sfmount
vboxmanage sharedfolder add "%vername%" ^
  --name %~1 --hostpath "%cd%\%~1" --automount
