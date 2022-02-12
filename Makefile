include go.mk

.PHONY: build-changelog
build-changelog:  ## Build the changelog with git-chglog
	@which git-chglog > /dev/null || (echo "Installing git-chglog@v0.15.1 ..."; go install github.com/git-chglog/git-chglog/cmd/git-chglog@v0.15.1 && echo -e "Installation complete!\n")
	mkdir -p ./build
	git-chglog -o ./build/CHANGELOG.md

.PHONY: clean
clean:  ## Clean build bundles
	-rm -rf ./build/bundles
