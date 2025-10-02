package main

func main() {
	var cfg ConfigWorker
	HandlerUserInput(&cfg)
	ParseConfig(cfg.PathJsonCfg)
}
