package core

import (
	"os"
	"strconv"
)

func (s *Server) generateConfig() error {
	return os.WriteFile("server.properties", []byte(`enable-jmx-monitoring=false
rcon.port=25575
level-seed=
gamemode=`+s.gamemode+`
enable-command-block=false
enable-query=false
generator-settings={}
enforce-secure-profile=true
level-name=`+s.worldName+`
motd=`+s.serverName+`
query.port=25565
pvp=`+bool2string(s.pvp)+`
generate-structures=`+bool2string(s.structures)+`
max-chained-neighbor-updates=1000000
difficulty=`+s.difficulty+`
network-compression-threshold=256
max-tick-time=60000
require-resource-pack=false
use-native-transport=true
max-players=`+strconv.FormatInt(s.maxPlayers, 10)+`
online-mode=`+bool2string(s.pvp)+`
enable-status=true
allow-flight=false
initial-disabled-packs=
broadcast-rcon-to-ops=false
view-distance=`+strconv.FormatInt(s.viewDistance, 10)+`
server-ip=
resource-pack-prompt=
allow-nether=true
server-port=`+strconv.FormatInt(s.port, 10)+`
enable-rcon=false
sync-chunk-writes=true
op-permission-level=4
prevent-proxy-connections=false
hide-online-players=false
resource-pack=
entity-broadcast-range-percentage=100
simulation-distance=`+strconv.FormatInt(s.viewDistance, 10)+`
rcon.password=
player-idle-timeout=0
force-gamemode=true
rate-limit=0
hardcore=`+bool2string(s.hardcore)+`
white-list=true
broadcast-console-to-ops=false
spawn-npcs=`+bool2string(s.npcs)+`
spawn-animals=`+bool2string(s.animals)+`
log-ips=true
function-permission-level=2
initial-enabled-packs=vanilla
level-type=minecraft\:normal
text-filtering-config=
spawn-monsters=`+bool2string(s.monsters)+`
enforce-whitelist=false
spawn-protection=`+strconv.FormatInt(s.spawnProtection, 10)+`
resource-pack-sha1=
max-world-size=29999984`), 0775)
}
