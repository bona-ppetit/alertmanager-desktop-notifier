MAIN         = cmd/alertmanager-notifier/alertmanager-notifier.go

GO           ?= go


ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

.PHONY: alertmanager-notifier debug clean

all: alertmanager-notifier

alertmanager-notifier: ${MAIN}
	$(GO) build $<

debug: ${MAIN}
	$(GO) build -o alertmanager-notifier_$@ \
        -ldflags="-X 'main.debug=true' -X 'github.com/bona-ppetit/alertmanager-desktop-notifier/internal/alertparse.debug=true'" \
        $<

tests: $(wildcard test/*.go)
	$(foreach test,\
	$^,\
	go run $(test);)

install: alertmanager-notifier
	install -m 0555 $< /usr/bin

clean:
	@rm -f alertmanager-notifier
	@rm -f alertmanager-notifier_debug
