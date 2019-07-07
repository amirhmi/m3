	// Raise fd limits to nr_open system limit
	result, err := xos.RaiseProcessNoFileToNROpen()
	if err != nil {
		logger.Warn("unable to raise rlimit to no file fds limit",
			zap.Error(err))
	} else {
		logger.Info("raised rlimit no file fds limit",
			zap.Bool("required", result.RaisePerformed),
			zap.Uint64("sysNROpenValue", result.NROpenValue),
			zap.Uint64("noFileMaxValue", result.NoFileMaxValue),
			zap.Uint64("noFileCurrValue", result.NoFileCurrValue))