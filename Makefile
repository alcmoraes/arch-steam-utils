install-aans:
	mkdir -p ~/.local/share/kservices5/ServiceMenus
	mkdir -p ~/.bin
	go build -o ~/.bin/aans add-as-no-steam/main.go
	cp add-as-no-steam/steam-context-menu.desktop ~/.local/share/kservices5/ServiceMenus/SimpleSteamShortcutAdder.desktop
	chmod + ~/.bin/aans