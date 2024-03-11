# Mapped Diagnostic Context (MDC) for Go

This module exposes a simple API for manipulating a map of key-value pairs
that can be used to enrich log messages with pieces of information that are
not available in the scope of the log message itself.

This module is inspired by the [Mapped Diagnostic Context][mdc] in the
[Slf4J][slf4j] library for Java. This module uses [timandy/routine], which is
essentially provides a ThreadLocal implementation for Go's goroutines. Version
0.6.0 is used.

See [example/main.go](example/main.go) for a simple example on how to use.

> [!CAUTION]
> 
> Be sure to check out the [support grid][support-grid] before you take it in use.

<!-- @formatter:off -->
<!-- Links -->
[mdc]: https://logback.qos.ch/manual/mdc.html
[slf4j]: https://slf4j.org/
[timandy/routine]: https://github.com/timandy/routine
[support-grid]: https://github.com/timandy/routine/tree/c7d040ea0795aad2075f959fbca7a51291f00716?tab=readme-ov-file#support-grid
<!-- @formatter:on -->
