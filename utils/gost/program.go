package gost

import (
	"net/http"
	"os"
	"vorker/conf"

	"github.com/go-gost/core/logger"
	"github.com/go-gost/x/config"
	"github.com/go-gost/x/config/parsing"
	xmetrics "github.com/go-gost/x/metrics"
	"github.com/go-gost/x/registry"
	"github.com/judwhite/go-svc"
)

type program struct{}

func (p *program) Init(env svc.Environment) error {
	cfg := &config.Config{}
	if cfgFile != "" && conf.AppConfigInstance.RunMode == "master" {
		if err := cfg.ReadFile(cfgFile); err != nil {
			return err
		}
	}

	cmdCfg, err := buildConfigFromCmd(services, nodes)
	if err != nil {
		return err
	}
	cfg = p.mergeConfig(cfg, cmdCfg)

	if len(cfg.Services) == 0 && apiAddr == "" {
		if err := cfg.Load(); err != nil {
			return err
		}
	}

	if v := os.Getenv("GOST_API"); v != "" {
		cfg.API = &config.APIConfig{
			Addr: v,
		}
	}
	if v := os.Getenv("GOST_LOGGER_LEVEL"); v != "" {
		cfg.Log = &config.LogConfig{
			Level: v,
		}
	}
	if v := os.Getenv("GOST_PROFILING"); v != "" {
		cfg.Profiling = &config.ProfilingConfig{
			Addr: v,
		}
	}
	if v := os.Getenv("GOST_METRICS"); v != "" {
		cfg.Metrics = &config.MetricsConfig{
			Addr: v,
		}
	}

	if apiAddr != "" {
		cfg.API = &config.APIConfig{
			Addr: apiAddr,
		}
	}
	if debug {
		if cfg.Log == nil {
			cfg.Log = &config.LogConfig{}
		}
		cfg.Log.Level = string(logger.DebugLevel)
	}
	if metricsAddr != "" {
		cfg.Metrics = &config.MetricsConfig{
			Addr: metricsAddr,
		}
	}

	logger.SetDefault(logFromConfig(cfg.Log))

	if outputFormat != "" {
		if err := cfg.Write(os.Stdout, outputFormat); err != nil {
			return err
		}
		os.Exit(0)
	}

	parsing.BuildDefaultTLSConfig(cfg.TLS)

	config.Set(cfg)

	return nil
}

func (p *program) Start() error {
	log := logger.Default()
	cfg := config.Global()

	if cfg.API != nil {
		s, err := buildAPIService(cfg.API)
		if err != nil {
			return err
		}
		go func() {
			defer s.Close()
			log.Info("api service on ", s.Addr())
			log.Fatal(s.Serve())
		}()
	}
	if cfg.Profiling != nil {
		go func() {
			addr := cfg.Profiling.Addr
			if addr == "" {
				addr = ":6060"
			}
			log.Info("profiling server on ", addr)
			log.Fatal(http.ListenAndServe(addr, nil))
		}()
	}

	if cfg.Metrics != nil {
		xmetrics.Init(xmetrics.NewMetrics())
		if cfg.Metrics.Addr != "" {
			s, err := buildMetricsService(cfg.Metrics)
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				defer s.Close()
				log.Info("metrics service on ", s.Addr())
				log.Fatal(s.Serve())
			}()
		}
	}

	for _, svc := range buildService(cfg) {
		svc := svc
		go func() {
			svc.Serve()
		}()
	}

	return nil
}

func (p *program) Stop() error {
	for name, srv := range registry.ServiceRegistry().GetAll() {
		srv.Close()
		logger.Default().Debugf("service %s shutdown", name)
	}
	return nil
}

func (p *program) mergeConfig(cfg1, cfg2 *config.Config) *config.Config {
	if cfg1 == nil {
		return cfg2
	}
	if cfg2 == nil {
		return cfg1
	}

	cfg := &config.Config{
		Services:   append(cfg1.Services, cfg2.Services...),
		Chains:     append(cfg1.Chains, cfg2.Chains...),
		Hops:       append(cfg1.Hops, cfg2.Hops...),
		Authers:    append(cfg1.Authers, cfg2.Authers...),
		Admissions: append(cfg1.Admissions, cfg2.Admissions...),
		Bypasses:   append(cfg1.Bypasses, cfg2.Bypasses...),
		Resolvers:  append(cfg1.Resolvers, cfg2.Resolvers...),
		Hosts:      append(cfg1.Hosts, cfg2.Hosts...),
		Ingresses:  append(cfg1.Ingresses, cfg2.Ingresses...),
		Recorders:  append(cfg1.Recorders, cfg2.Recorders...),
		Limiters:   append(cfg1.Limiters, cfg2.Limiters...),
		CLimiters:  append(cfg1.CLimiters, cfg2.CLimiters...),
		RLimiters:  append(cfg1.RLimiters, cfg2.RLimiters...),
		TLS:        cfg1.TLS,
		Log:        cfg1.Log,
		API:        cfg1.API,
		Metrics:    cfg1.Metrics,
		Profiling:  cfg1.Profiling,
	}
	if cfg2.TLS != nil {
		cfg.TLS = cfg2.TLS
	}
	if cfg2.Log != nil {
		cfg.Log = cfg2.Log
	}
	if cfg2.API != nil {
		cfg.API = cfg2.API
	}
	if cfg2.Metrics != nil {
		cfg.Metrics = cfg2.Metrics
	}
	if cfg2.Profiling != nil {
		cfg.Profiling = cfg2.Profiling
	}

	return cfg
}
