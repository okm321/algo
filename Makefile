.PHONY: .check-abc-num
.check-abc-num:
ifndef ABC_NUM
	$(error ABC_NUM is required.)
endif

.PHONY: abc
abc: .check-abc-num
	@mkdir -p abc/$(ABC_NUM)
	@for problem in A B C D E F; do \
		mkdir -p abc/$(ABC_NUM)/$$problem; \
		cp template.go abc/$(ABC_NUM)/$$problem/main.go; \
		echo "Created abc/$(ABC_NUM)/$$problem/main.go"; \
	done
	@echo "# AtCoder Beginner Contest $(ABC_NUM)" > abc/$(ABC_NUM)/README.md
	@echo "" >> abc/$(ABC_NUM)/README.md
	@echo "[トップページ](https://atcoder.jp/contests/abc$(ABC_NUM))" >> abc/$(ABC_NUM)/README.md
	@echo "" >> abc/$(ABC_NUM)/README.md
	@echo "## 問題" >> abc/$(ABC_NUM)/README.md
	@echo "" >> abc/$(ABC_NUM)/README.md
	@for problem in A B C D E F; do \
		p=`echo $$problem | tr 'A-Z' 'a-z'`; \
		echo "- [$$problem - 問題](https://atcoder.jp/contests/abc$(ABC_NUM)/tasks/abc$(ABC_NUM)_$$p)" >> abc/$(ABC_NUM)/README.md; \
	done
	@echo "abc/$(ABC_NUM)/README.md を作成しました"
	@echo "完了しました！"
