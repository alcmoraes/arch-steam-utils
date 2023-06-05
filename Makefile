install-aans:
	mkdir -p ~/.local/share/kservices5/ServiceMenus
	mkdir -p ~/.bin
	go build -o ~/.bin/aans add-as-non-steam/main.go
	cp add-as-non-steam/steam-context-menu.desktop ~/.local/share/kservices5/ServiceMenus/SimpleSteamShortcutAdder.desktop
	chmod + ~/.bin/aans