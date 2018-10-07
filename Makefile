krs_version := 0.1
git_version := `git rev-parse HEAD`
main_dir := `pwd`

.PHONY: gbuild gclean cbuild cpush release

###############################################################################
# commands related to Go testing and builds creating binaries
gbuild :
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.releaseVersion=$(krs_version)" -o ./out/krs_macos *.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.releaseVersion=$(krs_version)" -o ./out/krs_linux *.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.releaseVersion=$(krs_version)" -o ./out/krs_windows *.go

gclean :
	@rm out/krs_*

###############################################################################
# commands related to container image builds
crelease : cbuild cpush

cbuild :
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.releaseVersion=$(krs_version)" -o ./out/krs_linux *.go
	@docker build --build-arg krsv=$(krs_version) -t quay.io/mhausenblas/krs:$(krs_version) .
	@rm out/krs_linux

cpush :
	@docker push quay.io/mhausenblas/krs:$(krs_version)
