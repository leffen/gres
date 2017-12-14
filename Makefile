bump:
	bump_version patch gres.go 

push: 
	git push origin master
	git push origin --tags --all

release: test bump push


test:
	go test ./... -cover -bench=. -test.benchtime=3s;
