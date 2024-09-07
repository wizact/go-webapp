ENV := $(if $(ENV),$(ENV),$(shell echo 'development'))

.PHONY: templ css air vite temple

ec:
	@echo $(ENV)

air: css-$(ENV) temple vite clean
	@echo "Starting air, templ, and css..."
	npx vite build
	air

css-development:
	npx tailwindcss -i views/css/global.css -o public/styles/global.css

css-production:
	npx tailwindcss -i views/css/global.css -o public/styles/global.css --minify

temple:
	templ generate --watch --proxy=http://localhost:1323 & \

vite:
	npx vite build

# The clean target removes the tmp directory used by air
clean:
	@rm -rf ./tmp
