.PHONY: templ css air vite

air: css
	@echo "Starting air, templ, and css..."
	npx tailwindcss -i views/css/global.css -o public/styles/global.css & \
	templ generate --watch --proxy=http://localhost:1323 & \
	npx vite build & \
	air

css: vite
	npx tailwindcss -i views/css/global.css -o public/styles/global.css

vite: clean
	npx vite build

# The clean target removes the tmp directory used by air
clean:
	@rm -rf ./tmp
