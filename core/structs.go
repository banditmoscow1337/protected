package core

var (
	versionListURL = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

type backup struct {
	hash      string
	timestamp int64
}

type client struct {
	nick, ip, city string
}

type Server struct {
	backups []int64

	version,
	worldName,
	serverName,
	gamemode,
	difficulty string

	memory struct {
		min,
		max int64
	}

	viewDistance,
	spawnProtection,
	port,
	maxPlayers int64

	pvp,
	online,
	structures,
	npcs,
	animals,
	monsters,
	hardcore bool
}
