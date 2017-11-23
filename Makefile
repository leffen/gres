bump:
	bump_version patch gres.go 

push: 
	git push origin --tags

release: test bump push
	git push --tags

test:
	go test . -v
