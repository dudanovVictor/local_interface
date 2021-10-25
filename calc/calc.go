package calc

type SimpleLog interface {
	Logf(s string, a ...interface{})
}

func Calculate(log SimpleLog) error {
	for i := 0; i < 5; i++ {
		log.Logf("number %d\n", i)
	}
	return nil
}
