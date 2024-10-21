ENV := $(if $(ENV),$(ENV),$(shell echo 'development'))

#.PHONY: templ css air vite temple refresh
.PHONY: live/clean live/vite live/server live/templ live/tailwind live/sync_assets live build/clean build

ec:
	@echo $(ENV)

# The clean target removes the tmp directory used by air
live/clean:
	@rm -rf ./tmp

live/vite:
	npx vite build

live/server: live/clean
	air

live/templ:
	templ generate --watch --proxy=http://localhost:1323 -v

live/tailwind:
	npx tailwindcss -i views/css/global.css -o assets/styles/global.css --watch

live/sync_assets:
	go run github.com/cosmtrek/air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

live:
	make -j5 live/templ live/server live/tailwind live/vite live/sync_assets


build/clean:
	@rm -rf ./out

build: live/clean build/clean
	mkdir -p ./out/
	templ generate -v
	npx tailwindcss -i views/css/global.css -o assets/styles/global.css --minify
	npx vite build
	cp -r ./public/** ./assets/
	cp -rp ./assets ./out
	go build -tags 'prod' -o ./out/main ./
