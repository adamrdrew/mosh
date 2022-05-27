$fso = new-object -ComObject scripting.filesystemobject

if ( !$fso.FolderExists("mosh_tmp") ) {
    $fso.CreateFolder("mosh_tmp")
}

$Env:MOSH_CONFIG_DIR='mosh_tmp'
$Env:MOSH_LOG_DIR='mosh_tmp'
$Env:MOSH_PID_DIR='mosh_tmp'
$Env:MOSH_PORT='9777'
$Env:MOSH_CACHE_DIR='mosh_tmp/cache'

go run .\winmoshd.go